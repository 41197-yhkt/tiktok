package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `gorm:"column:name"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

type Video struct {
	gorm.Model
	AuthorId int64  `gorm:"column:author_id"`
	PlayUrl  string `gorm:"column:play_url"`
	CoverUrl string `gorm:"column:cover_url"`
	Title    string `gorm:"column:title"`
}

type UserFavorite struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}

type Comment struct {
	gorm.Model
	UserId  int64  `gorm:"column:user_id"`
	VideoId int64  `gorm:"column:video_id"`
	Content string `gorm:"column:content"`
}
