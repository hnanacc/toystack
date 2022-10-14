package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("Hello, world\n\n")

	hn, _ := os.Hostname()
	print("Hostname", hn)
	print("CurrentProcessId", os.Getpid())
	print("MaxProcessId", "undefined")
}

func print(k string, v interface{}) {
	fmt.Printf("%20s: %v\n", k, v)
}