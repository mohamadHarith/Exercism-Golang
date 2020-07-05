package raindrops

import (
	"math"
	"strconv"
)

func Convert(input int) (output string) {

	if math.Mod(float64(input), 3) == 0 {
		output += "Pling"
	}
	if math.Mod(float64(input), 5) == 0 {
		output += "Plang"
	}
	if math.Mod(float64(input), 7) == 0 {
		output += "Plong"
	}
	if (math.Mod(float64(input), 3) != 0) && (math.Mod(float64(input), 5) != 0) && (math.Mod(float64(input), 7) != 0) {
		output = strconv.Itoa(input)
	}

	return

}
