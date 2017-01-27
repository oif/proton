package core

import (
	"github.com/BurntSushi/toml"
)

type ProtonConfig struct {
	TCP   TCPConfig   `toml:"tcp"`
	UDP   UDPConfig   `toml:"udp"`
	Proxy ProxyConfig `toml:"proxy"`
	Cache CacheConfig `toml:"cache"`
}

type TCPConfig struct{}

type UDPConfig struct{}

type ProxyConfig struct{}

type CacheConfig struct{}

func GetProtonConfig() (ProtonConfig, error) {
	var c ProtonConfig
	_, err := toml.DecodeFile("proton.toml", &c)
	return c, err
}
