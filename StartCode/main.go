package main

import (
	"log"
	"time"

	"github.com/algosup/2022-2023-project-1-smart-signage-Project-3-group/StartCode/hw/led"
	"github.com/algosup/2022-2023-project-1-smart-signage-Project-3-group/StartCode/hw/led/fakeled"
)

func main() {
	var light led.LED

	light = fakeled.NewFake()
	ledStr := light.String()
	println(ledStr)

	// blinkSimple(l)
	//blinkWithGoroutine(l)
}

func blinkSimple(l led.LED) {
	//Make the LED blink every 100ms
	for {
		time.Sleep(time.Millisecond * 100)
		l.Toggle()
	}
}

func blinkWithGoroutine(l led.LED) {
	ledCtrl := make(chan bool, 10)

	go func() {
		for {
			select {
			case <-ledCtrl:
				l.Toggle()
			}
		}
	}()
	var ledState bool

	for {
		time.Sleep(time.Millisecond * 100)

		ledState = !ledState
		log.Println(ledState)

		ledCtrl <- ledState

	}
}

func BrightnessLow(l led.LED) {
	//make the LED blink every 100ms to decrease its brightness
	for {
		time.Sleep(time.Millisecond * 100)
		l.Toggle()
		time.Sleep(time.Millisecond * 100)
		l.Toggle()

	}
}

func BrightnessHigh(l led.LED) {
	//set the LED's state to on and increase its brightness
}
