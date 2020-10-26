package coinproblem

import (
	"math"
)

// Solve solves the coins problem with dynamic programming.
// Coins can be arbitrary. If the value can not be expressed
// by any sums of given coins, (-1) is returned.
// The complexity of the algorithm is O(nm) where n is the
// given value and the m is the number of coins that can be used.
func Solve(value int, coins []int) int {

	// catch impossible situation
	if value < 0 {
		return -1
	}

	// init memory for dyn. program.
	m := make([]float64, value+1)
	m[0] = 0

	// iterate over coins
	for v := 1; v <= value; v++ {

		m[v] = math.Inf(1)

		// for each coin
		for _, c := range coins {

			// after coin added, the sum will not be higher
			if v-c >= 0 {
				m[v] = math.Min(m[v], m[v-c]+1)
			}
		}
	}

	return int(m[value])
}
