package service

import (
	"context"
	"douyin/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	data, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, repository.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	userId := CheckToken(token)
	if userId == 0 || userId == -1 {
		c.JSON(http.StatusOK, repository.Response{
			StatusCode: 1,
			StatusMsg:  "Error",
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./public/", finalName)
	//上传视频
	err2 := c.SaveUploadedFile(data, saveFile)
	if err2 != nil {
		c.JSON(http.StatusOK, repository.Response{
			StatusCode: 1,
			StatusMsg:  err2.Error(),
		})
		return
	}

	//视频信息添加到数据库，保存到redis
	playUrl := "http://10.180.22.64:8080/static/" + finalName
	newVideo := repository.VideoDao{AuthorId: userId, PlayUrl: playUrl}
	db.Create(&newVideo)
	rdb.ZAdd(context.Background(), "feed", redis.Z{Score: float64(time.Now().Unix()), Member: newVideo.Id})
	c.JSON(http.StatusOK, repository.Response{
		StatusCode: 0,
		StatusMsg:  "视频上传成功",
	})
}

func PublishList(token string, userId string) repository.VideoListResponse {
	check := CheckToken(token)
	if check == 0 || check == -1 {
		return repository.VideoListResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "用户未登陆或不存在"},
		}
	}

	authorId, _ := strconv.ParseInt(userId, 10, 64)
	var videos []repository.VideoDao
	db.Where("author_id = ?", authorId).Find(&videos)
	videoList := VideoDaoToVideoList(check, videos)

	return repository.VideoListResponse{
		Response: repository.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videoList,
	}
}
