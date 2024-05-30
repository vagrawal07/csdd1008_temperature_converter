package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// FahrenheitToCelsius converts a Fahrenheit temperature to Celsius.
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// CelsiusToFahrenheit converts a Celsius temperature to Fahrenheit.
func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter temperature (e.g., 0 F or 0 C) or type 'exit' to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting the program.")
			break
		}

		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Invalid input format. Please enter the temperature followed by the unit (e.g., 0 F or 0 C).")
			continue
		}

		temperature, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Invalid temperature value. Please enter a valid number.")
			continue
		}

		unit := strings.ToUpper(parts[1])
		switch unit {
		case "F":
			celsius := FahrenheitToCelsius(temperature)
			fmt.Printf("%.2f F is %.2f C\n", temperature, celsius)
		case "C":
			fahrenheit := CelsiusToFahrenheit(temperature)
			fmt.Printf("%.2f C is %.2f F\n", temperature, fahrenheit)
		default:
			fmt.Println("Invalid unit. Please use 'C' for Celsius or 'F' for Fahrenheit.")
		}
	}
}
