package main


import (
	"fmt"
)

func main(){
	matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[i]); j++ {
        fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
    }
}

	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

for i, row := range matrix1 {
    for j, val := range row {
        fmt.Printf("matrix1[%d][%d] = %d\n", i, j, val)
    }
}
}
