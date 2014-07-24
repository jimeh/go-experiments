package enumerators

func RecursiveCombination(pool []int, length int) chan []int {
	results := make(chan []int)

	go func() {
		size := len(pool)
		for i := 0; i < size; i++ {
			if length == 1 {
				results <- []int{pool[i]}
			} else {
				next := RecursiveCombination(pool[i+1:len(pool)], length-1)

				for result := range next {
					res := []int{pool[i]}
					if result != nil {
						for _, item := range result {
							res = append(res, item)
						}
					}
					results <- res
				}
			}
		}
		close(results)
	}()

	return results
}
