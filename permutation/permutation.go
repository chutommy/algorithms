package permutation

// Generate generates all permutations of the given array.
// The complexity of the algorithm is of O(n!), where n
// is the lenght of the given array.
func Generate(arr []int) [][]int {

	// generate permutations
	var perms [][]int
	gen(0, []int{}, arr, &perms)

	return perms
}

// gen generates all permutations for inp and stores it in out.
// Arguments k and set should have zero value since they are
// only for correct recursion's functionality purpose.
func gen(k int, set, inp []int, out *[][]int) {

	if k == len(inp) {

		// generate subset
		subset := make([]int, len(set))
		for pos, val := range set {
			subset[pos] = inp[val]
		}

		// store subset
		*out = append(*out, subset)

	} else {
		set = append(set, k)
		gen(k+1, set, inp, out)
		set = set[:len(set)-1]
		gen(k+1, set, inp, out)
	}
}
