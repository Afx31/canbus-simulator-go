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
	Rpm					uint16
	Speed				uint16
	Gear				uint8
	Voltage			uint8
}

type Frame661 struct {
	Iat					uint16
	Ect					uint16
}

type Frame662 struct {
	Tps					uint16
	Map					uint16
}

type Frame663 struct {
	Inj					float64
	Ign					uint16
}

type Frame664 struct {
	LambdaRatio uint16
}

type Frame667 struct {
	OilTemp			uint16
	OilPressure	uint16
}

var (
	frame660 = Frame660 {
		Rpm: 100,
		Speed: 0,
		Gear: 0,
		Voltage: 138,
	}

	frame661 = Frame661 {
		Iat: 10,
		Ect: 40,
	}

	frame662 = Frame662 {
		Tps: 0,
		Map: 0,
	}

	frame663 = Frame663 {
		Inj: 4,
		Ign: 5,
	}

	frame664 = Frame664 {
		LambdaRatio: 46686,
	}

	frame667 = Frame667 {
		OilTemp: 10,
		OilPressure: 10,
	}
)

func incrementFrameData(frame *can.Frame) {
	// https://go.dev/tour/moretypes/7

	switch (frame.ID) {
	case 660:
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
	
	case 661:
		if (frame661.Iat == 60) {
			frame661.Iat = 10
		}
		if (frame661.Ect == 140) {
			frame661.Ect = 0
		}

    frame661.Iat += 5
    frame661.Ect += 10

    binary.BigEndian.PutUint16(frame.Data[0:2], frame661.Iat)
    binary.BigEndian.PutUint16(frame.Data[2:4], frame661.Ect)

  case 662:
    if (frame662.Tps == 100) {
      frame662.Tps = 0
    }
    if (frame662.Map == 100) {
      frame662.Map = 0
    }

    frame662.Tps += 10
    frame662.Map += 5

    binary.BigEndian.PutUint16(frame.Data[0:2], frame662.Tps)
    binary.BigEndian.PutUint16(frame.Data[2:4], frame662.Map)
    
  case 663:
    if (frame663.Inj == 20) {
      frame663.Inj = 0
    }
    if (frame663.Ign == 20) {
      frame663.Ign = 0
    }

    frame663.Inj += 1
    frame663.Ign += 1

    binary.BigEndian.PutUint16(frame.Data[0:2], uint16(frame663.Inj))
    binary.BigEndian.PutUint16(frame.Data[2:4], frame663.Ign)

  case 664:
    if (frame664.LambdaRatio == 50000) {
      frame664.LambdaRatio = 46686
    }

    frame664.LambdaRatio += 1

    binary.BigEndian.PutUint16(frame.Data[0:2], frame664.LambdaRatio)

  case 667:
    if (frame667.OilTemp == 150) {
      frame667.OilTemp = 0
    }
    if (frame667.OilPressure == 100) {
      frame667.OilPressure = 0
    }

    frame667.OilTemp += 10
    frame667.OilPressure += 5

    binary.BigEndian.PutUint16(frame.Data[0:2], frame667.OilTemp)
    binary.BigEndian.PutUint16(frame.Data[2:4], frame667.OilPressure)
	}
}

func main() {
	conn, _ := socketcan.DialContext(context.Background(), "can", "vcan0")

	frame660 := can.Frame{
		ID: 660,
		Length: 6,
		Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
	}
	frame661 := can.Frame{
		ID: 661,
		Length: 4,
		Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
	}
	frame662 := can.Frame{
		ID: 662,
		Length: 4,
		Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
	}
	frame663 := can.Frame{
		ID: 663,
		Length: 4,
		Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
	}
  frame664 := can.Frame{
    ID: 664,
    Length: 2,
    Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
  }
  frame667 := can.Frame{
    ID: 667,
    Length: 4,
    Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
  }
	
	tx := socketcan.NewTransmitter(conn)

	ticker := time.NewTicker(time.Duration(10) * time.Millisecond)
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
				incrementFrameData(&frame661)
				_ = tx.TransmitFrame(context.Background(), frame661)
				fmt.Println("Sent 661: ", frame661)
				break

			case 2:
				incrementFrameData(&frame662)
				_ = tx.TransmitFrame(context.Background(), frame662)
				fmt.Println("Sent 662: ", frame662)
				break

			case 3:
				incrementFrameData(&frame663)
				_ = tx.TransmitFrame(context.Background(), frame663)
				fmt.Println("Sent 663: ", frame663)
				break

      case 4:
        incrementFrameData(&frame664)
        _ = tx.TransmitFrame(context.Background(), frame664)
        fmt.Println("Sent 664: ", frame664)
        break

      case 5:
        incrementFrameData(&frame667)
        _ = tx.TransmitFrame(context.Background(), frame667)
        fmt.Println("Sent 667: ", frame667)
        break
      }

			if (counter == 5) {
				counter = 0
			} else {
				counter++
			}
		}
	}
}