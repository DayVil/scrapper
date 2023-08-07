package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type validProxy struct {
	Proxies []string
	MU      sync.Mutex
}

func tryProxyHTTP(proxy string, site string, retries int, timeout time.Duration, proxies *validProxy, wg *sync.WaitGroup) {
	defer wg.Done()
	proxyURL, err := url.Parse("http://" + proxy)
	if err != nil {
		log.Println(err)
		return
	}

	transport := http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := http.Client{
		Transport: &transport,
		Timeout:   timeout,
	}

	alive := false
	for i := 0; i <= retries; i++ {
		resp, err := client.Get(site)
		if err != nil {
			fmt.Printf("%s:\t looking at %s:\t TIMEOUT\n", proxy, site)
			continue
		}

		if resp.StatusCode == 200 {
			defer resp.Body.Close()
			alive = true
			break
		}
		fmt.Printf("%s:\t looking at %s:\t FAILED\n", proxy, site)
	}

	if alive {
		proxies.MU.Lock()
		fmt.Printf("%s:\t looking at %s:\t SUCCESS\n", proxy, site)
		proxies.Proxies = append(proxies.Proxies, proxy)
		proxies.MU.Unlock()
	}
}

func tryProxysHTTP(proxyList []string, website string, retries int, timeout time.Duration) []string {
	var wg sync.WaitGroup

	validProxies := validProxy{Proxies: make([]string, 0), MU: sync.Mutex{}}
	for _, proxy := range proxyList {
		wg.Add(1)
		go tryProxyHTTP(proxy, website, retries, timeout, &validProxies, &wg)
	}
	wg.Wait()

	return validProxies.Proxies
}
