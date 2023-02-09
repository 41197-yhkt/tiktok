package service

import (
	dal "github.com/41197-yhkt/tiktok-video/gen/dal"
	query "github.com/41197-yhkt/tiktok-video/gen/dal/query"
)

var q *query.Query

func Init() {
	q = query.Use(dal.DB.Debug())
}
