package main

import (
	"fmt"
	"os"
	"time"

	rpio "github.com/sabhiram/go-rpio"
)

const (
	cFreqMultiplier = 20 * 10
)

func fatalOnErr(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())

		os.Exit(1)
	}
}

type PWMServo struct {
	pin rpio.Pin
}

func NewPWMServo(bcmpid uint8) (*PWMServo, error) {
	p := rpio.Pin(bcmpid)
	p.Mode(rpio.Pwm)
	p.Freq(50 * cFreqMultiplier)
	return &PWMServo{pin: p}, nil
}

// SetAngle sets the servo angle to between 0 and 180 degrees.
func (s *PWMServo) SetAngle(angle float32) error {
	if angle < 0.0 || angle > 180.0 {
		return fmt.Errorf("invalid angle (%d) [0 <= angle <= 180]", angle)
	}

	// DutyCycle of 1.0ms / 20ms corresponds to 0 deg
	// 				1.5ms / 20ms corresponds to 90 deg
	//				2.0ms / 20ms corresponds to 180 deg
	s.pin.DutyCycle(uint32(((1.0+(angle/180.0))/20.0)*cFreqMultiplier), cFreqMultiplier)
	return nil
}

func main() {
	fatalOnErr(rpio.Open())
	defer rpio.Close()

	servo, err := NewPWMServo(19)
	fatalOnErr(err)

	for i := 0; i < 180; i += 1 {
		servo.SetAngle(float32(i))
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("Bye!")
}
