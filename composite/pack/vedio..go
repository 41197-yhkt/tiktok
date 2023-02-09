package pack

import (
	"github.com/41197-yhkt/tiktok/composite/gen/dal/model"
	"github.com/41197-yhkt/tiktok/composite/kitex_gen/composite"
	"github.com/41197-yhkt/tiktok/user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok/video/kitex_gen/video"
)

// 打包 video
// TODO: author 不应该是 *model.User 而是 *User.User
func Video(video *video.Video, author *user.User) *composite.Video {
	if video == nil || author == nil {
		return nil
	}

	return &composite.Video{
		Id: video.Id,
		Author: &composite.User{
			Id:            int64(author.Id),
			Name:          author.Name,
			FollowCount:   author.FollowCount,
			FollowerCount: author.FollowerCount,
			// TODO: 这里也不应该自己传参，后续从 user 那直接拿
			IsFollow: false,
		},
		PlayUrl:  video.PlayUrl,
		CoverUrl: video.CoverUrl,
		Title:    video.Title,
	}
}

// 打包 video list
// TODO: authors 不应该是 []*model.User 而是 []*User.User
func Videos(videos []*video.Video, authors []*user.User) []*composite.Video {
	res := make([]*composite.Video, 0)
	for i := 0; i < len(videos); i++ {
		if v := Video(videos[i], authors[i]); v != nil {
			res = append(res, v)
		}
	}
	return res
}

func VideoAuthorIds(videos []*model.Video) []int64 {
	res := make([]int64, 0)
	for _, v := range videos {
		if v != nil {
			res = append(res, v.AuthorId)
		}
	}
	return res
}

func VideoIds(videos []*model.Video) []int64 {
	res := make([]int64, 0)
	for _, v := range videos {
		if v != nil {
			res = append(res, int64(v.ID))
		}
	}
	return res
}
