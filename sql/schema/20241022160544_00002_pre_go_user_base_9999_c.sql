-- +goose Up
-- +goose StatementBegin
-- Need to create table 00002_pre_go_user_base_9999
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_base_9999` (
    `user_id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_account` VARCHAR(255) NOT NULL,
    `user_password` VARCHAR(255) NOT NULL,
    `user_salt` VARCHAR(255) NOT NULL,
    `user_login_time` TIMESTAMP NULL DEFAULT NULL,
    `user_login_ip` VARCHAR(255) NULL DEFAULT NULL,
    `user_logout_time` TIMESTAMP NULL DEFAULT NULL,
    `user_status` TINYINT(1) NOT NULL DEFAULT 0,
    `user_created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `user_updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `user_account_idx` (`user_account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
