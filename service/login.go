package service

import (
	"context"
	"crypto/md5"
	"douyin/repository"
	"encoding/hex"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

func Login(username string, password string) repository.LoginResponse {
	var user repository.UserDao
	err := db.Where("username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return repository.LoginResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  "用户不存在"},
		}
	}

	if check(password, user.Password) == false {
		return repository.LoginResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  "密码错误"},
		}
	}

	token := "token:" + username
	ctx := context.Background()
	rdb.Set(ctx, token, user.Id, 0)

	return repository.LoginResponse{
		Response: repository.Response{
			StatusCode: 0,
			StatusMsg:  "登陆成功",
		},
		UserId: user.Id,
		Token:  token,
	}
}

func UserInfo(userId string, token string) repository.UserResponse {
	exist := CheckToken(token)
	if exist == 0 || exist == -1 {
		//token鉴权失败？
	}
	id, _ := strconv.ParseInt(userId, 10, 64)
	var user repository.UserDao
	db.Where("id = ?", id).Find(&user)
	return repository.UserResponse{
		Response: repository.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		User: repository.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		},
	}
}

func Register(username string, password string) repository.RegisterResponse {
	var user repository.UserDao
	err := db.Where("username = ?", username).First(&user).Error
	if err != gorm.ErrRecordNotFound {
		return repository.RegisterResponse{
			Response: repository.Response{
				StatusCode: 1,
				StatusMsg:  "用户已存在",
			},
		}
	}

	name := "user_" + username
	token := "token:" + username
	encodePassword := md5Encode(password) //加密存储密码
	newUser := repository.UserDao{Username: username, Password: encodePassword, Name: name}
	db.Create(&newUser)
	//自动登陆
	rdb.Set(context.Background(), token, user.Id, 0)

	return repository.RegisterResponse{
		Response: repository.Response{
			StatusCode: 0,
			StatusMsg:  "用户注册成功",
		},
		UserId: newUser.Id,
		Token:  token,
	}
}

// md5加密与验证
func md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func check(content string, encrypted string) bool {
	return strings.EqualFold(md5Encode(content), encrypted)
}
