package main

import (
	"fmt"
	"time"

	"github.com/DayVil/scrapper/src/proxy"
	"github.com/DayVil/scrapper/src/proxy/protocols"
)

const timeout = time.Second * 6
const retries = 2

func main() {
	var websites []protocols.ProxySites = proxy.GetProxySourcesDefault()
	printWebsites(websites)
}

func printWebsites(websites []protocols.ProxySites) {
	for _, list := range websites {
		fmt.Println(list)
	}
}
