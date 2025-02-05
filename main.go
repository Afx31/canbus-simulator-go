package main

import (
	"context"
	"fmt"
	"time"

	"go.einride.tech/can"
	"go.einride.tech/can/pkg/socketcan"
)

func main() {
	conn, _ := socketcan.DialContext(context.Background(), "can", "vcan0")

	frame660 := can.Frame{
		ID: 660,
		Length: 8,
		Data: [8]byte { 00, 11, 22, 33, 44, 55, 66, 77 },
	}
	frame661 := can.Frame{
		ID: 661,
		Length: 8,
		Data: [8]byte { 00, 11, 22, 33, 44, 55, 66, 77 },
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
				_ = tx.TransmitFrame(context.Background(), frame660)
				fmt.Println("Sent 660: ", frame660)
				break
			case 1:
				_ = tx.TransmitFrame(context.Background(), frame661)
				fmt.Println("Sent 661: ", frame661)
				break
			case 2:
				_ = tx.TransmitFrame(context.Background(), frame662)
				fmt.Println("Sent 662: ", frame662)
				break
			case 3:
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