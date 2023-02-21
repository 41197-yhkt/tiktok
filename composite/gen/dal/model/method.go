package model

import (
	"time"

	"gorm.io/gen"
)

type UserMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)
}

type VideoMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)

	//sql(select * from @@table where updated_at < @lastTime order by updated_at limit @limit)
	FindByUpdatedtime(lastTime time.Time, limit int) ([]*gen.T, error)

	//sql(updata @@table set favorite_count = favorite_count+1 where id = @vedioId)
	IncreaseFavoriteCount(vedioId int64) error

	//sql(updata @@table set favorite_count = favorite_count-1 where id = @vedioId)
	DecreaseFavoriteCount(vedioId int64) error

	//sql(updata @@table set comment_count = comment_count+1 where id = @vedioId)
	IncreaseCommentCount(vedioId int64) error

	//sql(updata @@table set comment_count = comment_count-1 where id = @vedioId)
	DecreaseCommentCount(vedioId int64) error
}

type UserFavoriteMethod interface {
	//sql(select video_id from @@table where user_id = @userId)
	FindByUserid(userId int64) ([]*gen.T, error)

	//sql(select user_id from @@table where video_id = @videoId)
	FindByVideoid(videoId int64) ([]*gen.T, error)

	//sql(select * from @@table where video_id = @videoId and user_id = @userId)
	FindByUseridAndVideoid(userId, videoId int64) error

	//sql(select count(*) from @@table where video_id = @videoId)
	CountByVideoid(videoId int64) (int64, error)

	//sql(delete from @@table where user_id = @userId and video_id = @videoId)
	DeleteByUseridAndVideoid(userId, videoId int64) error
}

type CommentMethod interface {
	//where(id=@id)
	FindByID(id int64) (gen.T, error)

	//where(user_id = @userId)
	FindByUserid(userId int64) ([]gen.T, error)

	//where(video_id = @videoId)
	FindByVideoid(videoId int64) ([]*gen.T, error)

	//where(user_id = @userId and video_id = @videoId)
	FindByUseridAndVideoid(userId, videoId int64) ([]gen.T, error)

	// sql(delete from @@table where id = @id)
	DeleteById(id int64) error

	//sql(select count(*) from @@table where video_id = @videoId)
	CountByVideoid(videoId int64) (int64, error)

	//sql(SELECT LAST_INSERT_ID())
	LastInsertID() (uint, error)
}
