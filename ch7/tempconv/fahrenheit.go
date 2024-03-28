package main

type Fahrenheit float64

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
