package main

import (
	"fmt"
	"sync"
	"time"
)

// MAPSTART OMIT
type Greetings struct {
	mu sync.Mutex
	m  map[string]int
}

// MAPEND OMIT

// STARTCODE OMIT
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	m := map[string]int{"hello": 1}      // Start at 1
	go incrementByN(&wg, m, "hello", 10) // 11
	go incrementByN(&wg, m, "hello", 5)  // 16

	wg.Wait()
	fmt.Printf("Expected Value: (%d) Actual Value: (%d) \n", 16, m["hello"])
}

func incrementByN(wg *sync.WaitGroup, m map[string]int, key string, n int) {
	v := m[key]
	time.Sleep(10 * time.Millisecond)
	m[key] = v + n
	wg.Done()
}

// ENDCODE OMIT
