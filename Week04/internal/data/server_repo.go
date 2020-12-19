package data

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"week04/internal/pkg/code"
)

//var _ biz.AssetRepo = (*assetRepo)(nil)

//type assetRepo struct{}

type CloudSpecs struct {
	Id       string
	HostName string
	Address  string
	SshPort  int
	Cores    int
	Memory   int
	DiskSize int
	Region   string
	OsType   string
}

func (sc *CloudSpecs) SaveCloudServerInfo(db *sql.DB) error {
	insertStr := fmt.Sprintf(`insert into conf_hosts (id, host_name, address, port, cores, memory, disk_size, region, os) value (%s, %s, %s, %d, %d, %d, %d, %s, %s)`, sc.Id, sc.HostName, sc.Address, sc.SshPort, sc.Cores, sc.Memory, sc.DiskSize, sc.Region, sc.OsType)
	if _, err := db.Exec(insertStr); err != nil {
		return errors.Wrapf(code.StoreFail, fmt.Sprintf("sql: %v", err))
	}
	return nil
}
