package service

import (
	"douyin/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Comment(c *gin.Context) {
	token := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	var response repository.CommentResponse
	check := CheckToken(token)
	if check == 0 || check == -1 {
		c.JSON(http.StatusOK, repository.CommentResponse{Response: repository.Response{StatusCode: 1, StatusMsg: "用户未登陆或不存在"}})
		return
	}
	var user repository.UserDao
	db.Where("id = ?", check).Find(&user)
	video := repository.VideoDao{Id: videoId}
	if actionType == "1" {
		//添加评论
		commentText := c.Query("comment_text")
		commentRecord := repository.CommentDao{UserId: check, VideoId: videoId, Content: commentText}
		db.Create(&commentRecord)
		//更新视频评论数量
		db.Model(&video).Update("comment_count", gorm.Expr("comment_count + ?", 1))
		response = repository.CommentResponse{
			Response: repository.Response{StatusCode: 0, StatusMsg: "评论成功"},
			Comment: repository.Comment{
				Id:         commentRecord.Id,
				User:       UserDaoToUser(user),
				Content:    commentText,
				CreateDate: commentRecord.CreatedAt.String(),
			},
		}
	} else {
		//删除评论
		commentIdStr := c.Query("comment_id")
		commentId, _ := strconv.ParseInt(commentIdStr, 10, 64)
		var comment repository.CommentDao
		db.Where("id = ?", commentId).Find(&comment)
		db.Delete(&comment)
		db.Model(&video).Update("comment_count", gorm.Expr("comment_count - ?", 1))
		response = repository.CommentResponse{
			Response: repository.Response{StatusCode: 0, StatusMsg: "删除评论成功"},
			Comment: repository.Comment{
				Id:         comment.Id,
				User:       UserDaoToUser(user),
				Content:    comment.Content,
				CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			},
		}
	}
	c.JSON(http.StatusOK, response)
}

func CommentList(token string, videoIdStr string) repository.CommentListResponse {
	check := CheckToken(token)
	if check == 0 || check == -1 {
		return repository.CommentListResponse{Response: repository.Response{StatusCode: 1, StatusMsg: "用户未登陆或不存在"}}
	}
	videoId, _ := strconv.ParseInt(videoIdStr, 10, 64)
	var commentDaoList []repository.CommentDao
	db.Where("video_id = ?", videoId).Find(&commentDaoList)
	commentList := CommentDaoListToCommentList(commentDaoList)
	return repository.CommentListResponse{
		Response:    repository.Response{StatusCode: 0, StatusMsg: "success"},
		CommentList: commentList,
	}
}
