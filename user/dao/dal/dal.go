package dal

import (
	"fmt"
	"sync"

	"github.com/41197-yhkt/tiktok-user/dao/dal/model"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DB = ConnectDB().Debug()
		_ = DB.AutoMigrate(&model.User{}, &model.UserRelation{})
	})
}

func ConnectDB() (conn *gorm.DB) {
	dsn := "root:123456@tcp(localhost:6666)/tiktok_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
