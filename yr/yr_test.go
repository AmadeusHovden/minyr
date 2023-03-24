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
		{filename: "kjevik-temp-celsius-20220318-20230318.csv", want: 16756},
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

/* func TestKonverterGrader(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
	}

	for _, tc := range tests {
		got := yr.KonverterGrader(tc.input)
		if !reflect.DeepEqual(tc.want, got[0]) {
			t.Errorf("expected: %v, got: %v", tc.want, got[0])
		}
	}
}
*/

func TestKonv(t *testing.T) {

}

func TestCelsiusGjennomsnitt(t *testing.T) {
	want := 8.56
	got, err := yr.CelsiusGjennomsnitt()
	if err != nil {
		t.Fatalf("CelsiusGjennomsnitt() feilet med %v", err)
	}
	if got != want {
		t.Errorf("CelsiusGjennomsnitt() = %v; want %v", got, want)
	}
}
