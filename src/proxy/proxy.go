package proxy

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/DayVil/scrapper/src/checker"
	"github.com/DayVil/scrapper/src/proxy/protocols"
	"github.com/DayVil/scrapper/src/scraper"
)

// GetProxySources returns a list of proxy sites from the specified URL.
// The list contains sites that support the specified protocol.
//
// The function returns a list of proxy sites.
func GetProxySources(url string, proto protocols.Protocols) []protocols.ProxySites {
	sites := make([]protocols.ProxySites, 0)

	response, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return sites
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "HTTP GET request failed with status:", response.Status)
		return sites
	}

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		sites = append(sites, protocols.ProxySites{
			Url:      line,
			Protocol: proto,
		})
	}

	return sites
}

// GetProxySourcesFile returns a list of proxy sites from a file.
func GetProxySourcesFile(path string, proto protocols.Protocols) []protocols.ProxySites {
	sites := make([]protocols.ProxySites, 0)

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
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

// Get default sources for each protocol.
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

// removeDuplicate takes a slice of any type and returns a slice of the same type
// containing only the unique elements in the original slice.
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

// GetProxys returns a list of proxy addresses from the specified sites.
func GetProxys(websites []protocols.ProxySites) []protocols.ProxyAdrr {
	addrr := scraper.SearchSite(websites)
	addrr = removeDuplicate(addrr)
	return addrr
}

// GetDefaultProxies returns a list of proxy addresses from the default sites.
func GetDefaultProxies() []protocols.ProxyAdrr {
	websites := GetProxySourcesDefault()
	proxys := GetProxys(websites)
	validProxies := checker.TryProxiesDefaultW(proxys)

	return validProxies
}

// TrySite returns a list of proxy addresses from the specified test website.
func TrySite(url string) []protocols.ProxyAdrr {
	websites := GetProxySourcesDefault()
	proxys := GetProxys(websites)
	return checker.TryProxies(proxys, protocols.TIMEOUT, protocols.RETRIES, url)
}
