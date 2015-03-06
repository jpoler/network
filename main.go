package main

import (
	"fmt"
	"github.com/jpoler/network/ping"
)

func main() {
	err := ping.ICMPEcho("74.125.226.55")
	if err != nil {
		fmt.Println(err)
	}
}
