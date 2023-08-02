package main

import (
	"fmt"
	"log"

	"github.com/DayVil/scrapper/src/proxy"
)

func main() {
	proxyList, err := proxy.GetProxys()
	if err != nil {
		log.Println(err)
		return
	}

	for _, line := range proxyList {
		fmt.Println(line)
	}

	fmt.Println("Amount of Proxies: ", len(proxyList))
}
