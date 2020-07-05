package isogram

import "strings"

// IsIsogram :
func IsIsogram(str string) bool {

	m := make(map[byte]uint)
	s := strings.ToLower(str)

	for i := range s {
		m[s[i]]++
		if m[s[i]] > 1 && string(s[i]) != "-" && string(s[i]) != " " {
			return false
		}
	}
	return true
}
