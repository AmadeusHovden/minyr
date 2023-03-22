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

	// Skriver konverterte temperaturer til en ny fil.
	err = yr.SkrivLinjer(convertedTemperatures, "KONVERTERT kjevik.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Suksess! Nå velg convert, average, eller exit.")

	yr.GjsnittTemp()
}
