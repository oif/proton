package core

import (
	"time"
)

// 服务统计
type ProtonStat struct {
	StartAt      time.Time // 服务启动时间
	ResolveCount uint64    // 解析数
	HitCount     uint64    // hit 数
}

// 构造
func NewProtonStat() *ProtonStat {
	return &ProtonStat{
		StartAt:      time.Now(),
		ResolveCount: 0,
		HitCount:     0,
	}
}

// 增加 1 解析数
func (s *ProtonStat) Resolve() *ProtonStat {
	s.ResolveCount++
	return s
}

// 增加 1 命中
func (s *ProtonStat) Hit() *ProtonStat {
	s.HitCount++
	return s
}
