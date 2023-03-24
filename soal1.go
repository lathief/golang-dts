package main

import (
	"fmt"
	"time"
)
func printFruit(index int, fruit []string) {
	fmt.Println(index ,fruit)

}
func main() {
	var secret = []interface{}{
		[]string{"coba1", "coba2", "coba3", "coba4"},
		[]string{"bisa1", "bisa2", "bisa3", "bisa4"},
	}
	for i := 0; i < 4; i++ {

		go printFruit(i, secret[0].([]string))
		go printFruit(i, secret[1].([]string))
	}
	time.Sleep(time.Second)
}