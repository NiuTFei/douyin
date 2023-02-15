package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, service.Login(username, password))
}

func UserInfo(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	c.JSON(http.StatusOK, service.UserInfo(userId, token))
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, service.Register(username, password))
}
