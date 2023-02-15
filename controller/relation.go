package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserIdStr := c.Query("to_user_id")
	actionType := c.Query("action_type")
	c.JSON(http.StatusOK, service.RelationAction(token, toUserIdStr, actionType))
}

func FollowList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	token := c.Query("token")
	c.JSON(http.StatusOK, service.FollowList(userIdStr, token))
}

func FollowerList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	token := c.Query("token")
	c.JSON(http.StatusOK, service.FollowerList(userIdStr, token))
}

func FriendList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	token := c.Query("token")
	c.JSON(http.StatusOK, service.FriendList(userIdStr, token))
}
