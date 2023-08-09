package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DayVil/scrapper/src/proxy"
)

const timeout = time.Second * 6
const retries = 2

func main() {
	// proxyList, err := proxy.GetProxyListFromFile("./config/websource/http.txt")
	proxyList, err := proxy.GetDefaultProxys()
	if err != nil {
		log.Fatal(err)
		return
	}

	proxies := proxy.TryProxys(proxyList, "https://www.amazon.de/", retries, timeout)

	printProxys(proxies)
}

func printProxys(proxyList []string) {
	for _, list := range proxyList {
		fmt.Println(list)
	}
}
