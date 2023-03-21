package conv

var Fahrenheit float64
var Celsius float64
var Kelvin float64

// Konverterer Farhenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {
	Celsius = (Fahrenheit - 32.0) * 5.0 / 9.0
	return (value - 32.0) * 5.0 / 9.0
}

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(value float64) float64 {
	Fahrenheit = Celsius*(9.0/5.0) + 32.0

	return value*(9.0/5.0) + 32.0
}

// Konverterer Celsius til kelvin
func CelsiusToKelvin(value float64) float64 {
	Kelvin = Celsius + 273.15

	return value + 273.15
}

// Konverterer Kelvin til celsius
func KelvinToCelsius(value float64) float64 {
	Celsius = Kelvin - 273.15

	return value - 273.15
}

// Konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) float64 {
	Fahrenheit = (Kelvin * 9.0 / 5.0) - 460.0

	return (value * 9.0 / 5.0) - 460.0
}

// Konverterer Fahrenheit til kelvin
func FahrenheitToKelvin(value float64) float64 {
	Kelvin = (Fahrenheit + 460) * 5.0 / 9.0

	return (value + 460) * 5.0 / 9.0
}
