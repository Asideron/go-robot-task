package main

func recFindPathsNumber(i, j int, calls *int) int {
	*calls++

	if i < 1 || j < 1 {
		return 0
	}
	if i == 1 || j == 1 {
		return 1
	}

	return recFindPathsNumber(i-1, j, calls) + recFindPathsNumber(i, j-1, calls)
}

func RecursiveSolution(n, m int) (steps int, calls int) {
	steps = recFindPathsNumber(n, m, &calls)
	return steps, calls
}
