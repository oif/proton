package core

import (
	"github.com/BurntSushi/toml"
)

// ProtonConfig the whole config
type ProtonConfig struct {
	TCP   TCPConfig   `toml:"tcp"`   // TCP 配置
	UDP   UDPConfig   `toml:"udp"`   // UDP 配置
	Proxy ProxyConfig `toml:"proxy"` // 代理配置
	Cache CacheConfig `toml:"cache"` // 缓存配置
}

// TCPConfig for TCP service config
type TCPConfig struct{}

// UDPConfig for UDP service config
type UDPConfig struct{}

// ProxyConfig a import config to make sure the connection between service and Google DNS
type ProxyConfig struct{}

// CacheConfig for resolve result cache
type CacheConfig struct{}

// GetProtonConfig 解析 toml 配置
func GetProtonConfig() (ProtonConfig, error) {
	var c ProtonConfig
	_, err := toml.DecodeFile("proton.toml", &c)
	return c, err
}
