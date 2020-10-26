package coinproblem

import (
	"fmt"
	"math"
)

// SolveOptimal finds the smallest number of coins with given sum.
// Coins can be arbitrary. If the value can not be expressed
// by any sums of given coins, (-1) is returned.
// The complexity of the algorithm is O(nm) where n is the
// given value and the m is the number of coins that can be used.
func SolveOptimal(value int, coins []int) int {

	// catch impossible situation
	if value < 0 {
		return -1
	}

	// init memory
	m := make([]float64, value+1)
	m[0] = 0

	// build memory
	for v := 1; v <= value; v++ {

		m[v] = math.Inf(1)

		// iterate over coins
		for _, c := range coins {

			// filter
			if v-c >= 0 {
				m[v] = math.Min(m[v], m[v-c]+1)
			}
		}
	}

	return int(m[value])
}

// SolveCounts finds the total number of ways to produce sum
// of the given value using the providied coins.  Coins can
// be arbitrary. The complexity of the algorithm is O(nm) where n is the
// given value and the m is the number of coins that can be used.
func SolveCounts(value int, coins []int) int {

	// init memory
	m := make([]int, value+1)
	m[0] = 1

	// build memory
	for v := 1; v <= value; v++ {

		// iterate over coins
		for _, c := range coins {

			// filter
			if v-c >= 0 {
				m[v] += m[v-c]
			}
		}
	}

	fmt.Println(m)
	return m[value]
}
