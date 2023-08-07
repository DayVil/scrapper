package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DayVil/scrapper/src/proxy"
)

func main() {
	// startTime := time.Now()
	timeout := time.Second * 20
	proxies, err := proxy.GetProxys("./config/websource/http.txt", "https://www.amazon.de/", 3, timeout)
	if err != nil {
		log.Println(err)
		return
	}
	// elapsed := time.Since(startTime)

	printProxys(proxies)
	// fmt.Println("\nTime taken: " + elapsed.String())
}

func printProxys(proxyList []string) {
	for _, list := range proxyList {
		fmt.Println(list)
	}
}