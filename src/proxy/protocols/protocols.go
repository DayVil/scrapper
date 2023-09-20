package protocols

import "time"

// Different Protocols enum
type Protocols int8

const (
	HttpProt Protocols = iota
	Socks4
	Socks5
)

type ProxySites struct {
	Url      string
	Protocol Protocols
}

type ProxyAdrr string

const (
	TIMEOUT   = 10 * time.Second
	CHECKSITE = "http://pool.proxyspace.pro/judge.php"
	RETRIES   = 3
)
