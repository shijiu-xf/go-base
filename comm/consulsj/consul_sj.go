package consulsj

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/v3/config"
)

func GetConsulConfig(addr string, prefix string) (config.Config, error) {
	source := consul.NewSource(
		// 设置consul 地址
		consul.WithAddress(addr),
		// 设置config 的前缀
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	newConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	err = newConfig.Load(source)
	if err != nil {
		return nil, err
	}
	return newConfig, nil
}
