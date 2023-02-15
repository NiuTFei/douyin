package repository

import (
	"time"
)

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type VideoDao struct {
	Id            int64     `gorm:"column:id"`
	AuthorId      int64     `gorm:"column:author_id"`
	PlayUrl       string    `gorm:"column:play_url"`
	CoverUrl      string    `gorm:"cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count"`
	IsFavorite    bool      `gorm:"column:is_favorite"`
	Title         string    `gorm:"column:title"`
	CreatedAt     time.Time `gorm:"column:create_time"`
}

func (VideoDao) TableName() string {
	return "video"
}
