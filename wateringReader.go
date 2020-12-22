package main

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
)

var (
	pin = rpio.Pin(21)
)

func main() {
	if err := rpio.Open(); err != nil {
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Input()

	for true {
		res := pin.Read()
		fmt.Println(res)
	}
}
