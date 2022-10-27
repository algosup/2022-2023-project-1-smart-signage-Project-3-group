//go:build bluepill

package lorae5

import (
	"log"
	"machine"
)

type LoRaE5 struct {
	//LoRaE5 struct with the chosen UART and its state

	uart *machine.UART
	on   bool
}

func NewLorae5() *LoRaE5 {
	//allow a connection with the lorae5 and configure the chosen UART as its output
	uart := machine.UART2
	uart.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.A2, RX: machine.A3})

	return &LoRaE5{uart: uart}
}

func (l *LoRaE5) Join() error {
	//allow the bluepill to send commands with a message to the lorae5
	_, err := l.uart.Write([]byte("AT+JOIN\r\n"))
	msg1 := ""

	if err != nil {
		return err
	}

	for {
		if l.uart.Buffered() > 0 {
			rb, err := l.uart.ReadByte()
			if err != nil {
				println("Error: " + err.Error())
				continue
				// return ""
			}
			msg1 += string(rb)
			if msg1[len(msg1)-1] == '\n' {
				break
			}
		}
	}
	println(msg1)
	l.uart.Write([]byte("AT+MSG=\"Hello Thomas\"\r\n"))
	return nil
}

func (l *LoRaE5) Reader(data []byte) (n int, err error) {
	// Read from the RX buffer
	l.uart.Read(data)
	return n, err

}

func (l *LoRaE5) Writer(data []byte) (n int, err error) {
	// Write data from the UART
	l.uart.Write(data)
	return n, err

}

func SendNotification(data []byte) {
	//Send notification to the user depending on the AT command received
	if string(data) == "AT+" {
		log.Println("+AT OK")
	}
}
