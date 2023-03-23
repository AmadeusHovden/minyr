package main

import (
	"bufio"
	"fmt"
	"log"
	"minyr/yr"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Velg convert, average, eller exit: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "convert" {
			convertedTemperatures, err := yr.KonverterGrader()
			if err != nil {
				log.Fatal(err)
			}

			err = yr.SkrivLinjer(convertedTemperatures, "KONVERTERT kjevik.txt")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Konvertering fullført!")

		} else if text == "average" {
			average, err := yr.GjsnittTemp()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Gjennomsnittstemperaturen er: %.2f grader Celsius\n", average)

		} else if text == "exit" {
			break
		} else {
			fmt.Println("Ugyldig kommando!")
		}
	}
}
