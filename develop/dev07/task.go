package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	wg := sync.WaitGroup{}
	res := make(chan interface{})

	wg.Add(len(channels))
	doneChan := func(ch <-chan interface{}) {
		for v := range ch {
			res <- v
		}
		wg.Done()
	}

	for _, ch := range channels {
		go doneChan(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(5*time.Second),
		sig(8*time.Second),
	)

	fmt.Printf("done after %v", time.Since(start))
}
