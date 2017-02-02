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

// BasicConfig with addr and port
type BasicConfig struct {
	Addr string // address
	Port uint   // port
}

// TCPConfig for TCP service config
type TCPConfig struct {
	BasicConfig
}

// UDPConfig for UDP service config
type UDPConfig struct {
	BasicConfig
}

// ProxyConfig a important config to make sure the connection between service and Google DNS
type ProxyConfig struct {
	Protocol string
	BasicConfig
}

// CacheConfig for resolve result cache
type CacheConfig struct {
	Size int // 缓存大小
}

// GetProtonConfig 解析 toml 配置
func GetProtonConfig(configPath string) (ProtonConfig, error) {
	var c ProtonConfig
	_, err := toml.DecodeFile(configPath, &c)
	return c, err
}
