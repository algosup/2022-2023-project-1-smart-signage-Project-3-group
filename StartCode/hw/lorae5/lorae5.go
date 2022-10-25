//go:build bluepill

package lorae5

import {
	"machine"
	"log"
}

type LoRaE5 struct {
	uart machine.UART2
	on   bool
}

func NewLorae5() *LoRaE5 {

	serial2 := machine.UART2
	serial2.Configure(machine.UARTConfig{
		BaudRate: 9600,
		RX:       machine.A3,
		TX:       machine.A2,
	})

	lorae5 := LoRaE5{
		uart: serial2,
	}

	return &LoRaE5
}

func (l *LoRaE5) Reader() (n int, err error) {
	l.uart.Read()

}

func (l *LoRaE5) Writer() (n int, err error) {
	l.uart.Write()

}

func SendNotification(data []byte) string{
	if data == "AT+"
	{
		log.Println("+AT OK")
	}
}
