package structures

import "sync"

type ValidProxy struct {
	Proxies []string
	MU      sync.Mutex
}
