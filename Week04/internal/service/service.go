package service

import (
	"context"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"log"
	"week04/api"
	"week04/internal/biz"
)

var Provider = wire.NewSet(NewCloudService, wire.Bind(new(api.DemoBMServer), new(*CloudService)))

type CloudService struct {
	smd *biz.CloudServerRepo
}

func NewCloudService(csu *biz.CloudServerRepo) (*CloudService, func(), error) {
	return &CloudService{csu}, func() {}, nil
}

func (receiver *CloudService) CreateCloudServer(ctx context.Context, cloudServerReq *api.CloudServerReq) (cloudServerReply *api.CloudServerResp) {
	s := new(biz.ServerSpecs)
	cloudServerReply = new(api.CloudServerResp)
	s.HostName = cloudServerReq.HostName
	s.Port = cloudServerReq.Port
	s.Cores = cloudServerReq.Cores
	s.Memory = cloudServerReq.Memory
	s.DiskSize = cloudServerReq.DiskSize
	if !receiver.smd.IsValidModel(s.Cores, s.Memory) {
		cloudServerReply.Message = "传入型号无效"
		return
	}
	result, err := receiver.smd.Create(s)
	if err != nil {
		log.Println(errors.Cause(err))
		cloudServerReply.Message = "购买失败"
		return
	}
	cloudServerReply.Uuid = result["uuid"]
	cloudServerReply.Address = result["address"]
	return
}
