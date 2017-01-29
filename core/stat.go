package core

import (
	"time"
)

// ProtonStat a struct for cache statistics data
type ProtonStat struct {
	StartAt      time.Time // 服务启动时间
	ResolveCount uint64    // 解析数
	HitCount     uint64    // hit 数
}

// NewProtonStat return ProtonStat with default value
func NewProtonStat() *ProtonStat {
	return &ProtonStat{
		StartAt:      time.Now(),
		ResolveCount: 0,
		HitCount:     0,
	}
}

// Resolve will +1 when get resolve request
func (s *ProtonStat) Resolve() *ProtonStat {
	s.ResolveCount++
	return s
}

// Hit will +1 when hit cache
func (s *ProtonStat) Hit() *ProtonStat {
	s.HitCount++
	return s
}
