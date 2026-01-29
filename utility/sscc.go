package utility

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const ssccBodyLen = 17

var ssccPrefixRE = regexp.MustCompile(`^\d{7,12}$`)

// Sscc returns the 18-digit SSCC (17-digit body + check digit) without the AI "00".
func Sscc(code string) (out string, err error) {
	if len(code) != ssccBodyLen {
		return "", fmt.Errorf("wrong lenght code %s", code)
	}
	// Validate that code contains only digits
	for i, ch := range code {
		if ch < '0' || ch > '9' {
			return "", fmt.Errorf("invalid character '%c' at position %d", ch, i)
		}
	}
	sum := 0
	for i := range code {
		n := code[i] - '0'
		if i%2 == 0 {
			n *= 3
			sum += int(n)
		} else {
			sum += int(n)
		}
	}
	return fmt.Sprintf("%s%d", code, roundUp(sum)-sum), nil
}

func roundUp(val int) int {
	return 10 * ((val + 9) / 10)
}

// GenerateSSCC builds and returns the 20-digit string with AI "00" prefixed
// from a 7–12 digit prefix and non-negative sequence i.
func GenerateSSCC(i int, prefix string) (string, error) {
	if i < 0 {
		return "", fmt.Errorf("invalid i: must be non-negative")
	}
	// Must be 7–12 digits
	if !ssccPrefixRE.MatchString(prefix) {
		return "", fmt.Errorf("invalid SSCC prefix %q: must be 7–12 digits", prefix)
	}
	prefixLength := len(prefix)
	number := strconv.Itoa(i)
	if len(number) > ssccBodyLen-prefixLength {
		return "", fmt.Errorf("invalid i number: must be %d digits", ssccBodyLen-prefixLength)
	}
	padding := strings.Repeat("0", ssccBodyLen-prefixLength-len(number))
	code := prefix + padding + number
	sscc, err := Sscc(code)
	if err != nil {
		return "", fmt.Errorf("sscc returned error for code %s: %w", code, err)
	}
	return "00" + sscc, nil
}
