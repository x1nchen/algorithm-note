package main

import (
	"fmt"
)

func main() {
	tStr := "int"
	for _, r := range tStr {
		fmt.Println(r)
	}

	for i:=0; i<len(tStr); i++ {
		fmt.Println(tStr[i])
	}
}