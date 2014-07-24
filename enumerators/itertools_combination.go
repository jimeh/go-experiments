package enumerators

// ItertoolsCombination is an alternative implementation of Combination, and
// is mostly a rip off of a combinations function that itself is based on
// Python's itertools: http://play.golang.org/p/JEgfXR2zSH
func ItertoolsCombination(pool []int, length int) chan []int {
	results := make(chan []int)

	go func() {
		size := len(pool)

		if length > size || length == 0 {
			close(results)
			return
		}

		indices := make([]int, length)
		for i := range indices {
			indices[i] = i
		}

		result := make([]int, length)
		copy(result, pool)

		results <- result

		lastResult := make([]int, length)
		copy(lastResult, result)

		for {
			result := make([]int, length)
			copy(result, lastResult)

			i := length - 1
			for i >= 0 && indices[i] == i+size-length {
				i--
			}

			if i < 0 {
				close(results)
				return
			}

			indices[i]++
			for j := i + 1; j < length; j++ {
				indices[j] = indices[j-1] + 1
			}

			for i < len(indices) {
				result[i] = pool[indices[i]]
				i++
			}

			results <- result
			copy(lastResult, result)
		}
	}()

	return results
}
