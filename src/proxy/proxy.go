package proxy

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/DayVil/scrapper/src/scrape"
)

func getUrlSources(path string) ([]string, error) {
	httpSources := make([]string, 0)
	content, err := os.ReadFile(path)
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

func GetProxyList(path string) ([]string, error) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	scrapingSites, err := getUrlSources(path)
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

func GetProxys(path string, websiteTry string, retries uint, timeout time.Duration) ([]string, error) {
	proxyList, err := GetProxyList(path)
	if err != nil {
		return nil, err
	}

	proxyList = tryProxysHTTP(proxyList, websiteTry, int(retries), timeout)

	return proxyList, nil
}
