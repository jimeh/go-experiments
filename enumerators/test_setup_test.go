package enumerators

import . "gopkg.in/check.v1"
import "testing"

// Test Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func collectResults(results chan []int) [][]int {
	output := [][]int{}
	for result := range results {
		if result != nil {
			output = append(output, result)
		}
	}
	return output
}
