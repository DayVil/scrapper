package checker

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/DayVil/scrapper/src/proxy/protocols"
)

const (
	TIMEOUT   = 10 * time.Second
	CHECKSITE = "http://pool.proxyspace.pro/judge.php"
	RETRIES   = 3
)

type validProxyMU struct {
	Elements []protocols.ProxyAdrr
	MU       sync.Mutex
}

func TryProxiesDefaultW(addrs []protocols.ProxyAdrr) []protocols.ProxyAdrr {
	return TryProxiesDefault(addrs, CHECKSITE)
}

func TryProxiesDefault(addrs []protocols.ProxyAdrr, website string) []protocols.ProxyAdrr {
	return TryProxies(addrs, TIMEOUT, RETRIES, website)
}

func TryProxies(addrs []protocols.ProxyAdrr, timeout time.Duration, retries uint, testSite string) []protocols.ProxyAdrr {
	validProxys := make([]protocols.ProxyAdrr, 0)
	validMU := validProxyMU{
		Elements: validProxys,
		MU:       sync.Mutex{},
	}
	wg := sync.WaitGroup{}

	for _, addr := range addrs {
		addrURL, err := url.Parse(string(addr))
		if err != nil {
			fmt.Println(err)
			continue
		}

		transport := http.Transport{
			Proxy: http.ProxyURL(addrURL),
		}

		client := http.Client{
			Transport: &transport,
			Timeout:   timeout,
		}

		wg.Add(1)
		go tryProxy(&client, addr, retries, &validMU, &wg, testSite)
	}
	wg.Wait()

	return validMU.Elements
}

func tryProxy(client *http.Client, proxy protocols.ProxyAdrr, retries uint, validProxy *validProxyMU, wg *sync.WaitGroup, testSite string) {
	defer wg.Done()

	for i := 0; i < int(retries); i++ {
		resp, err := client.Get(testSite)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Status Code for " + testSite + " is " + strconv.Itoa(resp.StatusCode) + " for following proxy: " + string(proxy) + ", retry: " + strconv.Itoa(resp.StatusCode))
			continue
		}

		validProxy.MU.Lock()
		validProxy.Elements = append(validProxy.Elements, proxy)
		validProxy.MU.Unlock()
		fmt.Println("Added " + proxy)
		break
	}
}
