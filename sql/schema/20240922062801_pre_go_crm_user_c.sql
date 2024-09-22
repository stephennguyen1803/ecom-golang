-- +goose Up
-- +goose StatementBegin
CREATE Table `pre_go_crm_user_c` (
	`usr_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
	`usr_email` varchar(30) NOT NULL Default '' comment 'Email',
	`usr_phone` varchar(15) not null default '' comment 'Phone Number',
	`usr_username` varchar(30) not null default '' comment 'Username',
	`usr_password` varchar(32) not null default '' comment 'Password',
	`usr_created_at` int(11)	not null default '0' comment 'Created Time',
	`usr_updated_at` int(11) not null default '0' comment 'Updated Time',
	`usr_created_ip_at` varchar(12) NOt null default '' comment 'Creation IP',
	`usr_last_login_at` int(11) not null default '0' comment 'Last Login Time',
	`usr_last_login_ip` varchar(12) not null default '' comment 'Last Login IP',
	`usr_login_times` int(11) not null default '0' comment 'Login Times',
	`usr_status` tinyint(1) not null default '0' comment 'Status 1:enable, 0:disable, -1:deleted',
	PRIMARY KEY (`usr_id`),
	KEY `idx_email` (`usr_email`),
	KEY `idx_phone` (`usr_phone`),
	KEY `idx_username` (`usr_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 Comment='Account';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop Table `pre_go_crm_user_c`;
-- +goose StatementEnd
