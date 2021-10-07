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

func TestPrefix(t *testing.T) {
	data := []string{
		"ttt", "", "ttt",
		"ttt.hhh", "ttt", "hhh",
		"aaa.bb.cc", "aaa.bb", "cc",
		"aaa.bb..cc", "aaa.bb.", "cc",
		".aaa.bb..cc", ".aaa.bb.", "cc",
		"aaa..bb.cc", "aaa..bb", "cc",
		".aa", "", "aa",
	}

	for i := 0; i < len(data); i += 3 {

		p, k := getPrefix(data[i])
		if p != data[i+1] {
			t.Fatalf("%s : expected prefix : %s but got %s", data[i], data[i+1], p)
		}
		if k != data[i+2] {
			t.Fatalf("%s : expected short key : %s but got %s", data[i], data[i+2], k)
		}

	}
}
