package main

import (
	"fmt"
	"log"

	"github.com/DayVil/scrapper/src/config"
)

func main() {
	cfg, err := config.NewCfg()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, line := range cfg.HttpSources {
		fmt.Println(line)
	}
}
