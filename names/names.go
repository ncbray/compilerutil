package names

import (
	"strings"
	"unicode"
)

func SplitSnakeCase(input string) []string {
	if input == "" {
		return []string{}
	}
	return strings.Split(input, "_")
}

func JoinSnakeCase(parts []string, all_caps bool) string {
	output := strings.Join(parts, "_")
	if all_caps {
		output = strings.ToUpper(output)
	} else {
		output = strings.ToLower(output)
	}
	return output
}

func SplitCamelCase(input string) []string {
	if input == "" {
		return []string{}
	}
	runes := []rune(input)
	output := []string{}
	i := 0
	for i < len(runes) {
		start := i
		i++
		if unicode.IsUpper(runes[start]) {
			if i < len(runes) && unicode.IsUpper(runes[i]) {
				// Upper case word.
				i++
				for i < len(runes) && unicode.IsUpper(runes[i]) {
					i++
				}
				// Is the last character the start of a lower Capitalized word?
				if i < len(runes) && unicode.IsLower(runes[i]) {
					i--
				}
			} else {
				// Capitalized word.
				for i < len(runes) && unicode.IsLower(runes[i]) {
					i++
				}
			}
		} else if unicode.IsLower(runes[start]) {
			// Lowercase word.
			for i < len(runes) && unicode.IsLower(runes[i]) {
				i++
			}
		} else if unicode.IsDigit(runes[start]) {
			// Number.
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				i++
			}
		} else {
			// Unknown
			for i < len(runes) && !unicode.IsLower(runes[i]) && !unicode.IsUpper(runes[i]) && !unicode.IsDigit(runes[i]) {
				i++
			}
		}
		output = append(output, string(runes[start:i]))
	}
	return output
}

func Capitalize(input string) string {
	if len(input) == 0 {
		return ""
	}
	runes := []rune(input)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func JoinCamelCase(parts []string, inital_caps bool) string {
	if len(parts) == 0 {
		return ""
	}

	output := ""
	if inital_caps {
		output = Capitalize(parts[0])
	} else {
		output = strings.ToLower(parts[0])
	}
	for i := 1; i < len(parts); i++ {
		output += Capitalize(parts[i])
	}
	return output
}
