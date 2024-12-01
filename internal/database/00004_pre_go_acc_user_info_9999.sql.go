// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 00004_pre_go_acc_user_info_9999.sql

package database

import (
	"context"
	"database/sql"
)

const addUserHaveUserId = `-- name: AddUserHaveUserId :execresult
INSERT INTO ` + "`" + `pre_go_acc_user_info_9999` + "`" + ` (
    user_id,
    user_account, 
    user_nickname, 
    user_avatar, 
    user_state, 
    user_mobile,
    user_gender,
    user_birthday,
    user_email, user_is_authentication) 
 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type AddUserHaveUserIdParams struct {
	UserID               uint64
	UserAccount          string
	UserNickname         sql.NullString
	UserAvatar           sql.NullString
	UserState            uint8
	UserMobile           sql.NullString
	UserGender           sql.NullInt16
	UserBirthday         sql.NullTime
	UserEmail            sql.NullString
	UserIsAuthentication uint8
}

func (q *Queries) AddUserHaveUserId(ctx context.Context, arg AddUserHaveUserIdParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addUserHaveUserId,
		arg.UserID,
		arg.UserAccount,
		arg.UserNickname,
		arg.UserAvatar,
		arg.UserState,
		arg.UserMobile,
		arg.UserGender,
		arg.UserBirthday,
		arg.UserEmail,
		arg.UserIsAuthentication,
	)
}

const editUserByUserId = `-- name: EditUserByUserId :execresult
Update ` + "`" + `pre_go_acc_user_info_9999` + "`" + ` 
SET user_nickname = ?, user_avatar = ?, user_state = ?, user_mobile = ?,
user_gender = ?, user_birthday = ?, user_email = ? , updated_at = NOW()
where user_id = ? And user_is_authentication = 1
`

type EditUserByUserIdParams struct {
	UserNickname sql.NullString
	UserAvatar   sql.NullString
	UserState    uint8
	UserMobile   sql.NullString
	UserGender   sql.NullInt16
	UserBirthday sql.NullTime
	UserEmail    sql.NullString
	UserID       uint64
}

func (q *Queries) EditUserByUserId(ctx context.Context, arg EditUserByUserIdParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, editUserByUserId,
		arg.UserNickname,
		arg.UserAvatar,
		arg.UserState,
		arg.UserMobile,
		arg.UserGender,
		arg.UserBirthday,
		arg.UserEmail,
		arg.UserID,
	)
}

const getUser = `-- name: GetUser :one
SELECT user_id, user_account, user_email, user_state, user_mobile,
user_is_authentication, created_at, updated_at from ` + "`" + `pre_go_acc_user_9999` + "`" + `
where user_id = ? LIMIT 1
`

type GetUserRow struct {
	UserID               uint64
	UserAccount          string
	UserEmail            sql.NullString
	UserState            uint8
	UserMobile           sql.NullString
	UserIsAuthentication uint8
	CreatedAt            sql.NullTime
	UpdatedAt            sql.NullTime
}

func (q *Queries) GetUser(ctx context.Context, userID uint64) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i GetUserRow
	err := row.Scan(
		&i.UserID,
		&i.UserAccount,
		&i.UserEmail,
		&i.UserState,
		&i.UserMobile,
		&i.UserIsAuthentication,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT user_id, user_account, user_email, user_state, user_mobile,
user_is_authentication, created_at, updated_at from ` + "`" + `pre_go_acc_user_9999` + "`" + `
where user_id IN(?)
`

type GetUsersRow struct {
	UserID               uint64
	UserAccount          string
	UserEmail            sql.NullString
	UserState            uint8
	UserMobile           sql.NullString
	UserIsAuthentication uint8
	CreatedAt            sql.NullTime
	UpdatedAt            sql.NullTime
}

func (q *Queries) GetUsers(ctx context.Context, userID uint64) ([]GetUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.UserID,
			&i.UserAccount,
			&i.UserEmail,
			&i.UserState,
			&i.UserMobile,
			&i.UserIsAuthentication,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
