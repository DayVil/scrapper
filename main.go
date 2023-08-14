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
	var proxyAddresses []protocols.ProxyAdrr = proxy.GetProxys(websites)
	// var validProxys []protocols.Proxy = proxy.TryProxies(proxyAddresses)
	printWebsites(websites)
	printSlice(proxyAddresses)
}

func printWebsites(websites []protocols.ProxySites) {
	for _, list := range websites {
		fmt.Println(list)
	}
}

func printSlice[T ~string](elements []T) {
	for _, element := range elements {
		fmt.Println(element)
	}
}
