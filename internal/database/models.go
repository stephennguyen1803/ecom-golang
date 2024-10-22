// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

type PreGoAccUserBase9999 struct {
	UserID         uint64
	UserAccount    string
	UserPassword   string
	UserSalt       string
	UserLoginTime  sql.NullTime
	UserLoginIp    sql.NullString
	UserLogoutTime sql.NullTime
	UserStatus     bool
	UserCreatedAt  time.Time
	UserUpdatedAt  time.Time
}

type PreGoAccUserVerify9999 struct {
	VerifyID        uint64
	VerifyOtp       string
	VerifyKey       string
	VerifyKeyHash   string
	VerifyType      string
	IsVerfified     bool
	IsDeleted       bool
	VerifyCreatedAt time.Time
	VerifyUpdatedAt time.Time
}

// Account
type PreGoCrmUserC struct {
	// Account ID
	UsrID uint32
	// Email
	UsrEmail string
	// Phone Number
	UsrPhone string
	// Username
	UsrUsername string
	// Password
	UsrPassword string
	// Created Time
	UsrCreatedAt int32
	// Updated Time
	UsrUpdatedAt int32
	// Creation IP
	UsrCreatedIpAt string
	// Last Login Time
	UsrLastLoginAt int32
	// Last Login IP
	UsrLastLoginIp string
	// Login Times
	UsrLoginTimes int32
	// Status 1:enable, 0:disable, -1:deleted
	UsrStatus bool
}
