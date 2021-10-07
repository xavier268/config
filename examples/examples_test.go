package config_test

import (
	"fmt"

	"github.com/xavier/config"
)

func ExampleConfig_ex1() {
	c := config.New()
	fmt.Println(c.Get("version"))
	c.Set("version", "2.3.4")
	fmt.Println(c.Get("version"))
	// Output: 2.3.4
}

func ExampleConfig_ex2() {
	c := config.New("nonexistant.conf", "./example.conf")

	c.Set("hello", "world")
	fmt.Println(c.Get("VERSION"))
	fmt.Println(c.Get("VERSion"))
	fmt.Println(c.Get("www.vvv"))
	fmt.Println(c.Get("another.k"))
	fmt.Println(c.Get("hello"))
	// Output:
	// 333.55.6  build 245
	//
	// "456"
	// a long line with a single quote (") and a second = sign
	// world
}
