package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func query() {
	timeNow := time.Now()
	resp, err := http.Get("http://manage.fastly.com")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Status: %s, Response Time: %s\n", resp.Status, time.Since(timeNow))
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			query()
		}()
	}
	wg.Wait()
}
