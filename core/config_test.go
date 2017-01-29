package core

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	config, err := GetProtonConfig("../proton.toml")
	if err != nil {
		t.Errorf("%v\n", err.Error())
	} else {
		if config.Proxy.Port != 6152 {
			t.Errorf("fail to read config, get port %d, should be 6152", config.Proxy.Port)
		}
	}
}
