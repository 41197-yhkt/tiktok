package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id   int64  `gorm:"primary_key"`
	Name string `gorm:"column:name"`
}

type Video struct {
	gorm.Model
	Id            uint      `gorm:"primary_key"`
	AuthorId      int64     `gorm:"column:author_id"`
	PlayUrl       string    `gorm:"column:play_url"`
	CoverUrl      string    `gorm:"column:cover_url"`
	FavoriteCount string    `gorm:"column:favorite_count"`
	CommentCount  string    `gorm:"column:comment_count"`
	Title         string    `gorm:"column:title"`
	Created_at    time.Time `gorm:"column:created_at"`
	Updated_at    time.Time `gorm:"column:updated_at"`
}

type UserFavorite struct {
	gorm.Model
	Id      uint  `gorm:"primary_key"`
	UserId  int64 `gorm:"column:user_id"`
	VideoId int64 `gorm:"column:video_id"`
}
