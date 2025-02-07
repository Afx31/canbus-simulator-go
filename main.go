package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"

	"go.einride.tech/can"
	"go.einride.tech/can/pkg/socketcan"
)

type Frame660 struct {
	Rpm				uint16
	Speed			uint16
	Gear			uint8
	Voltage		uint8
}

var (
	frame660 = Frame660 {
		Rpm: 100,
		Speed: 0,
		Gear: 0,
		Voltage: 138,
	}
)

func incrementFrameData(frame *can.Frame) {
	// https://go.dev/tour/moretypes/7


  // Assume its for 660 only
	
	// Reset everything back to min values. Temp for now
	if (frame660.Rpm == 9000) {
		frame660.Rpm = 100
	}
	if (frame660.Speed == 300) {
		frame660.Speed = 0
	}
	if (frame660.Gear == 6) {
		frame660.Gear = 0
	}
	if (frame660.Voltage == 150) {
		frame660.Voltage = 138
	}

	frame660.Rpm += 100
	frame660.Speed += 5
	frame660.Gear += 1
	frame660.Voltage += 1

  binary.BigEndian.PutUint16(frame.Data[0:2], frame660.Rpm)
  binary.BigEndian.PutUint16(frame.Data[2:4], frame660.Speed)
  frame.Data[4] = frame660.Gear
  frame.Data[5] = frame660.Voltage
}

func main() {
	conn, _ := socketcan.DialContext(context.Background(), "can", "vcan0")

	frame660 := can.Frame{
		ID: 660,
		Length: 8,
		Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
	}
	frame661 := can.Frame{
		ID: 0x661,
		Length: 4,
		Data: [8]byte { 0xAA, 0xBB, 0xCC, 0xDD },
	}
	frame662 := can.Frame{
		ID: 662,
		Length: 8,
		Data: [8]byte { 00, 11, 22, 33, 44, 55, 66, 77 },
	}
	frame663 := can.Frame{
		ID: 663,
		Length: 8,
		Data: [8]byte { 00, 11, 22, 33, 44, 55, 66, 77 },
	}
	
	tx := socketcan.NewTransmitter(conn)

	ticker := time.NewTicker(time.Duration(100) * time.Millisecond)
	defer ticker.Stop()

	counter := 0

	for {
		select {
		case <-ticker.C:

			switch (counter) {
			case 0:
				incrementFrameData(&frame660)
				_ = tx.TransmitFrame(context.Background(), frame660)
				fmt.Println("Sent 660: ", frame660)
				break
			case 1:
				// incrementFrameData(&frame661)
				_ = tx.TransmitFrame(context.Background(), frame661)
				fmt.Println("Sent 661: ", frame661)
				break
			case 2:
				// incrementFrameData(&frame662)
				_ = tx.TransmitFrame(context.Background(), frame662)
				fmt.Println("Sent 662: ", frame662)
				break
			case 3:
				// incrementFrameData(&frame663)
				_ = tx.TransmitFrame(context.Background(), frame663)
				fmt.Println("Sent 663: ", frame663)
				break
			}

			if (counter == 3) {
				counter = 0
			} else {
				counter++
			}
		}
	}
}