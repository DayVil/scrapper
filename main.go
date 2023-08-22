package main

import (
	"fmt"

	"github.com/DayVil/scrapper/src/proxy"
	"github.com/DayVil/scrapper/src/proxy/protocols"
)

func main() {
	// var websites []protocols.ProxySites = proxy.GetProxySourcesDefault()
	// var proxyAddresses []protocols.ProxyAdrr = proxy.GetProxys(websites)
	// var validProxys []protocols.ProxyAdrr = checker.TryProxiesDefaultW(proxyAddresses)
	// printWebsites(websites)
	// printSlice(validProxys)

	defaultproxies := proxy.GetDefaultProxies()
	printSlice(defaultproxies)

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
