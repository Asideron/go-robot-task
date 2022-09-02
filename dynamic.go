package main

func dynFindPathsNumber(i, j int, stepsInfo *[][]int, calls *int) int {
	*calls++

	if i < 1 || j < 1 {
		return 0
	}
	if i == 1 && j == 1 {
		return 1
	}

	if (*stepsInfo)[i][j] != 0 {
		return (*stepsInfo)[i][j]
	}

	(*stepsInfo)[i][j] = dynFindPathsNumber(i-1, j, stepsInfo, calls) + dynFindPathsNumber(i, j-1, stepsInfo, calls)
	return (*stepsInfo)[i][j]
}

func DynamicSolution(n, m int) (steps int, calls int) {
	matrix := make([][]int, n+1)
	for row := range matrix {
		matrix[row] = make([]int, m+1)
	}

	steps = dynFindPathsNumber(n, m, &matrix, &calls)
	return steps, calls
}
