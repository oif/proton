package core

import (
	"time"
)

type ProtonStat struct {
	StartAt      time.Time // 服务启动时间
	ResolveCount uint64    // 解析数
	HitCount     uint64    // hit 数
}

func NewProtonStat() *ProtonStat {
	return &ProtonStat{
		StartAt:      time.Now(),
		ResolveCount: 0,
		HitCount:     0,
	}
}

func (s *ProtonStat) Resolve() *ProtonStat {
	s.ResolveCount++
	return s
}

func (s *ProtonStat) Hit() *ProtonStat {
	s.HitCount++
	return s
}
