package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

/*
var data = `
blog: xiaorui.cc
best_authors: ["fengyun","lee","park"]
desc:
  counter: 521
  plist: [3, 4]
`
*/

type T struct {
	Blog    string
	Authors []string `yaml:"best_authors,flow"`
	Desc    struct {
		Counter int   `yaml:"Counter"`
		Plist   []int `yaml:",flow"`
	}
	Remote_Url []string `yaml:"remote_url"`
	Tag        string
}

func main() {

	t := T{}

	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	err = yaml.Unmarshal(data, &t)
	//err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
	t.Blog = "this is Blog"
	t.Authors = append(t.Authors, "myself")
	t.Desc.Counter = 99
	t.Tag = "syslog"
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
}
