package service

import (
	"github.com/41197-yhkt/tiktok/tiktok-composite/gen/dal"
	"github.com/41197-yhkt/tiktok/tiktok-composite/gen/dal/query"
)

var q *query.Query

func Init() {
	q = query.Use(dal.DB.Debug())
}
