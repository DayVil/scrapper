package proxy

import (
	"os"
	"strings"
	"sync"

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

func GetProxyList() ([]string, error) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	scrapingSites, err := getUrlSources()
	if err != nil {
		return nil, err
	}

	proxies := make([]string, 0)
	for _, site := range scrapingSites {
		wg.Add(1)
		go scrape.GetProxyListIpV4(site, &proxies, &wg, &mutex)
	}
	wg.Wait()

	proxies = removeDuplicateEntries(proxies)

	return proxies, nil
}

func GetProxys() ([]string, error) {
	proxyList, err := GetProxyList()
	if err != nil {
		return nil, err
	}

	return proxyList, nil
}
