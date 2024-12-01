// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"
)

// Account
type GoCrmUser struct {
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

type PreGoAccUser9999 struct {
	// User ID
	UserID uint64
	// User account
	UserAccount string
	// User nickname
	UserNickname sql.NullString
	// User avatar
	UserAvatar sql.NullString
	// User state: 0-Locked ,  1-Activated ,  2-Not Activated
	UserState uint8
	// Mobile phone number
	UserMobile sql.NullString
	// User gender: 0-Secret ,  1-Male ,  2-Female
	UserGender sql.NullInt16
	// User birthday
	UserBirthday sql.NullTime
	// User email address
	UserEmail sql.NullString
	// Authentication status: 0-Not Authenticated ,  1-Pending ,  2-Authenticated ,  3-Failed
	UserIsAuthentication uint8
	// Record creation time
	CreatedAt sql.NullTime
	// Record update time
	UpdatedAt sql.NullTime
}

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

type PreGoAccUserInfo9999 struct {
	// User ID
	UserID uint64
	// User account
	UserAccount string
	// User nickname
	UserNickname sql.NullString
	// User avatar
	UserAvatar sql.NullString
	// User state: 0-Locked, 1-Activated, 2-Not Activated
	UserState uint8
	// Mobile phone number
	UserMobile sql.NullString
	// User gender: 0-Secret, 1-Male, 2-Female
	UserGender sql.NullInt16
	// User birthday
	UserBirthday sql.NullTime
	// User email address
	UserEmail sql.NullString
	// Authentication status: 0-Not Authenticated, 1-Pending, 2-Authenticated, 3-Failed
	UserIsAuthentication uint8
	// Record creation time
	CreatedAt time.Time
	// Record updated time
	UpdatedAt time.Time
}

type PreGoAccUserVerify9999 struct {
	VerifyID        uint64
	VerifyOtp       string
	VerifyKey       string
	VerifyKeyHash   string
	VerifyType      sql.NullInt32
	IsVerfified     sql.NullBool
	IsDeleted       sql.NullBool
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
