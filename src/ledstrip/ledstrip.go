package ledstrip

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
	"tinygo.org/x/drivers/ws2812"
)

func Init_rpio() {

	fmt.Println("Initializing rpio")
	err := rpio.Open()
	if err != nil {
		log.Printf("Fatal gpio error %v\n", err)
	}
	defer rpio.Close()
	pin := rpio.Pin(17)
	pin.Output()
	pin.Write(rpio.High)
	fmt.Println("5 blink confirmation")
	time.Sleep(500)
	for x := 0; x < 5; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 2)
	}
}

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

var ledStrip ws2812.Device

// func (d *ws2812.Device) Init(pixels int) {
func Init_ledstrip(pixels int) {
	// n.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledStrip = ws2812.New(5)
	fmt.Println(ledStrip.Pin)
	// for i := 0; i < pixels; i++ {
	// 	n.NeoPixels = append(n.NeoPixels, &NeoPixel{RGBW{0x00, 0x00, 0x00, 0x00}})
	// }

}

func setColor(ledCount int, color RGBW) {
	buf := new(bytes.Buffer)
	for i := 0; i < ledCount; i++ {
		buf.Write(color.GBRASlice())
	}
	WriteColor(buf, color)
}

func WriteColor(buf *bytes.Buffer, color RGBW) {
	// buf.Write(color.GBRASlice())
	// for _, color := range Colors {
	// }
	// ledStrip.Write(buf.Bytes())
	ledStrip.Write(buf.Bytes())

}
func (ledstrip.WriteColor) WriteColor(colors [][]RGBW) {
	buf := new(bytes.Buffer)
	for _, color := range Colors {
		buf.Write(color.GBRASlice())
	}
	ledStrip.Write(buf.Bytes())

}
