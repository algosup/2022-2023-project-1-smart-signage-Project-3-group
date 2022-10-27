package main

import (
	goio "io"
	"log"
	"time"

	"github.com/algosup/2022-2023-project-1-smart-signage-Project-3-group/StartCode/hw/led"
	"github.com/algosup/2022-2023-project-1-smart-signage-Project-3-group/StartCode/hw/led/realled"
)

func main() {

	//Main program
	//connect to a real LED and create a 45 seconds loop with different orders

	var light = realled.NewReal()

	for {
		//Switch on the LEDS for 15 seconds
		light.On()
		time.Sleep(15 * time.Second)
		//Switch off the LEDS for 15 seconds
		light.Off()
		time.Sleep(15 * time.Second)
		//Switch on the LEDS for 15 seconds and reduce its brightness
		light.On()
		BrightnessLow(light)
		time.Sleep(15 * time.Second)
	}

}

func SerialSend(serial goio.Writer, atCommand string) error {
	//Send an AT command to the lorae5

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
	//blink a led every 100ms by using a goroutine
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
	//get the LED's state

}
