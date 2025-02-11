package utils

import "strings"

// used to clean strings of clutter
func CleanInput(text string) []string {
	clean_text := strings.TrimSpace(text)
	clean_text = strings.ToLower(clean_text)
	clean_text_arr := strings.Fields(clean_text)

	return clean_text_arr
}
