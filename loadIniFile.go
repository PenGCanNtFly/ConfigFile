package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/larspensjo/config"
)

var (
	configFile = flag.String("configFile", "config.ini", "General configration file")
)

var TOPIC = make(map[string]string)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	fmt.Println(cfg)

	if cfg.HasSection("topicArr") {
		section, err := cfg.SectionOptions("topicArr")
		if err != nil {
			for _, v := range section {
				options, err := cfg.String("topicArr", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}

	fmt.Println(TOPIC)
	fmt.Println(TOPIC["debug"])
}
