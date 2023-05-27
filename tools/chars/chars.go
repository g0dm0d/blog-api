package chars

import (
	"regexp"
	"strings"
)

// ToLatin Func takes the text as input and returns the processed text as output. Where extra spaces and non-Latin characters are removed from the text
func ToLatin(str string) string {
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
	str = regexp.MustCompile("\\s+").ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, " ", "-")
	return str
}
