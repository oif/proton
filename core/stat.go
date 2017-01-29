package core

import (
	"time"
)

type ProtonStat struct {
	StartAt      time.Time
	ResolveCount uint64
	HitCount     uint64
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
