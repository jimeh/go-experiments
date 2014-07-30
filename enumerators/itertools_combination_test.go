package enumerators

import . "gopkg.in/check.v1"

type ItertoolsCombinationSuite struct {
	results [][]int
}

var _ = Suite(&ItertoolsCombinationSuite{})

/*
   Tests
*/

func (s *ItertoolsCombinationSuite) TestItertoolsCombination1(c *C) {
	results := ItertoolsCombination([]int{1, 2, 3, 4, 5}, 0)
	c.Assert(results, DeepEquals, [][]int{})
}

func (s *ItertoolsCombinationSuite) TestItertoolsCombination2(c *C) {
	results := ItertoolsCombination([]int{1, 2, 3, 4, 5}, 1)
	c.Assert(results, DeepEquals, [][]int{{1}, {2}, {3}, {4}, {5}})
}

func (s *ItertoolsCombinationSuite) TestItertoolsCombination3(c *C) {
	results := ItertoolsCombination([]int{10, 20, 30}, 2)
	c.Assert(results, DeepEquals, [][]int{{10, 20}, {10, 30}, {20, 30}})
}

func (s *ItertoolsCombinationSuite) TestItertoolsCombination4(c *C) {
	results := ItertoolsCombination([]int{1, 2, 3, 4, 5}, 3)
	c.Assert(results, DeepEquals, [][]int{
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

func (s *ItertoolsCombinationSuite) Benchmark(c *C) {
	var r [][]int
	for i := 0; i < c.N; i++ {
		r = ItertoolsCombination(
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			8,
		)
	}
	s.results = r
}
