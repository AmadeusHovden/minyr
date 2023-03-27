package yr_test

import (
	"bufio"
	"minyr/yr"
	"os"
	"strings"
	"testing"
)

func TestTellLinjer(t *testing.T) {
	type test struct {
		filename string
		want     int
	}
	tests := []test{
		{filename: "kjevik-temp-celsius-20220318-20230318.csv", want: 16756}, //funket
	}

	for _, tc := range tests {
		file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
		if err != nil {
			t.Errorf("could not open file %s: %v", tc.filename, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineCount := 0
		for scanner.Scan() {
			lineCount++
		}

		if lineCount != tc.want {
			t.Errorf("unexpected number of lines in file %s. got %d, want %d", tc.filename, lineCount, tc.want)
		}
	}
}

func TestKonverterGrader(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8°F"},
		{input: "Kjevik;SN39040;18.03.2022 01:50;0", want: "Kjevik;SN39040;18.03.2022 01:50;32.0°F"},
		{input: "Kjevik;SN39040;18.03.2022 01:50;-11", want: "Kjevik;SN39040;18.03.2022 01:50;12.2°F"},
	}

	_, err := yr.KonverterGrader()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	file, err := yr.OpenFil("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer yr.LukkFil(file)

	lines, err := yr.LesLinjer(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, tt := range tests {
		var found bool
		for _, line := range lines {
			if strings.Contains(line, tt.input) {
				found = true
				if !strings.Contains(line, tt.want) {
					t.Errorf("test failed: want %q, got %q", tt.want, line)
				}
				break
			}
		}
		if !found {
			t.Errorf("test failed: input %q not found in file", tt.input)
		}
	}
}

func TestKonverterGraderDataGyldig(t *testing.T) { //funket
	want := "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET); endringen er gjort av Amadeus Hovden"
	_, err := yr.KonverterGrader()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	file, err := yr.OpenFil("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer yr.LukkFil(file)

	lines, err := yr.LesLinjer(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(lines[len(lines)-1], want) {
		t.Errorf("test failed: want %q, got %q", want, lines[len(lines)-1])
	}
}

func TestCelsiusGjennomsnitt(t *testing.T) { //funket
	want := 8.56
	got, err := yr.CelsiusGjennomsnitt()
	if err != nil {
		t.Fatalf("CelsiusGjennomsnitt() feilet med %v", err)
	}
	if got != want {
		t.Errorf("CelsiusGjennomsnitt() = %v; want %v", got, want)
	}
}
