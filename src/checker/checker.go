package checker

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/DayVil/scrapper/src/proxy/protocols"
)

// validProxyMU is a mutex that protects the Elements field of validProxy.
type validProxyMU struct {
	Elements []protocols.ProxyAdrr
	MU       sync.Mutex
}

// TryProxiesDefaultW is a wrapper for TryProxiesDefault that uses the default website.
func TryProxiesDefaultW(addrs []protocols.ProxyAdrr) []protocols.ProxyAdrr {
	return TryProxiesDefault(addrs, protocols.CHECKSITE)
}

// TryProxiesDefault is a wrapper for TryProxies that uses the default timeout and retries.
func TryProxiesDefault(addrs []protocols.ProxyAdrr, website string) []protocols.ProxyAdrr {
	return TryProxies(addrs, protocols.TIMEOUT, protocols.RETRIES, website)
}

// TryProxies returns a list of valid proxies from the specified addresses.
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
			fmt.Fprintln(os.Stderr, err)
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

// tryProxy tries to connect to the specified proxy.
func tryProxy(client *http.Client, proxy protocols.ProxyAdrr, retries uint, validProxy *validProxyMU, wg *sync.WaitGroup, testSite string) {
	defer wg.Done()

	for i := 0; i < int(retries); i++ {
		resp, err := client.Get(testSite)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Fprintln(os.Stderr, "Status Code for "+testSite+" is "+strconv.Itoa(resp.StatusCode)+" for following proxy: "+string(proxy)+", retry: "+strconv.Itoa(resp.StatusCode))
			continue
		}

		validProxy.MU.Lock()
		validProxy.Elements = append(validProxy.Elements, proxy)
		validProxy.MU.Unlock()
		fmt.Fprintln(os.Stderr, "Added "+proxy)
		break
	}
}
