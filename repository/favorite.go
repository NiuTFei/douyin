package repository

type FavoriteDao struct {
	Id              int64 `gorm:"column:id"`
	UserId          int64 `gorm:"column:user_id"`
	FavoriteVideoId int64 `gorm:"column:favorite_video_id"`
	//IsDel           soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

func (FavoriteDao) TableName() string {
	return "favorite"
}
