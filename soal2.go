package main

import (
	"fmt"
	"sync"
	"time"
)
func printFruit(index *int, fruit []string, wg *sync.WaitGroup, mu *sync.Mutex){
	mu.Lock()
	defer mu.Unlock()	
	*index++
	fmt.Println(*index,fruit)	
	
	wg.Done()
}
func main() {
	time.Sleep(time.Second)
	var secret = []interface{}{
		[]string{"coba1", "coba2", "coba3", "coba4"},
		[]string{"bisa1", "bisa2", "bisa3", "bisa4"},
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	var j int
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go printFruit(&j, secret[0].([]string), &wg,&mu)
		wg.Add(2)
		go printFruit(&j, secret[1].([]string), &wg,&mu)
	}
	wg.Wait()
}