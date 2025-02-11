package main

import "fmt"

type Foot float64
type Meter float64

func (f Foot) String() string  { return fmt.Sprintf("%gfeet", f) }
func (m Meter) String() string { return fmt.Sprintf("%gmeters", m) }

func FoM(f Foot) Meter { return Meter(f * 0.3048) }

// FToC converts a Fahrenheit temperature to Celsius.
func MToF(m Meter) Foot { return Foot(m / 0.3048) }

func main() {
	fmt.Println(FoM(100))
	fmt.Println(MToF(100))
}
