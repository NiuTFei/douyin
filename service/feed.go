package service

import (
	"context"
	"douyin/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {
	latestTime := c.Query("latest_time")
	token := c.Query("token")
	check := CheckToken(token)
	if latestTime == "" {
		latestTime = strconv.FormatInt(time.Now().Unix(), 10)
	}
	//从redis中读取得到最近30个视频ID，并根据ID从mysql数据库查询并返回
	videoIds, _ := rdb.ZRevRange(context.Background(), "feed", 0, 29).Result()
	videos := make([]repository.VideoDao, len(videoIds))
	for i, stringId := range videoIds {
		id, _ := strconv.ParseInt(stringId, 10, 64)
		var video repository.VideoDao
		db.Where("id = ?", id).Find(&video)
		videos[i] = video
	}

	var videoList []repository.Video
	if token != "" && check > 0 {
		videoList = VideoDaoToVideoList(check, videos)
	} else {
		videoList = VideoDaoToVideoListWithNoToken(videos)
	}

	c.JSON(http.StatusOK, repository.FeedResponse{
		Response:  repository.Response{StatusCode: 0, StatusMsg: "success"},
		VideoList: videoList,
	})
}
