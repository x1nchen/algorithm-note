package main

import (
	"fmt"
	"time"
)

func hh(g *int) {
	fmt.Println(*g)
}

func main() {
	a := make(map[int]int)
	a[1] = 1
	a[2] = 2
	a[3] = 3
	for _, k := range a {
		fmt.Println(k)
		go hh(&k)
	}
	time.Sleep(3*time.Second)
}
