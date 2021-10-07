package config

import (
	"testing"
)

// compiler check interface satisfied
var _ Config = new(config)

func TestSetGet(t *testing.T) {
	c := New("nnnn")
	if c.Get("ttt") != "" {
		t.FailNow()
	}
	c.Set("ttt", "aaa")
	if c.Get("ttt") != "aaa" {
		t.FailNow()
	}
}
