-- name: GetUser :one
SELECT user_id, user_account, user_email, user_state, user_mobile,
user_is_authentication, created_at, updated_at from `pre_go_acc_user_9999`
where user_id = ? LIMIT 1;

-- name: GetUsers :many
SELECT user_id, user_account, user_email, user_state, user_mobile,
user_is_authentication, created_at, updated_at from `pre_go_acc_user_9999`
where user_id IN(?);

-- name: AddUserHaveUserId :execresult
INSERT INTO `pre_go_acc_user_info_9999` (
    user_id,
    user_account, 
    user_nickname, 
    user_avatar, 
    user_state, 
    user_mobile,
    user_gender,
    user_birthday,
    user_email, user_is_authentication) 
 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

 -- name: AddUser :execresult
 INSERT INTO `pre_go_acc_user_info_9999` (
    user_account, 
    user_nickname, 
    user_avatar, 
    user_state, 
    user_mobile,
    user_gender,
    user_birthday,
    user_email, user_is_authentication) 
 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);   

-- name: EditUserByUserId :execresult
Update `pre_go_acc_user_info_9999` 
SET user_nickname = ?, user_avatar = ?, user_state = ?, user_mobile = ?,
user_gender = ?, user_birthday = ?, user_email = ? , updated_at = NOW()
where user_id = ? And user_is_authentication = 1;