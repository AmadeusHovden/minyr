package yr_test

import (
	"bufio"
	"minyr/yr"
	"os"
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
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8°F"}, //funket
		{input: "Kjevik;SN39040;18.03.2022 01:50;0", want: "Kjevik;SN39040;18.03.2022 01:50;32.0°F"},
		{input: "Kjevik;SN39040;18.03.2022 01:50;-11", want: "Kjevik;SN39040;18.03.2022 01:50;12.2°F"},
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;", want: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Amadeus Hovden"},
	}
	got, err := yr.KonverterGrader()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i, tt := range tests {
		if i >= len(got) {
			t.Fatalf("not enough converted temperatures, got %d, want at least %d", len(got), i+1)
		}

		if got[i] != tt.want {
			t.Errorf("unexpected result for test %d:\ngot  %q\nwant %q", i+1, got[i], tt.want)
		}
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
