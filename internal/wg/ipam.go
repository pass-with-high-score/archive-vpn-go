package wg

import (
	"fmt"
	"sync"
)

type IPAM struct {
	mu     sync.Mutex
	nextIP int
}

func NewIPAM(start int) *IPAM {
	return &IPAM{
		nextIP: start,
	}
}

func (i *IPAM) Allocate() string {
	i.mu.Lock()
	defer i.mu.Unlock()

	ip := fmt.Sprintf("10.10.0.%d/32", i.nextIP)
	i.nextIP++

	return ip
}
