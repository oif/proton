package core

import (
	"testing"
)

func TestStatistics(t *testing.T) {
	s := NewProtonStat()
	s.Resolve().Hit()
	if s.ResolveCount != 1 {
		t.Errorf("add resolve count fail, should be 1 but %d given", s.ResolveCount)
	}
	if s.HitCount != 1 {
		t.Errorf("add hit count fail, should be 1 but %d given", s.ResolveCount)
	}
}
