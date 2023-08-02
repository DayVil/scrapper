package scrape

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

var re = regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}\b`)

// TODO: go routine
func GetProxyListIpV4(website string) []string {
	var proxyList []string

	u, err := url.Parse(website)
	if err != nil {
		log.Println(err)
		return proxyList
	}

	r, err := http.Get(u.String())
	if err != nil {
		log.Println(err)
		return proxyList
	}

	reader := r.Body
	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)

	for scanner.Scan() {
		line := scanner.Text()
		ip4 := re.FindAllString(line, -1)
		if len(ip4) != 0 {
			proxyList = append(proxyList, ip4...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	return proxyList
}
