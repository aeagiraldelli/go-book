package tempconv

import "fmt"

type (
	Celcius     float64
	Fahreinheit float64
)

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func CToF(c Celcius) Fahreinheit {
	return Fahreinheit(c*9/5 + 32)
}

func FToC(f Fahreinheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahreinheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
