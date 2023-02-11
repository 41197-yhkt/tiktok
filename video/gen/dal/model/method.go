package model

import (
	"gorm.io/gen"
)

type UserFavoriteMethod interface {
	// //sql(select video_id from @@table where user_id = @userId)
	// FindByUserid(userId int64) (gen.T, error)

	// //sql(select user_id from @@table where video_id = @videoId)
	// FindByVideoid(videoId int64) (gen.T, error)

	//sql(select * from @@table where video_id = @videoId and user_id = @userId)
	FindByUseridAndVideoid(userId, videoId int64) error

	//sql(select count(*) from @@table where video_id = @videoId)
	CountByVideoid(videoId int64) (int64, error)

	//sql(select count(*) from @@table where user_id = @userId)
	CountByUserid(userId int64) (int64, error)
}

type UserMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)
}

type VideoMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)

	//sql(select * from @@table where author = @Authorid)
	FindByAuthorId(Authorid int) ([]*gen.T, error)
}
