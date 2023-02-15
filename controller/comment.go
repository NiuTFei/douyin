package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommentAction(c *gin.Context) {
	service.Comment(c)
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	c.JSON(http.StatusOK, service.CommentList(token, videoId))
}
