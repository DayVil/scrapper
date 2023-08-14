package protocols

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
