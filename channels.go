package main

import (
	"context"
	"fmt"
)

func NewLoadBalancer(ctx context.Context, ips []string) <-chan string {
	ch := make(chan string)
	if len(ips) == 0 {
		close(ch)
		return ch
	}

	go func() {
		defer close(ch)
		for {
			for _, ip := range ips {
				select {
				case <-ctx.Done():
					return
				case ch <- ip:
				}
			}
		}
	}()
	return ch
}

func main() {
	ips := []string{"1", "2"}
	ctx, cancel := context.WithCancel(context.Background())
	lb := NewLoadBalancer(ctx, ips)
	fmt.Println(<-lb)
	fmt.Println(<-lb)
	fmt.Println(<-lb)
	cancel()
}
