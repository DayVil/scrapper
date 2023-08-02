package proxy

import (
	"os"
	"strings"

	"github.com/DayVil/scrapper/src/scrape"
)

func getUrlSources() ([]string, error) {
	httpSources := make([]string, 0)
	content, err := os.ReadFile("./config/websource/http.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	httpSources = append(httpSources, lines...)

	return httpSources, nil
}

func removeDuplicateEntries(addrs []string) []string {
	bucket := make(map[string]bool)
	clean := make([]string, 0)

	for _, entry := range addrs {
		if ok := bucket[entry]; !ok {
			bucket[entry] = true
			clean = append(clean, entry)
		}
	}

	return clean
}

func GetProxys() ([]string, error) {
	scrapingSites, err := getUrlSources()
	if err != nil {
		return nil, err
	}

	proxies := make([]string, 0)
	for _, site := range scrapingSites {
		proxyList := scrape.GetProxyListIpV4(site)
		proxies = append(proxies, proxyList...)
	}
	proxies = removeDuplicateEntries(proxies)

	return proxies, nil
}
