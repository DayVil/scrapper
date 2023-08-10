package main

import (
	"fmt"
	"time"

	"github.com/DayVil/scrapper/src/proxy"
	"github.com/DayVil/scrapper/src/proxy/protocols"
)

const (
	timeout = time.Second * 6
	retries = 2
)

func main() {
	var websites []protocols.ProxySites = proxy.GetProxySourcesDefault()
	// var proxyAddresses []protocols.Proxy = proxy.GetProxys(websites)
	// var validProxys []protocols.Proxy = proxy.TryProxies(proxyAddresses)
	printWebsites(websites)
}

func printWebsites(websites []protocols.ProxySites) {
	for _, list := range websites {
		fmt.Println(list)
	}
}
