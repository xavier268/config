package config

import (
	"fmt"
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

func TestParse(t *testing.T) {
	c := newConfig("./examples/test.conf")
	c.Get("VERSION")
	fmt.Println("Values parsed : \n", c.values)
}
