package ledstrip

import (
	"bytes"

	"tinygo.org/x/drivers/ws2812"
)

// func Init_rpio() {

// 	fmt.Println("Initializing rpio")
// 	err := rpio.Open()
// 	if err != nil {
// 		log.Printf("Fatal gpio error %v\n", err)
// 	}
// 	defer rpio.Close()
// 	pin := rpio.Pin(17)
// 	pin.Output()
// 	pin.Write(rpio.High)
// 	fmt.Println("5 blink confirmation")
// 	time.Sleep(500)
// 	for x := 0; x < 5; x++ {
// 		pin.Toggle()
// 		time.Sleep(time.Second / 2)
// 	}
// }

type MessageType int

const (
	Color MessageType = iota
	Clear
	Effect
)

type RGBW struct {
	R byte
	G byte
	B byte
	W byte
}

var (
	Colors = []RGBW{
		RGBW{0x64, 0x00, 0x00, 0x00}, // Red
		RGBW{0x00, 0x64, 0x00, 0x00}, // Green
		RGBW{0x00, 0x00, 0x64, 0x00}, // Blue
		RGBW{0x00, 0x00, 0x00, 0x64}, // White
		RGBW{0x00, 0x64, 0x64, 0x00}, // Gopher
		RGBW{0x64, 0x64, 0x00, 0x00}, // Yellow
		RGBW{0x64, 0x32, 0x00, 0x00}, // Orange
	}
)

func (c *RGBW) GBRASlice() []byte {
	return []byte{c.G, c.R, c.B, c.W}
}

type rgbledstrip struct {
	Device ws2812.Device
	Count  int
	Pin    uint8
}

// func (d *ws2812.Device) Init(pixels int) {
func Init_ledstrip(pixels int) *rgbledstrip {
	var strip rgbledstrip
	// n.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// var leds = new LedStrip;
	tmp_device := ws2812.NewWS2812(17)
	strip = rgbledstrip{Device: tmp_device}
	strip.Count = 0
	// pin := tmp_device.Pin
	// strip.pin = pin
	// // strip.pin = uint8()
	// fmt.Printf("Current pin:\t%f\n", strip.device.Pin)
	// for i := 0; i < pixels; i++ {
	// 	setColor(buf, 2, Colors[1])
	// 	// n.NeoPixels = append(n.NeoPixels, &NeoPixel{RGBW{0x00, 0x00, 0x00, 0x00}})
	// }
	// return strip
	return &strip
}

func (rgbstrip rgbledstrip) SetColor(buf *bytes.Buffer, ledCount int, color RGBW) {
	// buf := new(bytes.Buffer)
	for i := 0; i < ledCount; i++ {
		buf.Write(color.GBRASlice())
	}
	writeColor(rgbstrip, buf, color)
}

func writeColor(strip rgbledstrip, buf *bytes.Buffer, color RGBW) {
	// buf.Write(color.GBRASlice())
	// for _, color := range Colors {
	// }
	// ledStrip.Write(buf.Bytes())
	// _, err := strip.device.Write(buf.Bytes())
	// if err != nil {
	// 	fmt.Printf("Error writing colorbuffer to led(s):\n%v\n", err)
	// }

}

// func (ledstrip.WriteColor) WriteColor(colors [][]RGBW) {
// 	buf := new(bytes.Buffer)
// 	for _, color := range Colors {
// 		buf.Write(color.GBRASlice())
// 	}
// 	ledStrip.Write(buf.Bytes())

// }
