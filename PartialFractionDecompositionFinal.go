package main


import(

	"fmt"

)


type SVar struct {
	Multiplier complex128
	Exponent complex128

}

func(s SVar) ItemType() string{
	return "s"
}

type Parenthesis struct {
	Items []SVar
	Exponent complex128
}

func(p Parenthesis) ItemType() string{
	return "p"
}

type Numerator struct {
	ParenthesisTerms []Parenthesis

}

type Denominator struct {
	ParenthesisTerms []Parenthesis
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

	cTermNum := SVar{complex(5, 0), complex(0, 0)}

	sTermNum1 := SVar{complex(12, 0), complex(2, 0)}

	sTermNum2 := SVar{complex(-1, 0), complex(3, 0)}

	numeratorTerm := Numerator{[]Parenthesis{Parenthesis{[]SVar{cTermNum, sTermNum1, sTermNum2}, complex(1, 0)}}}


	//SUBTLE THING TO NOTICE
	//the way this is input is (s)^2 notice it is s inside the parenthesis squared
	//NOT (s^2) this is due to the nature of how the program performs partial fraction
	//decomposition
	sTermDenomP1 := SVar{complex(1, 0), complex(1, 0)}

	pTermDenom1 := Parenthesis{[]SVar{sTermDenomP1}, complex(2, 0)}


	sTermDenomP2 := SVar{complex(1, 0), complex(1, 0)}

	cTermDenomP2 := SVar{complex(-9, 0), complex(0, 0)}

	pTermDenom2 := Parenthesis{[]SVar{sTermDenomP2, cTermDenomP2}, complex(1, 0)}
	
	sTermDenomP3 := SVar{complex(1, 0), complex(1, 0)}

	cTermDenomP3 := SVar{complex(-1, 0), complex(0, 0)}

	pTermDenom3 := Parenthesis{[]SVar{sTermDenomP3, cTermDenomP3}, complex(1, 0)}


	denominatorTerm := Denominator{[]Parenthesis{pTermDenom1, pTermDenom2, pTermDenom3}}

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


	//assume only one term in numerator for now
	highestPowerLeftSide := GetHighestPowerOfSInSliceEquationItems(leftSide.ParenthesisTerms[0].Items)


	fmt.Println("HIGHEST POWER OF S", highestPowerLeftSide)

	//length is up to the highest power and the constant term is the + 1
	solutionSlice := make([]complex128, int(highestPowerLeftSide) + 1)

	for i := 0; i < len(solutionSlice); i++ {
		solutionSlice[i] = 0
	}

	for i := 0; i < len(leftSide.ParenthesisTerms[0].Items); i++ {

		exponentVal := real(leftSide.ParenthesisTerms[0].Items[i].Exponent)
		indexToPlugInto := int(float64(highestPowerLeftSide) - exponentVal)
		solutionSlice[indexToPlugInto] = leftSide.ParenthesisTerms[0].Items[i].Multiplier
				
		

	}

	fmt.Println("left side equation")
	fmt.Println(solutionSlice)

	denominatorTerms := fractionInput.DenominatorData

	// multiplyingTermCountDenominator := len(denominatorTerms.Items)


	// //these will be the top terms in partial fraction decomposition, IE: as+b, c, ds+e
	// //the values depend on the denominator beneath them
	numeratorsForSystem :=  [][]SVar{}

	denominatorsForSystem := []Parenthesis{}

	for i := 0; i <len(denominatorTerms.ParenthesisTerms); i++ {

		fmt.Printf("%#v\n", denominatorTerms.ParenthesisTerms[i])

	}

	

	for i := 0; i < len(denominatorTerms.ParenthesisTerms); i++ {

		pItem := denominatorTerms.ParenthesisTerms[i]

		highestPowerInParenthesis := GetHighestPowerOfSInSliceEquationItems(pItem.Items)

		fmt.Println("highest s power", highestPowerInParenthesis)

		parenthesisOuterPower := int(real(pItem.Exponent))

		fmt.Println("exponent per cycle", parenthesisOuterPower)

		for i := 0; i < parenthesisOuterPower; i++ {
			
			fmt.Println("cycle", i)
			//if less than 2 power plug in S^0
			if(highestPowerInParenthesis < 2){
				numeratorsForSystem = append(numeratorsForSystem, []SVar{SVar{complex(1, 0), complex(0, 0)}})
				denominatorsForSystem = append(denominatorsForSystem, FoilOutParenthesisToSomePower(pItem, (i+1)))
			
				fmt.Println("case 1")

			//if the highest power is 2 but it is a lone term, then same deal S^0
			}else if(highestPowerInParenthesis == 2 && len(pItem.Items) == 1){
				numeratorsForSystem = append(numeratorsForSystem, []SVar{SVar{complex(1, 0), complex(0, 0)}})	
				denominatorsForSystem = append(denominatorsForSystem, FoilOutParenthesisToSomePower(pItem, (i+1)))
	
				fmt.Println("case 2")			

			//if quadratic then plug in S^1 + S^0
			}else if(highestPowerInParenthesis == 2 && len(pItem.Items) == 1){
				numeratorsForSystem = append(numeratorsForSystem, []SVar{SVar{complex(1, 0), complex(1, 0)}, SVar{complex(1, 0), complex(0, 0)}})		
				denominatorsForSystem = append(denominatorsForSystem, FoilOutParenthesisToSomePower(pItem, (i+1)))
			
				fmt.Println("case 3")

			}
		}

		

	}

	if(len(numeratorsForSystem) != len(denominatorsForSystem)){
		panic("mismatch in number of numerators and denominators for partial fraction decomposition CreateSystemOfEquationsForFractionToUndergoPartialFractionDecomposition()")
	}

	fmt.Println("numerator and denominator")
	for i := 0; i < len(numeratorsForSystem); i++ {
		fmt.Println("Numerator")
		fmt.Printf("%#v\n", numeratorsForSystem[i])
		fmt.Println("Denominator")
		fmt.Printf("%#v\n", denominatorsForSystem[i])
	}



	


	// //IMPORTANT THING TO NOTE... I DID NOT KNOW THIS BEFORE, BUT FOR PARTIAL FRACTION 
	// //DECOMPOSITION, IF YOU HAVE A LONE S^2 IN THE DENOMINATOR THE TOP VALUE IS A 
	// //IF YOU HAVE S^2 + 9 OR ANY CONSTANT THEN THE TOP VALUE IS AS + B 
	// //SO ONCE A SQUARED VALUE GETS ANOTHER TERM THEN AND ONLY THEN IS IT QUADRATIC
	// //NOTE.... THIS FUNCTION WILL ESNURE THAT ONLY POWER 2 FACTORS ARE USED... FOR NOW...

	

	


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



func GetHighestPowerOfSInSliceEquationItems(items []SVar) int {

	highestPower := 0

	for i := 0; i < len(items); i++ {
		sVariable := items[i]

		sVariablePower := int(real(sVariable.Exponent))
		if(sVariablePower > highestPower){
			highestPower = sVariablePower
		}
		
	}

	return highestPower

}





func FoilOutParenthesisToSomePower(parentheisData Parenthesis, timesToFoil int) Parenthesis {


	if(timesToFoil == 1){
		return Parenthesis{parentheisData.Items, complex(1,0)}
	}

	parentItems := parentheisData.Items

	//for first cycle parent multiplies parent
	multiplyingItems := parentheisData.Items

	var result []SVar

	for timesToFoil > 1 {

		result = MultiplyTwoParenthesisData(parentItems, multiplyingItems)
	
		multiplyingItems = result

		timesToFoil--

	}

	return Parenthesis{result, complex(1, 0)}

}

func MultiplyTwoParenthesisData(items1 []SVar, items2 []SVar) []SVar {

	returnSlice := []SVar{}

	for i := 0; i < len(items1); i++ {

		currentItem := items1[i]

		for j := 0; j < len(items2); j++ {

			newItem := MultiplyTwoSVars(currentItem, items2[j])

			returnSlice = append(returnSlice, newItem)

		}

	}

	returnSlice = SimplifyLikeTermsInParenthesis(returnSlice)

	return returnSlice

}


func MultiplyTwoSVars(sVar1 SVar, sVar2 SVar) SVar {

	return SVar{(sVar1.Multiplier*sVar2.Multiplier),  (sVar1.Exponent + sVar2.Exponent)}

}


func SimplifyLikeTermsInParenthesis(items []SVar) []SVar {

	powerMap := make(map[complex128][]complex128) 

	for i := 0; i < len(items); i++ {
		if(powerMap[items[i].Exponent] == nil){
			powerMap[items[i].Exponent] = []complex128{items[i].Multiplier}
		}else{
			powerMap[items[i].Exponent] = append(powerMap[items[i].Exponent], items[i].Multiplier)
		}
	}


	returnSlice := []SVar{}

	for exponent, multipliers := range powerMap {

		multiplierSum := complex(0, 0)

		for i := 0; i < len(multipliers); i++ {
			multiplierSum = multiplierSum + multipliers[i]
		}


		returnSlice = append(returnSlice, SVar{multiplierSum, exponent})

	}

	return returnSlice


}



// func PrettyPrintNumerator(num Numerator) {

// 	SVarsString := ""

// 	for i := 0; i < len(num.Items); i++ {
// 		if(i == 0){
// 			SVarsString = SVarsString + num.Items[i]
// 		}else{

// 			SVarsString = SVarsString + " + (" + num.Items[i].Multiplier + ")* S ^ " + num.Items[i].Exponent + " "
// 		}
		
// 	}


// 	fmt.Println("Numerator", SVarsString)
// }

// func PrettyPrintDenominator(denom Denominator){

// 	SVarsString := ""

// 	for i := 0; i < len(denom.ParenthesisTerms); i++ {

// 		SVarsString = SVarsString + "( "

// 		for j := 0; j < len(denom.Items); j++ {
// 			if(j == 0){
// 				SVarsString = SVarsString + denom.Items[j].Multiplier + ")* S ^ " + denom.Items[j].Exponent + " "
// 			}else{
// 				SVarsString = SVarsString + " + (" + denom.Items[j].Multiplier + ")* S ^ " + denom.Items[j].Exponent + " "
// 			}
			
// 		}

// 		SVarsString = SVarsString + ") "
// 	}

// 	fmt.Println("Denominator", SVarsString)
		
// }

// func PrettyPrintAddedNumerator(sVarSlices [][]SVar) {

// 	SVarsString := ""

// 	for i := 0; i < len(sVarSlices); i++ {

// 		SVarsString = SVarsString + "( "

// 		currentSlice := sVarSlices[i]

// 		for j := 0; j < len(currentSlice); j++ {
// 			if(j == 0){
// 				SVarsString = SVarsString + "(" + fmt.Sprintf("%#v", currentSlice[j].Multiplier) + ")* S ^ " + fmt.Sprintf("%#v", currentSlice[j].Exponent) + " "
// 			}else{
// 				SVarsString = SVarsString + " + (" + fmt.Sprintf("%#v", currentSlice[j].Multiplier) + ")* S ^ " + fmt.Sprintf("%#v", currentSlice[j].Exponent) + " "
// 			}
			
			
// 		}

// 		SVarsString = SVarsString + ") "
// 	}

// 	fmt.Println("Numerator", SVarsString)
// }

// func PrettyPrintAddedDenominator(denom []Parenthesis) {
// 	SVarsString := ""


// 	for i := 0; i < len(denom); i++ {

// 		SVarsString = SVarsString + "( "

// 		currentSlice := denom[i].Items

// 		for j := 0; j < len(currentSlice); j++ {
// 			if(j == 0){
// 				SVarsString = SVarsString + "(" + fmt.Sprintf("%#v", currentSlice[j].Multiplier) + ")* S ^ " + fmt.Sprintf("%#v", currentSlice[j].Exponent) + " "
// 			}else{
// 				SVarsString = SVarsString + " + (" + fmt.Sprintf("%#v", currentSlice[j].Multiplier) + ")* S ^ " + fmt.Sprintf("%#v", currentSlice[j].Exponent) + " "
// 			}
			
			
// 		}

// 		SVarsString = SVarsString + ") "
// 	}

// 	fmt.Println("Denominator", SVarsString)
// }

func PrettyPrintParenthesis(p Parenthesis) {

}
func PrettyPrintSVar(s SVar) {

}









