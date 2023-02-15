package service

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var rdb *redis.Client

func Init() error {
	var err error
	dsn := "root:Fei0628.@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //关闭默认事务
		PrepareStmt:            true, //缓存预编译语句
	})

	rdb = redis.NewClient(&redis.Options{
		Addr:     "47.109.77.112:6379",
		Password: "990628", // no password set
		DB:       0,        // use default DB
	})

	return err
}
