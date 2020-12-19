package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"week04/internal/biz"
	"week04/internal/pkg/code"
)

var (
	Db *sql.DB
	Rd *redis.Client
)

type CreateAssetReq struct {
	Id       string `json:"id"`
	HostName string `json:"host_name"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Cores    int    `json:"cores"`
	Memory   int    `json:"memory"`
	DiskSize int    `json:"disk_size"`
	Region   string `json:"region"`
	Os       string `json:"os"`
}

func CreateServer(ctx context.Context, c *gin.Context) {

	req := new(CreateAssetReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(400, map[string]string{"errMsg": "参数格式错误"})
		return
	}
	result, err := biz.BuyCloudServer(Db, req.Id, req.HostName, req.Address, req.Region, req.Os, req.Port, req.Cores, req.Memory, req.DiskSize)
	if err != nil && errors.Is(err, code.StoreFail) {
		c.JSON(200, map[string]string{"errMsg": "购买服务器失败"})
	}
	c.JSON(200, result)
	return
}
