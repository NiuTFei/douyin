package repository

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentDao struct {
	Id        int64          `gorm:"column:id"`
	UserId    int64          `gorm:"column:user_id"`
	VideoId   int64          `gorm:"column:video_id"`
	Content   string         `gorm:"column:content"`
	CreatedAt time.Time      `gorm:"column:create_time"`
	Deleted   gorm.DeletedAt `gorm:"column:deleted_time"`
}

func (CommentDao) TableName() string {
	return "comment"
}
