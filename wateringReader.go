package main

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	pin      = rpio.Pin(21)
	relayPin = rpio.Pin(20)
)

func main() {
	if err := rpio.Open(); err != nil {
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Input()
	relayPin.Output()

	for true {
		readResult := pin.Read()
		fmt.Println(time.Now().String(), ": ", readResult)
		if readResult != 1 {
			relayPin.High()
			time.Sleep(10 * time.Second)
			relayPin.Low()
		}
		time.Sleep(60 * time.Minute)
	}
}
