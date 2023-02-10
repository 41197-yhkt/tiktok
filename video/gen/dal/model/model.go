package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"column:name"`
}

type Video struct {
	gorm.Model
	AuthorId      int64  `gorm:"column:author"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount string `gorm:"column:favorite_count"`
	CommentCount  string `gorm:"column:comment_count"`
	Title         string `gorm:"column:title"`
}

type UserFavorite struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}
