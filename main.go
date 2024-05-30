package main

import "fmt"

// FahrenheitToCelsius converts a Fahrenheit temperature to Celsius.
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// CelsiusToFahrenheit converts a Celsius temperature to Fahrenheit.
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}u

func main() {
	fmt.Println("Created Project Structure")
}
