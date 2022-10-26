//go:build bluepill

package lorae5

import (
	"log"
	"machine"
)

type LoRaE5 struct {
	uart *machine.UART
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

	return &lorae5
}

func (l *LoRaE5) Reader(data []byte) (n int, err error) {
	l.uart.Read(data)
	return n, err

}

func (l *LoRaE5) Writer(data []byte) (n int, err error) {
	l.uart.Write(data)
	return n, err

}

func SendNotification(data []byte) {
	if string(data) == "AT+" {
		log.Println("+AT OK")
	}
}
