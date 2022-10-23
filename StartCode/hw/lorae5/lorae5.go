//go:build bluepill

package lorae5

import "machine"

type LoRaE5 struct {
}

func New() *LoRaE5 {
	uart := machine.UART2{}
	_ = uart
	return &LoRaE5{}
}

func (l *LoRaE5) SendNotification() {
}
