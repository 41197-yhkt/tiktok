package pack

import (
	"github.com/41197-yhkt/tiktok-composite/gen/dal/model"
)

func VideoAndVideoAuthorIds(userfavorites []*model.UserFavorite) ([]int64, []int64) {
	videoIds, videoAuthorIds := make([]int64, 0), make([]int64, 0)
	if len(userfavorites) == 0 {
		return videoIds, videoAuthorIds
	}
	for _, userfavorite := range userfavorites {
		videoIds = append(videoIds, userfavorite.UserId)
		videoAuthorIds = append(videoAuthorIds, userfavorite.VideoId)
	}
	return videoIds, videoAuthorIds
}
