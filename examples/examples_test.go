package config_test

import (
	"fmt"
	"os"

	"github.com/xavier268/config"
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
	fmt.Println(c.Get("this"))
	// Output:
	// 333.55.6  build 245
	//
	// "456"
	// a long line with a single quote (") and a second = sign
	// world
	// is it
}

func ExampleConfig_ex3() {
	c := config.New("./example.conf")

	c.Set("VERSION", "2") // that should overwite the file content, even if lazylily loaded later ...
	fmt.Println(c.Get("VERSION"))
	// Output: 2
}

func ExampleConfig_ex4() {
	out := "./out.conf"
	c := config.New("./example.conf")
	c.Set("hello", "world")
	err := c.Save(out)
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	err = os.Remove(out)
	if err != nil {
		panic(err)
	}
	// Output:
	// VERSION=333.55.6  build 245
	// [another]
	// k=a long line with a single quote (") and a second = sign
	// []
	// hello=world
	// this=is it
	// [www]
	// vvv="456"
}
