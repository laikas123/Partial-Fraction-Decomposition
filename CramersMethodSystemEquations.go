package main


import(

	"fmt"

)


type Matrix struct {

	Rows [][]complex128

}

type Equation struct {

    Items []complex128 `json:"eqtn"`

}

type Determinant struct {

	Row1 []complex128
	Row2 []complex128
	Row3 []complex128

}



func main() {
	
	
	equation1 := []complex128{12,   0,   0,  0,  1,   0,  0, 0,  7}
	equation2 := []complex128{6,  12,   0,  0,  3,  1,  0, 0,  12}
	equation3 := []complex128{11,  6, 12,  0,   2, 3, 1, 0,  15}
	equation4 := []complex128{2,  11,   6, 12, 9, 2, 3, 1, 14}
	equation5 := []complex128{1,   2, 11,  6,  9, 9, 2, 3, 1}
	equation6 := []complex128{0,  1,   2, 11,   0,  9, 9, 2, 2}
	equation7 := []complex128{0,   0,   1, 2,    0,   0, 9, 9, 1}
	equation8 := []complex128{0,   0,   0,  1,    0,   0,  0, 9, 9}
	
	matrixSystem := Matrix{[][]complex128{equation1, equation2, equation3, equation4, equation5, equation6, equation7, equation8}}

	lengthOfRow := len(equation1)

	matrixSystemWithoutLastColumn := Matrix{[][]complex128{equation1[0:lengthOfRow-1], equation2[0:lengthOfRow-1], equation3[0:lengthOfRow-1], equation4[0:lengthOfRow-1], equation5[0:lengthOfRow-1], equation6[0:lengthOfRow-1], equation7[0:lengthOfRow-1], equation8[0:lengthOfRow-1]}}

	solutionColumn := []complex128{}

	//append the last item in each row, as that is the solution column
	for i := 0; i < len(matrixSystem.Rows); i++ {
		solutionColumn = append(solutionColumn, matrixSystem.Rows[i][(len(matrixSystem.Rows[i]) - 1) ])
	}

	

	mainDeterminant := solveNbyNMatrixDeterminant(matrixSystemWithoutLastColumn, 1)

	// fmt.Println(mainDeterminant)

	mainSummation := complex(0, 0)

	for i := 0; i < len(mainDeterminant); i++ {
		mainSummation = mainSummation + mainDeterminant[i]
	}

	fmt.Println(mainSummation)

	solutionsForVariables := []complex128{}

	for i := 0; i < len(matrixSystem.Rows[0])-1; i++ {

		swappedColumnMatrixDeterminant := solveNbyNMatrixDeterminant(SwapSolutionColumnToSpecifiedColumn(matrixSystemWithoutLastColumn, i, solutionColumn), 1)

		innerSummation := complex(0, 0)

		for i := 0; i < len(swappedColumnMatrixDeterminant); i++ {
			innerSummation = innerSummation + swappedColumnMatrixDeterminant[i]
		}		

		solutionsForVariables = append(solutionsForVariables, innerSummation/mainSummation)
	}

	fmt.Println(solutionsForVariables)

}


func SwapSolutionColumnToSpecifiedColumn(matrixInput Matrix, specifiedColumn int, solutionColumn []complex128) Matrix {

	if(len(solutionColumn) !=  len(matrixInput.Rows)){
		panic("err length solution column != length matrix input SwapSolutionColumnToSpecifiedColumn()")
	}  

	cleanCopyToReturn := CleanCopyMatrix(matrixInput)

	for i := 0; i < len(cleanCopyToReturn.Rows); i++ {
		cleanCopyToReturn.Rows[i][specifiedColumn] = solutionColumn[i]		
	}

	return cleanCopyToReturn

}






func solveNbyNMatrixDeterminant(matrixInput Matrix, multiplier complex128) []complex128 {

	matrixData := matrixInput.Rows

	for j := 0; j < len(matrixData); j++ {
		// fmt.Println(matrixData[j])
	}

	

	returnValues := []complex128{}

	doNotPassForward := false

	if(len(matrixData[0]) != len(matrixData)){
		panic("error not square matrix")
	}

	if(len(matrixData[0]) < 3 || len(matrixData) < 3){
		panic("matrix smaller than length 3")
	}


	if(len(matrixData[0]) == 3 && len(matrixData) == 3){
		doNotPassForward = true
	}

	if(!doNotPassForward){


		yCursor := 0

		for i := 0; i < len(matrixData[0]); i++ {
			passForwardMultiplier := complex(1, 0)

			if(i % 2 == 0){
				passForwardMultiplier = complex(1, 0) 
			}else{
				passForwardMultiplier = complex(-1, 0) 
			}

			yCursor = 0

			passForwardMultiplier = passForwardMultiplier * matrixData[0][i] * multiplier


			matrixToPassForward := Matrix{[][]complex128{}}

			for y := 0; y < len(matrixData); y++ {
			
				if(y == 0){
					continue
				}else{
					matrixToPassForward.Rows = append(matrixToPassForward.Rows, []complex128{})
				}

			for x := 0; x < len(matrixData[0]); x++ {

				if(x == i){
					continue
				}

				matrixToPassForward.Rows[yCursor] = append(matrixToPassForward.Rows[yCursor], matrixData[y][x])

			}

				yCursor++ 
			}

				//fmt.Println("passforward", matrixToPassForward)

				returnValues = append(returnValues, solveNbyNMatrixDeterminant(matrixToPassForward, (passForwardMultiplier))...)						


			}

					

	}else{

		if(len(matrixData) != 3){
			panic("error with passforward matrix")
		}

		returnVal := multiplier * calculateDeterminantIndividual(Determinant{matrixData[0], matrixData[1], matrixData[2]})


		// fmt.Println("return val", returnVal)

		returnValues = append(returnValues, returnVal)
	}


	return returnValues
	



}





func calculateDeterminantIndividual(determinantInput Determinant) complex128{


	// fmt.Println("determinant received", determinantInput)


	row1 := determinantInput.Row1
	row2 := determinantInput.Row2
	row3 := determinantInput.Row3

	if(len(row1) != 3 || len(row2) != 3 || len(row3) != 3 ){
		panic("invalid determinant value")
	}


	iHat := row1[0]*((row2[1]*row3[2])-(row3[1]*row2[2]))

	jHat := (-1)*row1[1]*((row2[0]*row3[2])-(row3[0]*row2[2]))

	kHat := row1[2]*((row2[0]*row3[1])-(row3[0]*row2[1]))

	return iHat + jHat + kHat

}

func CleanCopyMatrix(matrixInput Matrix) Matrix {

	rowsToCopy := matrixInput.Rows

	copiedReturnRows := [][]complex128{}

	for i := 0; i < len(rowsToCopy); i++ {

		newRow := make([]complex128, len(rowsToCopy[i]))

		itemsCopied := copy(newRow, rowsToCopy[i])

		if(itemsCopied != len(rowsToCopy[i])){
			panic("error copying CleanCopyMatrix()")
		}

		copiedReturnRows = append(copiedReturnRows, newRow)

	}

	if(len(copiedReturnRows) != len(rowsToCopy)){
		panic("not all rows copied CleanCopyMatrix()")
	}

	return Matrix{copiedReturnRows}


}


func ReturnItemType(item ItemInEquation) string {

	switch item.ItemType() {
		case "s":
			return "s variable"
		case "c":
			return "constant"	
		default:
			panic("unknown item type")
	}

	return "error"

}