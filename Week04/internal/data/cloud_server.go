package data

import (
	"github.com/pkg/errors"
	"time"
	"week04/internal/biz"
)

var _ biz.ServerRepository = (*Dao)(nil)

type Hosts struct {
	Id         int32     `gorm:"primary_key;column:id"`
	HostName   string    `gorm:"column:host_name"`
	Address    string    `gorm:"column:private_ip"`
	Port       int32     `gorm:"column:ssh_port"`
	Cores      int32     `gorm:"column:cpu"`
	Memory     int32     `gorm:"column:max_memory"`
	AppId      int       `gorm:"column:app_id"`
	EnvId      int       `gorm:"column:env_id"`
	Describe   string    `gorm:"column:describe"`
	HostType   int       `gorm:"column:host_type"`
	ServerId   string    `gorm:"column:server_id"`
	ManageId   string    `gorm:"column:jump_server_id"`
	DiskSize   int       `gorm:"column:disk_size"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	AgentName  string    `gorm:"column:salt_name"`
	OwnerId    int       `gorm:"column:owner_service_id"`
}

func (receiver *Dao) SaveServer(s *biz.ServerSpecs) error {
	host := new(Hosts)
	host.HostName = s.HostName
	host.Address = s.Address
	host.Port = s.Port
	host.Memory = s.Memory
	host.Cores = s.Cores
	host.AppId = 6
	host.EnvId = 1
	host.ServerId = s.UUID
	if err := receiver.db.Table("conf_hosts").Create(host).Error; err != nil {
		return errors.Wrapf(err, "主机信息存储失败")
	}
	return nil
}
