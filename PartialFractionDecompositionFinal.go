package main


import(

	"fmt"

)

type ItemInEquation interface {
	ItemType() string
}

type SVar struct {
	Multiplier complex128
	Exponent complex128

}

func(s SVar) ItemType() string{
	return "s"
}

type Constant struct {
	Value complex128
}

func(c Constant) ItemType() string{
	return "c"
}

type Parenthesis struct {
	Items []ItemInEquation
}

func(p Parenthesis) ItemType() string{
	return "p"
}

type Numerator struct {
	Items []ItemInEquation

}

type Denominator struct {
	Items []ItemInEquation
}

type Fraction struct {
	NumeratorData Numerator
	DenominatorData Denominator
}

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

	cTermNum := Constant{complex(5, 0)}

	sTermNum1 := SVar{complex(12, 0), complex(2, 0)}

	sTermNum2 := SVar{complex(-1, 0), complex(3, 0)}

	numeratorTerm := Numerator{[]ItemInEquation{cTermNum, sTermNum1, sTermNum2}}

	sTermDenomP1 := SVar{complex(1, 0), complex(2, 0)}

	pTermDenom1 := Parenthesis{[]ItemInEquation{sTermDenomP1}}


	sTermDenomP2 := SVar{complex(1, 0), complex(1, 0)}

	cTermDenomP2 := Constant{complex(-9, 0)}

	pTermDenom2 := Parenthesis{[]ItemInEquation{sTermDenomP2, cTermDenomP2}}
	
	sTermDenomP3 := SVar{complex(1, 0), complex(1, 0)}

	cTermDenomP3 := Constant{complex(-1, 0)}

	pTermDenom3 := Parenthesis{[]ItemInEquation{sTermDenomP3, cTermDenomP3}}


	denominatorTerm := Denominator{[]ItemInEquation{pTermDenom1, pTermDenom2, pTermDenom3}}

	fractionTerm := Fraction{numeratorTerm, denominatorTerm}

	CreateSystemOfEquationsForFractionToUndergoPartialFractionDecomposition(fractionTerm)

}

//for now it is assumed that only real number values are exponents 
//and also that there are no fractional exponents
//this will be the icing on the cake once that is added however...

//this function presumes the denominator contains only multiplying terms, a.k.a. parenthesis * parenthesis (no limit as to how many multiplying terms)
//note it presumes that every term is seperated by parenthesis
func CreateSystemOfEquationsForFractionToUndergoPartialFractionDecomposition(fractionInput Fraction) Matrix{

	leftSide := fractionInput.NumeratorData

	highestPowerLeftSide := GetHighestPowerOfSInSliceEquationItems(leftSide.Items)


	fmt.Println("HIGHEST POWER OF S", highestPowerLeftSide)

	//length is up to the highest power and the constant term is the + 1
	solutionSlice := make([]complex128, int(highestPowerLeftSide) + 1)

	for i := 0; i < len(solutionSlice); i++ {
		solutionSlice[i] = 0
	}

	for i := 0; i < len(leftSide.Items); i++ {

		//TODO PARENTHESIS ARE ALSO A VALID TYPE

		switch leftSide.Items[i].(type){
			case SVar:
				sVariable, ok := leftSide.Items[i].(SVar)
				if(ok){
					exponentVal := real(sVariable.Exponent)
					indexToPlugInto := int(float64(highestPowerLeftSide) - exponentVal)
					solutionSlice[indexToPlugInto] = sVariable.Multiplier
				}
			case Constant:
				constantVal, ok := leftSide.Items[i].(Constant)
				if(ok){
					solutionSlice[len(solutionSlice)-1] = constantVal.Value
				}
			default:
				panic("Unknown item type")
				
		}

	}

	fmt.Println(solutionSlice)

	denominatorTerms := fractionInput.Denominator 

	multiplyingTermCountDenominator := len(denominatorTerms.Items)


	// //these will be the top terms in partial fraction decomposition, IE: as+b, c, ds+e
	// //the values depend on the denominator beneath them
	numeratorsForSystem :=  [][]ItemInEquation

	for i := 0; i < len(denominatorTerms.Items); i++ {

		pItem, ok := denominatorTerms.Items[i].(Parenthesis)

		if(ok){
			highestPowerInParenthesis := GetHighestPowerOfSInSliceEquationItems(pItem.Items)

			if(highestPowerInParenthesis < 2){
				numeratorsForSystem = append(numeratorsForSystem, )
			}

		}

	}

	// //IMPORTANT THING TO NOTE... I DID NOT KNOW THIS BEFORE, BUT FOR PARTIAL FRACTION 
	// //DECOMPOSITION, IF YOU HAVE A LONE S^2 IN THE DENOMINATOR THE TOP VALUE IS A 
	// //IF YOU HAVE S^2 + 9 OR ANY CONSTANT THEN THE TOP VALUE IS AS + B 
	// //SO ONCE A SQUARED VALUE GETS ANOTHER TERM THEN AND ONLY THEN IS IT QUADRATIC
	// //NOTE.... THIS FUNCTION WILL ESNURE THAT ONLY POWER 2 FACTORS ARE USED... FOR NOW...

	

	for i := 0; i < len



	return Matrix{}

}


func SolveSystemOfEquations(matrixSystem Matrix) []complex128{
	
	//TODO THIS IS ASSOCIATED WITH THE TODO BELOW
	//lengthOfRow := len(matrixSystem.Rows[0])

	matrixSystemWithoutLastColumn := CleanCopyMatrix(matrixSystem)


	//TODO... YOU NEED TO FIX THIS LINE
	// matrixSystemWithoutLastColumn := Matrix{[][]complex128{matrixSystemWithoutLastColumn[0][0:lengthOfRow-1], matrixSystemWithoutLastColumn[1][0:lengthOfRow-1], matrixSystemWithoutLastColumn[2][0:lengthOfRow-1], matrixSystemWithoutLastColumn[3][0:lengthOfRow-1], matrixSystemWithoutLastColumn[4][0:lengthOfRow-1], matrixSystemWithoutLastColumn[5][0:lengthOfRow-1], matrixSystemWithoutLastColumn[6][0:lengthOfRow-1], matrixSystemWithoutLastColumn[7][0:lengthOfRow-1]}}

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

	return solutionsForVariables

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



func GetHighestPowerOfSInSliceEquationItems(items []ItemInEquation) int {

	highestPower := -1

	for i := 0; i < len(items); i++ {
		sVariable, ok := items[i].(SVar)

		if(ok){
			sVariablePower := int(real(sVariable.Exponent))
			if(sVariablePower > highestPower){
				highestPower = sVariablePower
			}
		}
	}

	return highestPower

}

