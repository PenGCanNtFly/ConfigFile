package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var data = `
[[conf]]
[[conf.a1]]
min = 10
max = 10240
[[conf.a2]]
min = 5000
max = 10240
[[conf.a3]]
"yangqi01.youku.com"=[0 ,100]
`

type MinMax struct {
	Min int
	Max int
}

type OneLevelConf struct {
	A1 []MinMax
	A2 []MinMax
	A3 []map[string][]int64
}

type CONF struct {
	Conf []OneLevelConf
}

func main() {
	var conf CONF
	err := toml.Unmarshal([]byte(data), &conf)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(conf)
	fmt.Println(conf.Conf)
}
