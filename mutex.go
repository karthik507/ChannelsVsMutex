package main

import (
	"fmt"
	"sync"
)

type LoadBalancer struct {
	mu   sync.Mutex
	ips  []string
	next int
}

func NewLoadBalancer(ips []string) *LoadBalancer {
	return &LoadBalancer{
		ips: ips,
	}
}

func (lb *LoadBalancer) Next() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if len(lb.ips) == 0 {
		return ""
	}

	ip := lb.ips[lb.next]
	lb.next = (lb.next + 1) % len(lb.ips)
	return ip
}

func main() {
	lb := NewLoadBalancer([]string{"10.12.1.1", "10.12.1.2"})
	fmt.Println(lb.Next())
	fmt.Println(lb.Next())
	fmt.Println(lb.Next())
}
