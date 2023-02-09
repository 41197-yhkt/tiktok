package dal

import (
	"sync"

	"github.com/41197-yhkt/tiktok/composite/gen/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func Init() {
	once.Do(func() {
		DB = ConnctDB().Debug()
		_ = DB.AutoMigrate(&model.UserFavorite{}, &model.Comment{})
	})
}

func ConnctDB() (conn *gorm.DB) {
	conn, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:6666)/tiktok_db?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return conn
}
