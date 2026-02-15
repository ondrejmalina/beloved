package cfg_test

import (
	"testing"
	"reflect"


	"github.com/ondrejmalina/beloved/internal/cfg"

)

func TestCreateConfig(t *testing.T) {
	want := cfg.Config{"~/.config", []string{}}
	got := cfg.CreateConfig("~/.config")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got = %s, want = %s", got, want)
	}
}
