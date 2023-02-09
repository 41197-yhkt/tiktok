package model

import "gorm.io/gen"

type UserMethod interface {
	//where("user_name=@user_name")
	FindByUserName(user_name string) (gen.T, error)
	//where("id=@id")
	FindByUserID(id uint) (gen.T, error)
	//update @@table
	//	{{set}}
	//		update_time=now(),
	//		{{if follow_count > 0}}
	//			follow_count=@follow_count
	//		{{end}}
	//	{{end}}
	// where id=@id
	UpdateUserFollowCount(id uint, follow_count int) error
	//update @@table
	//	{{set}}
	//		update_time=now(),
	//		{{if follower_count != 0}}
	//			follower_count=@follower_count
	//		{{end}}
	//	{{end}}
	// where id=@id
	UpdateUserFollowerCount(id uint, follower_count int) error
}

type UserRelationMethod interface {
	//where("id=@id")
	FindByID(id uint) (gen.T, error)
	//where("follow_from=@follow_from")
	FindByFollowFrom(follow_from uint) ([]gen.T, error)
	//where("follow_to=@follow_to")
	FindByFollowTo(follow_to uint) ([]gen.T, error)
	//where("follow_from=@follow_from and follow_to=@follow_to")
	FindByFollowFromAndFollowTo(follow_from uint, follow_to uint) (gen.T, error)
	// //select follow_to from @@table where (follow_from=@follow_from")
	// getFollowIdListByFollowFrom(follow_to uint) ([]int, error)
	// //select follow_from from @@table where (follow_to=@follow_to")
	// getFollowerIdListByFollowTo(follow_from uint) ([]int, error)
}
