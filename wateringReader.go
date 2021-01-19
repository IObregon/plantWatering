package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	pin      = rpio.Pin(21)
	relayPin = rpio.Pin(20)
)

func main() {
	f, err := os.OpenFile("/home/pi/repositories/plantWatering/text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if err := rpio.Open(); err != nil {
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Input()
	relayPin.Output()

	for true {
		readResult := pin.Read()
		f.WriteString(fmt.Sprint(time.Now().String(), ": ", readResult, "\n"))

		if readResult == 1 {
			relayPin.High()
			time.Sleep(5 * time.Second)
			relayPin.Low()
			time.Sleep(120 * time.Minute)
		}
		time.Sleep(60 * time.Minute)
	}
}
