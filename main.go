package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"time"

	"go.einride.tech/can"
	"go.einride.tech/can/pkg/socketcan"
)

const (
	// This is purely for spamming data, more realistic hz for CAN should be about 100hz
	SETTINGS_HZ = 100
	SETTINGS_ECU = "S300"
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
	Mil					uint8
	Vts					uint8
	Cl					uint8
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
type Frame665 struct {
	Knock				uint16
}
type Frame666 struct {
	TargetCamAngle	float64
	ActualCamAngle	float64
}
type Frame667 struct {
	OilTemp			uint16
	OilPressure	uint16
	//Analog2			uint16
	//Analog3			uint16
}
// type Frame668 struct {
// 	Analog4			uint16
// 	Analog5			uint16
// 	Analog6			uint16
// 	Analog7			uint16
// }
type Frame669S300 struct {
	Frequency			uint8
	Duty					float64
	Content				float64
}
type Frame669KPRO struct {
	Frequency				uint8
	EthanolContent	float64
	FuelTemperature	uint16
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

	frame665 = Frame665 {
		Knock: 0,
	}

	frame666 = Frame666 {
		TargetCamAngle: 1.0,
		ActualCamAngle: 1.0,
	}

	frame667 = Frame667 {
		OilTemp: 10,
		OilPressure: 10,
	}

	// frame668 = Frame668 {
	// 	Analog4: 0,
	// 	Analog5: 0,
	// 	Analog6: 0,
	// 	Analog7: 0,
	// }

	frame669S300 = Frame669S300 {
		Frequency: 1,
		Duty: 1.0,
		Content: 1.0,
	}

	frame669KPRO = Frame669KPRO {
		Frequency: 1,
		EthanolContent: 1.0,
		FuelTemperature: 1.0,
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
		if (SETTINGS_ECU == "KPRO") {
			if (frame661.Mil == 1) {
				frame661.Mil = 0
			} else {
				frame661.Mil = 1
			}

			if (frame661.Vts == 1) {
				frame661.Vts = 0
			} else {
				frame661.Vts = 1
			}

			if (frame661.Cl == 1) {
				frame661.Cl = 0
			} else {
				frame661.Cl = 1
			}
		}

    frame661.Iat += 5
    frame661.Ect += 10

    binary.BigEndian.PutUint16(frame.Data[0:2], frame661.Iat)
    binary.BigEndian.PutUint16(frame.Data[2:4], frame661.Ect)
		
		if (SETTINGS_ECU == "KPRO") {
			frame.Data[5] = frame661.Mil
			frame.Data[6] = frame661.Vts
			frame.Data[7] = frame661.Cl
		}

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
    if (frame663.Inj == 11900) {
      frame663.Inj = 11200
    }
    if (frame663.Ign == 20) {
      frame663.Ign = 0
    }

    frame663.Inj += 100
    frame663.Ign += 1

    binary.BigEndian.PutUint16(frame.Data[0:2], uint16(frame663.Inj))
    binary.BigEndian.PutUint16(frame.Data[2:4], frame663.Ign)

  case 664:
    if (frame664.LambdaRatio == 50000) {
      frame664.LambdaRatio = 46686
    }

    frame664.LambdaRatio += 1

    binary.BigEndian.PutUint16(frame.Data[0:2], frame664.LambdaRatio)

	case 665:
		if (SETTINGS_ECU == "KPRO") {
			if (frame665.Knock == 10) {
				frame665.Knock = 0
			}
			
			frame665.Knock += 1

			binary.BigEndian.PutUint16(frame.Data[0:2], frame665.Knock)
		}

	case 666:
		if (SETTINGS_ECU == "KPRO") {
			if (frame666.TargetCamAngle == 10.0) {
				frame666.TargetCamAngle = 0.0
			}
			if (frame666.ActualCamAngle == 10.0) {
				frame666.ActualCamAngle = 0.0
			}

			frame666.TargetCamAngle += 0.5
			frame666.ActualCamAngle += 0.5

			binary.BigEndian.PutUint16(frame.Data[0:2], uint16(frame666.TargetCamAngle))
			binary.BigEndian.PutUint16(frame.Data[2:4], uint16(frame666.ActualCamAngle))
		}

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

	// case 668:

	case 669:
		if (SETTINGS_ECU == "S300") {
			if (frame669S300.Frequency == 20) {
				frame669S300.Frequency = 0
			}
			if (frame669S300.Duty == 20.0) {
				frame669S300.Duty = 0.0
			}
			if (frame669S300.Content == 20.0) {
				frame669S300.Content = 0.0
			}

			frame669S300.Frequency += 2
			frame669S300.Duty += 0.5
			frame669S300.Content += 0.2

			frame.Data[0] = frame669S300.Frequency
			frame.Data[1] = byte(frame669S300.Duty)
			frame.Data[2] = byte(frame669S300.Content)

		} else if (SETTINGS_ECU == "KPRO") {
			if (frame669KPRO.Frequency == 20) {
				frame669KPRO.Frequency = 0
			}
			if (frame669KPRO.EthanolContent == 20.0) {
				frame669KPRO.EthanolContent = 0.0
			}
			if (frame669KPRO.FuelTemperature == 40) {
				frame669KPRO.FuelTemperature = 0
			}

			frame669KPRO.Frequency += 2
			frame669KPRO.EthanolContent += 0.5
			frame669KPRO.FuelTemperature += 2

			frame.Data[0] = frame669KPRO.Frequency
			frame.Data[1] = byte(frame669KPRO.EthanolContent)
			binary.BigEndian.PutUint16(frame.Data[2:4], frame669KPRO.FuelTemperature)
		}
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
	frame665 := can.Frame{
		ID: 665,
		Length: 2,
		Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
	}
	frame666 := can.Frame{
		ID: 666,
		Length: 2,
		Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
	}
  frame667 := can.Frame{
    ID: 667,
    Length: 4,
    Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
  }
	// frame668 := can.Frame{
  //   ID: 668,
  //   Length: 4,
  //   Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
  // }
	frame669S300 := can.Frame{
    ID: 669,
    Length: 4,
    Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
  }
	frame669KPRO := can.Frame{
    ID: 669,
    Length: 4,
    Data: [8]byte { 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00 },
  }

	ticker := time.NewTicker(time.Second / time.Duration(SETTINGS_HZ))
	defer ticker.Stop()

	tx := socketcan.NewTransmitter(conn)
	counter := 0

	for {
		select {
		case <-ticker.C:
			switch (counter) {
			case 0:
				incrementFrameData(&frame660)
				_ = tx.TransmitFrame(context.Background(), frame660)
				fmt.Println("Sent 660: ", frame660)

			case 1:
				incrementFrameData(&frame661)
				_ = tx.TransmitFrame(context.Background(), frame661)
				fmt.Println("Sent 661: ", frame661)

			case 2:
				incrementFrameData(&frame662)
				_ = tx.TransmitFrame(context.Background(), frame662)
				fmt.Println("Sent 662: ", frame662)

			case 3:
				incrementFrameData(&frame663)
				_ = tx.TransmitFrame(context.Background(), frame663)
				fmt.Println("Sent 663: ", frame663)

      case 4:
        incrementFrameData(&frame664)
        _ = tx.TransmitFrame(context.Background(), frame664)
        fmt.Println("Sent 664: ", frame664)

			case 5:
				if (SETTINGS_ECU == "KPRO") {
					incrementFrameData(&frame665)
					_ = tx.TransmitFrame(context.Background(), frame665)
					fmt.Println("Sent 665: ", frame665)
				}

			case 6:
				if (SETTINGS_ECU == "KPRO") {
        	incrementFrameData(&frame666)
        	_ = tx.TransmitFrame(context.Background(), frame666)
        	fmt.Println("Sent 666: ", frame666)
				}

      case 7:
        incrementFrameData(&frame667)
        _ = tx.TransmitFrame(context.Background(), frame667)
        fmt.Println("Sent 667: ", frame667)

			// case 8:
      //   incrementFrameData(&frame668)
      //   _ = tx.TransmitFrame(context.Background(), frame668)
      //   fmt.Println("Sent 668: ", frame668)

			case 9:
				if (SETTINGS_ECU == "S300") {
					incrementFrameData(&frame669S300)
					_ = tx.TransmitFrame(context.Background(), frame669S300)
					fmt.Println("Sent 669: ", frame669S300)
				} else if (SETTINGS_ECU == "KPRO") {
					incrementFrameData(&frame669KPRO)
					_ = tx.TransmitFrame(context.Background(), frame669KPRO)
					fmt.Println("Sent 669: ", frame669KPRO)
				}				
      }

			if (counter == 9) {
				counter = 0
			} else {
				counter++
			}
		}
	}
}