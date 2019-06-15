package main

import (
	"fmt"
	"github.com/sabhiram/go-rpio"
	"os"
	"time"
)

func fatalOnErr(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())

		os.Exit(1)
	}
}
func main() {
	fmt.Printf("Yo!\n")
	fatalOnErr(rpio.Open())
	defer rpio.Close()
	pin := rpio.Pin(19)
	pin.Mode(rpio.Pwm)
	pin.Freq(50 * 200)
	for i := 0; i < 24; i += 1 {
		pin.DutyCycle(uint32(i), 200)
		time.Sleep(200 * time.Millisecond)
	}

}
