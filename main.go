package main

import (
	"fmt"

	"github.com/DayVil/scrapper/src/proxy"
	"github.com/DayVil/scrapper/src/proxy/protocols"
)

func main() {
	// defaultproxies := proxy.GetDefaultProxies()
	// printSlice(defaultproxies)

	testSite := proxy.TrySite("https://www.stade.de/")
	printSlice(testSite)

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
