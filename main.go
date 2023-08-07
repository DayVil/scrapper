package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DayVil/scrapper/src/proxy"
)

func main() {
	timeout := time.Second * 20
	proxies, err := proxy.GetProxys("./config/websource/http.txt", "https://www.amazon.de/", 3, timeout)
	if err != nil {
		log.Println(err)
		return
	}

	printProxys(proxies)
}

func printProxys(proxyList []string) {
	for _, list := range proxyList {
		fmt.Println(list)
	}
}