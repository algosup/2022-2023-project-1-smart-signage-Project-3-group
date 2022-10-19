package main

import (
	"log"
	"time"

	"github.com/tuxago/go/dolanor/tinygo/hw/led"
	"github.com/tuxago/go/dolanor/tinygo/hw/led/fakeled"
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
	for {
		time.Sleep(time.Millisecond * 100)
		l.Toggle()
		time.Sleep(time.Millisecond * 100)
		l.Toggle()

	}

}
