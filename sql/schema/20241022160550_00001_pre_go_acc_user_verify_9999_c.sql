-- +goose Up
-- +goose StatementBegin
-- Need code create schema for table 00001_pre_go_acc_user_verify_9999
CREATE TABLE IF NOT EXISTS `pre_go_acc_user_verify_9999` (
    `verify_id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `verify_otp` VARCHAR(255) NOT NULL,
    `verify_key` VARCHAR(255) NOT NULL,
    `verify_key_hash` VARCHAR(255) NOT NULL,
    `verify_type` VARCHAR(255) NOT NULL,
    `is_verfified` TINYINT(1) NOT NULL DEFAULT 0,
    `is_deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `verify_created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `verify_updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`verify_id`),
    UNIQUE KEY `verify_key_hash_idx` (`verify_key_hash`),
    UNIQUE KEY `verify_key_idx` (`verify_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_acc_user_verify_9999`;
-- +goose StatementEnd
