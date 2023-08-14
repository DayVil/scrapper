package proxy

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/DayVil/scrapper/src/proxy/protocols"
	"github.com/DayVil/scrapper/src/scraper"
)

func GetProxySources(url string, proto protocols.Protocols) []protocols.ProxySites {
	sites := make([]protocols.ProxySites, 0)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return sites
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP GET request failed with status:", response.Status)
		return sites
	}

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()
		sites = append(sites, protocols.ProxySites{
			Url:      line,
			Protocol: proto,
		})
	}

	return sites
}

func GetProxySourcesFile(path string, proto protocols.Protocols) []protocols.ProxySites {
	sites := make([]protocols.ProxySites, 0)

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sites = append(sites, protocols.ProxySites{
			Url:      line,
			Protocol: proto,
		})
	}

	return sites
}

func GetProxySourcesDefault() []protocols.ProxySites {
	sites := make([]protocols.ProxySites, 0)

	httpResponse := GetProxySources("https://raw.githubusercontent.com/DayVil/scrapper/main/config/websource/http.txt", protocols.HttpProt)
	socks4Response := GetProxySources("https://raw.githubusercontent.com/DayVil/scrapper/main/config/websource/socks4.txt", protocols.Socks4)
	socks5Response := GetProxySources("https://raw.githubusercontent.com/DayVil/scrapper/main/config/websource/socks5.txt", protocols.Socks5)

	sites = append(sites, httpResponse...)
	sites = append(sites, socks4Response...)
	sites = append(sites, socks5Response...)

	return sites
}

func removeDuplicate[T comparable](elements []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}

	for _, item := range elements {
		if _, value := allKeys[item]; !value {
			list = append(list, item)
		}
	}

	return list
}

// TODO: To implement
func GetProxys(websites []protocols.ProxySites) []protocols.ProxyAdrr {
	addrr := scraper.SearchSite(websites)
	addrr = removeDuplicate(addrr)
	return addrr
}

// TODO: To implement
func TryProxies(proxyAddresses []protocols.ProxyAdrr) []protocols.ProxyAdrr {
	return nil
}
