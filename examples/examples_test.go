package config_test

import (
	"fmt"

	"github.com/xavier/config"
)

type dummy struct{}

var _ = new(dummy)

func ExampleConfig_ex1() {
	c := config.New()
	fmt.Println(c.Get("version"))
	c.Set("version", "2.3.4")
	fmt.Println(c.Get("version"))
	// Output: 2.3.4
}

func ExampleConfig_ex2() {
	fmt.Println("Hello world")
	// Output: Hello world
}

func ExampleConfig_ex3() {
	c := config.New("./test.conf")
	fmt.Println(c.Get("VERSION"))
	fmt.Println(c.Get("version"))
	fmt.Println(c.Get("www.vvv"))
	// Output: 333.55.6  build 245
	//
	// "456"
}
