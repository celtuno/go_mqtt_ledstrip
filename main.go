// package main

// import (
// "fmt"
// )

// func main() {

// 	fmt.Println("Hello Pi world")
// }

package main

import (
	"fmt"
	"log"
	"time"

	"machine"

	"github.com/stianeikeland/go-rpio/v4"
	"tinygo.org/x/drivers/ws2812"
)

var neo machine.Pin = machine.WS2812

func main() {

	fmt.Println("Hello Pi world")
	err := rpio.Open()
	if err != nil {
		log.Printf("Fatal gpio error %v\n", err)
	}
	defer rpio.Close()

	pin := rpio.Pin(17)
	pin.Output()
	pin.Write(rpio.High)
	fmt.Println("Wait 2")
	time.Sleep(2000)
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}

	// led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	// tmp := machine
	// dev := ws2812.NewWS2812(LED)
	// fmt.Println(dev.Pin)
	// pin.DutyCycle()

	// leds := ledstrip.Init_ledstrip(5)

	// fmt.Printf("Test%v", leds.Pin)

}

// type MessageType int

// const (
// 	Color MessageType = iota
// 	Clear MessageType = iota
// )

// type RGBW struct {
// 	R byte
// 	G byte
// 	B byte
// 	W byte
// }

// var ledStrip ws2812.Device

// var (
// 	Colors = []RGBW{
// 		RGBW{0x64, 0x00, 0x00, 0x00}, // Red
// 		RGBW{0x00, 0x64, 0x00, 0x00}, // Green
// 		RGBW{0x00, 0x00, 0x64, 0x00}, // Blue
// 		RGBW{0x00, 0x00, 0x00, 0x64}, // White
// 		RGBW{0x00, 0x64, 0x64, 0x00}, // Gopher
// 		RGBW{0x64, 0x64, 0x00, 0x00}, // Yellow
// 		RGBW{0x64, 0x32, 0x00, 0x00}, // Orange
// 	}
// )

// func (c *RGBW) GBRASlice() []byte {
// 	return []byte{c.G, c.R, c.B, c.W}
// }

// func (d *ws2812.Device) Init(pixels int) {
// func Init(pixels int) {
// 	// n.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
// 	ledStrip = ws2812.New(5)
// 	fmt.Println(ledStrip.Pin)
// 	// for i := 0; i < pixels; i++ {
// 	// 	n.NeoPixels = append(n.NeoPixels, &NeoPixel{RGBW{0x00, 0x00, 0x00, 0x00}})
// 	// }

// }
// func setColors(colors []int) {
// 	buf := new(bytes.Buffer)
// 	for _, color := range Colors {
// 		buf.Write(color.GBRASlice())
// 	}
// 	ledStrip.Write(buf.Bytes())

// }
