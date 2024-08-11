package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id;type:int;notnull;primarykey;autoIncrement;comment:'Primary Key is ID'"` // ID
	RoleName string `gorm:"column:role_name"`                                                                // Role name
	RoleNote string `gorm:"column:role_note;type:text"`                                                      // Role note
}

func (r *Role) TableName() string {
	return "go_db_role"
}
