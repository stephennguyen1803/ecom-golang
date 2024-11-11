-- name: GetAccUserByAccount :one
SELECT user_id, user_account, user_email, user_state, user_mobile,
user_is_authentication from `pre_go_acc_user_9999`
where user_account = ?;

-- name: GetAccUserByUserId :one
SELECT user_id, user_account, user_email, user_state, user_mobile,
user_is_authentication from `pre_go_acc_user_9999`
where user_id = ?;

-- name: GetAccUserByUserEmail :one
SELECT user_id, user_account, user_email, user_state, user_mobile,
user_is_authentication from `pre_go_acc_user_9999`
where user_email = ?;

-- name: InsertAccUser :exec
INSERT INTO pre_go_acc_user_9999 (
    user_account, 
    user_nickname, 
    user_avatar, 
    user_state, 
    user_mobile,
    user_gender,
    user_birthday,
    user_email,
    user_is_authentication,
    created_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW());