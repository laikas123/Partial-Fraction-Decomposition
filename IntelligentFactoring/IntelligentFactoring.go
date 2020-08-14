package main

import (

	"fmt"
	"math/cmplx"
	"sort"
	"strings"
	"strconv"
)

type Container struct {
	Parent [][]complex128
	Children []Container
}


//TODO, when multiple variables get involved a third index needs to be added to the float slice
//which will allow the third index to essentially span the alphabet 0-25 A-Z for variable names
//for now since this is only being used for inverse laplace transform of one variable, everything is
//assumed to be 's'	



//BIG TODO, OBVIOUSLY AT SOME POINT A NUMBER WILL GET ZEROED OUT, IT'S IMPORTANT THAT IT DOESN'T GET TREATED AS A PARENTHESIS
//OTHERWISE THAT COULD MESS UP THE FLOW OF THINGS, BEGIN ADDING CHECKS FOR SUMMATIONS THAT RESULT IN 0

func main() {
	

	// equation := [][]complex128{gOP(), gNum(2, 3, 3, 0), gCP(3) }

	equation := [][]complex128{gOP(), gOP(), gNum(3,0), gCP(1), gOP(), gNum(1, 1, 2, 0), gCP(2), gNum(3, 0), gCP(1)}       

	fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))

	panic("test")




	// foiledEquation = FoilAllNeighboringParenthesis(foiledEquation)

	// fmt.Println(DecodeFloatSliceToEquation(foiledEquation))

	// foiledEquation = FoilAllNeighboringParenthesis(foiledEquation)

	// fmt.Println(DecodeFloatSliceToEquation(foiledEquation))

	// foiledEquation = FoilAllNeighboringParenthesis(foiledEquation)

	// fmt.Println(DecodeFloatSliceToEquation(foiledEquation))

	// foiledEquation = RemoveUnusedParenthesis(foiledEquation)

	// fmt.Println(DecodeFloatSliceToEquation(foiledEquation))


	// foiledEquation = FoilAllNeighboringParenthesis(foiledEquation)

	// fmt.Println(DecodeFloatSliceToEquation(foiledEquation))
	

}



func DecodeFloatSliceToEquation(equationInput [][]complex128 ) string {

//	CheckEquationForSyntaxErrors(equation)

	equation := CleanCopyEntire2DComplex128Slice(equationInput)
	
	equationString := ""

	previousTerm := []complex128{}

	for i := 0; i < len(equation); i++ {
		
		currentItem := equation[i]

		firstIndex := currentItem[0]
		secondIndex := currentItem[1]


		if(IsOP(firstIndex, secondIndex)){
			equationString += "( "
			
		}else if(IsCP(currentItem)){
			equationString += " )"
			if(currentItem[2] != 0 && currentItem[2] != 1) {
				equationString += "^" + strconv.FormatFloat(real(currentItem[2]), 'f', -1, 64) + " "
			}
			
		}else if(IsNumber(firstIndex)){

			if(IsCP(previousTerm)){
				equationString += "+"
			}

			for i := 0; i < len(currentItem); i = i+2{
				multiplier := currentItem[i]
				exponent := currentItem[i+1]

				if((i < len(currentItem) - 2) && exponent == 1){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64)  + "s " + " + "
				}else if((i < len(currentItem) - 2) && exponent == 0){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) + " + "
				}else if((i < len(currentItem) - 2) && exponent != 0 && exponent != 1){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) + "s^" + strconv.FormatFloat(real(exponent),'f', -1, 64) + " + "
				}else if((i == len(currentItem) - 2) && exponent == 1){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) + "s " 
				}else if((i == len(currentItem) - 2) && exponent == 0){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) 
				}else if((i ==  len(currentItem) - 2) && exponent != 0 && exponent != 1){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) + "s^" + strconv.FormatFloat(real(exponent),'f', -1, 64) 
				}
			}
			
		}else{
			panic("unknown equation item DecodeFloatSliceToEquation()")	
		}
	
		previousTerm = currentItem

	}

	

	return equationString

}



// func DecodeFloatSliceToEquation(equationInput [][]complex128 ) string {

// //	CheckEquationForSyntaxErrors(equation)

// 	equation := CleanCopyEntire2DComplex128Slice(equationInput)
	
// 	equationString := ""

// 	depthLevel := 0

// 	for i := 0; i < len(equation); i++ {
		
// 		currentItem := equation[i]

// 		for j := 0; j < len(currentItem); j = (j+2) {
// 			firstIndex := currentItem[j]
// 			secondIndex := currentItem[j+1]

// 			firstIndexString := strconv.FormatFloat(real(firstIndex), 'f', -1, 64)

// 			secondIndexString := strconv.FormatFloat(real(secondIndex), 'f', -1, 64)
			


// 			if(IsOP(firstIndex, secondIndex)){
// 				equationString += "( "
// 				depthLevel++ 
// 			}else if(IsCP(currentItem)){
// 				equationString += " )"
// 				if(currentItem[j+2] != 0 && currentItem[j+2] != 1){
// 					equationString += "^" + strconv.FormatFloat(real(currentItem[2]), 'f', -1, 64) + " "
// 				}
// 				depthLevel--
// 			}else if(IsNumber(firstIndex)){
// 				if(real(firstIndex) != 1){
// 					if(real(secondIndex) == 0){
// 						equationString += firstIndexString + " + "
// 					}else if(real(secondIndex) == 1){
// 						equationString += firstIndexString + "S + "
// 					}else{
// 						equationString += firstIndexString + "S^" + secondIndexString + " + "
// 					}
// 				}else{
// 					if(real(secondIndex) == 0){
// 						equationString += firstIndexString + " + "
// 					}else if(real(secondIndex) == 1){
// 						equationString += "S + "
// 					}else{
// 						equationString += "S^" + secondIndexString + " +"
// 					}
// 				}
				
// 			}else{
// 				panic("unknown equation item DecodeFloatSliceToEquation()")	
// 			}
// 			// fmt.Println(equationString)
// 			// fmt.Println(depthLevel)

// 			if(IsCP(currentItem)){
// 				break	
// 			}

// 		}

// 	}

// 	okToRemovePlus := false

// 	plusIndicesToRemove := []int{}

// 	trackCurrent := -1

// 	for i := 0; i < len(equationString); i++ {
// 		if(rune(equationString[i]) == ' '){
// 			continue
// 		}else if(rune(equationString[i]) == ')' && !okToRemovePlus){
// 			okToRemovePlus= true
// 		}else if(rune(equationString[i]) == '+'){
// 			if(okToRemovePlus){
// 				trackCurrent = i
// 			}
// 		}else if(rune(equationString[i]) == ')' && okToRemovePlus){
// 			if(okToRemovePlus){
// 				plusIndicesToRemove = append(plusIndicesToRemove, trackCurrent)
// 			}
// 			okToRemovePlus = false
// 		}else{
// 			okToRemovePlus = false
// 		}

// 	}

// 	for i := 0; i < len(plusIndicesToRemove); i++ {

// 		if(plusIndicesToRemove[i] < len(equationString) && plusIndicesToRemove[i] > 0 ){



// 			equationString = equationString[0:plusIndicesToRemove[i]] + equationString[(plusIndicesToRemove[i] + 1):len(equationString)]

// 			for j := 0; j < len(plusIndicesToRemove); j++ {
// 				plusIndicesToRemove[j]--
// 			}


// 		}
// 	}

// 	//strings.TrimSpace(equationString)


// 	return equationString

// }

//g stands for generate

func gNum(nums ...complex128) []complex128 {

	if( (len(nums)%2) != 0){
		panic("error, invalid amount of numbers")
	}

	returnSlice := []complex128{}

	for i := 0; i < len(nums); i = (i + 2) {



		returnSlice = append(returnSlice, nums[i])
		returnSlice = append(returnSlice, nums[i+1])


	}

	return returnSlice

}

func gOP() []complex128 {
	return []complex128{0, 0}
}

func gCP(exponent complex128) []complex128 {
	return []complex128{0, 1, exponent}
}




// finds "(" followed by numbers followed by ")^x" where x is some power greater than 1

func FoilOutParenthesisRaisedToExponent(equationInput [][]complex128) [][]complex128 {

	CheckEquationForSyntaxErrors(equationInput, "FoilOutParenthesisRaisedToExponent()")

	equation := CleanCopyEntire2DComplex128Slice(equationInput)

	numbersHolder := []complex128{}

	indexOpener := -1

	indexCloser := -1

	exponentCloser := -1

	foundValid := false

	for i := 0; i < len(equation); i ++ {

		if(foundValid){
			break
		}

		if(IsOP(equation[i][0], equation[i][1])){

			indexOpener = i

			fmt.Println("character")

			checkingIfValid := true

			sawOneInt := false

			cursor := i

			numbersHolder = []complex128{}
		

			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}

				if(!sawOneInt){
					if(IsNumber(equation[cursor][0])){
						numbersHolder = append(numbersHolder, equation[cursor]...)
						sawOneInt = true
					}else{
						//there was no integer after the opening parenthesis, not valid
						checkingIfValid = false
						break
					}
				}else if(sawOneInt){
					if(IsNumber(equation[cursor][0])){
						//there should only be one set of numbers inside parenthesis this should not be possible
						panic("interesting case, should not get here FoilOutParenthesisRaisedToExponent()")
						continue
					}else if(IsOP(equation[cursor][0], equation[cursor][1])){
						checkingIfValid = false
						break
					}else if IsCP(equation[cursor]){
						if(real(equation[cursor][2]) > 1){
							fmt.Println(equation[cursor:len(equation)])
							fmt.Println("getting here")
							indexCloser = cursor
							//TODO ALSO ADD FUNCTIONALITY FOR FRACTIONAL EXPONENTS
							exponentCloser = int(real(equation[cursor][2]))
							checkingIfValid = false
							foundValid = true
						}
						break
					}
				}

			}

		}


	}

	if(foundValid){

		fmt.Println("NUMBERS FOUND", numbersHolder)
		
		sliceToInsert := MultiplyParenthesisGivenExponent(numbersHolder, exponentCloser)

		slicesToInsert :=  [][]complex128{gOP(), sliceToInsert, gCP(1)}

		returnEquation := [][]complex128{}

		returnEquation = append(returnEquation, equation[0:indexOpener]...)
		
		returnEquation = append(returnEquation, slicesToInsert...)
		
		returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)
	
		returnEquation = CleanCopyEntire2DComplex128Slice(returnEquation)

		fmt.Println("RETURN EQUATION", strings.ReplaceAll(DecodeFloatSliceToEquation(returnEquation), " ", ""))		

		//recursive call if there was a change will check if more foils possible
		return FoilOutParenthesisRaisedToExponent(returnEquation)

	}else{

		//if no foils possible return input
		return equation
	}	


	




}



func MultiplyParenthesisGivenExponent(numbers []complex128, exponent int) []complex128{

	leftTerm := numbers
	rightTerm := numbers

	timesToFoil := exponent -1

	allNumbersFromFoil := []complex128{}

	
	for timesToFoil > 0 {

		fmt.Println("left term", leftTerm)
		fmt.Println("right term", rightTerm)
		fmt.Println("allNumbersFromFoil", allNumbersFromFoil)

		resultCurrentLoop := []complex128{}



		for i := 0; i < len(leftTerm); i = (i+2) {

			currentNumberMultiplier := leftTerm[i]

			currentNumberExponent := leftTerm[i+1]

			for j := 0; j < len(rightTerm); j = (j+2) {

				foilNumberMultiplier := rightTerm[j]
				foilNumberExponent := rightTerm[j+1]

				
				resultCurrentLoop = append(resultCurrentLoop, currentNumberMultiplier*foilNumberMultiplier)
				resultCurrentLoop = append(resultCurrentLoop, currentNumberExponent+foilNumberExponent)
			}
		}

		fmt.Println("result from current loop", resultCurrentLoop)

		fmt.Println("allNumbersFromFoil", allNumbersFromFoil)

		allNumbersFromFoil = resultCurrentLoop
		
		//this sets the left term equal to the result from the previous cycle
		leftTerm = resultCurrentLoop
		timesToFoil--
		

	}


	mapOfExponents := make(map[complex128][]complex128)

	for i := 0; i < len(allNumbersFromFoil); i = (i+2) {

		if(mapOfExponents[allNumbersFromFoil[i+1]] == nil){
			mapOfExponents[allNumbersFromFoil[i+1]] = []complex128{allNumbersFromFoil[i]} 
		}else{
			mapOfExponents[allNumbersFromFoil[i+1]] = append(mapOfExponents[allNumbersFromFoil[i+1]], allNumbersFromFoil[i])
		}

	}

	mapOfSimpliefiedMultipliersForExponents := make(map[complex128]complex128)

	for exponent, multipliers := range mapOfExponents {

		summationMultipliers := complex(0, 0)

		for i := 0; i < len(multipliers); i++ {
			summationMultipliers += multipliers[i]
		}

		mapOfSimpliefiedMultipliersForExponents[exponent] = summationMultipliers

	}


	sliceOfExponentsFloat := []float64{}
	sliceOfExponentsComplex := []complex128{}

	for exponent, _ := range mapOfSimpliefiedMultipliersForExponents {
		sliceOfExponentsFloat = append(sliceOfExponentsFloat, real(exponent))
		sliceOfExponentsComplex = append(sliceOfExponentsComplex, exponent)
	}


	copyOfExponentsFloat := make([]float64, len(sliceOfExponentsFloat))

	itemsCopied := copy(copyOfExponentsFloat, sliceOfExponentsFloat)

	if(itemsCopied != len(sliceOfExponentsFloat)){
		panic("invalid copy MultiplyParenthesisGivenExponent()")
	}else{

		sort.Sort(sort.Reverse(sort.Float64Slice((sliceOfExponentsFloat))))

		if(len(copyOfExponentsFloat) != len(sliceOfExponentsFloat)){
			panic("reverse changed the length of one of the float slices MultiplyParenthesisGivenExponent()")
		}

		newIndices := []int{}

		for i := 0; i < len(copyOfExponentsFloat); i++ {

			currentValue := copyOfExponentsFloat[i]

			for j := 0; j < len(sliceOfExponentsFloat); j++ {

				if(sliceOfExponentsFloat[j] == currentValue){
					newIndices = append(newIndices, j)
				}


			}

		}

		copyOfExponentsComplex := make([]complex128, len(sliceOfExponentsComplex))


		for i := 0; i < len(newIndices); i++ {
			copyOfExponentsComplex[newIndices[i]] = sliceOfExponentsComplex[i]
		}

		returnNumbers := []complex128{}

		for i := 0; i < len(copyOfExponentsComplex); i++ {
			returnNumbers = append(returnNumbers, mapOfSimpliefiedMultipliersForExponents[copyOfExponentsComplex[i]])
			returnNumbers = append(returnNumbers, copyOfExponentsComplex[i])
		}


		return returnNumbers

	}




} 



func FoilNeighborParenthesis(equationInput [][]complex128) [][]complex128 {

	CheckEquationForSyntaxErrors(equationInput, "FoilNeighborParenthesis()")

	equation := CleanCopyEntire2DComplex128Slice(equationInput)

	numbersHolderFirstNeighbor := []complex128{}
	numbersHolderSecondNeighbor := []complex128{}

	indexOpener := -1

	indexCloser := -1

	foundValid := false

	secondTerm := false

	for i := 0; i < len(equation); i ++ {

		if(foundValid){
			break
		}

		if(IsOP(equation[i][0], equation[i][1])){

			indexOpener = i

			checkingIfValid := true

			sawOneInt := false

			cursor := i

			numbersHolderFirstNeighbor = []complex128{}
		
			numbersHolderSecondNeighbor = []complex128{}


			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}

				if(!sawOneInt && !secondTerm){
					if(IsNumber(equation[cursor][0])){
						numbersHolderFirstNeighbor = append(numbersHolderFirstNeighbor, equation[cursor]...)
						sawOneInt = true
					}else{
						//there was no integer after the opening parenthesis, not valid
						checkingIfValid = false
						break
					}
				}else if(sawOneInt && !secondTerm){
					if(IsNumber(equation[cursor][0])){
						//there should only be one set of numbers inside parenthesis this should not be possible
						panic("interesting case, should not get here FoilNeighborParenthesis()")
						continue
					}else if(IsOP(equation[cursor][0], equation[cursor][1])){
						checkingIfValid = false
						break
					}else if IsCP(equation[cursor]){
						if(real(equation[cursor][2]) == 1){

							//TODO ALSO ADD FUNCTIONALITY FOR FRACTIONAL EXPONENTS
							// checkingIfValid = false
							if((cursor+1) >= len(equation)){
								break
							}else if(IsOP(equation[cursor+1][0], equation[cursor+1][1])){
								cursor++
								secondTerm = true
								sawOneInt = false
							}else{
								break
							}
							
							
						}
					}
				}else if(!sawOneInt && secondTerm){
					if(IsNumber(equation[cursor][0])){
						numbersHolderSecondNeighbor = append(numbersHolderSecondNeighbor, equation[cursor]...)
						sawOneInt = true
					}else{
						//there was no integer after the opening parenthesis, not valid
						checkingIfValid = false
						break
					}
				}else if(sawOneInt && secondTerm){
					if(IsNumber(equation[cursor][0])){
						//there should only be one set of numbers inside parenthesis this should not be possible
						panic("interesting case, should not get here FoilNeighborParenthesis()")
						continue
					}else if(IsOP(equation[cursor][0], equation[cursor][1])){
						checkingIfValid = false
						break
					}else if IsCP(equation[cursor]){
						if(real(equation[cursor][2]) == 1){
							indexCloser = cursor
							//TODO ALSO ADD FUNCTIONALITY FOR FRACTIONAL EXPONENTS
							checkingIfValid = false
							foundValid = true
						}
						break
					}
				}

			}

		}


	}

	if(foundValid){

		sliceToInsert := MultiplyNeighboringParenthesis(numbersHolderFirstNeighbor, numbersHolderSecondNeighbor)

		slicesToInsert :=  [][]complex128{gOP(), sliceToInsert, gCP(1)}

		returnEquation := [][]complex128{}

		returnEquation = append(returnEquation, equation[0:indexOpener]...)
		
		returnEquation = append(returnEquation, slicesToInsert...)
		
		returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)

		returnEquation = CleanCopyEntire2DComplex128Slice(returnEquation)

		fmt.Println("RETURN EQUATION", strings.ReplaceAll(DecodeFloatSliceToEquation(returnEquation), " ", ""))		

		//recursive call if there was a change will check if more foils possible
		return FoilNeighborParenthesis(returnEquation)

	}else{

		//if no foils possible return input
		return equation
	}	



}




func MultiplyNeighboringParenthesis(numbers1 []complex128, numbers2 []complex128) []complex128{

	leftTerm := numbers1
	rightTerm := numbers2


	allNumbersFromFoil := []complex128{}


	fmt.Println("left term", leftTerm)
	fmt.Println("right term", rightTerm)
	fmt.Println("allNumbersFromFoil", allNumbersFromFoil)

	resultCurrentLoop := []complex128{}



	for i := 0; i < len(leftTerm); i = (i+2) {

		currentNumberMultiplier := leftTerm[i]

		currentNumberExponent := leftTerm[i+1]

		for j := 0; j < len(rightTerm); j = (j+2) {

			foilNumberMultiplier := rightTerm[j]
			foilNumberExponent := rightTerm[j+1]

			
			resultCurrentLoop = append(resultCurrentLoop, currentNumberMultiplier*foilNumberMultiplier)
			resultCurrentLoop = append(resultCurrentLoop, currentNumberExponent+foilNumberExponent)
		}
	}

	fmt.Println("result from current loop", resultCurrentLoop)

	fmt.Println("allNumbersFromFoil", allNumbersFromFoil)

	allNumbersFromFoil = resultCurrentLoop
	
	
	mapOfExponents := make(map[complex128][]complex128)

	for i := 0; i < len(allNumbersFromFoil); i = (i+2) {

		if(mapOfExponents[allNumbersFromFoil[i+1]] == nil){
			mapOfExponents[allNumbersFromFoil[i+1]] = []complex128{allNumbersFromFoil[i]} 
		}else{
			mapOfExponents[allNumbersFromFoil[i+1]] = append(mapOfExponents[allNumbersFromFoil[i+1]], allNumbersFromFoil[i])
		}

	}

	mapOfSimpliefiedMultipliersForExponents := make(map[complex128]complex128)

	for exponent, multipliers := range mapOfExponents {

		summationMultipliers := complex(0, 0)

		for i := 0; i < len(multipliers); i++ {
			summationMultipliers += multipliers[i]
		}

		mapOfSimpliefiedMultipliersForExponents[exponent] = summationMultipliers

	}


	sliceOfExponentsFloat := []float64{}
	sliceOfExponentsComplex := []complex128{}

	for exponent, _ := range mapOfSimpliefiedMultipliersForExponents {
		sliceOfExponentsFloat = append(sliceOfExponentsFloat, real(exponent))
		sliceOfExponentsComplex = append(sliceOfExponentsComplex, exponent)
	}


	copyOfExponentsFloat := make([]float64, len(sliceOfExponentsFloat))

	itemsCopied := copy(copyOfExponentsFloat, sliceOfExponentsFloat)

	if(itemsCopied != len(sliceOfExponentsFloat)){
		panic("invalid copy MultiplyParenthesisGivenExponent()")
	}else{

		sort.Sort(sort.Reverse(sort.Float64Slice((sliceOfExponentsFloat))))

		if(len(copyOfExponentsFloat) != len(sliceOfExponentsFloat)){
			panic("reverse changed the length of one of the float slices MultiplyParenthesisGivenExponent()")
		}

		newIndices := []int{}

		for i := 0; i < len(copyOfExponentsFloat); i++ {

			currentValue := copyOfExponentsFloat[i]

			for j := 0; j < len(sliceOfExponentsFloat); j++ {

				if(sliceOfExponentsFloat[j] == currentValue){
					newIndices = append(newIndices, j)
				}


			}

		}

		copyOfExponentsComplex := make([]complex128, len(sliceOfExponentsComplex))


		for i := 0; i < len(newIndices); i++ {
			copyOfExponentsComplex[newIndices[i]] = sliceOfExponentsComplex[i]
		}

		returnNumbers := []complex128{}

		for i := 0; i < len(copyOfExponentsComplex); i++ {
			returnNumbers = append(returnNumbers, mapOfSimpliefiedMultipliersForExponents[copyOfExponentsComplex[i]])
			returnNumbers = append(returnNumbers, copyOfExponentsComplex[i])
		}


		return returnNumbers

	}




} 



func FactorQuadraticsWithABCAllPresent(equationInput [][]complex128)[][]complex128 {

	
	CheckEquationForSyntaxErrors(equationInput, "FactorQuadraticsWithABCAllPresent()")

	equation := CleanCopyEntire2DComplex128Slice(equationInput)

	numbersHolder := []complex128{}

	indexOpener := -1

	indexCloser := -1

	foundValid := false

	aTerm := complex(0, 0)

	bTerm := complex(0, 0)

	cTerm := complex(0, 0)

	for i := 0; i < len(equation); i ++ {

		if(foundValid){
			break
		}

		if(IsOP(equation[i][0], equation[i][1])){

			indexOpener = i

			checkingIfValid := true

			sawOneIntAndIsValidQuadratic := false

			cursor := i

			numbersHolder = []complex128{}
		

			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}

				if(!sawOneIntAndIsValidQuadratic){
					if(IsNumber(equation[cursor][0]) && len(equation[cursor]) == 6){
						numbersHolder = append(numbersHolder, equation[cursor]...)
						secondDegreeExponentPresent := false
						firstDegreeExponentPresent := false
						zeroDegreeExponentPresent := false

						for k := 0; k < len(equation[cursor]); k = (k+2) {

							if(real(equation[cursor][k+1]) == 2){
								secondDegreeExponentPresent = true
								aTerm = equation[cursor][k]
							}else if(real(equation[cursor][k+1]) == 1){
								firstDegreeExponentPresent = true
								bTerm = equation[cursor][k]
							}else if(real(equation[cursor][k+1]) == 0){
								zeroDegreeExponentPresent = true
								cTerm = equation[cursor][k]
							}

						}						

						if(secondDegreeExponentPresent && firstDegreeExponentPresent && zeroDegreeExponentPresent){
							sawOneIntAndIsValidQuadratic = true
						}else{
							break
						}

						
					}else{
						//there was no integer after the opening parenthesis, not valid
						checkingIfValid = false
						break
					}
				}else if(sawOneIntAndIsValidQuadratic){
					if(IsNumber(equation[cursor][0])){
						//there should only be one set of numbers inside parenthesis this should not be possible
						panic("interesting case, should not get here FactorQuadraticsWithABCAllPresent()")
						continue
					}else if(IsOP(equation[cursor][0], equation[cursor][1])){
						checkingIfValid = false
						break
					}else if IsCP(equation[cursor]){
						indexCloser = cursor
						//TODO ALSO ADD FUNCTIONALITY FOR FRACTIONAL EXPONENTS
						checkingIfValid = false
						foundValid = true
						break
					}
				}

			}

		}


	}

	if(foundValid){

		if(aTerm == 0 || bTerm == 0 || cTerm == 0 ){
			panic("too many or too few values for a value for quadratic formula")
		}else{

			root1, root2 := QuadraticFormula(aTerm, bTerm, cTerm)

			fmt.Println("root1", root1, "root2", root2)

			root1Slice := []complex128{complex(1, 0), complex(1,0), -1*root1, complex(0,0)}

			root2Slice := []complex128{complex(1, 0), complex(1,0), -1*root2, complex(0,0)}

			slicesToInsert :=  [][]complex128{gOP(), root1Slice, gCP(1), gOP(), root2Slice, gCP(1)}			

			returnEquation := [][]complex128{}

			returnEquation = append(returnEquation, equation[0:indexOpener]...)
			
			returnEquation = append(returnEquation, slicesToInsert...)
			
			returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)
		
			returnEquation = CleanCopyEntire2DComplex128Slice(returnEquation)

			fmt.Println("RETURN EQUATION", strings.ReplaceAll(DecodeFloatSliceToEquation(returnEquation), " ", ""))		

			//recursive call if there was a change will check if more foils possible
			return FactorQuadraticsWithABCAllPresent(returnEquation)
		}
	}else{

		//if no foils quadratic factors return input
		return equation
	}	



}



func FactorQuadraticsWithABOnlyPresent(equationInput [][]complex128)[][]complex128 {

	
	CheckEquationForSyntaxErrors(equationInput, "FactorQuadraticsWithABOnlyPresent()")

	equation := CleanCopyEntire2DComplex128Slice(equationInput)

	numbersHolder := []complex128{}

	indexOpener := -1

	indexCloser := -1

	foundValid := false

	aTerm := complex(0, 0)

	bTerm := complex(0, 0)


	for i := 0; i < len(equation); i ++ {

		if(foundValid){
			break
		}

		if(IsOP(equation[i][0], equation[i][1])){

			indexOpener = i

			checkingIfValid := true

			sawOneIntAndIsValidQuadratic := false

			cursor := i

			numbersHolder = []complex128{}
		

			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}

				if(!sawOneIntAndIsValidQuadratic){
					if(IsNumber(equation[cursor][0]) && len(equation[cursor]) == 4){
						numbersHolder = append(numbersHolder, equation[cursor]...)
						secondDegreeExponentPresent := false
						firstDegreeExponentPresent := false
						
						for k := 0; k < len(equation[cursor]); k = (k+2) {

							if(real(equation[cursor][k+1]) == 2){
								secondDegreeExponentPresent = true
								aTerm = equation[cursor][k]
							}else if(real(equation[cursor][k+1]) == 1){
								firstDegreeExponentPresent = true
								bTerm = equation[cursor][k]
							}

						}						

						if(secondDegreeExponentPresent && firstDegreeExponentPresent){
							sawOneIntAndIsValidQuadratic = true
						}else{
							break
						}

						
					}else{
						//there was no integer after the opening parenthesis, not valid
						checkingIfValid = false
						break
					}
				}else if(sawOneIntAndIsValidQuadratic){
					if(IsNumber(equation[cursor][0])){
						//there should only be one set of numbers inside parenthesis this should not be possible
						panic("interesting case, should not get here FactorQuadraticsWithABOnlyPresent()")
						continue
					}else if(IsOP(equation[cursor][0], equation[cursor][1])){
						checkingIfValid = false
						break
					}else if IsCP(equation[cursor]){
						indexCloser = cursor
						//TODO ALSO ADD FUNCTIONALITY FOR FRACTIONAL EXPONENTS
						checkingIfValid = false
						foundValid = true
						break
					}
				}

			}

		}


	}

	if(foundValid){

		if(aTerm == 0 || bTerm == 0 ){
			panic("too many or too few values for a value for quadratic formula")
		}else{

			//scale so A is 1

			aTermBeforeScale := aTerm

			aTerm := aTerm/aTermBeforeScale

			bTerm := bTerm/aTermBeforeScale

			completingTheSquareTerm := cmplx.Pow((bTerm/2), 2)

			cTerm := completingTheSquareTerm

			root1, root2 := QuadraticFormula(aTerm, bTerm, cTerm)

			fmt.Println("root1", root1, "root2", root2)

			scaleDownASlice := []complex128{aTermBeforeScale, complex(0,0)}

			root1Slice := []complex128{complex(1, 0), complex(1,0), -1*root1, complex(0,0)}

			root2Slice := []complex128{complex(1, 0), complex(1,0), -1*root2, complex(0,0)}

			completingTheSquareTermSlice := []complex128{(-1*cTerm*aTermBeforeScale), complex(0, 0)}

			slicesToInsert :=  [][]complex128{gOP(), gOP(), scaleDownASlice, gCP(1), gOP(), root1Slice, gCP(1), gOP(), root2Slice, gCP(1), completingTheSquareTermSlice, gCP(1)}			

			returnEquation := [][]complex128{}

			returnEquation = append(returnEquation, equation[0:indexOpener]...)
			
			returnEquation = append(returnEquation, slicesToInsert...)
			
			returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)
		
			returnEquation = CleanCopyEntire2DComplex128Slice(returnEquation)

			fmt.Println("RETURN EQUATION", strings.ReplaceAll(DecodeFloatSliceToEquation(returnEquation), " ", ""))		

			//recursive call if there was a change will check if more foils possible
			return FactorQuadraticsWithABOnlyPresent(returnEquation)
		}
	}else{

		//if no foils quadratic factors return input
		return equation
	}	



}




func FactorQuadraticsWithACOnlyPresent(equationInput [][]complex128)[][]complex128 {

	
	CheckEquationForSyntaxErrors(equationInput, "FactorQuadraticsWithACOnlyPresent()")

	equation := CleanCopyEntire2DComplex128Slice(equationInput)

	numbersHolder := []complex128{}

	indexOpener := -1

	indexCloser := -1

	foundValid := false

	aTerm := complex(0, 0)

	cTerm := complex(0, 0)


	for i := 0; i < len(equation); i ++ {

		if(foundValid){
			break
		}

		if(IsOP(equation[i][0], equation[i][1])){

			indexOpener = i

			checkingIfValid := true

			sawOneIntAndIsValidQuadratic := false

			cursor := i

			numbersHolder = []complex128{}
		

			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}

				if(!sawOneIntAndIsValidQuadratic){
					if(IsNumber(equation[cursor][0]) && len(equation[cursor]) == 4){
						numbersHolder = append(numbersHolder, equation[cursor]...)
						secondDegreeExponentPresent := false
						zeroDegreeExponentPresent := false
						

						for k := 0; k < len(equation[cursor]); k = (k+2) {

							if(real(equation[cursor][k+1]) == 2){
								secondDegreeExponentPresent = true
								aTerm = equation[cursor][k]
							}else if(real(equation[cursor][k+1]) == 0){
								zeroDegreeExponentPresent = true
								cTerm = equation[cursor][k]
							}

						}						

						if(secondDegreeExponentPresent && zeroDegreeExponentPresent){
							sawOneIntAndIsValidQuadratic = true
						}else{
							break
						}

						
					}else{
						//there was no integer after the opening parenthesis, not valid
						checkingIfValid = false
						break
					}
				}else if(sawOneIntAndIsValidQuadratic){
					if(IsNumber(equation[cursor][0])){
						//there should only be one set of numbers inside parenthesis this should not be possible
						panic("interesting case, should not get here FactorQuadraticsWithACOnlyPresent()")
						continue
					}else if(IsOP(equation[cursor][0], equation[cursor][1])){
						checkingIfValid = false
						break
					}else if IsCP(equation[cursor]){
						indexCloser = cursor
						//TODO ALSO ADD FUNCTIONALITY FOR FRACTIONAL EXPONENTS
						checkingIfValid = false
						foundValid = true
						break
					}
				}

			}

		}


	}

	if(foundValid){

		if(aTerm == 0 || cTerm == 0 ){
			panic("too many or too few values for a value for quadratic formula")
		}else{

			//scale so A is 1

			aTermBeforeScale := aTerm

			aTerm = aTerm/aTermBeforeScale

			cTerm = cTerm/aTermBeforeScale

			fmt.Println("c term", cTerm)

			scaleDownASlice := []complex128{aTermBeforeScale, complex(0,0)}

			root1Slice := []complex128{complex(1, 0), complex(1,0), cmplx.Pow(cTerm, complex(0.5, 0)), complex(0, 0)}

			root2Slice := []complex128{complex(1, 0), complex(1,0), (-1* cmplx.Pow(cTerm, complex(0.5, 0))), complex(0, 0)}

			slicesToInsert :=  [][]complex128{gOP(), scaleDownASlice, gCP(1), gOP(), root1Slice, gCP(1), gOP(), root2Slice, gCP(1)}			

			returnEquation := [][]complex128{}

			returnEquation = append(returnEquation, equation[0:indexOpener]...)
			
			returnEquation = append(returnEquation, slicesToInsert...)
			
			returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)
		
			returnEquation = CleanCopyEntire2DComplex128Slice(returnEquation)

			//recursive call if there was a change will check if more foils possible
			return FactorQuadraticsWithACOnlyPresent(returnEquation)
		}
	}else{

		//if no foils quadratic factors return input
		return equation
	}	



}




func IsOP(num1 complex128, num2 complex128) bool {
	if(real(num1) == 0 && real(num2) == 0){
		return true
	}else{
		return false
	}
}
func IsCP(nums []complex128) bool {
	if(len(nums) == 3){
		return true
	}else{
		return false
	}
}
func IsNumber(num1 complex128) bool {
	if(num1 != 0){
		return true
	}else{
		return false
	}
}


func CheckEquationForSyntaxErrors(equation [][]complex128, parentFunction string) {

	fmt.Println("Parent Function", parentFunction)

	depthLevel := 0

	fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))	

	for i := 0; i < len(equation); i++ {

		currentItem := equation[i]

		for j := 0; j < len(currentItem); j = (j+2) {

			firstIndex := currentItem[j]
			secondIndex := currentItem[j+1]



			if(i == 0 && !IsOP(firstIndex, secondIndex)){
				panic("Syntax Error first item must be ( CheckEquationForSyntaxErrors()")
			}

			if(IsOP(firstIndex, secondIndex)){
				depthLevel++ 
			}else if(IsCP(currentItem)){
				depthLevel--
			}else if(IsNumber(firstIndex)){

			}else{
				fmt.Println(currentItem)
				panic("Syntax Error unknown item type CheckEquationForSyntaxErrors()")
			}

			//this would occur if there's more closing parenthesis than opening
			if(depthLevel == -1){
				panic("Syntax Error too many ) for the number of ( CheckEquationForSyntaxErrors()")
			}

			if(len(currentItem) == 3){
				break	
			}
		}
	}

	if(depthLevel != 0){
		panic("Syntax Error not all ( items were closed properly CheckEquationForSyntaxErrors()")
	}

}


func RemoveUnusedParenthesis(equationInput [][]complex128) [][]complex128 {


	equation := CleanCopyEntire2DComplex128Slice(equationInput)


	equation = RemoveLastItemIfItIsOpeningParenthesis(equation)

	equation = RemoveExcessParenthesisViaDepthCheck(equation)

	if(!(TwoEquationsAreExactlyIdentical(equationInput, equation))){
		return RemoveUnusedParenthesis(equation)
	}else{
		return equation
	}

	
	




}






func QuadraticFormula(a complex128, b complex128, c complex128) (complex128, complex128) {

	squareRoot := cmplx.Pow((cmplx.Pow(b, 2)) - (4*a*c), 0.5)

	fmt.Println(a, b, c, "abc")

	denominator := 2 * a

	positiveSquareRootResult := ((-1*b) + squareRoot)/(denominator)
	negativeSquareRootResult := ((-1*b) - squareRoot)/(denominator)

	return positiveSquareRootResult, negativeSquareRootResult


}


func CleanCopyEntire2DComplex128Slice(equationToCopy [][]complex128) [][]complex128 {

	returnEquation := [][]complex128{}

	for i := 0; i < len(equationToCopy); i++ {

		cleanCopyToAppend := make([]complex128, len(equationToCopy[i]))

		itemsCopied := copy(cleanCopyToAppend, equationToCopy[i])

		if(itemsCopied != len(equationToCopy[i])){
			panic("invalid copy CleanCopyEntire2DComplex128Slice()")
		}else{
			returnEquation = append(returnEquation, cleanCopyToAppend)
		}


	}

	return returnEquation

}


//returns the number of direct children for the opening parenthesis, 
//where a direct child is a set of closing parenthesis where the parenthesis are 
//one level of depth deeper...
//for instance... ( ( (3S+2)(3S + 5)(20S) ) )
//the outermost parenthesis have one child only since there's only one set of closing parenthesis
//the second level in however has 3 direct children...
//this function focuses on places where there is a double set of parenthesis that can be removed
//
//the other int returned is the index of the closing parenthesis for this opening parent parenthesis
func GetDirectChildCountOfParenthesisCurrentOpeningParenthesis(startIndex int, equation [][]complex128) (int, int) {

	depth := 0

	canCountOpener := true
	canCountCloser := false

	openerCount := 0
	closerCount := 0

	indexOfCloserToParent := -1

	//start query at 1 past the start index since the start index is the parent
	for i := 1; i < len(equation); i++ {

		firstIndex := equation[i][0]
		secondIndex := equation[i][1]

		if(IsOP(firstIndex, secondIndex)){
			depth++
			if(depth == 1 && canCountOpener){
				openerCount++ 
				canCountOpener = false
				canCountCloser = true
			}
		}else if(IsCP(equation[i])){
			depth--
			if(depth == 0 && canCountCloser){
				closerCount++ 
				canCountOpener = true
				canCountCloser = false
			}else if(depth == -1){
				indexOfCloserToParent = i
			}
		}

	}


	if(openerCount != closerCount){
		fmt.Println("opener count", openerCount, "closer count ", closerCount)
		panic("syntax error not all parenthesis closed GetDirectChildCountOfParenthesisCurrentOpeningParenthesis")
	}else{
		return openerCount, indexOfCloserToParent
	}




}



func RemoveLastItemIfItIsOpeningParenthesis(equation [][]complex128) [][]complex128{

	cleanCopyToReturn := CleanCopyEntire2DComplex128Slice(equation)


	for (IsOP(cleanCopyToReturn[len(cleanCopyToReturn)-1][0], cleanCopyToReturn[len(cleanCopyToReturn)-1][1])) {
		cleanCopyToReturn = cleanCopyToReturn[0:(len(cleanCopyToReturn)-1)]
		
	}

	return cleanCopyToReturn


}



func RemoveExcessParenthesisViaDepthCheck(equationInput [][]complex128) [][]complex128 {

	equation := CleanCopyEntire2DComplex128Slice(equationInput)

	depthLevel := 0


	for i := 0; i < len(equation); i++ {

		currentItem := equation[i]

		for j := 0; j < len(currentItem); j = (j+2) {

			firstIndex := currentItem[j]
			secondIndex := currentItem[j+1]



			if(i == 0 && !IsOP(firstIndex, secondIndex)){
				panic("Syntax Error first item must be ( RemoveUnusedParenthesis()")
			}

			if(IsOP(firstIndex, secondIndex)){
				depthLevel++ 
			}else if(IsCP(currentItem)){
				depthLevel--

				if(depthLevel == -1){
					equation = append(equation[0:i], equation[(i+1):len(equation)]...) 
					return equation
				}
				break	
			}else if(IsNumber(firstIndex)){

			}else{
				fmt.Println(currentItem)
				panic("Syntax Error unknown item type RemoveUnusedParenthesis()")
			}

		}
	
	}

		

	return equation

}



func TwoEquationsAreExactlyIdentical(equation1 [][]complex128, equation2 [][]complex128) bool {

	if(len(equation1) != len(equation2)){
		return false
	}

	

	for i := 0; i < len(equation1); i++ {
		currentItemEquation1 := equation1[i]
		currentItemEquation2 := equation2[i]

		if(len(currentItemEquation1) != len(currentItemEquation2)){
			return false
		}

		for j := 0; j < len(currentItemEquation1); j = j+2  {


			firstIndexEquation1 := currentItemEquation1[j]
			firstIndexEquation2 := currentItemEquation2[j]

			if(firstIndexEquation1 != firstIndexEquation2){
				return false
			}

			secondIndexEquation1 := currentItemEquation1[j+1]
			secondIndexEquation2 := currentItemEquation2[j+1]

			if(secondIndexEquation1 != secondIndexEquation2){
				return false
			}

			if(len(currentItemEquation1) == 3){
				
				if(currentItemEquation1[j+2] != currentItemEquation2[j+2]){
					return false
				}else{
					break
				}
			}

		}


	}

	return true

}


func RemoveParenthesisWith0DirectChildren(equationInput [][]complex128) [][]complex128 {

	equation := CleanCopyEntire2DComplex128Slice(equationInput)

	breakAll := false

	for i := 0; i < len(equation); i++ {

		if(breakAll){
			break
		}

		currentItem := equation[i]

		for j := 0; j < len(currentItem); j = (j+2) {

			firstIndex := currentItem[j]
			secondIndex := currentItem[j+1]



			if(i == 0 && !IsOP(firstIndex, secondIndex)){
				panic("Syntax Error first item must be ( RemoveParenthesisWith0DirectChildren()")
			}

			if(IsOP(firstIndex, secondIndex)){
				if(i+1 >= len(equation)){
					return equation
				}else if(IsCP(equation[i+1])){
					//this is if the last 2 items are "( )"
					if((i + 2) >= len(equation)){
						equation = equation[0:i]
					}else{
						equation = append(equation[0:i], equation[(i+2):len(equation)]...)
						breakAll = true
						break
					}
				}	
			}

			if(IsCP(currentItem)){
				break
			}
		
		}
	
	}	

	if(!TwoEquationsAreExactlyIdentical(equation, equationInput)){
		return  RemoveParenthesisWith0DirectChildren(equation)
	}else{
		return equation
	}
	


}


func CreateATreeFromCurrentEquation(equation [][]complex128) []Container {

	// CheckEquationForSyntaxErrors(equation, "CreateATreeFromCurrentEquation()")



	breakAll := false
	

	

	equations := []Container{}

	for i := 0; i < len(equation); i ++ {



		if(breakAll){
			break
		}

		currentItem := equation[i]

		if(IsOP(currentItem[0], currentItem[1])){

			depth := 1

			foundCP := false

			cursor := i

			openerIndex := i

			for !foundCP {

				cursor++


				if(IsCP(equation[cursor])){
					depth--
				}else if(IsOP(equation[cursor][0], equation[cursor][1])){
					depth++
				}

				if(depth == 0){
					i = cursor
					cleanCopyToAppend := CleanCopyEntire2DComplex128Slice(equation)
					dataToAddToTree := cleanCopyToAppend[openerIndex+1:cursor]
					sliceForChildren := []Container{}
					equations = append(equations, Container{dataToAddToTree, sliceForChildren})
					
					break
				}

				if((cursor + 1) >= len(equation)){
					breakAll = true
					break
				}


			}

		}
		

	}

	

	return equations



}




func CreateEntireTreeForEquation(equation [][]complex128) []Container{

	currentContainer := CreateATreeFromCurrentEquation(equation)

	previousParentContainers := [][]Container{}

	layer := 0

	xPosForLayer := []int{0}

	

	for layer > -1 {

		fmt.Println("layer", layer)
		fmt.Println("xpos for layer", xPosForLayer[layer])
		fmt.Println("previousParentContainers", previousParentContainers)
		fmt.Println("lengthpreviousParentContainers", len(previousParentContainers))

		for(xPosForLayer[layer] >= len(currentContainer)){
			xPosForLayer[layer] = 0
			layer--
			if(layer == -1){
				break
			}
			if(len(previousParentContainers) != 1){
				currentContainer = previousParentContainers[len(previousParentContainers)-2]
			}
			previousParentContainers = previousParentContainers[0:len(previousParentContainers)-1]
			xPosForLayer = xPosForLayer[0:len(xPosForLayer)-1]
			
			
			fmt.Println("was here1")
		}

		fmt.Println("was here2")
		newContainers := CreateATreeFromCurrentEquation(currentContainer[xPosForLayer[layer]].Parent)

		if(len(newContainers) != 0){
			currentContainer[xPosForLayer[layer]].Children = newContainers
			xPosForLayer[layer]++
			previousParentContainers = append(previousParentContainers, currentContainer)
			xPosForLayer = append(xPosForLayer, 0)
			layer++
			currentContainer = currentContainer[xPosForLayer[layer]].Children
		}else{
			xPosForLayer[layer]++
		}

	}


	return currentContainer	

}


// func CreateEntireTreeForEquation(equation [][]complex128) []Container {

// 	treeSlice := []Container{Container{}}

// 	treeSliceData, _ := CreateATreeFromCurrentEquation(equation)

// 	treeSlice[0] =  Container{[][]complex128{}, treeSliceData}

// 	xPosAtLayer := []int{0}

// 	layer := 0

// 	currentContainer := treeSlice[xPosAtLayer[layer]]

// 	for layer > -1 {

// 		fmt.Println(treeSlice)

// 		// containerCursor := xPosAtLayer[layer]

// 		currentContainer = currentContainer.ChildrenEquations[xPosAtLayer[layer]]

// 		newToAppend, dataAdded := CreateATreeFromCurrentEquation(currentContainer.Equation)

// 		fmt.Println(newToAppend)

// 		if(dataAdded){
// 			currentContainer.ChildrenEquations[xPosAtLayer[layer]].ChildrenEquations = newToAppend
// 		}

// 		if(len(currentContainer.ChildrenEquations[xPosAtLayer[layer]].ChildrenEquations) != 0){

// 			currentContainer = currentContainer.ChildrenEquations[xPosAtLayer[layer]].ChildrenEquations[0]

// 			// currentContainer = currentContainer.ChildrenEquations[0]

// 			xPosAtLayer = append(xPosAtLayer, 0)

// 			layer++
// 		}else{
// 			xPosAtLayer[layer]++
// 			if(xPosAtLayer[layer] == len(currentContainer.ChildrenEquations)){
// 				xPosAtLayer[layer] = 0
// 				layer--
// 			}
// 		}

// 	}

// 	return treeSlice

// }






























