package cfg_test

import (
	"reflect"
	"testing"

	"github.com/ondrejmalina/beloved/internal/cfg"
)

func TestNewConfig(t *testing.T) {
	got := cfg.New("~/.config/beloved")
	want := &cfg.Config{Path: "~/.config/beloved"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got = %s, want = %s", got, want)
	}
}

func TestLoadDefaultConfig(t *testing.T) {
	config := cfg.New("")
	err := config.Load()
	if err != nil {
		t.Error("failed to create or read default cfg: %w", err)
	}
}

func TestLoadNonDefaultConfig(t *testing.T) {
	config := cfg.New("/Users/ondrejmalina/Downloads/bl.cfg")
	err := config.Load()
	if err != nil {
		t.Error("failed to create or read default cfg: %w", err)
	}
}
