package main

import (
	helpers "AoG/helpers"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getPassportStrings(input []string) []string {
	var result []string
	current := ""
	for _, line := range input {
		if len(line) == 0 {
			current = strings.TrimSpace(current)
			result = append(result, current)
			current = ""
			continue
		}
		current += fmt.Sprintf(" %s", line)
	}
	current = strings.TrimSpace(current)
	result = append(result, current)
	return result
}

type Passport map[string]string

func parsePassport(input string) Passport {
	result := make(Passport)
	entries := strings.Split(input, " ")
	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		splits := strings.Split(entry, ":")
		result[splits[0]] = splits[1]
	}
	return result
}

func validatePassportKeys(input string) bool {
	pp := parsePassport(input)
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // ignore cid
	for _, r := range required {
		_, ok := pp[r]
		if !ok {
			return false
		}
	}
	return true
}

type validationFunc func(string) bool

type passportValidator struct {
	validators map[string]validationFunc
}

func validStrInt(in string, min, max int) bool {
	value, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return value >= min && value <= max
}

func createPPValidator() passportValidator {
	var validator passportValidator
	validator.validators = make(map[string]validationFunc)
	//byr
	validator.validators["byr"] = func(in string) bool {
		return validStrInt(in, 1920, 2002)
	}
	//iyr
	validator.validators["iyr"] = func(in string) bool {
		return validStrInt(in, 2010, 2020)
	}
	//eyr
	validator.validators["eyr"] = func(in string) bool {
		return validStrInt(in, 2020, 2030)
	}

	//hgt
	validator.validators["hgt"] = func(in string) bool {
		numPart := in[:len(in)-2]
		unit := in[len(in)-2:]
		if unit == "cm" {
			return validStrInt(numPart, 150, 193)
		} else if unit == "in" {
			return validStrInt(numPart, 59, 76)
		}
		return false
	}

	//hcl
	validator.validators["hcl"] = func(in string) bool {
		if len(in) != 7 {
			return false
		}
		firstChar := in[0]
		if firstChar != '#' {
			return false
		}
		for _, char := range in[1:] {
			if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f')) {
				return false
			}
		}
		return true
	}

	//ecl
	validator.validators["ecl"] = func(in string) bool {
		options := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, o := range options {
			if in == o {
				return true
			}
		}
		return false
	}

	//pid
	validator.validators["pid"] = func(in string) bool {
		if len(in) != 9 {
			return false
		}
		_, err := strconv.Atoi(in)
		if err != nil {
			log.Fatal(err)
			return false
		}
		return true
	}

	//cid
	validator.validators["cid"] = func(in string) bool {
		return true
	}

	return validator
}

func (p *passportValidator) validatePassport(input string) bool {
	if !validatePassportKeys(input) {
		return false
	}

	pp := parsePassport(input)
	for key, value := range pp {
		v, ok := p.validators[key]
		if !ok {
			continue
		}
		valid := v(value)
		if !valid {
			return false
		}
	}
	return true
}

func Day4Part1() {
	lines := helpers.ReadFileLines("./input/04.txt")
	passPorts := getPassportStrings(lines)
	count := 0
	for _, p := range passPorts {
		if validatePassportKeys(p) {
			count++
		}
	}
	fmt.Printf("%d valid passports found.\n", count)
}

func Day4Part2() {
	lines := helpers.ReadFileLines("./input/04.txt")
	passPorts := getPassportStrings(lines)
	validator := createPPValidator()
	count := 0
	for _, p := range passPorts {
		if validator.validatePassport(p) {
			count++
		}
	}
	fmt.Printf("valid passports found: %d / %d\n", count, len(passPorts))
}
