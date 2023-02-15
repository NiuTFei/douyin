package service

import (
	"douyin/repository"
	"gorm.io/gorm"
	"strconv"
)

func FavoriteAction(token string, videoId string, actionType string) repository.Response {
	//actionType 1-点赞 2-取消点赞
	check := CheckToken(token)
	if check == 0 || check == -1 {
		return repository.Response{
			StatusCode: 1, StatusMsg: "用户未登陆或不存在",
		}
	}

	favoriteVideoId, _ := strconv.ParseInt(videoId, 10, 64)
	isFavorite, _ := strconv.ParseInt(actionType, 10, 64)
	video := repository.VideoDao{Id: favoriteVideoId}

	if isFavorite == 1 {
		//创建点赞记录
		newFavoriteRecord := repository.FavoriteDao{UserId: check, FavoriteVideoId: favoriteVideoId}
		db.Create(&newFavoriteRecord)
		//更新视频点赞字段
		db.Model(&video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		return repository.Response{StatusCode: 0, StatusMsg: "点赞成功"}
	}

	//取消点赞
	var deleteFavoriteRecord repository.FavoriteDao
	db.Where("user_id = ? and favorite_video_id = ?", check, favoriteVideoId).Delete(&deleteFavoriteRecord)
	db.Model(&video).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
	return repository.Response{StatusCode: 0, StatusMsg: "取消点赞"}
}

func FavoriteList(userId string, token string) repository.FavoriteListResponse {
	check := CheckToken(token)
	if check == 0 || check == -1 {
		return repository.FavoriteListResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "用户未登陆或不存在"},
		}
	}
	id, _ := strconv.ParseInt(userId, 10, 64)

	var favoriteVideosDao []repository.FavoriteDao
	db.Where("user_id = ?", id).Find(&favoriteVideosDao)
	//Todo 考虑查询结果为空

	favoriteVideos := make([]repository.VideoDao, len(favoriteVideosDao))
	for i, favoriteVideoDao := range favoriteVideosDao {
		var favoriteVideo repository.VideoDao
		db.Where("id = ?", favoriteVideoDao.FavoriteVideoId).Find(&favoriteVideo)
		favoriteVideos[i] = favoriteVideo
	}

	favoriteVideoList := VideoDaoToVideoList(check, favoriteVideos)
	return repository.FavoriteListResponse{
		Response: repository.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: favoriteVideoList,
	}
}
