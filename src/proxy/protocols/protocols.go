package protocols

// Different Protocols enum
type Protocols int8

const (
	HttpProt Protocols = iota
	Socks5
	Socks4
)

type ProxySites struct {
	Url string
	Protocol Protocols
}
