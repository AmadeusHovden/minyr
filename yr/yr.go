package yr

import (
	"bufio"
	"fmt"
	"log"
	"minyr/conv-kopi"
	"os"
	"strconv"
	"strings"
	//"io"
)

func openFil(filename string) (*os.File, error) { // funksjon for å åpne fil
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func lesLinjer(file *os.File) ([]string, error) { // funksjon for å lese fil
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Navn") || strings.HasPrefix(line, "Data") {
			continue // returnerer alle linjer utenom de som starter på navn og data.
		}
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func lukkFil(file *os.File) { //funksjon for å lukke fila.
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func SkrivLinjer(lines []string, filename string) error { //funksjon for å skrive linjene i fila
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer lukkFil(file)

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprint(writer, "Navn;Stasjon;Tid(norsk normaltid);Lufttemperatur") //skriver i første linje
	fmt.Fprintln(writer, "")                                               //setter det etter på neste linje.

	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	fmt.Fprint(writer, "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);endringen er gjort av Amadeus Hovden")
	return nil
}

func CelsiusToFahrenheit(celsius float64) float64 { //funksjon for konvertere gradene. Hentet fra conv
	return conv.CelsiusToFahrenheit(celsius)
}

func KonverterGrader() ([]string, error) { // funksjon for å åpne og konvertere gradene, og ignorer første linje.
	file, err := openFil("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		return nil, err
	}

	defer lukkFil(file)
	lines, err := lesLinjer(file)
	if err != nil {
		return nil, err
	}

	convertedTemperatures := make([]string, 0, len(lines)-1) // ikke ta med header linja

	for i, line := range lines {
		if i == 0 {
			continue // ignorer header linja
		}

		fields := strings.Split(line, ";")
		if i == 0 {
			continue // skip header linja
		}
		if len(fields) != 4 {
			return nil, fmt.Errorf("unexpected number of fields in line %d: %d", i, len(fields))
		}

		timestamp := fields[0]
		temperatureCelsius, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return nil, fmt.Errorf("could not parse temperature in line %d: %s", i, err)
		}
		temperatureFahrenheit := temperatureCelsius*(9.0/5.0) + 32.0

		convertedTemperature := fmt.Sprintf("%s;%s;%.2fF", strings.Join(fields[:3], ";"), timestamp, temperatureFahrenheit)
		convertedTemperatures = append(convertedTemperatures, convertedTemperature)
	}

	return convertedTemperatures, nil
}
