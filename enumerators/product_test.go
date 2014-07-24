package enumerators

import . "gopkg.in/check.v1"

type ProductSuite struct{}

var _ = Suite(&ProductSuite{})

/*
   Tests
*/

func (s *ProductSuite) TestProduct1(c *C) {
	results := Product([]int{1}, []int{10, 20})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{
		{1, 10},
		{1, 20},
	})
}

func (s *ProductSuite) TestProduct2(c *C) {
	results := Product([]int{1, 2}, []int{10, 20})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{
		{1, 10}, {1, 20},
		{2, 10}, {2, 20},
	})
}

func (s *ProductSuite) TestProduct3(c *C) {
	results := Product([]int{1}, []int{10, 20}, []int{100})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{
		{1, 10, 100},
		{1, 20, 100},
	})
}

func (s *ProductSuite) TestProduct4(c *C) {
	results := Product(
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

func (s *ProductSuite) TestProduct5(c *C) {
	results := Product([]int{1}, []int{}, []int{100, 200})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{})
}

func (s *ProductSuite) TestProduct6(c *C) {
	results := Product([]int{1}, []int{10}, []int{100})
	output := collectResults(results)

	c.Assert(output, DeepEquals, [][]int{{1, 10, 100}})
}

/*
   Benchmarks
*/

func (s *ProductSuite) Benchmark(c *C) {
	for i := 0; i < c.N; i++ {
		results := Product(
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{10, 20, 30, 40, 50},
			[]int{100, 200, 300},
			[]int{1000, 2000, 3000, 4000, 5000},
		)
		collectResults(results)
	}
}
