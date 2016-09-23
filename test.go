package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var data = `
a: 1
b: 2
`

type T struct {
	F int `yaml:"a,omitempty"`
	B int
}

func main() {
	var t T

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)
	fmt.Println(t.B)
}
