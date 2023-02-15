package controller

import (
	"douyin/repository"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	service.MessageClient(token, toUserId, content)
	c.JSON(http.StatusOK, repository.Response{StatusCode: 0, StatusMsg: "消息发送成功"})
}

func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	c.JSON(http.StatusOK, service.MessageChat(token, toUserId))
}
