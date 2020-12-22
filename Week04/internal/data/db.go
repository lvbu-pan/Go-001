package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"time"
)

import _ "github.com/jinzhu/gorm/dialects/postgres"

var Provider = wire.NewSet(NewREDIS, NewPostgres, NewDao)

//type Dao interface {
//	biz.ServerRepository
//}

type Dao struct {
	db *gorm.DB
	rd *redis.Client
}

type dbCfg struct {
	port     int
	host     string
	userName string
	password string
	instance string
}

type redisCfg struct {
	port     int
	db       int
	timeout  int
	host     string
	password string
}

// database数据初始化，本该从配置文件读取
var dc = dbCfg{
	port:     1923,
	host:     "10.0.16.239",
	userName: "devops",
	password: "devops123456",
	instance: "devops",
}

// redis数据初始化，本该从配置文件读取
var rc = redisCfg{
	port:     6379,
	db:       0,
	timeout:  10,
	host:     "10.0.16.239",
	password: "youknowthat",
}

func NewDao(db *gorm.DB, client *redis.Client) (*Dao, func(), error) {
	return &Dao{db, client}, func() {}, nil
}

func NewPostgres() (*gorm.DB, func(), error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", dc.host, dc.port, dc.userName, dc.instance, dc.password))
	if err != nil {
		fmt.Println("cc")
		return nil, nil, err
	}
	clean := func() { _ = db.Close() }
	return db, clean, nil
}

func NewREDIS() (*redis.Client, func(), error) {
	rd := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", rc.host, rc.port),
		OnConnect:   nil,
		Password:    rc.password,
		DB:          rc.db,
		DialTimeout: time.Duration(rc.timeout) * time.Second,
	})
	_, err := rd.Ping(context.Background()).Result()
	if err != nil {
		return nil, nil, err
	}
	clean := func() { _ = rd.Close() }
	return rd, clean, nil
}
