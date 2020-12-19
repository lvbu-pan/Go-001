package biz

import (
	"database/sql"
	"week04/internal/data"
)

//需要创建的服务器规格

func BuyCloudServer(db *sql.DB, id, name, address, region, os string, port, cores, memory, size int) (map[string]string, error) {
	sp := new(data.CloudSpecs)
	sp.Id = id
	sp.HostName = name
	sp.Address = address
	sp.Region = region
	sp.OsType = os
	sp.SshPort = port
	sp.Cores = cores
	sp.Memory = memory
	sp.DiskSize = size
	err := sp.SaveCloudServerInfo(db)
	if err != nil {
		return nil, err
	}
	return map[string]string{"message": "ok"}, nil
}
