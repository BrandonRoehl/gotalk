package gotalk_test

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

var urls = []string{
	"http://www.jamf.com/",
	"http://www.golang.org/",
	"http://www.google.com/",
	"http://www.bing.com/",
	"http://brandonroehl.org/",
}

func BenchmarkSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, url := range urls {
			// Fetch the URL.
			http.Get(url)
		}
	}
}

func BenchmarkConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for _, url := range urls {
			wg.Add(1)
			// Launch a goroutine to fetch the URL.
			go func(url string) {
				// Decrement the counter when the goroutine completes.
				defer wg.Done()
				// Fetch the URL.
				http.Get(url)
			}(url)
		}
		// Wait for all HTTP fetches to complete.
		wg.Wait()
	}
}

func fibonacciN(n int) int {
	x, y := 0, 1
	for i := 0; i < n-1; i++ {
		x, y = y, x+y
	}
	return x
}

func TestMutex(*testing.T) {
	// Go routines run in the same context so have access to the same variables
	// and memory
	var (
		wg sync.WaitGroup
		n  int
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		n = fibonacciN(10)
	}()
	fmt.Print("The 10th fibonacci number is ")
	wg.Wait() // wait for the go routine
	fmt.Println(n)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestChannels(*testing.T) {
	// Example from the go tour
	// https://tour.golang.org/concurrency/4
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// This way we are reading the response from the other task
	// over the channel and then we know exactly what we are trying to read
	// This allows the other process to go as fast as it can and "return" stuff
	// to the main thread without a mutex or setting a variable
	for i := range c {
		fmt.Println(i)
	}
}
