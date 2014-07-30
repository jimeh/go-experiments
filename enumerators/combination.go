package enumerators

// Combination returns a channel which will receive one by one all
// combinations of length num of elements from pool.
//
// Based on the C source code of Ruby's Array#combination method:
// http://www.ruby-doc.org/core-1.9.3/Array.html#method-i-combination
func Combination(pool []int, num int) [][]int {
	results := [][]int{}

	size := len(pool)

	if num <= 0 || size < num {
		return results
	} else if num == 1 {
		for _, item := range pool {
			results = append(results, []int{item})
		}
		return results
	} else {
		stack := make([]int, num+1)
		lastChosen := make([]int, num)
		lev := 0

		stack[0] = -1
		for {
			chosen := make([]int, num)
			copy(chosen, lastChosen)

			chosen[lev] = pool[stack[lev+1]]

			for lev++; lev < num; lev++ {
				stack[lev+1] = stack[lev] + 1
				chosen[lev] = pool[stack[lev+1]]
			}
			results = append(results, chosen)
			copy(lastChosen, chosen)

			for {
				if lev == 0 {
					return results
				}

				stack[lev]++
				lev--

				if stack[lev+1]+num != size+lev+1 {
					break
				}
			}
		}
	}
}
