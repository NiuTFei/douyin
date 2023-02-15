package repository

type RelationDao struct {
	Id         int64
	FromUserId int64
	ToUserId   int64
}

func (RelationDao) TableName() string {
	return "relation"
}
