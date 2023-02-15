package service

import (
	"context"
	"douyin/repository"
	"gorm.io/gorm"
	"strconv"
)

func CheckToken(token string) int64 {
	ctx := context.Background()
	result, err := rdb.Get(ctx, token).Result()
	if err != nil {
		return 0 //未登录
	}
	id, _ := strconv.ParseInt(result, 10, 64)

	var user repository.UserDao
	err2 := db.First(&user, id).Error
	if err2 == gorm.ErrRecordNotFound {
		return -1 //用户不存在
	}

	return user.Id //用户存在，返回用户ID
}

func VideoDaoToVideoList(userId int64, videos []repository.VideoDao) []repository.Video {
	videoList := make([]repository.Video, len(videos)) //需要给切片预分配空间
	for i, video := range videos {
		var user repository.UserDao
		var isFavorite bool
		var favoriteDao repository.FavoriteDao
		db.Where("id = ?", video.AuthorId).Find(&user)
		//需要判断该视频是否被登陆用户点赞过,即在favorite表中查询有没有userid-videoID记录
		db.Where("user_id = ? and favorite_video_id = ?", userId, video.Id).Find(&favoriteDao)

		if favoriteDao.Id > 0 {
			isFavorite = true
		} else {
			isFavorite = false
		}
		videoList[i] = repository.Video{
			Id:            video.Id,
			Author:        UserDaoToUserWithToken(userId, user),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
	}
	return videoList
}

func VideoDaoToVideoListWithNoToken(videos []repository.VideoDao) []repository.Video {
	videoList := make([]repository.Video, len(videos)) //需要给切片预分配空间
	for i, video := range videos {
		var user repository.UserDao
		db.Where("id = ?", video.AuthorId).Find(&user)
		videoList[i] = repository.Video{
			Id:            video.Id,
			Author:        UserDaoToUser(user),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		}
	}
	return videoList
}

func UserDaoToUser(userDao repository.UserDao) repository.User {

	user := repository.User{
		Id:            userDao.Id,
		Name:          userDao.Name,
		FollowCount:   userDao.FollowCount,
		FollowerCount: userDao.FollowerCount,
		IsFollow:      userDao.IsFollow,
	}
	return user
}

func UserDaoToUserWithToken(userId int64, userDao repository.UserDao) repository.User {
	var relation repository.RelationDao
	var isFollow bool
	if userId == userDao.Id {
		isFollow = false
	} else {
		err := db.Where("from_user_id = ? and to_user_id = ?", userId, userDao.Id).First(&relation).Error
		if err == gorm.ErrRecordNotFound {
			isFollow = false
		} else {
			isFollow = true
		}
	}
	user := repository.User{
		Id:            userDao.Id,
		Name:          userDao.Name,
		FollowCount:   userDao.FollowCount,
		FollowerCount: userDao.FollowerCount,
		IsFollow:      isFollow,
	}
	return user
}

func CommentDaoListToCommentList(commentDaoList []repository.CommentDao) []repository.Comment {
	commentList := make([]repository.Comment, len(commentDaoList)) //需要给切片预分配空间
	for i, commentDao := range commentDaoList {
		var user repository.UserDao
		db.Where("id = ?", commentDao.UserId).Find(&user)
		commentList[i] = repository.Comment{
			Id:         commentDao.Id,
			User:       UserDaoToUser(user),
			Content:    commentDao.Content,
			CreateDate: commentDao.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	return commentList
}

// FollowListToUserList 关注-->User
func FollowListToUserList(userId int64, followList []repository.RelationDao) []repository.User {
	userList := make([]repository.User, len(followList))
	for i, relation := range followList {
		var userDao repository.UserDao
		db.Where("id = ?", relation.ToUserId).Find(&userDao)
		userList[i] = UserDaoToUserWithToken(userId, userDao)
	}
	return userList
}

// FollowerListToUserList 粉丝-->User
func FollowerListToUserList(userId int64, followList []repository.RelationDao) []repository.User {
	userList := make([]repository.User, len(followList))
	for i, relation := range followList {
		var userDao repository.UserDao
		db.Where("id = ?", relation.FromUserId).Find(&userDao)
		userList[i] = UserDaoToUserWithToken(userId, userDao)
	}
	return userList
}

func MessageDaoListToMessageList(messageDaoList []repository.MessageDao) []repository.Message {
	messageList := make([]repository.Message, len(messageDaoList))
	for i, messageDao := range messageDaoList {
		messageList[i] = repository.Message{
			Id:         messageDao.Id,
			Content:    messageDao.Content,
			CreateTime: messageDao.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	return messageList
}
