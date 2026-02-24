package cfg_test

import (
	"testing"

	"github.com/ondrejmalina/beloved/internal/cfg"
)

func TestLoadDefaultConfig(t *testing.T) {
	cfg, err := cfg.Init()
	t.Log(cfg)
	if err != nil {
		t.Error(err)
	}
}

func TestWriteDefaultConfig(t *testing.T) {
	cfg, err := cfg.Init()
	if err != nil {
		t.Error(err)
	}
	t.Log(cfg)
	
	n, err := cfg.Add("/workspaces/beloved/")
	if err != nil {
		t.Error(err)
	}
	t.Log(n)
}
