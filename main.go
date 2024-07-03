package main

var (
	height int
	width  int
)

func main() {
	intro()
	height, width := readTwoNums()

	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}
	mode := chooseMode()
	if mode == 1 {
		height, width = randomMap()
	} else {
		height, width = customMap()
	}
}
