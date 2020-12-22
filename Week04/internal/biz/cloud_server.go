package biz

type ServerSpecs struct {
	UUID     string
	HostName string
	Address  string
	Port     int32
	Cores    int32
	Memory   int32
	DiskSize int32
}

type ServerRepository interface {
	SaveServer(specs *ServerSpecs) error
}

type CloudServerRepo struct {
	repo ServerRepository
}

func NewCloudServerRepo(repository ServerRepository) (*CloudServerRepo, func(), error) {
	return &CloudServerRepo{repo: repository}, func() {}, nil
}

func (receiver *CloudServerRepo) Create(s *ServerSpecs) (map[string]string, error) {
	//生成服务器
	s.UUID = "tacos"
	s.Address = "192.168.10.100"
	err := receiver.repo.SaveServer(s)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"uuid":    s.UUID,
		"address": s.Address,
	}, nil
}

//判断服务器型号是否有效
func (receiver *CloudServerRepo) IsValidModel(core, mem int32) bool {
	if core <= 0 || mem <= 0 {
		return false
	}
	return true
}
