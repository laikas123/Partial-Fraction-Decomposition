package main

import (

	"fmt"
	"math/cmplx"
	"reflect"
	"sort"
	"strings"
	"strconv"
)

type Container struct {
	Parent *[][]complex128
	Children []*Container
}

type Operators struct {
	Data [][]complex128
}

type Factor struct {
	Data [][]complex128
}

type Fraction struct {
	Numerator []Factor
	Denominator []Factor
}

//TODO, when multiple variables get involved a third index needs to be added to the float slice
//which will allow the third index to essentially span the alphabet 0-25 A-Z for variable names
//for now since this is only being used for inverse laplace transform of one variable, everything is
//assumed to be 's'	



//BIG TODO, OBVIOUSLY AT SOME POINT A NUMBER WILL GET ZEROED OUT, IT'S IMPORTANT THAT IT DOESN'T GET TREATED AS A PARENTHESIS
//OTHERWISE THAT COULD MESS UP THE FLOW OF THINGS, BEGIN ADDING CHECKS FOR SUMMATIONS THAT RESULT IN 0

func main() {
	

	// equation := [][]complex128{gOP(), gNum(2, 3, 3, 0), gCP(3) }

	// equation := Create2DEquationFromSliceInputs(gOP(), gOP(), gNum(3,0, 3), gCP(1, 3), gOP(), gNum(1, 1, 2, 0, 3), gCP(2, 3), gNum(3, 0, 3), gCP(1, 3))

	// fmt.Println(DecodeFloatSliceToEquation(equation))

	panic("test")
	

}



func DecodeFloatSliceToEquation(equationInput [][]complex128 ) string {

//	CheckEquationForSyntaxErrors(equation)

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)
	
	equationString := ""

	

	for i := 0; i < len(equation); i++ {
		
		currentItem := equation[i]



		firstIndex := currentItem[0]
		secondIndex := currentItem[1]


		if(IsOP(firstIndex, secondIndex)){
			equationString += "("
			
		}else if(IsCP(currentItem)){
			equationString += " )"
			if(currentItem[2] != 0 && currentItem[2] != 1) {

				equationString += "^" + strconv.FormatFloat(real(currentItem[2]), 'f', -1, 64) + " " + GetStringForCodeOfCP(real(currentItem[4])) 
			}else{

				equationString += GetStringForCodeOfCP(real(currentItem[4])) + " "
			}
			
		}else if(IsNumber(firstIndex)){


			for i := 0; i < len(currentItem); i = i+2{
				multiplier := currentItem[i]
				exponent := currentItem[i+1]

				if((i < len(currentItem) - 2) && exponent == 1){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64)  + "s " + " "
				}else if((i < len(currentItem) - 2) && exponent == 0){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) + "  "
				}else if((i < len(currentItem) - 2) && exponent != 0 && exponent != 1){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) + "s^" + strconv.FormatFloat(real(exponent),'f', -1, 64) 
				}else if((i == len(currentItem) - 2) && exponent == 1){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) + "s " 
				}else if((i == len(currentItem) - 2) && exponent == 0){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) +"~"
				}else if((i ==  len(currentItem) - 2) && exponent != 0 && exponent != 1){
					equationString += strconv.FormatFloat(real(multiplier),'f', -1, 64) + "s^" + strconv.FormatFloat(real(exponent),'f', -1, 64) 
				}
			}
			
		}else if(IsOperator(currentItem)){
			equationString += " " + GetStringForCodeOfCP(real(currentItem[1])) + " "
		}else{
			panic("unknown equation item DecodeFloatSliceToEquation()")	
		}
	
		

	}

	
	// equationString = RemoveOperatorsBetweenTwoClosingParenthesisAndRemoveSpaces(equationString)

	return equationString

}


func gPlus() []complex128{
	return []complex128{complex(0,0), complex(1,0)}
}
func gPMinus() []complex128{
	return []complex128{complex(0,0), complex(2,0)}
}
func gPMultiply() []complex128{
	return []complex128{complex(0,0), complex(3,0)}
}
func gDivide() []complex128{
	return []complex128{complex(0,0), complex(4,0)}
}

func gNum(nums ...complex128) [][]complex128 {

	if(len(nums) == 2 && nums[0] != 0){
		return [][]complex128{[]complex128{nums[0], nums[1]}}
	}


	if( (len(nums)%3) != 2 || len(nums) < 5){
		panic("error, invalid amount of numbers gNum()")
	}


	returnSlice := [][]complex128{}

	// returnSlice = append(returnSlice, gOP())
	

	for i := 0; i < len(nums); i = i+3 {

		returnSlice = append(returnSlice, []complex128{nums[i], nums[i+1]})
		if((i+2) < len(nums)){
			returnSlice = append(returnSlice, []complex128{complex(0, 0), nums[i+2]})
		}
		
	}

	// returnSlice = append(returnSlice, gCP(1, 1))
	

	return returnSlice

}

func gOP() []complex128 {
	return []complex128{0, 0}
}

//Operators are how the items within the parenthesis should interact with
//the next neighbor...
//codes are
//1 = add
//2 = subtract
//3 = multiply
//4 = divide
func gCP(exponent complex128, operator complex128) []complex128 {
	return []complex128{0, 1, exponent, 0, operator}
}



func SimplifyInnerParenthesis(equationInput [][]complex128) [][]complex128 {

	CheckEquationForSyntaxErrors(equationInput, "SimplifyInnerParenthesis()")

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

	numbersHolder := [][]complex128{}

	operatorsHolder := [][]complex128{}

	indexOpener := -1

	indexCloser := -1

	foundValid := false

	fmt.Println("SimplifyInnerParenthesis() equation", DecodeFloatSliceToEquation(equation))

	for i := 0; i < len(equation); i ++ {

		if(foundValid){
			break
		}

		if(IsOP(equation[i][0], equation[i][1])){

			indexOpener = i

			fmt.Println("character")

			checkingIfValid := true

			sawOneNumber := false


			//these two bools are used to make sure
			//numbers and symbols alternate
			indexShouldBeNumber := true
			indexShouldBeOperator := false

			cursor := i

			//set these to null before each attempt to not have lingering data
			numbersHolder = [][]complex128{}		
			operatorsHolder = [][]complex128{}

			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}

				if(!sawOneNumber){

					for(!IsCP(equation[cursor]) && !IsOP(equation[cursor][0], equation[cursor][1])){


						if(cursor >= len(equation)){
							checkingIfValid = false
							foundValid = false
							break
						}

						if(IsNumber(equation[cursor][0]) && indexShouldBeNumber){
							numbersHolder = append(numbersHolder, equation[cursor])
							indexShouldBeNumber = false
							indexShouldBeOperator = true
						}else if(IsOperator(equation[cursor]) && indexShouldBeOperator){
							operatorsHolder = append(operatorsHolder, equation[cursor])
							indexShouldBeNumber = true
							indexShouldBeOperator = false
						}else{
							checkingIfValid = false
							foundValid = false
							break	
						}

						cursor++

					}

					//make sure what broke the loop was a closing parenthesis
					if(IsCP(equation[cursor])){
						indexCloser = cursor
						checkingIfValid = false
						foundValid = true
						break
					}else{
						checkingIfValid = false
						foundValid = false
						break	
					}

			}

		}


	}
}

	if(foundValid){

		
		fmt.Println("index opener", indexOpener, "indexCloser", indexCloser)

		fmt.Println("numbers found", numbersHolder)
		fmt.Println("operators found", operatorsHolder)

		
		for i := 0; i < len(operatorsHolder); i++ {

			//if it equals multiply
			if(real(operatorsHolder[i][1]) == 3){

				leftNum := numbersHolder[i]
				rightNum := numbersHolder[i+1]

				result := MultiplyTwoAdjacentNumbers(leftNum, rightNum)

				numbersHolder[i] = result

				numbersHolder = append(numbersHolder[0:i+1], numbersHolder[(i+2):len(numbersHolder)]...)

				operatorsHolder = append(operatorsHolder[0:i], operatorsHolder[(i+1):len(operatorsHolder)]...)
				

			//if it equals divide
			}else if(real(operatorsHolder[i][1]) == 4){

				leftNum := numbersHolder[i]
				rightNum := numbersHolder[i+1]

				result := DivideTwoAdjacentNumbers(leftNum, rightNum)

				numbersHolder[i] = result

				numbersHolder = append(numbersHolder[0:i+1], numbersHolder[(i+2):len(numbersHolder)]...)
				
				
				operatorsHolder = append(operatorsHolder[0:i], operatorsHolder[i+1:len(operatorsHolder)]...)
				
			}

		}

		fmt.Println("pre numbersHolder", numbersHolder, "pre operators holder", operatorsHolder)

		numbersHolder, operatorsHolder = SortParenthesisContainingOnlyPlusAndMinusBySExponent(numbersHolder, operatorsHolder)

		fmt.Println("post numbersHolder", numbersHolder, "post operators holder", operatorsHolder)

		for i := 0; i < len(operatorsHolder); i++ {



			//if it equals multiply
			if(real(operatorsHolder[i][1]) == 1 && TwoAdjacentNumbersCanAddOrSubtract(numbersHolder[i], numbersHolder[i+1])){

				leftNum := numbersHolder[i]
				rightNum := numbersHolder[i+1]

				result := AddTwoAdjacentNumbers(leftNum, rightNum)

				numbersHolder[i] = result

				numbersHolder = append(numbersHolder[0:i+1], numbersHolder[(i+2):len(numbersHolder)]...)

				
				operatorsHolder = append(operatorsHolder[0:i], operatorsHolder[(i+1):len(operatorsHolder)]...)
				

			//if it equals divide
			}else if(real(operatorsHolder[i][1]) == 2 && TwoAdjacentNumbersCanAddOrSubtract(numbersHolder[i], numbersHolder[i+1])){

				leftNum := numbersHolder[i]
				rightNum := numbersHolder[i+1]

				result := SubtractTwoAdjacentNumbers(leftNum, rightNum)

				numbersHolder[i] = result

				numbersHolder = append(numbersHolder[0:i+1], numbersHolder[(i+2):len(numbersHolder)]...)
			
				operatorsHolder = append(operatorsHolder[0:i], operatorsHolder[i+1:len(operatorsHolder)]...)
				
			}

		}


		fmt.Println("numbers after", numbersHolder)
		fmt.Println("operators after", operatorsHolder)

		resultOfOperations := [][]complex128{}

		//if there is no operators then the parenthesis were simplified to one number
		if(len(operatorsHolder) == 0){
			resultOfOperations = numbersHolder
		}


		for i := 0; i < len(operatorsHolder); i++ {

			resultOfOperations = append(resultOfOperations, numbersHolder[i])

			resultOfOperations = append(resultOfOperations, operatorsHolder[i])

			if(i == len(operatorsHolder) -1 ){
				resultOfOperations = append(resultOfOperations, numbersHolder[i + 1])
			}

		}

		returnEquation := [][]complex128{}

		returnEquation = append(equation[0:indexOpener+1], resultOfOperations...)
		returnEquation = append(returnEquation, equation[indexCloser:len(equation)]...)

		if(InnerParenthesisCanBeSimplifiedFurther(numbersHolder, operatorsHolder)){
			return SimplifyInnerParenthesis(returnEquation)
		}else{
			return returnEquation
		}

		

	}else{

		fmt.Println("nothing valid")
		
		return equation
	}	
}


// finds "(" followed by numbers followed by ")^x" where x is some power greater than 1
func FoilOutParenthesisRaisedToExponent(equationInput [][]complex128) [][]complex128 {

	CheckEquationForSyntaxErrors(equationInput, "FoilOutParenthesisRaisedToExponent()")

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

	equation = SimplifyInnerParenthesis(equation)

	numbersHolder := [][]complex128{}

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

			sawOneNumber := false


			//these two bools are used to make sure
			//numbers and symbols alternate
			indexShouldBeNumber := true
			indexShouldBeOperator := false

			cursor := i

			//set these to null before each attempt to not have lingering data
			numbersHolder = [][]complex128{}		
			

			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}

				if(!sawOneNumber){

					for(!IsCP(equation[cursor]) && !IsOP(equation[cursor][0], equation[cursor][1])){


						if(cursor >= len(equation)){
							checkingIfValid = false
							foundValid = false
							break
						}

						if(IsNumber(equation[cursor][0]) && indexShouldBeNumber){
							numbersHolder = append(numbersHolder, equation[cursor])
							indexShouldBeNumber = false
							indexShouldBeOperator = true
						}else if(IsOperator(equation[cursor]) && indexShouldBeOperator){
							numbersHolder = append(numbersHolder, equation[cursor])
							indexShouldBeNumber = true
							indexShouldBeOperator = false
						}else{
							checkingIfValid = false
							foundValid = false
							break	
						}

						cursor++

					}

					//make sure what broke the loop was a closing parenthesis
					if(IsCP(equation[cursor])){
						if(real(equation[cursor][2]) > 1){
							indexCloser = cursor
							exponentCloser = int(real(equation[cursor][2]))
							checkingIfValid = false
							foundValid = true
							break
						}else{
							checkingIfValid = false
							foundValid = false
							break	
						}
					}else{
						checkingIfValid = false
						foundValid = false
						break	
					}

			}

		}


	}
}













	if(foundValid){

		fmt.Println("NUMBERS FOUND", numbersHolder)
		
		exponentiationResult := numbersHolder

		//this is the times to perform the exponentiation of the parenthesis
		for exponentCloser > 1 {

			exponentiationResult = MultiplyNeighboringParenthesis(exponentiationResult, numbersHolder)

			exponentCloser--
		}


		// sliceToInsert := MultiplyParenthesisGivenExponent(numbersHolder, exponentCloser)

		slicesToInsert :=  Create2DEquationFromSliceInputs(gOP(), exponentiationResult, gCP(1, 3))

		returnEquation := [][]complex128{}

		returnEquation = append(returnEquation, equation[0:indexOpener]...)
		
		returnEquation = append(returnEquation, slicesToInsert...)
		
		returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)
	
		returnEquation = CleanCopyEntire2Dcomplex128Slice(returnEquation)

		fmt.Println("RETURN EQUATION", DecodeFloatSliceToEquation(returnEquation))		

		//recursive call if there was a change will check if more foils possible
		return FoilOutParenthesisRaisedToExponent(returnEquation)

	}else{

		fmt.Println("no valid found")
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

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

	equation = SimplifyInnerParenthesis(equation)

	numbersHolderFirstNeighbor := [][]complex128{}
	numbersHolderSecondNeighbor := [][]complex128{}



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

			sawOneNumber := false

			cursor := i

			//set these to null before each attempt to not have lingering data
			numbersHolderFirstNeighbor = [][]complex128{}
			numbersHolderSecondNeighbor = [][]complex128{}
			fmt.Println("should be only once here")

			//these two bools are used to make sure
			//numbers and symbols alternate
			indexShouldBeNumber := true
			indexShouldBeOperator := false

			fmt.Println("OP")


			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}



				if(!sawOneNumber && !secondTerm){

					for(!IsCP(equation[cursor]) && !IsOP(equation[cursor][0], equation[cursor][1])){

						fmt.Println("NUM")

						if(cursor >= len(equation)){
							checkingIfValid = false
							foundValid = false
							break
						}

						if(IsNumber(equation[cursor][0]) && indexShouldBeNumber){
							numbersHolderFirstNeighbor = append(numbersHolderFirstNeighbor, equation[cursor])
							indexShouldBeNumber = false
							indexShouldBeOperator = true
						}else if(IsOperator(equation[cursor]) && indexShouldBeOperator){
							numbersHolderFirstNeighbor = append(numbersHolderFirstNeighbor, equation[cursor])
							// operatorsHolderFirstNeighbor = append(operatorsHolderFirstNeighbor, equation[cursor])
							indexShouldBeNumber = true
							indexShouldBeOperator = false
						}else{
							checkingIfValid = false
							foundValid = false
							break	
						}

						cursor++

					}

					//make sure what broke the loop was a closing parenthesis
					if(IsCP(equation[cursor])){
						if(real(equation[cursor][2]) == 1){
							fmt.Println("VALID CP")
							if((cursor+1) >= len(equation)){
								break
							}else if(IsOP(equation[cursor+1][0], equation[cursor+1][1])){
								fmt.Println("VALID OP")
								cursor++
								//reset the booleans for second parenthesis
								indexShouldBeNumber = true
								indexShouldBeOperator = false
								secondTerm = true
								sawOneNumber = false
							}else{
								break
							}
						}
					}else{
						checkingIfValid = false
						foundValid = false
						break	
					}

			}else if(!sawOneNumber && secondTerm){

				for(!IsCP(equation[cursor]) && !IsOP(equation[cursor][0], equation[cursor][1])){

					fmt.Println("NUM 2")

					if(cursor >= len(equation)){
						checkingIfValid = false
						foundValid = false
						break
					}

					if(IsNumber(equation[cursor][0]) && indexShouldBeNumber){
						numbersHolderSecondNeighbor = append(numbersHolderSecondNeighbor, equation[cursor])
						indexShouldBeNumber = false
						indexShouldBeOperator = true
					}else if(IsOperator(equation[cursor]) && indexShouldBeOperator){
						// operatorsHolderSecondNeighbor = append(operatorsHolderSecondNeighbor, equation[cursor])
						numbersHolderSecondNeighbor = append(numbersHolderSecondNeighbor, equation[cursor])
						indexShouldBeNumber = true
						indexShouldBeOperator = false
					}else{
						checkingIfValid = false
						foundValid = false
						break	
					}

					cursor++

				}

				//make sure what broke the loop was a closing parenthesis
				if(IsCP(equation[cursor]) ){
					if(real(equation[cursor][2]) == 1){
						fmt.Println("VALID CP")
						indexCloser = cursor
						checkingIfValid = false
						foundValid = true
						break
					}else{
						checkingIfValid = false
						foundValid = false
						break
					}	
				}else{
					checkingIfValid = false
					foundValid = false
					break	
				}

			}		
				

		}

	}


	}

	if(foundValid){

	

		// numbersHolderFirstNeighbor = SimplifyInnerParenthesis(numbersHolderFirstNeighbor)
		// numbersHolderSecondNeighbor = SimplifyInnerParenthesis(numbersHolderSecondNeighbor)

		// numbersHolderFirstNeighborWithParenthesis := Create2DEquationFromSliceInputs(gOP(), numbersHolderFirstNeighbor, gCP(1, 3))
		// numbersHolderSecondNeighborWithParenthesis := Create2DEquationFromSliceInputs(gOP(), numbersHolderSecondNeighbor, gCP(1, 3))

		// numbersHolderFirstNeighborWithParenthesis = SimplifyInnerParenthesis(numbersHolderFirstNeighborWithParenthesis)
		// numbersHolderSecondNeighborWithParenthesis = SimplifyInnerParenthesis(numbersHolderSecondNeighborWithParenthesis)

		// numbersHolderFirstNeighbor = numbersHolderFirstNeighborWithParenthesis[1:(len(numbersHolderFirstNeighborWithParenthesis)-1)]
		// numbersHolderSecondNeighbor = numbersHolderSecondNeighborWithParenthesis[1:(len(numbersHolderSecondNeighborWithParenthesis)-1)]

		fmt.Println("first number", numbersHolderFirstNeighbor)
		fmt.Println("second number", numbersHolderSecondNeighbor)

		// panic("got here")

		sliceToInsert := MultiplyNeighboringParenthesis(numbersHolderFirstNeighbor, numbersHolderSecondNeighbor)

		slicesToInsert :=  Create2DEquationFromSliceInputs(gOP(), sliceToInsert, gCP(1, 3))

		returnEquation := [][]complex128{}

		returnEquation = append(returnEquation, equation[0:indexOpener]...)
		
		returnEquation = append(returnEquation, slicesToInsert...)
		
		returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)

		returnEquation = CleanCopyEntire2Dcomplex128Slice(returnEquation)

		fmt.Println("RETURN EQUATION", DecodeFloatSliceToEquation(returnEquation))		

		//recursive call if there was a change will check if more foils possible
		return FoilNeighborParenthesis(returnEquation)

	}else{

		//if no foils possible return input
		return equation
	}	



}




func MultiplyNeighboringParenthesis(numbers1 [][]complex128, numbers2 [][]complex128) [][]complex128{

	leftTerm := numbers1
	rightTerm := numbers2


	allNumbersFromFoil := [][]complex128{}


	fmt.Println("left term", leftTerm)
	fmt.Println("right term", rightTerm)
	fmt.Println("allNumbersFromFoil", allNumbersFromFoil)

	resultCurrentLoop := [][]complex128{}


	for i := 0; i < len(leftTerm); i = (i+2) {

		currentNumberMultiplier := leftTerm[i][0]

		currentNumberExponent := leftTerm[i][1]

		//if its not the first term then each number has a plus
		//or a minus behind it, if there's a minus then the number needs to 
		//turned negative
		if(i != 0){
			operator := leftTerm[i-1]
			//minus sign check
			if(real(operator[1]) == 2){
				currentNumberMultiplier = currentNumberMultiplier*-1
			}
		}
		

		for j := 0; j < len(rightTerm); j = (j+2) {

			foilNumberMultiplier := rightTerm[j][0]
			foilNumberExponent := rightTerm[j][1]

			//if its not the first term then each number has a plus
			//or a minus behind it, if there's a minus then the number needs to 
			//turned negative
			if(j != 0){
				operator := leftTerm[j-1]
				//minus sign check
				if(real(operator[1]) == 2){
					foilNumberMultiplier = foilNumberMultiplier*-1
				}
			}
			
			resultCurrentLoop = append(resultCurrentLoop, []complex128{currentNumberMultiplier*foilNumberMultiplier, currentNumberExponent+foilNumberExponent})

		}
	}

	fmt.Println("result from current loop", resultCurrentLoop)

	fmt.Println("allNumbersFromFoil", allNumbersFromFoil)

	allNumbersFromFoil = resultCurrentLoop
		
	mapOfExponents := make(map[complex128][]complex128)

	for i := 0; i < len(allNumbersFromFoil); i++ {

		if(mapOfExponents[allNumbersFromFoil[i][1]] == nil){
			mapOfExponents[allNumbersFromFoil[i][1]] = []complex128{allNumbersFromFoil[i][0]} 
		}else{
			mapOfExponents[allNumbersFromFoil[i][1]] = append(mapOfExponents[allNumbersFromFoil[i][1]], allNumbersFromFoil[i][0])
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

		returnNumbers := [][]complex128{}

		for i := 0; i < len(copyOfExponentsComplex); i++ {
			returnNumbers = append(returnNumbers, []complex128{mapOfSimpliefiedMultipliersForExponents[copyOfExponentsComplex[i]], copyOfExponentsComplex[i]})

			if(i < (len(copyOfExponentsComplex)-1)){
				returnNumbers = append(returnNumbers, []complex128{complex(0,0), complex(1,0)})
			}


		}


		return returnNumbers

	}




} 



func FactorQuadraticsWithABCAllPresent(equationInput [][]complex128)[][]complex128 {

	
	CheckEquationForSyntaxErrors(equationInput, "FactorQuadraticsWithABCAllPresent()")

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

	equation = SimplifyInnerParenthesis(equation)

	numbersHolder := [][]complex128{}

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

			cursor := i

			numbersHolder = [][]complex128{}
	
			indexShouldBeNumber := true
			indexShouldBeOperator := false		

			secondDegreeExponentPresent := false
			firstDegreeExponentPresent := false
			zeroDegreeExponentPresent := false


			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}


			
					for(!IsCP(equation[cursor]) && !IsOP(equation[cursor][0], equation[cursor][1])) {
						numbersHolder = append(numbersHolder, equation[cursor])
						
						if(IsNumber(equation[cursor][0]) && indexShouldBeNumber){

							fmt.Println("number", equation[cursor])

							if(real(equation[cursor][1]) == 2){
								secondDegreeExponentPresent = true
								aTerm = equation[cursor][0]
							}else if(real(equation[cursor][1]) == 1){
								firstDegreeExponentPresent = true
								bTerm = equation[cursor][0]
							}else if(real(equation[cursor][1]) == 0){
								zeroDegreeExponentPresent = true
								cTerm = equation[cursor][0]
							}

							indexShouldBeNumber = false
							indexShouldBeOperator = true

							cursor++

						}else if(IsOperator(equation[cursor]) && indexShouldBeOperator){
							
							fmt.Println("operator", equation[cursor])
							//only plus or minuses are a valid quadratic, it's possible multipliers could be valid
							//however the inner parenthesis should already be simplified since that function is called
							//at the start of this function
							if(equation[cursor][1] == 1 || equation[cursor][1] == 2){
									indexShouldBeNumber = true
									indexShouldBeOperator = false
									cursor++
							}else{
								checkingIfValid = false
								break
							}
						}else{
							checkingIfValid = false
							break
						}

					
				}
				if IsCP(equation[cursor]){
						if(secondDegreeExponentPresent && firstDegreeExponentPresent && zeroDegreeExponentPresent){
							indexCloser = cursor
							checkingIfValid = false
							foundValid = true
							break
						}else{
							checkingIfValid = false
							break
						}
						
				}
			}

		


	}
}
	if(foundValid){

		fmt.Println("a", aTerm, "b", bTerm, "c", cTerm)

		if(aTerm == 0 || bTerm == 0 || cTerm == 0 ){
			panic("too many or too few values for a value for quadratic formula")
		}else{

			root1, root2 := QuadraticFormula(aTerm, bTerm, cTerm)

			fmt.Println("root1", root1, "root2", root2)

			root1Slice := Create2DEquationFromSliceInputs(gOP(), gNum(1, 1, 1, -1*root1, 0), gCP(1, 1))

			root2Slice := Create2DEquationFromSliceInputs(gOP(), gNum(1, 1, 1, -1*root2, 0), gCP(1, 1))

//			root2Slice := []complex128{complex(1, 0), complex(1,0), complex(0,0), complex(1,0), -1*root2, complex(0,0)}

			slicesToInsert :=  Create2DEquationFromSliceInputs(root1Slice, root2Slice)

			returnEquation := [][]complex128{}

			returnEquation = append(returnEquation, equation[0:indexOpener]...)
			
			returnEquation = append(returnEquation, slicesToInsert...)
			
			returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)
		
			returnEquation = CleanCopyEntire2Dcomplex128Slice(returnEquation)

			fmt.Println("RETURN EQUATION", DecodeFloatSliceToEquation(returnEquation))		

			// panic("ok")

			//recursive call if there was a change will check if more foils possible
			return FactorQuadraticsWithABCAllPresent(returnEquation)
		}
	}else{

		//if no foils quadratic factors return input
		fmt.Println("RETURN EQUATION", DecodeFloatSliceToEquation(equation))		

		return equation
	}	



}



func FactorQuadraticsWithABOnlyPresent(equationInput [][]complex128)[][]complex128 {
	
	CheckEquationForSyntaxErrors(equationInput, "FactorQuadraticsWithABCAllPresent()")

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

	equation = SimplifyInnerParenthesis(equation)

	numbersHolder := [][]complex128{}

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

			cursor := i

			numbersHolder = [][]complex128{}
	
			indexShouldBeNumber := true
			indexShouldBeOperator := false		

			secondDegreeExponentPresent := false
			firstDegreeExponentPresent := false
			

			breakBecauseCIsPresent := false

			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}


			
					for(!IsCP(equation[cursor]) && !IsOP(equation[cursor][0], equation[cursor][1])) {
						numbersHolder = append(numbersHolder, equation[cursor])
						
						if(IsNumber(equation[cursor][0]) && indexShouldBeNumber){

							fmt.Println("number", equation[cursor])

							if(real(equation[cursor][1]) == 2){
								secondDegreeExponentPresent = true
								aTerm = equation[cursor][0]
							}else if(real(equation[cursor][1]) == 1){
								firstDegreeExponentPresent = true
								bTerm = equation[cursor][0]
							}else if(real(equation[cursor][1]) == 0){
								secondDegreeExponentPresent = false
								firstDegreeExponentPresent = false
								breakBecauseCIsPresent = true
								checkingIfValid = false
								break
							}

							indexShouldBeNumber = false
							indexShouldBeOperator = true

							cursor++

						}else if(IsOperator(equation[cursor]) && indexShouldBeOperator){
							
							fmt.Println("operator", equation[cursor])
							//only plus or minuses are a valid quadratic, it's possible multipliers could be valid
							//however the inner parenthesis should already be simplified since that function is called
							//at the start of this function
							if(equation[cursor][1] == 1 || equation[cursor][1] == 2){
									indexShouldBeNumber = true
									indexShouldBeOperator = false
									cursor++
							}else{
								checkingIfValid = false
								break
							}
						}else{
							checkingIfValid = false
							break
						}

					
				}
				if IsCP(equation[cursor]){

						if(breakBecauseCIsPresent){
							checkingIfValid = false
							foundValid = false
							break							
						}

						if(secondDegreeExponentPresent && firstDegreeExponentPresent){
							indexCloser = cursor
							checkingIfValid = false
							foundValid = true
							break
						}else{
							checkingIfValid = false
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

			// root1Slice := []complex128{complex(1, 0), complex(1,0), -1*root1, complex(0,0)}

			// root2Slice := []complex128{complex(1, 0), complex(1,0), -1*root2, complex(0,0)}


			root1Slice := Create2DEquationFromSliceInputs(gOP(), gNum(1, 1, 1, -1*root1, 0), gCP(1, 3))

			root2Slice := Create2DEquationFromSliceInputs(gOP(), gNum(1, 1, 1, -1*root2, 0), gCP(1, 3))

			completingTheSquareTermSlice := Create2DEquationFromSliceInputs(gOP(), gNum((-1*cTerm*aTermBeforeScale), complex(0, 0)), gCP(1, 3))


			slicesToInsert :=  Create2DEquationFromSliceInputs(gOP(), scaleDownASlice, root1Slice, root2Slice, gCP(1, 1), completingTheSquareTermSlice)			



			// completingTheSquareTermSlice := []complex128{(-1*cTerm*aTermBeforeScale), complex(0, 0)}

			// slicesToInsert :=  [][]complex128{gOP(), gOP(), scaleDownASlice, gCP(1, 3), gOP(), root1Slice, gCP(1, 3), gOP(), root2Slice, gCP(1, 3), completingTheSquareTermSlice, gCP(1, 3)}			

			returnEquation := [][]complex128{}

			returnEquation = append(returnEquation, equation[0:indexOpener]...)
			
			returnEquation = append(returnEquation, slicesToInsert...)
			
			returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)
		
			returnEquation = CleanCopyEntire2Dcomplex128Slice(returnEquation)

			fmt.Println("RETURN EQUATION", DecodeFloatSliceToEquation(returnEquation))		

			//recursive call if there was a change will check if more foils possible
			return FactorQuadraticsWithABOnlyPresent(returnEquation)
		}
	}else{

		//if no foils quadratic factors return input
		return equation
	}	



}




func FactorQuadraticsWithACOnlyPresent(equationInput [][]complex128)[][]complex128 {
	
	CheckEquationForSyntaxErrors(equationInput, "FactorQuadraticsWithABCAllPresent()")

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

	equation = SimplifyInnerParenthesis(equation)

	numbersHolder := [][]complex128{}

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

			cursor := i

			numbersHolder = [][]complex128{}
	
			indexShouldBeNumber := true
			indexShouldBeOperator := false		

			secondDegreeExponentPresent := false
			zeroDegreeExponentPresent := false

			breakBecauseBIsPresent := false


			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation
				}


			
					for(!IsCP(equation[cursor]) && !IsOP(equation[cursor][0], equation[cursor][1])) {
						numbersHolder = append(numbersHolder, equation[cursor])
						
						if(IsNumber(equation[cursor][0]) && indexShouldBeNumber){

							fmt.Println("number", equation[cursor])

							if(real(equation[cursor][1]) == 2){
								secondDegreeExponentPresent = true
								aTerm = equation[cursor][0]
							}else if(real(equation[cursor][1]) == 1){
								//set these false so that the program doesn't continue for these parenthesis
								//that is because the b term is present
								secondDegreeExponentPresent = false
								zeroDegreeExponentPresent = false
								breakBecauseBIsPresent = true
								checkingIfValid = false
								break
							}else if(real(equation[cursor][1]) == 0){
								zeroDegreeExponentPresent = true
								cTerm = equation[cursor][0]
							}

							indexShouldBeNumber = false
							indexShouldBeOperator = true

							cursor++

						}else if(IsOperator(equation[cursor]) && indexShouldBeOperator){
							


							fmt.Println("operator", equation[cursor])
							//only plus or minuses are a valid quadratic, it's possible multipliers could be valid
							//however the inner parenthesis should already be simplified since that function is called
							//at the start of this function
							if(equation[cursor][1] == 1 || equation[cursor][1] == 2){
									indexShouldBeNumber = true
									indexShouldBeOperator = false
									cursor++
							}else{
								checkingIfValid = false
								break
							}
						}else{
							checkingIfValid = false
							break
						}

					
				}
				if IsCP(equation[cursor]){

						if(breakBecauseBIsPresent){
							checkingIfValid = false
							foundValid = false
							break
						}

						if(secondDegreeExponentPresent && zeroDegreeExponentPresent){
							indexCloser = cursor
							checkingIfValid = false
							foundValid = true
							break
						}else{
							checkingIfValid = false
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

			sqrtCTerm := cmplx.Pow(cTerm, complex(0.5, 0))

			root1Slice := Create2DEquationFromSliceInputs(gOP(), gNum(1, 1, 1, sqrtCTerm, 0), gCP(1, 3))

			root2Slice := Create2DEquationFromSliceInputs(gOP(), gNum(1, 1, 1, -1*sqrtCTerm, 0), gCP(1, 3))

			slicesToInsert :=  Create2DEquationFromSliceInputs(gOP(), scaleDownASlice, gCP(1, 3), root1Slice, root2Slice)			

			returnEquation := [][]complex128{}

			returnEquation = append(returnEquation, equation[0:indexOpener]...)
			
			returnEquation = append(returnEquation, slicesToInsert...)
			
			returnEquation = append(returnEquation, equation[indexCloser+1:len(equation)]...)
		
			returnEquation = CleanCopyEntire2Dcomplex128Slice(returnEquation)

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
	if(len(nums) == 5){
		return true
	}else{
		return false
	}
}

func IsOperator(nums []complex128) bool {
	
	if(real(nums[1]) != 1 && real(nums[1]) != 2 && real(nums[1]) != 3 && real(nums[1]) != 4){
		panic("invalid operator IsOperator()")
	}

	if(len(nums) == 2){
		if(nums[0] == 0 && nums[1] != 0){
			return true
		}else{
			return false
		}
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

	fmt.Println(DecodeFloatSliceToEquation(equation))	

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

			}else if(IsOperator(currentItem)){
				
			}else{
				fmt.Println("current item", currentItem)
				panic("Syntax Error unknown item type CheckEquationForSyntaxErrors()")
			}

			//this would occur if there's more closing parenthesis than opening
			if(depthLevel == -1){
				panic("Syntax Error too many ) for the number of ( CheckEquationForSyntaxErrors()")
			}

			if(IsCP(currentItem)){
				break	
			}
		}
	}

	if(depthLevel != 0){
		panic("Syntax Error not all ( items were closed properly CheckEquationForSyntaxErrors()")
	}

}


func RemoveUnusedParenthesis(equationInput [][]complex128) [][]complex128 {


	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)


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


func CleanCopyEntire2Dcomplex128Slice(equationToCopy [][]complex128) [][]complex128 {

	returnEquation := [][]complex128{}

	for i := 0; i < len(equationToCopy); i++ {

		cleanCopyToAppend := make([]complex128, len(equationToCopy[i]))

		itemsCopied := copy(cleanCopyToAppend, equationToCopy[i])

		if(itemsCopied != len(equationToCopy[i])){
			panic("invalid copy CleanCopyEntire2Dcomplex128Slice()")
		}else{
			returnEquation = append(returnEquation, cleanCopyToAppend)
		}


	}

	return returnEquation

}




func RemoveLastItemIfItIsOpeningParenthesis(equation [][]complex128) [][]complex128{

	cleanCopyToReturn := CleanCopyEntire2Dcomplex128Slice(equation)


	for (IsOP(cleanCopyToReturn[len(cleanCopyToReturn)-1][0], cleanCopyToReturn[len(cleanCopyToReturn)-1][1])) {
		cleanCopyToReturn = cleanCopyToReturn[0:(len(cleanCopyToReturn)-1)]
		
	}

	return cleanCopyToReturn


}



func RemoveExcessParenthesisViaDepthCheck(equationInput [][]complex128) [][]complex128 {

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

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

			if(IsCP(currentItemEquation1)){
				
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

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

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


func CreateATreeFromCurrentEquation(equation [][]complex128) []*Container {

	// CheckEquationForSyntaxErrors(equation, "CreateATreeFromCurrentEquation()")



	breakAll := false
	

	

	equations := []*Container{}

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
					cleanCopyToAppend := CleanCopyEntire2Dcomplex128Slice(equation)
					dataToAddToTree := cleanCopyToAppend[openerIndex+1:cursor]
					fmt.Println("data added", dataToAddToTree)
					sliceForChildren := []*Container{}
					//creates a new struct with the parent equation and an empty slice for its children
					equations = append(equations, &Container{&dataToAddToTree, sliceForChildren})
					
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


func CreateEntireTreeForEquation(equation [][]complex128) []*Container {



	topLevelContainer := []*Container{}

	topLevelContainerChild := &Container{&equation, []*Container{}}

	topLevelContainer = append(topLevelContainer, topLevelContainerChild)

	//start at 0, append a new number for each new depth
	//remove the last number when coming back from a previous depth
	xPosForCurrentDepthChildren := []int{0} 

	depth := 0

	previousChildren := [][]*Container{}

	currentChildren := topLevelContainer

	currentChild := currentChildren[xPosForCurrentDepthChildren[depth]]

	for depth > -1 {




		//this gets the child container 
		currentChild = currentChildren[xPosForCurrentDepthChildren[depth]]		

		fmt.Println("depth", depth, "xpos", xPosForCurrentDepthChildren[depth])

		fmt.Println("current child", DecodeFloatSliceToEquation(*currentChild.Parent))

		//this generates new children from the child containers equation
		newChildren := CreateATreeFromCurrentEquation(*currentChild.Parent)

		for i := 0; i < len(newChildren); i++ {
			fmt.Println("new children", DecodeFloatSliceToEquation(*newChildren[i].Parent))
		}

		
	
		//if there was children to add immediately add them and move to the new depth
		if(len(newChildren) > 0){

		//	fmt.Println("hello 1")
		//	fmt.Println(topLevelContainer)

			currentChild.Children = newChildren
		//	fmt.Println("hello 2")

		//	fmt.Println(topLevelContainer)

			for i := 0; i < len(currentChild.Children); i++ {
				fmt.Println("new children copied", DecodeFloatSliceToEquation(*currentChild.Children[i].Parent))
			}



			// PrintTopLevelContainer(topLevelContainer)

			//increment the current depth horizontally before moving one deeper
			xPosForCurrentDepthChildren[depth]++
			//start at index 0 for the next depth
			xPosForCurrentDepthChildren = append(xPosForCurrentDepthChildren, 0)
			//move to the next depth for horizontal cursor
			depth++

			//keep track of the children to move back into 
			previousChildren = append(previousChildren, currentChildren)

			//move to the next depth of children
			currentChildren = currentChild.Children

			fmt.Println(topLevelContainer)

		//if there was no children to add remain at the current depth
		}else{

			//move horizontally at the current depth
			xPosForCurrentDepthChildren[depth]++

			//if the horizontal cursor is out of bounds for the current children
			//do this in a loop as its possible the previous depth was out of 
			//bounds as well
			for(xPosForCurrentDepthChildren[depth] >= len(currentChildren)){

				//reset the cursor for this depth
				xPosForCurrentDepthChildren[depth] = 0

				//remove the horizontal cursor for this depth as it no longer exists 
				xPosForCurrentDepthChildren = xPosForCurrentDepthChildren[0:(len(xPosForCurrentDepthChildren) - 1)]

				//decrement layers
				depth--

				//if the horizontal cursor was out of bounds at depth 0
				//then the task has ended
				if(depth == -1){
					break
				}

				//get the previous children at the previous depth
				currentChildren = previousChildren[len(previousChildren) - 1]

				//remove this depth of previous children as it is now the current children
				previousChildren = previousChildren[0:(len(previousChildren)-1)]

				if(xPosForCurrentDepthChildren[depth] < len(currentChildren)){
					currentChild = currentChildren[xPosForCurrentDepthChildren[depth]]
				}
				
				
			}



		}


	}

	fmt.Println(topLevelContainer[0].Children)

	return topLevelContainer


}



func IntelligentlyPrintTree(tree []*Container) {

	fmt.Println("PRINTING TREE")


	topLevelContainer := tree

	//start at 0, append a new number for each new depth
	//remove the last number when coming back from a previous depth
	xPosForCurrentDepthChildren := []int{0} 

	depth := 0

	previousChildren := [][]*Container{}

	currentChildren := topLevelContainer

	currentChild := currentChildren[xPosForCurrentDepthChildren[depth]]

	for depth > -1 {


		fmt.Println("depth", depth)
		fmt.Println("equation at depth", depth, "x pos", xPosForCurrentDepthChildren[depth],  "equation", DecodeFloatSliceToEquation(*currentChildren[xPosForCurrentDepthChildren[depth]].Parent))
		// for i := 0; i < len(currentChildren); i++ {
		// 	fmt.Println(DecodeFloatSliceToEquation(*currentChildren[i].Parent))
		// }

		//this gets the child container 
		currentChild = currentChildren[xPosForCurrentDepthChildren[depth]]		

		//this generates new children from the child containers equation
		newChildren := currentChild.Children

	
		//if there was children to add immediately add them and move to the new depth
		if(len(newChildren) > 0){


			xPosForCurrentDepthChildren[depth]++
			//start at index 0 for the next depth
			xPosForCurrentDepthChildren = append(xPosForCurrentDepthChildren, 0)
			//move to the next depth for horizontal cursor
			depth++

			//keep track of the children to move back into 
			previousChildren = append(previousChildren, currentChildren)

			//move to the next depth of children
			currentChildren = currentChild.Children

			

		//if there was no children to add remain at the current depth
		}else{

			//move horizontally at the current depth
			xPosForCurrentDepthChildren[depth]++

			//if the horizontal cursor is out of bounds for the current children
			//do this in a loop as its possible the previous depth was out of 
			//bounds as well
			for(xPosForCurrentDepthChildren[depth] >= len(currentChildren)){

				//reset the cursor for this depth
				xPosForCurrentDepthChildren[depth] = 0

				//remove the horizontal cursor for this depth as it no longer exists 
				xPosForCurrentDepthChildren = xPosForCurrentDepthChildren[0:(len(xPosForCurrentDepthChildren) - 1)]

				//decrement layers
				depth--

				//if the horizontal cursor was out of bounds at depth 0
				//then the task has ended
				if(depth == -1){
					break
				}

				//get the previous children at the previous depth
				currentChildren = previousChildren[len(previousChildren) - 1]

				//remove this depth of previous children as it is now the current children
				previousChildren = previousChildren[0:(len(previousChildren)-1)]

				if(xPosForCurrentDepthChildren[depth] < len(currentChildren)){
					currentChild = currentChildren[xPosForCurrentDepthChildren[depth]]
				}
				
				
			}



		}


	}



}



func GetStringForCodeOfCP(code float64) string {

	switch code{
		case float64(1):
			return "+"
		case float64(2):
			return "-"
		case float64(3):
			return "*"
		case float64(4):
			return "/"
		//the last number for parenthesis has no operator, hence the 5
		case float64(5):
			return ""	
		default:
			panic("unkown operator code GetStringForCodeOfCP()")
	}

	return "error"

}


func RemoveOperatorsBetweenTwoClosingParenthesisAndRemoveSpaces(equationStringInput string) string {

	doneWorking := false

	cursor := 0

	equationString := equationStringInput

	equationString = strings.ReplaceAll(equationString, " ", "")

	threeAtATime := []rune{}

	for !doneWorking {

		if(cursor + 2 >= len(equationString)){
			doneWorking = true
			break
		}

		threeAtATime = []rune{rune(equationString[cursor]), rune(equationString[cursor+1]), rune(equationString[cursor+2])}

		if(threeAtATime[0] == ')' && (threeAtATime[1] == '+' || threeAtATime[1] == '-' || threeAtATime[1] == '*' || threeAtATime[1] == '/') && threeAtATime[2] == ')'){
			equationString = equationString[0:cursor+1] + equationString[cursor+2:len(equationString)]
		}else{
			cursor++
		}




	}

	if(rune(equationString[len(equationString)-1]) == '*' || rune(equationString[len(equationString)-1]) == '+' || rune(equationString[len(equationString)-1]) == '-' || rune(equationString[len(equationString)-1]) == '/'   ){
		equationString = equationString[0:len(equationString)-1]
	}

	return equationString

}




func Create2DEquationFromSliceInputs(inputSlices ...interface{}) [][]complex128 {

	returnEquation := [][]complex128{}


	for i := 0; i < len(inputSlices); i++ {

		currentInputSlice := inputSlices[i]

		reflectType := reflect.TypeOf(currentInputSlice)



		switch reflectType.Kind() {
			case reflect.Slice:
				elementType := reflectType.Elem()
				switch elementType.Kind(){
					case reflect.Slice:
					
						valueOf2DSlice := reflect.ValueOf(currentInputSlice)

						//REPLACETYPECASE
						typecomplex128 := reflect.TypeOf([][]complex128{})

						firstConversion2DSlice := valueOf2DSlice.Convert(typecomplex128)

						//REPLACETYPECASE
						finalConversion2DSlice := firstConversion2DSlice.Interface().([][]complex128)

						returnEquation = append(returnEquation, finalConversion2DSlice...)

						// for i := 0; i < len(finalConversion2DSlice); i++ {

						// 	firstElementInner := finalConversion2DSlice[i]

						// 	interfaceSliceInner := []interface{}{}

						// 	for j := 0; j < len(firstElementInner); j++ {
						// 		interfaceSliceInner = append(interfaceSliceInner, firstElementInner[j])
						// 	}

						// 	inputOptions = append(inputOptions, interfaceSliceInner)
						// }

					case reflect.Complex128:
						
						valueOf2DSlice := reflect.ValueOf(currentInputSlice)

						//REPLACETYPECASE
						typecomplex128 := reflect.TypeOf([]complex128{})

						firstConversion2DSlice := valueOf2DSlice.Convert(typecomplex128)

						//REPLACETYPECASE
						finalConversion2DSlice := firstConversion2DSlice.Interface().([]complex128)


						returnEquation = append(returnEquation, finalConversion2DSlice)						

					default:
						panic("error, not a valid 2D slice, outer type is a slice, but inner type is not")	
			}
		default:
			panic("error, not a valid 2D slice, outermost type is not of any slice")	
			
	}

	}

	return returnEquation

}



func TwoAdjacentNumbersCanAddOrSubtract(number1 []complex128, number2 []complex128) bool {

	//check the exponents match
	if(number1[1] == number2[1]){
		return true
	}else{
		 return false
	}


}

func AddTwoAdjacentNumbers(number1 []complex128, number2 []complex128) []complex128 {

	if(!TwoAdjacentNumbersCanAddOrSubtract(number1, number2)){
		panic("error invalid add")
	}


	if(number1[0]+number2[0] == 0){
		return []complex128{complex(0, 0), complex(0, 0)}
	}else{
		return []complex128{number1[0]+number2[0], number1[1]}
	}

}

func SubtractTwoAdjacentNumbers(number1 []complex128, number2 []complex128) []complex128 {

	if(!TwoAdjacentNumbersCanAddOrSubtract(number1, number2)){
		panic("error invalid add")
	}


	if(number1[0]-number2[0] == 0){
		return []complex128{complex(0, 0), complex(0, 0)}
	}else{
		return []complex128{number1[0]-number2[0], number1[1]}
	}
}


func MultiplyTwoAdjacentNumbers(number1 []complex128, number2 []complex128) []complex128 {


	
	//if the exponents add to 0 the result is 1
	if(number1[1] + number2[1] == 0){
		return []complex128{complex(1, 0), complex(0,0)}
	}else{
		return []complex128{number1[0]*number2[0], number1[1]+number2[1]} 
	}


}

func DivideTwoAdjacentNumbers(number1 []complex128, number2 []complex128) []complex128 {

	
	//if the exponents subtract to 0 the result is 1
	if(number1[1] - number2[1] == 0){
		return []complex128{complex(1, 0), complex(0,0)}
	}else{
		return []complex128{number1[0]/number2[0], number1[1]-number2[1]} 
	}


}


//returns the new and improved or returns exactly the same as input depending if anything was removed
//this function so far only works for formats such as (s^2 + s + 2) nothing complicated like ((s+1)^2 - (3s^2))
func FactorNumeratorAndDenonminatorRemoveLikeFactors(numeratorInput [][]complex128, denominatorInput [][]complex128) ([]Factor, []Factor) {

	
	numerator := TryAllABCAndACOnlyFactorMethodsOnEquation(numeratorInput)

	denominator := TryAllABCAndACOnlyFactorMethodsOnEquation(denominatorInput)

	fmt.Println("numerator", DecodeFloatSliceToEquation(numerator))
	fmt.Println("denominator", DecodeFloatSliceToEquation(denominator))

	// var firstFactorNumerator [][]complex128
	// var endIndexFactorNumerator int

	factorsNumerator := []Factor{}

	numeratorFactor, endIndexFactorNumerator := GetFirstFactorFromEquation(numerator)
	//secondFactorNumerator, _ :=  GetFirstFactorFromEquation(numerator[endIndexFactorNumerator+1:len(numerator)])



	//while the numerator still has factors keep appending them to the factors slice
	for len(numeratorFactor) != 0 {

		factorsNumerator = append(factorsNumerator, Factor{numeratorFactor})

		fmt.Println("end index", endIndexFactorNumerator)

		numerator = numerator[endIndexFactorNumerator+1:len(numerator)]

		numeratorFactor, endIndexFactorNumerator = GetFirstFactorFromEquation(numerator)

	}


	for i := 0; i < len(factorsNumerator); i++ {
		fmt.Println("numerator factor ", i, " ", DecodeFloatSliceToEquation(factorsNumerator[i].Data) )
	}




	factorsDenominator := []Factor{}

	denominatorFactor, endIndexFactorDenominator := GetFirstFactorFromEquation(denominator)
	//secondFactorDenominator, _ :=  GetFirstFactorFromEquation(Denominator[endIndexFactorDenominator+1:len(Denominator)])



	//while the Denominator still has factors keep appending them to the factors slice
	for len(denominatorFactor) != 0 {

		factorsDenominator = append(factorsDenominator, Factor{denominatorFactor})

		fmt.Println("end index", endIndexFactorDenominator)

		denominator = denominator[endIndexFactorDenominator+1:len(denominator)]

		denominatorFactor, endIndexFactorDenominator = GetFirstFactorFromEquation(denominator)

	}


	for i := 0; i < len(factorsDenominator); i++ {
		fmt.Println("Denominator factor ", i, " ", DecodeFloatSliceToEquation(factorsDenominator[i].Data) )
	}

	restrictedIndicesNumeratorsFactors := []int{}

	restrictedIndicesDenominatorFactors := []int{}

	//each inner slice is the i and j pair for matching factors
	numeratorDenominatorIndicesForMatch := [][]int{}

	for i := 0; i < len(factorsNumerator); i++ {
		if(!(IsRestrictedIndex(i, restrictedIndicesNumeratorsFactors))){
			factorNumerator := factorsNumerator[i]
		

			for j := 0; j < len(factorsDenominator); j++ {
				if(!(IsRestrictedIndex(j, restrictedIndicesDenominatorFactors))){
					factorDenominator := factorsDenominator[j]

					if(TwoFactorsAreEqual(factorNumerator, factorDenominator)){
						fmt.Println(i, j, " two factors equal", DecodeFloatSliceToEquation(factorNumerator.Data), DecodeFloatSliceToEquation(factorDenominator.Data))
						restrictedIndicesNumeratorsFactors = append(restrictedIndicesNumeratorsFactors, i)
						restrictedIndicesDenominatorFactors = append(restrictedIndicesDenominatorFactors, j)
						numeratorDenominatorIndicesForMatch = append(numeratorDenominatorIndicesForMatch, []int{i, j})
						break
					}


				}

			}

		}

	}

	newNumeratorFactors := []Factor{}

	newDenominatorFactors := []Factor{}

	for i := 0; i < len(numeratorDenominatorIndicesForMatch); i++ {

		numeratorIndex := numeratorDenominatorIndicesForMatch[i][0]

		numeratorItemData := factorsNumerator[numeratorIndex].Data

		denominatorIndex := numeratorDenominatorIndicesForMatch[i][1]

		denominatorItemData := factorsDenominator[denominatorIndex].Data


		//get the power of the closing parenthesis for the factor for the numerator
		powerParenthesisNumerator := numeratorItemData[len(numeratorItemData)-1][2]

		//get the power of the closing parenthesis for the factor for the denominator
		powerParenthesisDenominator := denominatorItemData[len(numeratorItemData)-1][2]

		fmt.Println("numerator factor", DecodeFloatSliceToEquation(numeratorItemData), "numerator factor power", powerParenthesisNumerator)

		fmt.Println("denominator factor", DecodeFloatSliceToEquation(denominatorItemData), "denominator factor power", powerParenthesisDenominator)

		numeratorMinusDenominator := powerParenthesisNumerator - powerParenthesisDenominator

		if(real(numeratorMinusDenominator) > 0){

			
			//set the power of the numerators closing parenthesis equal to the difference
			numeratorItemData[len(numeratorItemData)-1][2] = numeratorMinusDenominator

			newNumeratorFactors = append(newNumeratorFactors, Factor{numeratorItemData})

			// //remove the term from the denominator
			// factorsDenominator = append(factorsDenominator[0:denominatorIndex], factorsDenominator[(denominatorIndex+1):len(factorsDenominator)]


		}else if (real(numeratorMinusDenominator) < 0){

			//set the power of the denominator closing parenthesis equal to the difference
			//since the difference is negative but the denominator already implies negative
			//reset the sign to return accurate data
			denominatorItemData[len(numeratorItemData)-1][2] = (-1*numeratorMinusDenominator)

			newDenominatorFactors = append(newDenominatorFactors, Factor{denominatorItemData})



			//remove the term from the numerator
			// factorsNumerator = append(factorsNumerator[0:numeratorIndex], factorNumerator[(numeratorIndex+1):len(factorNumerator)]


		}else if(real(numeratorMinusDenominator) == 0){

			newNumeratorFactors = append(newNumeratorFactors, Factor{Create2DEquationFromSliceInputs(gOP(), gNum(1, 0), gCP(1, 1))} )
			newDenominatorFactors = append(newDenominatorFactors, Factor{Create2DEquationFromSliceInputs(gOP(), gNum(1, 0), gCP(1, 1))})

		}else{
			panic("there isn't a 4th case... FactorNumeratorAndDenonminatorRemoveLikeFactors")
		}

	}

	


	//for terms that did not have any matches, they need to be added back into the final result
	//they were located at the non restricted indices
	for i := 0; i < len(factorsNumerator); i++ {

		fmt.Println("num factor",  DecodeFloatSliceToEquation(factorsNumerator[i].Data))

		if(!IsRestrictedIndex(i, restrictedIndicesNumeratorsFactors)){
			newNumeratorFactors = append(newNumeratorFactors, factorsNumerator[i])
		}

	}

	for i := 0; i < len(factorsDenominator); i++ {

		fmt.Println("denom factor",  DecodeFloatSliceToEquation(factorsNumerator[i].Data))

		if(!IsRestrictedIndex(i, restrictedIndicesDenominatorFactors)){
			newDenominatorFactors = append(newDenominatorFactors, factorsDenominator[i])
		}

	}

	// for i := 0; i < len(newNumeratorFactors); i++ {
	// 	fmt.Println("New Numerator factor ", i, " ", DecodeFloatSliceToEquation(newNumeratorFactors[i].Data) )
	// }

	// for i := 0; i < len(newDenominatorFactors); i++ {
	// 	fmt.Println("New Denominator factor ", i, " ", DecodeFloatSliceToEquation(newDenominatorFactors[i].Data) )
	// }



	return newNumeratorFactors, newDenominatorFactors


}

//the reason only ABC and AB methods are tried is they don't yield a complete the square odd 
//factor
func TryAllABCAndACOnlyFactorMethodsOnEquation(equationInput [][]complex128) [][]complex128 {



	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

	equationCopy := CleanCopyEntire2Dcomplex128Slice(equationInput)

	equation = FactorQuadraticsWithABCAllPresent(equation)

	if(!TwoEquationsAreExactlyIdentical(equation, equationCopy)){
		fmt.Println("abc factorization TryAllABCAndABOnlyFactorMethodsOnEquation()")
		return equation
	}


	equation = FactorQuadraticsWithACOnlyPresent(equation)

	if(!TwoEquationsAreExactlyIdentical(equation, equationCopy)){
		
		fmt.Println("ac only factorization TryAllABCAndABOnlyFactorMethodsOnEquation()")
		return equation
	}


	


	return equation




}

func InnerParenthesisCanBeSimplifiedFurther(numbersHolder [][]complex128, operatorsHolder [][]complex128) bool {

	fmt.Println("inner parenthesis values", numbersHolder, "inner parenthesis operators", operatorsHolder)



	for i := 0; i < len(operatorsHolder); i++ {

		if(i <= len(operatorsHolder) -1){		
			leftNum := numbersHolder[i]
			rightNum := numbersHolder[i+ 1]


			

			if(real(operatorsHolder[i][1]) == 3 || real(operatorsHolder[i][1]) == 4){
				fmt.Println("true")
				return true
			}else{
				if(TwoAdjacentNumbersCanAddOrSubtract(leftNum, rightNum)){
					fmt.Println("true")
					return true
				}
			}
		}

	}

	fmt.Println("false")
	return false




}


func CleanCopyEntire1Dcomplex128Slice(sliceToCopy []complex128) []complex128 {

	returnCopy := make([]complex128, len(sliceToCopy))

	itemsCopied := copy(returnCopy, sliceToCopy)

	if(itemsCopied != len(sliceToCopy)){
		panic("error copying slice CleanCopyEntire1Dcomplex128Slice()")
	}else{
		return returnCopy
	}

}

//returns the factor and the ending index of where the factor ended
func GetFirstFactorFromEquation(equation [][]complex128) ([][]complex128, int) {

	endingIndex := 0

	foundValid := false

	numbersHolder := [][]complex128{}		

	for i := 0; i < len(equation); i ++ {

		if(foundValid){
			break
		}

		if(IsOP(equation[i][0], equation[i][1])){



			checkingIfValid := true

			sawOneNumber := false

			//set this to null before each attempt to not have lingering data
			numbersHolder = [][]complex128{}


			numbersHolder = append(numbersHolder, equation[i])

			fmt.Println("numbers holder", numbersHolder)

			//these two bools are used to make sure
			//numbers and symbols alternate
			indexShouldBeNumber := true
			indexShouldBeOperator := false

			cursor := i

			
			

			for checkingIfValid {

				cursor++

				//cursor is out of bounds, nothing to check
				if(cursor >= len(equation)){
					return equation, 0
				}

				if(!sawOneNumber){

					for(!IsCP(equation[cursor]) && !IsOP(equation[cursor][0], equation[cursor][1])){


						if(cursor >= len(equation)){
							checkingIfValid = false
							foundValid = false
							break
						}

						if(IsNumber(equation[cursor][0]) && indexShouldBeNumber){
							numbersHolder = append(numbersHolder, equation[cursor])
							indexShouldBeNumber = false
							indexShouldBeOperator = true
						}else if(IsOperator(equation[cursor]) && indexShouldBeOperator){
							numbersHolder = append(numbersHolder, equation[cursor])
							indexShouldBeNumber = true
							indexShouldBeOperator = false
						}else{
							checkingIfValid = false
							foundValid = false
							break	
						}

						cursor++

					}

					//make sure what broke the loop was a closing parenthesis
					if(IsCP(equation[cursor])){
						endingIndex = cursor
						numbersHolder = append(numbersHolder, equation[cursor])
						checkingIfValid = false
						foundValid = true
						break
					}else{
						checkingIfValid = false
						foundValid = false
						break	
					}

			}

		}


	}
}

	fmt.Println("numbers holder return ", numbersHolder)

	//if its only the opening parenthesis null it
	if(len(numbersHolder) == 1){
		numbersHolder = [][]complex128{}
	}

	return numbersHolder, endingIndex

}




func TwoFactorsAreEqual(factor1 Factor, factor2 Factor) bool {

	factor1Data := factor1.Data
	factor2Data := factor2.Data

	//if the outer lenght of the whole 2d slice differs return false
	if(len(factor1Data) != len(factor2Data)){
		return false
	}

	for i := 0 ; i <len(factor1Data); i++ {
		//the closing parenthesis should not be counted for factors,
		//factors are equal even if the power of their closing parethesis differs
		if(i != len(factor1Data)-1){
			currentItemFactor1 := factor1Data[i]

			currentItemFactor2 := factor2Data[i]

			//if the inner length of two 1d slices at the same index differs return false
			if(len(currentItemFactor1) != len(currentItemFactor2)){
				return false
			}

			for j := 0; j < len(currentItemFactor1); j++ {

				fmt.Println("compare items", currentItemFactor1, currentItemFactor2)

				
					if(currentItemFactor1[j] != currentItemFactor2[j]){
						return false
					}
				

			}

		}


	}

	return true


}





func IsRestrictedIndex(indexToCheck int, restrictedIndices []int ) bool {

	for i := 0; i < len(restrictedIndices); i++ {
		if(restrictedIndices[i] == indexToCheck){
			return true
		}
	}

	return false

}

//deleting this...
func GatherFactorsInSeriesThatMultiplyOrDivideEachOther(equationInput [][]complex128) ([][]Factor, [][]float64) {

	equation := CleanCopyEntire2Dcomplex128Slice(equationInput)

	equationFactor, endIndex := GetFirstFactorFromEquation(equation)

	factors := []Factor{}

	for len(equationFactor) > 0 {

		factors = append(factors, Factor{equationFactor})

		equation = equation[(endIndex+1):len(equation)]

		equationFactor, endIndex = GetFirstFactorFromEquation(equation)

	}

	operatorsForEachFactor := []float64{}


	for i := 0; i < len(factors); i++ {

		operatorsForEachFactor = append(operatorsForEachFactor, GetOperatorAppendedToClosingParenthesisOfFactor(factors[i]))

	}

	if(len(factors) != len(operatorsForEachFactor)){
		panic("not one operator per factor GatherFactorsInSeriesThatMultiplyOrDivideEachOther()")
	}

	factorsAndOperatorsCursor := 0

	factorsInSeriesWithEachOther := [][]Factor{}

	operatorsForEachSeries := [][]float64{}

	for factorsAndOperatorsCursor < len(factors) {

		currentSeries := []Factor{}

		currentSeriesOperators := []float64{}

		inSeries := true

		for inSeries {

			currentSeries = append(currentSeries, factors[factorsAndOperatorsCursor])

			currentOperator := operatorsForEachFactor[factorsAndOperatorsCursor]

			factorsAndOperatorsCursor++

			fmt.Println("decoded series", DecodeFloatSliceToEquation(factors[factorsAndOperatorsCursor].Data))

			if(currentOperator == 3 || currentOperator == 4){
				inSeries = true
				currentSeriesOperators = append(currentSeriesOperators, currentOperator)
			}else {
				inSeries = false
			}

			if(factorsAndOperatorsCursor >= len(factors)){
				inSeries = false
			}
		}

		fmt.Println(currentSeries, currentSeriesOperators)

		//there should always be one less operator than number in the seris
		if((len(currentSeries) - len(currentSeriesOperators)) != 1){
			panic("mismatch in numbers of items in series and numbers of operators for series")
		}

		factorsInSeriesWithEachOther = append(factorsInSeriesWithEachOther, currentSeries)

		operatorsForEachSeries = append(operatorsForEachSeries, currentSeriesOperators)

	}



	if(len(operatorsForEachSeries) != len(factorsInSeriesWithEachOther)){
		panic("not one row of operators per row of factors GatherFactorsInSeriesThatMultiplyOrDivideEachOther()")
	}

	//run every factor through simplifying its inner parenthesis, and foiling if the exponent is greater than 1
	//for the closing parenthesis
	for i := 0; i < len(factorsInSeriesWithEachOther); i++ {

		currentFactors := factorsInSeriesWithEachOther[i]

		for j := 0; j < len(currentFactors); j++ {

			factorsInSeriesWithEachOther[i][j] = Factor{SimplifyInnerParenthesis(factorsInSeriesWithEachOther[i][j].Data)}

			factorsInSeriesWithEachOther[i][j] = Factor{FoilOutParenthesisRaisedToExponent(factorsInSeriesWithEachOther[i][j].Data)}

		}

	}

	newFactorsInSeries := [][]Factor{}

	newOperatorsForEachSeries := [][]float64{}

	//here terms multiplying  can now be combined since everything has been foiled out
	//index to second to last item because each iteration gets current and next item
	//this also makes sure when indexing the operators slice it does not go out of bounds
	for i := 0; i < (len(factorsInSeriesWithEachOther)); i++ {

		currentFactors := factorsInSeriesWithEachOther[i]
		currentOperators := operatorsForEachSeries[i]

		newFactorsForRow := []Factor{}

		newOperatorsForRow := []float64{}

		//here terms multiplying  can now be combined since everything has been foiled out
		//index to second to last item because each iteration gets current and next item
		//this also makes sure when indexing the operators slice it does not go out of bounds
		for j := 0; j < (len(currentFactors)-1); j++ {

			if(currentOperators[i] == 3){
				newFactorsForRow = append(newFactorsForRow, Factor{MultiplyNeighboringParenthesis(currentFactors[j].Data, currentFactors[j+1].Data)})
				//iterate j once more so it skips the next index that was combined
				j++
			}else{
				newOperatorsForRow = append(newOperatorsForRow, 4)
				newFactorsForRow = append(newFactorsForRow, currentFactors[j])
			}

		}

		newFactorsInSeries  = append(newFactorsInSeries, newFactorsForRow)

		newOperatorsForEachSeries  = append(newOperatorsForEachSeries, newOperatorsForRow)

	}





	return newFactorsInSeries, newOperatorsForEachSeries

}


func SeperateBaseEquationIntoFractions(equationInput [][]complex128) []Fraction {


	return []Fraction{}

}

func GetPowerOfClosingParenthesisOfFactor(factor Factor) complex128 {

	return factor.Data[len(factor.Data)-1][2]	
}

func GetOperatorAppendedToClosingParenthesisOfFactor(factor Factor) float64 {

	return real(factor.Data[len(factor.Data)-1][4])

}



//returns a new numbers and operators holder for function SimplifyInnerParenthesi()
func SortParenthesisContainingOnlyPlusAndMinusBySExponent(numbersHolder [][]complex128, operatorsHolder [][]complex128) ([][]complex128, [][]complex128) {

	//double check there are no mutliplication or division operators present
	for i := 0; i < len(operatorsHolder); i++ {

		if(real(operatorsHolder[i][1]) == 3 || real(operatorsHolder[i][1]) == 4){
			panic("should not be any * or / operators SortParenthesisContainingOnlyPlusAndMinusBySExponent()")
		}

	}

	sExponentHolder := []float64{}

	for i := 0; i < len(numbersHolder); i++ {
		sExponentHolder = append(sExponentHolder, real(numbersHolder[i][1]))
	}


	//start at index 1 because each iteration also needs the previous item
	//starting at 0 would cause error
	sExponentCursor := 1

	oneSortOccurred := false

	needToSortAgain := true


	//since there are less operators than numbers, prepend a plus or minus for the first number depending
	//on its sign since it could get swappped and needs an operator to swap with it, 
	//at the end of the swapping the first operator is removed, and the sign of the number for
	//that sign is altered accordingly
	firstNumberPositive := (real(numbersHolder[0][0])) >= 0

	if(firstNumberPositive){
		operatorsHolder = append([][]complex128{[]complex128{complex(0, 0), complex(1, 0)}}, operatorsHolder...)
	}else{
		operatorsHolder = append([][]complex128{[]complex128{complex(0, 0), complex(2, 0)}}, operatorsHolder...)
	}


	for needToSortAgain {

			fmt.Println(sExponentHolder)
			fmt.Println(numbersHolder)
			fmt.Println(operatorsHolder)

			oneSortOccurred = false

			for sExponentCursor < len(sExponentHolder) {
				sPowerCurrent := sExponentHolder[sExponentCursor]
				sPowerPrevious := sExponentHolder[sExponentCursor-1]
			
				fmt.Println("previous power", sPowerPrevious, "current power", sPowerCurrent)

				if(sPowerCurrent > sPowerPrevious){
					sort.Float64Slice(sExponentHolder).Swap(sExponentCursor, sExponentCursor-1)
					
					//no swap function for type [][]complex128 so do it manually
					copyNumberSlice1 := numbersHolder[sExponentCursor]
					copyNumberSlice2 := numbersHolder[sExponentCursor-1]
					numbersHolder[sExponentCursor] = copyNumberSlice2
					numbersHolder[sExponentCursor-1] = copyNumberSlice1

					//no swap function for type [][]complex128 so do it manually
					copyOperatorSlice1 := operatorsHolder[sExponentCursor]
					copyOperatorSlice2 := operatorsHolder[sExponentCursor-1]
					operatorsHolder[sExponentCursor] = copyOperatorSlice2
					operatorsHolder[sExponentCursor-1] = copyOperatorSlice1

					oneSortOccurred = true
				}

				sExponentCursor++

				
			}

			sExponentCursor = 1

			// fmt.Println("stuck")

			if(oneSortOccurred){
				needToSortAgain = true
			}else{
				needToSortAgain = false
				
			}


	}

	firstOperatorIsPlusSign := (real(operatorsHolder[0][1]) == 1)

	//if it is a negative sign prepended to the first element
	//multiply by -1 to adjust for the removal, if its a plus sign
	//then nothing needs to happen
	if(!firstOperatorIsPlusSign){
		numbersHolder[0][0] = numbersHolder[0][0] * -1 
	}

	fmt.Println("operators holder", operatorsHolder)

	operatorsHolder = operatorsHolder[1:len(operatorsHolder)]

	fmt.Println("operators holder", operatorsHolder)

	return numbersHolder, operatorsHolder

}

//if true the return slice is only one long and is the combined fraction
//if false the return slice is 2 long and contains the original fractions
func CheckIfTwoFractionsCanCombineAndCombineIfSo(fraction1 Fraction, fraction2 Fraction) ([]Fraction) {

	originalFractions := []Fraction{fraction1, fraction2}

	if(len(fraction1.Denominator) != len(fraction2.Denominator)){
		return originalFractions
	}

	for i := 0; i < len(fraction1.Denominator); i++ {

		fraction1CurrentDenominatorData := fraction1.Denominator[i].Data
		fraction2CurrentDenominatorData := fraction2.Denominator[i].Data

		if(!TwoEquationsAreExactlyIdentical(fraction1CurrentDenominatorData, fraction2CurrentDenominatorData)){
			return originalFractions
		}

	}

	fmt.Println("fractions can combine!")

	return originalFractions

}


