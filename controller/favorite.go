package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	c.JSON(http.StatusOK, service.FavoriteList(userId, token))
}

func FavoriteAction(c *gin.Context) {
	videoId := c.Query("video_id")
	token := c.Query("token")
	actionType := c.Query("action_type")
	c.JSON(http.StatusOK, service.FavoriteAction(token, videoId, actionType))
}
