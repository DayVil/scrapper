package scraper

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"sync"

	"github.com/DayVil/scrapper/src/proxy/protocols"
)

var re = regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}\b`)

func SearchSite(websites []protocols.ProxySites) []protocols.ProxyAdrr {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	addrs := make([]protocols.ProxyAdrr, 0)

	for _, website := range websites {
		wg.Add(1)
		go searchSite(website, &addrs, &wg, &mutex)
	}
	wg.Wait()

	return addrs
}

func searchSite(website protocols.ProxySites, addrs *[]protocols.ProxyAdrr, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	res, err := http.Get(website.Url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("ERROR Code: " + strconv.Itoa(res.StatusCode) + " from " + website.Url)
		return
	}

	text, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	ips := re.FindAll(text, -1)
	for _, ip := range ips {
		mutex.Lock()
		switch website.Protocol {
		case protocols.HttpProt:
			*addrs = append(*addrs, protocols.ProxyAdrr("http://"+string(ip)))
		case protocols.Socks4:
		case protocols.Socks5:
			*addrs = append(*addrs, protocols.ProxyAdrr("socks5://"+string(ip)))
		}
		mutex.Unlock()
	}
}
