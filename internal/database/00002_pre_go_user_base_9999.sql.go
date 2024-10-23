// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 00002_pre_go_user_base_9999.sql

package database

import (
	"context"
	"database/sql"
)

const addUserBase = `-- name: AddUserBase :exec
INSERT INTO pre_go_acc_user_base_9999 (
    user_account, 
    user_password, 
    user_salt, 
    user_status, 
    user_created_at, 
    user_updated_at
) VALUES (?, ?, ?, ?, NOW(), NOW())
`

type AddUserBaseParams struct {
	UserAccount  string
	UserPassword string
	UserSalt     string
	UserStatus   bool
}

func (q *Queries) AddUserBase(ctx context.Context, arg AddUserBaseParams) error {
	_, err := q.db.ExecContext(ctx, addUserBase,
		arg.UserAccount,
		arg.UserPassword,
		arg.UserSalt,
		arg.UserStatus,
	)
	return err
}

const checkUserBaseExist = `-- name: CheckUserBaseExist :one
SELECT COUNT(*) 
FROM ` + "`" + `pre_go_acc_user_base_9999` + "`" + `
WHERE user_account = ?
`

func (q *Queries) CheckUserBaseExist(ctx context.Context, userAccount string) (int64, error) {
	row := q.db.QueryRowContext(ctx, checkUserBaseExist, userAccount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getOneUserInfo = `-- name: GetOneUserInfo :one
SELECT user_id, user_account, user_password, user_salt
FROM ` + "`" + `pre_go_acc_user_base_9999` + "`" + `
WHERE user_account = ?
`

type GetOneUserInfoRow struct {
	UserID       uint64
	UserAccount  string
	UserPassword string
	UserSalt     string
}

func (q *Queries) GetOneUserInfo(ctx context.Context, userAccount string) (GetOneUserInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getOneUserInfo, userAccount)
	var i GetOneUserInfoRow
	err := row.Scan(
		&i.UserID,
		&i.UserAccount,
		&i.UserPassword,
		&i.UserSalt,
	)
	return i, err
}

const getOneUserInfoAdmin = `-- name: GetOneUserInfoAdmin :one
SELECT user_id, user_account, user_password, 
       user_salt, user_login_time, user_login_ip, user_logout_time,
       user_status, user_created_at, user_updated_at
FROM ` + "`" + `pre_go_acc_user_base_9999` + "`" + `
WHERE user_account = ?
`

func (q *Queries) GetOneUserInfoAdmin(ctx context.Context, userAccount string) (PreGoAccUserBase9999, error) {
	row := q.db.QueryRowContext(ctx, getOneUserInfoAdmin, userAccount)
	var i PreGoAccUserBase9999
	err := row.Scan(
		&i.UserID,
		&i.UserAccount,
		&i.UserPassword,
		&i.UserSalt,
		&i.UserLoginTime,
		&i.UserLoginIp,
		&i.UserLogoutTime,
		&i.UserStatus,
		&i.UserCreatedAt,
		&i.UserUpdatedAt,
	)
	return i, err
}

const loginUserBase = `-- name: LoginUserBase :exec
UPDATE pre_go_acc_user_base_9999
SET user_login_time = NOW(),
    user_login_ip = ?,
    user_updated_at = NOW()
WHERE user_account = ? AND user_password = ?
`

type LoginUserBaseParams struct {
	UserLoginIp  sql.NullString
	UserAccount  string
	UserPassword string
}

func (q *Queries) LoginUserBase(ctx context.Context, arg LoginUserBaseParams) error {
	_, err := q.db.ExecContext(ctx, loginUserBase, arg.UserLoginIp, arg.UserAccount, arg.UserPassword)
	return err
}

const logoutUserBase = `-- name: LogoutUserBase :exec
UPDATE pre_go_acc_user_base_9999
SET user_logout_time = NOW(),
    user_updated_at = NOW()
WHERE user_account = ?
`

func (q *Queries) LogoutUserBase(ctx context.Context, userAccount string) error {
	_, err := q.db.ExecContext(ctx, logoutUserBase, userAccount)
	return err
}