package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	pin       = rpio.Pin(21)
	relay_pin = rpio.Pin(20)
)

func main() {
	if err := rpio.Open(); err != nil {
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Input()

	for true {
		readResult := pin.Read()
		fmt.Println(readResult)
		if readResult != 1 {
			relay_pin.High()
			time.Sleep(3 * time.Second)
			relay_pin.Low()
		}
		time.Sleep(5 * time.Second)
	}
}
