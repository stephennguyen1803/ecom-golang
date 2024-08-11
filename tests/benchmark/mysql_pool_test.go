package benchmark

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid;type:varchar(255);notnull;index:idx_uuid;unique;"` // UUID
	UserName string    `gorm:"column:user_name;type:varchar(255);"`                          // User name
	IsActive bool      `gorm:"column:is_active;type:boolean;default:true;"`                  // Is active
}

func (u *User) TableName() string {
	return "go_db_user_test"
}

func insertRecords(b *testing.B, db *gorm.DB) {
	user := User{
		UUID:     uuid.New(),
		UserName: "test",
		IsActive: true,
	}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err, "Failed to insert record")
	}
}

// go test -bench=. -benchmem
func BenchmarkMaxOpenConnIs1(b *testing.B) {
	dns := "root:admin@tcp(127.0.0.1:3306)/test_database?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		log.Fatal(err, "Failed to connect to database")
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{})
	}

	db.Migrator().CreateTable(&User{})
	db.AutoMigrate(&User{})

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err, "Failed to set pool")
	}

	// thiet lap 1 connect database
	sqlDb.SetMaxOpenConns(1)
	defer sqlDb.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecords(b, db)
		}
	})
}

// go test -bench=. -benchmem
func BenchmarkMaxOpenConnIs10(b *testing.B) {
	dns := "root:admin@tcp(127.0.0.1:3306)/test_database?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		log.Fatal(err, "Failed to connect to database")
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&User{})
	}

	db.Migrator().CreateTable(&User{})
	db.AutoMigrate(&User{})

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err, "Failed to set pool")
	}

	// thiet lap 10 connect database
	sqlDb.SetMaxOpenConns(10)
	defer sqlDb.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecords(b, db)
		}
	})
}
