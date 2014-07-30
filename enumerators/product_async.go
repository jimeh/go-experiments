package enumerators

// Product returns a channel which will receive one by one all combinations of
// ints from all given arrays/slices of ints.
//
// Based on the C source code of Ruby's Array#product method:
// http://www.ruby-doc.org/core-1.9.3/Array.html#method-i-product
func ProductAsync(pool ...[]int) chan []int {
	results := make(chan []int)

	go func() {
		defer close(results)

		size := len(pool)
		state := make([]int, size)

		for _, list := range pool {
			if len(list) == 0 {
				return
			}
		}

		for {
			result := make([]int, size)
			for i, index := range state {
				result[i] = pool[i][index]
			}
			results <- result

			point := size - 1
			state[point]++
			for state[point] == len(pool[point]) {
				state[point] = 0
				point--
				if point < 0 {
					return
				}
				state[point]++
			}
		}
	}()

	return results
}
