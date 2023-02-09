package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // ID uint CreatAt time.Time UpdateAt time.Time DeleteAt gorm.DeleteAt If it is repeated with the definition will be ignored
	Name          string `gorm:"column:user_name"`
	Password      string `gorm:"column:user_pwd_hash"`
	FollowCount   int    `gorm:"column:follow_count"`
	FollowerCount int    `gorm:"column:follower_count"`
}

type UserRelation struct {
	gorm.Model      // ID uint CreatAt time.Time UpdateAt time.Time DeleteAt gorm.DeleteAt If it is repeated with the definition will be ignored
	FollowFrom uint `gorm:"column:follow_from"`
	FollowTo   uint `gorm:"column:follow_to"`
}
