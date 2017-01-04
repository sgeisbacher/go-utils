package utils

import (
	"regexp"
	"strings"
)

var REGEX_SPECIAL_CHARS = regexp.MustCompile("[^a-zA-Z0-9-]")

func ToID(str string) string {
	replacer := strings.NewReplacer(" ", "-", "ä", "ae", "Ä", "Ae", "ö", "oe", "Ö", "Oe", "ü", "ue", "Ü", "Ue", "ß", "ss")
	id := replacer.Replace(str)
	id = REGEX_SPECIAL_CHARS.ReplaceAllString(id, "")
	id = strings.Trim(id, "-")
	return strings.ToLower(id)
}
