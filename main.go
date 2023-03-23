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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Velg convert, average, eller exit: ")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())

		if text == "convert" {
			_, err := os.Stat("KONV.csv")
			if err == nil {
				fmt.Print("Filen eksisterer allerede. Vil du generere filen på nytt? (j/n): ")
				scanner.Scan()
				answer := strings.ToLower(scanner.Text())
				if answer != "j" && answer != "n" {
					log.Fatal("Ugyldig svar")
				} else if answer == "n" {
					return
				}
			}

			convertedTemperatures, err := yr.KonverterGrader()
			if err != nil {
				log.Fatal(err)
			}

			err = yr.SkrivLinjer(convertedTemperatures, "KONV.csv")
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
