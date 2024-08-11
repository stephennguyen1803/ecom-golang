package po

//po => persistent object
import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid;type:varchar(255);notnull;index:idx_uuid;unique;"` // UUID
	UserName string    `gorm:"column:user_name;type:varchar(255);"`                          // User name
	IsActive bool      `gorm:"column:is_active;type:boolean;default:true;"`                  // Is active
	Roles    []Role    `gorm:"many2many:go_user_roles"`                                      // Role
}

func (u *User) TableName() string {
	return "go_db_user"
}
