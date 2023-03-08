package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s, sep string
	var index string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		index += sep + strconv.Itoa(i-1)
		sep = " "
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(index)
}
