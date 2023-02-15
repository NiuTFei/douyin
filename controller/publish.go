package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Publish(c *gin.Context) {
	service.Publish(c)
}

func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	c.JSON(http.StatusOK, service.PublishList(token, userId))
}
