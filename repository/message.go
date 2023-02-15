package repository

import "time"

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"from_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessageDao struct {
	Id         int64
	FromUserId int64
	ToUserId   int64
	Content    string
	CreatedAt  time.Time `gorm:"column:create_time"`
}

func (MessageDao) TableName() string {
	return "message"
}
