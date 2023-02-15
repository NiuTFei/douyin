package repository

import "time"

type UserDao struct {
	Id            int64     `gorm:"column:id"`
	Username      string    `gorm:"column:username"`
	Password      string    `gorm:"column:password"`
	Name          string    `gorm:"column:name"`
	FollowCount   int64     `gorm:"follow_count"`
	FollowerCount int64     `gorm:"follower_count"`
	IsFollow      bool      `gorm:"is_follow"`
	CreatedAt     time.Time `gorm:"column:create_time"`
	UpdatedAt     time.Time `gorm:"column:update_time"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func (UserDao) TableName() string {
	return "user"
}
