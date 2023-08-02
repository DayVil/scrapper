package proxy

import (
	"os"
	"strings"
	"github.com/DayVil/scrapper/src/scrape"
)

func getUrlSources() ([]string, error) {
	var httpSources []string
	content, err := os.ReadFile("./config/websource/http.txt")
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	httpSources = append(httpSources, lines...)

	return httpSources, nil
}

func GetProxys() ([]string, error) {
	scrapingSites, err := getUrlSources()
	if err != nil {
		return nil, err
	}

	var proxies []string
	for _, site := range scrapingSites {
		proxyList := scrape.GetProxyListIpV4(site)
		proxies = append(proxies, proxyList...)
	}

	return proxies, nil
}
