package wordcount

import (
	"fmt"
	"strings"
)

// Frequency :
type Frequency map[string]int

// WordCount :
func WordCount(str string) Frequency {
	fmt.Println(str)
	//  var a *string = &str
	// strings.ReplaceAll(*a, "\t", "")
	// strings.ReplaceAll(*a, "\n", "")
	// strings.ToLower(*a)

	result := make(Frequency)

	st := strings.ToLower(str)
	replacer := strings.NewReplacer(",", " ", ".", "", ";", "", "&", "", "@", "", "$", "", "%", "", "^", "", "&", "", ":", "")
	s := replacer.Replace(st)
	words := strings.Fields(s)
	fmt.Println(words)
	for _, v := range words {
		result[v]++
	}
	fmt.Println(result)
	return result
}
