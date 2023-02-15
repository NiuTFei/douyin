package main

import (
	"douyin/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(server *gin.Engine) {

	server.Static("/static", "./public")

	server.GET("hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	apiRouter := server.Group("/douyin")

	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.POST("/message/action/", controller.MessageAction)
	apiRouter.GET("/message/chat/", controller.MessageChat)

}
