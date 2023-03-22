package main

import (
	"fmt"
	"log"
	"minyr/yr"
)

func main() {
	// kaller på konverter funksjonen for å konvertere celsius til fahrenheut
	convertedTemperatures, err := yr.KonverterGrader()
	if err != nil {
		log.Fatal(err)
	}

	// Skriver komverterte temperaturer til en ny fil.
	err = yr.SkrivLinjer(convertedTemperatures, "KONVERTERT kjevik-temp-fahrenheit.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Temperaturer konvertert!")
}
