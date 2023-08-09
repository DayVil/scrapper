package proxy

import (
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/DayVil/scrapper/src/scrape"
)

func getUrlSources(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
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

func convertToIP4(websites []string) []string {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	proxies := make([]string, 0)
	for _, site := range websites {
		wg.Add(1)
		go scrape.GetProxyListIpV4(site, &proxies, &wg, &mutex)
	}
	wg.Wait()
	proxies = removeDuplicateEntries(proxies)
	return proxies
}

func GetDefaultProxys() ([]string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/DayVil/scrapper/main/config/websource/http.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	text := string(body)
	lines := strings.Split(text, "\n")
	proxies := convertToIP4(lines)
	return proxies, nil
}

func GetProxyListFromFile(path string) ([]string, error) {
	scrapingSites, err := getUrlSources(path)
	if err != nil {
		return nil, err
	}

	proxies := convertToIP4(scrapingSites)

	return proxies, nil
}

func TryProxys(proxyList []string, websiteTry string, retries uint, timeout time.Duration) []string {
	proxyList = tryProxysHTTP(proxyList, websiteTry, int(retries), timeout)

	return proxyList
}
