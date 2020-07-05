package hamming

import "errors"

//A function that calculates hamming distance
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("string length not equal")
	}

	var count int = 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
