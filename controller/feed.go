package controller

import (
	"douyin/service"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	service.Feed(c)
}
