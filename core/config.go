package core

import (
	"github.com/BurntSushi/toml"
)

type ProtonConfig struct {
	TCP   TCPConfig   `toml:"tcp"`   // TCP 配置
	UDP   UDPConfig   `toml:"udp"`   // UDP 配置
	Proxy ProxyConfig `toml:"proxy"` // 代理配置
	Cache CacheConfig `toml:"cache"` // 缓存配置
}

type TCPConfig struct{}

type UDPConfig struct{}

type ProxyConfig struct{}

type CacheConfig struct{}

// 解析 toml 配置
func GetProtonConfig() (ProtonConfig, error) {
	var c ProtonConfig
	_, err := toml.DecodeFile("proton.toml", &c)
	return c, err
}
