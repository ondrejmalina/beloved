package cfg_test

import (
	"testing"

	"github.com/ondrejmalina/beloved/internal/cfg"
)

func TestLoadDefaultConfig(t *testing.T) {
	cfg, err := cfg.Init()
	t.Log(cfg)
	if err != nil {
		t.Error("failed to create or read default cfg: %w", err)
	}
}
