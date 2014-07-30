package enumerators

import . "gopkg.in/check.v1"

type ProductAsyncSuite struct{}

var _ = Suite(&ProductAsyncSuite{})

/*
   Tests
*/

func (s *ProductAsyncSuite) TestProductAsync1(c *C) {
	results := ProductAsync([]int{1}, []int{10, 20})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{
		{1, 10},
		{1, 20},
	})
}

func (s *ProductAsyncSuite) TestProductAsync2(c *C) {
	results := ProductAsync([]int{1, 2}, []int{10, 20})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{
		{1, 10}, {1, 20},
		{2, 10}, {2, 20},
	})
}

func (s *ProductAsyncSuite) TestProductAsync3(c *C) {
	results := ProductAsync([]int{1}, []int{10, 20}, []int{100})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{
		{1, 10, 100},
		{1, 20, 100},
	})
}

func (s *ProductAsyncSuite) TestProductAsync4(c *C) {
	results := ProductAsync(
		[]int{1},
		[]int{10, 20},
		[]int{100},
		[]int{1000, 2000},
	)

	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{
		{1, 10, 100, 1000},
		{1, 10, 100, 2000},
		{1, 20, 100, 1000},
		{1, 20, 100, 2000},
	})
}

func (s *ProductAsyncSuite) TestProductAsync5(c *C) {
	results := ProductAsync([]int{1}, []int{}, []int{100, 200})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{})
}

func (s *ProductAsyncSuite) TestProductAsync6(c *C) {
	results := ProductAsync([]int{1}, []int{10}, []int{100})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{{1, 10, 100}})
}

/*
   Benchmarks
*/

func (s *ProductAsyncSuite) Benchmark(c *C) {
	for i := 0; i < c.N; i++ {
		results := ProductAsync(
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{10, 20, 30, 40, 50},
			[]int{100, 200, 300},
		)
		collectResults(results)
	}
}
