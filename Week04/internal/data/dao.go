package data

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

var Provider = wire.NewSet(NewREDIS, NewDB)

type Data struct {
	Rd *redis.Client
	Db *sql.DB
}

type dbCfg struct {
	port     int
	host     string
	userName string
	password string
	instance string
}

// database数据初始化，本该从配置文件读取
var dc = dbCfg{
	port:     3306,
	host:     "127.0.0.1",
	userName: "root",
	password: "123456",
	instance: "devops",
}

func NewDB() (*sql.DB, func(), error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dc.userName, dc.password, dc.host, dc.port, dc.instance))
	if err != nil {
		return nil, nil, err
	}
	clean := func() { _ = db.Close() }
	return db, clean, nil
}

type redisCfg struct {
	port     int
	db       int
	timeout  int
	host     string
	password string
}

// redis数据初始化，本该从配置文件读取
var rc = redisCfg{
	port:     3306,
	db:       0,
	timeout:  10,
	host:     "127.0.0.1",
	password: "tacoTuesday",
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
