package enumerators

import . "gopkg.in/check.v1"

type RecursiveCombinationSuite struct {
	results [][]int
}

var _ = Suite(&RecursiveCombinationSuite{})

/*
   Tests
*/

func (s *RecursiveCombinationSuite) TestRecursiveCombination1(c *C) {
	results := RecursiveCombination([]int{1, 2, 3, 4, 5}, 0)
	output := collectResults(results)
	c.Assert(output, DeepEquals, [][]int{})
}

func (s *RecursiveCombinationSuite) TestRecursiveCombination2(c *C) {
	results := RecursiveCombination([]int{1, 2, 3, 4, 5}, 1)
	output := collectResults(results)
	c.Assert(output, DeepEquals, [][]int{{1}, {2}, {3}, {4}, {5}})
}

func (s *RecursiveCombinationSuite) TestRecursiveCombination3(c *C) {
	results := RecursiveCombination([]int{1, 2, 3}, 2)
	output := collectResults(results)
	c.Assert(output, DeepEquals, [][]int{{1, 2}, {1, 3}, {2, 3}})
}

func (s *RecursiveCombinationSuite) TestRecursiveCombination4(c *C) {
	results := RecursiveCombination([]int{1, 2, 3, 4, 5}, 3)
	output := collectResults(results)
	c.Assert(output, DeepEquals, [][]int{
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 5},
		{1, 3, 4},
		{1, 3, 5},
		{1, 4, 5},
		{2, 3, 4},
		{2, 3, 5},
		{2, 4, 5},
		{3, 4, 5},
	})
}

/*
   Benchmarks
*/

func (s *RecursiveCombinationSuite) Benchmark(c *C) {
	var r [][]int
	for i := 0; i < c.N; i++ {
		results := RecursiveCombination(
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			8,
		)
		r = collectResults(results)
	}
	s.results = r
}
