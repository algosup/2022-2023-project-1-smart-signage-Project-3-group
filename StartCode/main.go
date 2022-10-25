package main

import (
	goio "io"
	"log"
	"time"

	"github.com/algosup/2022-2023-project-1-smart-signage-Project-3-group/StartCode/hw/io"
	"github.com/algosup/2022-2023-project-1-smart-signage-Project-3-group/StartCode/hw/led"
)

func main() {

	badSerial := io.NewSlowWriter(3)
	err := SerialSend(badSerial, "AT+JOIN")
	if err != nil {
		log.Fatal(err)
	}
	//var light led.LED

	//light = fakeled.NewFakeLED()
	//ledStr := light.String()
	//println(ledStr)

	// blinkSimple(l)
	//blinkWithGoroutine(l)
}

func SerialSend(serial goio.Writer, atCommand string) error {

	for len(atCommand) > 0 {

		n, err := serial.Write([]byte(atCommand))
		if err != nil {
			return err
		}
		atCommand = atCommand[n:]
		log.Println(n, atCommand)

	}

	return nil
}

func getLEDInfos(l led.LED) {
	//Check the state of the LEDs and gives some information
}

func sendATCommands() {
	//Send AT commands to the lorae5
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

func GetLedInfo(l led.LED) {

}
