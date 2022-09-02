package main

import "fmt"

type solutionFunc func(i, j int) (steps int, calls int)

func main() {
	run := func(fn solutionFunc, i, j int) {
		steps, calls := fn(i, j)
		fmt.Println("Steps: ", steps)
		fmt.Println("Calls: ", calls)
	}

	n, m := 10, 15

	fmt.Println("Dynamic solution")
	run(DynamicSolution, n, m)

	fmt.Println()

	fmt.Println("Recursive solution")
	run(RecursiveSolution, n, m)
}
