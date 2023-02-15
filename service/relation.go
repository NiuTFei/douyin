package service

import (
	"context"
	"douyin/repository"
	"fmt"
	"gorm.io/gorm"
	"strconv"
)

const (
	FromKey = "relation:from_"
	ToKey   = "relation:to_"
)

func RelationAction(token string, toUserIdStr string, actionType string) repository.Response {
	check := CheckToken(token)
	if check == 0 || check == -1 {
		return repository.Response{
			StatusCode: 1, StatusMsg: "用户未登陆或不存在",
		}
	}
	toUserId, _ := strconv.ParseInt(toUserIdStr, 10, 64)
	relation := repository.RelationDao{FromUserId: check, ToUserId: toUserId}
	fromUser := repository.UserDao{Id: check}
	toUser := repository.UserDao{Id: toUserId}

	if actionType == "1" {
		db.Create(&relation)
		//关注之后记录到redis   key:  from_id(id用户关注了哪些）   to_id（id用户被哪些用户关注）
		rdb.SAdd(context.Background(), FromKey+fmt.Sprintf("%d", check), toUserId)
		rdb.SAdd(context.Background(), ToKey+toUserIdStr, check)
		//更新用户信息，follow、follower数量
		db.Model(&fromUser).Update("follow_count", gorm.Expr("follow_count + ?", 1))
		db.Model(&toUser).Update("follower_count", gorm.Expr("follower_count + ?", 1))
		return repository.Response{
			StatusCode: 0, StatusMsg: "关注成功",
		}
	}
	db.Where("from_user_id = ? and to_user_id = ?", check, toUserId).Delete(&relation)
	rdb.SRem(context.Background(), FromKey+fmt.Sprintf("%d", check), toUserId)
	rdb.SRem(context.Background(), ToKey+toUserIdStr, check)
	db.Model(&fromUser).Update("follow_count", gorm.Expr("follow_count - ?", 1))
	db.Model(&toUser).Update("follower_count", gorm.Expr("follower_count - ?", 1))
	return repository.Response{
		StatusCode: 0, StatusMsg: "取消关注",
	}
}

func FollowList(userIdStr string, token string) repository.UserListResponse {
	check := CheckToken(token)
	if check == 0 || check == -1 {
		return repository.UserListResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "用户未登陆或不存在"},
		}
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	var followList []repository.RelationDao
	db.Where("from_user_id = ?", userId).Find(&followList)
	return repository.UserListResponse{
		Response: repository.Response{StatusCode: 0, StatusMsg: "success"},
		UserList: FollowListToUserList(userId, followList),
	}
}

func FollowerList(userIdStr string, token string) repository.UserListResponse {
	check := CheckToken(token)
	if check == 0 || check == -1 {
		return repository.UserListResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "用户未登陆或不存在"},
		}
	}
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	var followList []repository.RelationDao
	db.Where("to_user_id = ?", userId).Find(&followList)
	return repository.UserListResponse{
		Response: repository.Response{StatusCode: 0, StatusMsg: "success"},
		UserList: FollowerListToUserList(userId, followList),
	}
}

func FriendList(userIdStr string, token string) repository.UserListResponse {
	check := CheckToken(token)
	if check == 0 || check == -1 {
		return repository.UserListResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "用户未登陆或不存在"},
		}
	}
	//求关注 和 粉丝列表对交集，即为friend
	friendIdList := rdb.SInter(context.Background(), FromKey+userIdStr, ToKey+userIdStr).Val() //[]string
	friendList := make([]repository.User, len(friendIdList))
	for i, friendId := range friendIdList {
		var user repository.UserDao
		db.Where("id = ?", friendId).Find(&user)
		friendList[i] = UserDaoToUser(user)
	}
	return repository.UserListResponse{
		Response: repository.Response{StatusCode: 0, StatusMsg: "success"},
		UserList: friendList,
	}
}
