package main

import (
	// "math/rand"
	"fmt"
	// "math"
	// "time"
	"testing"
)





func TestMatrixOutcomes(t *testing.T){



	matrix1 := Matrix{[][]complex128{[]complex128{complex(-4, -1), complex(-4, -1), complex(-4, -1), complex(12, 3)}}, }

	// matrix1Old := CleanCopyMatrix(Matrix{matrix1})

	// matrix2 := Matrix{[][]complex128{[]complex128{complex(-2, 0), complex(-1, 0), complex(-4, 0), complex(870, 10)}}}


	// for i := 0; i < 1; i++ {

	// 		rand.Seed(time.Now().UnixNano())
	// 		//random number between rows 0 and 5
	// 		randomNumber := rand.Int()%(len(matrix1.Rows))

	// 		matrix1 = ZeroOutAllButSpecificIndex(randomNumber, matrix1)

	// 		matrix1 = GiveARandomRowSubstitutionRegardlessOfOutcome(matrix1)

	// 		for j := 0; j < len(matrix1.Rows); j++ {
	// 			matrix1 = ZeroOutAllButSpecificIndex(0, matrix1)
	// 			matrix1 = ScaleEachRowInMatrixByLowestNumberInRow(matrix1)
	// 		}

	// }

	// allPossible, _ := AllPossibleOutcomesTwoRows(matrix1[0], matrix2[0])

	// fmt.Println(AllPossibleOutcomesTwoRows(matrix1[0], matrix2[0]))

	// matrix1 = [][]complex128{allPossible[0]}

	fmt.Println(TwoMatricesEvaluateTheSame(matrix1,  []complex128{complex(-1, 0), complex(-1, 0), complex(-1, 0)}))


	if(false){
		t.Errorf("failure")
	}
}