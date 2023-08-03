package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DayVil/scrapper/src/proxy"
)

func main() {
	startTime := time.Now()
	_, err := proxy.GetProxys()
	if err != nil {
		log.Println(err)
		return
	}
	elapsed := time.Since(startTime)

	fmt.Println("\nTime taken: " + elapsed.String())
}
