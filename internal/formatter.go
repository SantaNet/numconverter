package internal

import (
	"fmt"
	"regexp"
	"strings"
)

func autoFormat(input string) int {
	binaryPattern := `^[01]+$`
	octalPattern := `^[0-7]+$`
	hexPattern := `^[0-9A-Fa-f]+$`

	binary := regexp.MustCompile(binaryPattern)
	octal := regexp.MustCompile(octalPattern)
	hex := regexp.MustCompile(hexPattern)

	if binary.MatchString(input) {
		return 2
	}
	if octal.MatchString(input) {
		return 8
	}
	if hex.MatchString(input) {
		return 16
	}
	return 0
}

func GetFormat(input string) (string, int, error) {
	if strings.HasPrefix(input, "0x") {
		return strings.TrimPrefix(input, "0x"), 16, nil
	}
	if strings.HasPrefix(input, "0b") {
		return strings.TrimPrefix(input, "0b"), 2, nil
	}
	if strings.HasPrefix(input, "0o") {
		return strings.TrimPrefix(input, "0o"), 8, nil
	}
	if num := autoFormat(input); num != 0 {
		return input, num, nil
	}
	return "", 0, fmt.Errorf("unknown format (use 0x, 0b, or 0o)")
}

func GetWhatFormat(input string) (string, int) {
	parts := strings.Fields(input)
	numStr := parts[0]
	toBase := 10 // по умолчанию decimal

	if len(parts) > 2 && parts[1] == "to" {
		switch parts[2] {
		case "hex":
			toBase = 16
		case "bin":
			toBase = 2
		case "oct":
			toBase = 8
		case "dec":
			toBase = 10
		}
	}
	return numStr, toBase
}
