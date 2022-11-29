package day04

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	content, err := os.ReadFile("example")
	if err != nil {
		log.Fatal(err)
	}

	res := Parse(content)

	var expected = map[string]string{
		"hcl": "#ae17e1",
		"iyr": "2013",
		"eyr": "2024",
		"ecl": "brn",
		"pid": "760753108",
		"byr": "1931",
		"hgt": "179cm",
	}

	if !reflect.DeepEqual(res[2], expected) {
		t.Fail()
	}
}

func TestValidExample(t *testing.T) {
	content, err := os.ReadFile("example")
	if err != nil {
		log.Fatal(err)
	}

	input := Parse(content)

	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	res := Valid(input, mandatory)

	if res != 2 {
		t.Errorf("Expected %d to be 2", res)
		t.Fail()
	}

}

func TestValid(t *testing.T) {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	input := Parse(content)

	mandatory := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	res := Valid(input, mandatory)

	if res != 247 {
		t.Errorf("Expected %d to be 2", res)
		t.Fail()
	}
}
