package main

import (

	"fmt"
	"strconv"
)


//[0, 0] = (
//[0, 1] = )
//[0, 2] = +
//[0, 3] = -
//[0, 4] = *
//[0, 5] = /
	

//TODO, when multiple variables get involved a third index needs to be added to the float slice
//which will allow the third index to essentially span the alphabet 0-25 A-Z for variable names
//for now since this is only being used for inverse laplace transform of one variable, everything is
//assumed to be 's'	


func main() {
	

	equation := [][]float64{gOP(), gOP(), gNumsAdded(2, 3, 3, 0), gCP(), gMultiply(), gOP(), gOP(), gNumsAdded(3, 2, 1, 1, 3, 0), gCP(), gMultiply(), gOP(), gNumsAdded(2, 1, 1, 0), gCP(), gCP(), gCP()}       

	fmt.Println(DecodeFloatSliceToEquation(equation))

}



func DecodeFloatSliceToEquation(equation [][]float64 ) string {

	CheckEquationForSyntaxErrors(equation)

	equationString := ""

	depthLevel := 0

	for i := 0; i < len(equation); i++ {
		
		currentItem := equation[i]

		for j := 0; j < len(currentItem); j = (j+2) {
			firstIndex := currentItem[j]
			secondIndex := currentItem[j+1]

			firstIndexString := strconv.FormatFloat(firstIndex, 'f', -1, 64)

			secondIndexString := strconv.FormatFloat(secondIndex, 'f', -1, 64)


			if(IsOP(firstIndex, secondIndex)){
				equationString += "( "
				depthLevel++ 
			}else if(IsCP(firstIndex, secondIndex)){
				equationString += " )"
				depthLevel--
			}else if(IsPlus(firstIndex, secondIndex)){
				equationString += " + "
			}else if(IsMinus(firstIndex, secondIndex)){
				equationString += " - "
			}else if(IsMultiply(firstIndex, secondIndex)){
				equationString += " * "
			}else if(IsDivide(firstIndex, secondIndex)){
				equationString += " / "
			}else if(IsNumber(firstIndex)){
				if(firstIndex != 1){
					if(secondIndex == 0){
						equationString += firstIndexString + " "
					}else if(secondIndex == 1){
						equationString += firstIndexString + "S "
					}else{
						equationString += firstIndexString + "S^" + secondIndexString + " "
					}
				}else{
					if(secondIndex == 0){
						equationString += firstIndexString + " "
					}else if(secondIndex == 1){
						equationString += "S "
					}else{
						equationString += "S^" + secondIndexString + " "
					}
				}
				
			}else{
				panic("unknown equation item DecodeFloatSliceToEquation()")	
			}
			fmt.Println(equationString)
			fmt.Println(depthLevel)

		}

	}

	return equationString

}

//g stands for generate

func gNum(nums ...float64) []float64 {

	if( (len(nums)%2) != 0){
		panic("error, invalid amount of numbers")
	}

	returnSlice := []float64{}

	for i := 0; i < len(nums); i = (i + 2) {



		returnSlice = append(returnSlice, nums[i])
		returnSlice = append(returnSlice, nums[i+1])


	}

	return returnSlice

}

func gNumsAdded(nums ...float64) []float64 {

	if( (len(nums)%2) != 0){
		panic("error, invalid amount of numbers")
	}

	returnSlice := []float64{}

	for i := 0; i < len(nums); i = (i + 2) {

		

		returnSlice = append(returnSlice, nums[i])
		returnSlice = append(returnSlice, nums[i+1])
		if(!((i+2) >= len(nums))){
			returnSlice = append(returnSlice, gPlus()...)
		}

	}

	return returnSlice

}


func gNumsSubtracted(nums ...float64) []float64 {

	if( (len(nums)%2) != 0){
		panic("error, invalid amount of numbers")
	}

	returnSlice := []float64{}

	for i := 0; i < len(nums); i = (i + 2) {

		

		returnSlice = append(returnSlice, nums[i])
		returnSlice = append(returnSlice, nums[i+1])

		if(!((i+2) >= len(nums))){
			returnSlice = append(returnSlice, gMinus()...)
		}

	}

	return returnSlice

}


func gNumsMultiplied(nums ...float64) []float64 {

	if( (len(nums)%2) != 0){
		panic("error, invalid amount of numbers")
	}

	returnSlice := []float64{}

	for i := 0; i < len(nums); i = (i + 2) {

		

		returnSlice = append(returnSlice, nums[i])
		returnSlice = append(returnSlice, nums[i+1])

		if(!((i+2) >= len(nums))){
			returnSlice = append(returnSlice, gMultiply()...)
		}

	}

	return returnSlice

}
func gNumsDivided(nums ...float64) []float64 {

	if( (len(nums)%2) != 0){
		panic("error, invalid amount of numbers")
	}

	returnSlice := []float64{}

	for i := 0; i < len(nums); i = (i + 2) {

		

		returnSlice = append(returnSlice, nums[i])
		returnSlice = append(returnSlice, nums[i+1])

		if(!((i+2) >= len(nums))){
			returnSlice = append(returnSlice, gDivide()...)
		}

	}

	return returnSlice

}
func gOP() []float64 {
	return []float64{0, 0}
}

func gCP() []float64 {
	return []float64{0, 1}
}

func gPlus() []float64 {
	return []float64{0, 2}
}

func gMinus() []float64 {
	return []float64{0, 3}
}

func gMultiply() []float64 {
	return []float64{0, 4}
}

func gDivide() []float64 {
	return []float64{0, 5}
}


// // //for two parenthesis to be eligible to foil..
// // //they need to be at the same depth level
// // //and the operator between them needs to be multiplication '*'

func FoilAllNeighboringParenthesis(equation [][]float64) [][]float64 {

	CheckEquationForSyntaxErrors(equation)

	depthLevel := 0

	foilStateMachine := 0

	foilStateMachinePrevious := 0

	numbersToFoilFirstTerm := []float64{}

	numbersToFoilSecondTerm := []float64{}

	previousTerm := []float64{}

	outerAndInnerMarkerForFoilStart := []int{}

	outerAndInnerMarkerForFoilEnd := []int{}

	foundFoil := false

	for i := 0; i < len(equation); i++ {

		if(foundFoil){
			break
		}

		currentItem := equation[i]

		for j := 0; j < len(currentItem); j = (j+2) {
			
			firstIndex := currentItem[j]
			secondIndex := currentItem[j+1]


			previousFirstIndex := previousTerm[0]
			previousSecondIndex := previousTerm[1]

			if(IsOP(firstIndex, secondIndex)){
				depthLevel++ 
				if(foilStateMachine == 0){
					foilStateMachine = 1
					outerAndInnerMarkerForFoilStart = []int{i, j}					
				}
				
				
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsCP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsPlus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMinus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMultiply(previousFirstIndex, previousSecondIndex)){
					if(foilStateMachine == 5){
						foilStateMachine = 6
					}
				}else if(IsDivide(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsNumber(previousFirstIndex)){
					//-1 is given for the previous term second index of the first
					//cycle since there is no previous term 
					//this is the only place in all if else needed to check
					//since first term is always opening parenthesis
					if(firstIndex == 0 && secondIndex != -1){

					}	
				}else{
					panic("unknown equation item FoilAllNeighboringParenthesis()")	
				}
			}else if(IsCP(firstIndex, secondIndex)){
				depthLevel--
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsCP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsPlus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMinus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMultiply(previousFirstIndex, previousSecondIndex)){

				}else if(IsDivide(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsNumber(previousFirstIndex)){
					if(foilStateMachine == 2){
						foilStateMachine = 4
					}
					if(foilStateMachine == 7){
						foilStateMachine = 9
						outerAndInnerMarkerForFoilEnd = []int{i, j}
					}
				}else{
					panic("unknown equation item FoilAllNeighboringParenthesis()")	
				}						
			}else if(IsPlus(firstIndex, secondIndex)){
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsCP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsPlus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMinus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMultiply(previousFirstIndex, previousSecondIndex)){

				}else if(IsDivide(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsNumber(previousFirstIndex)){

				}else{
					panic("unknown equation item FoilAllNeighboringParenthesis()")	
				}						
			
			}else if(IsMinus(firstIndex, secondIndex)){
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsCP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsPlus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMinus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMultiply(previousFirstIndex, previousSecondIndex)){

				}else if(IsDivide(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsNumber(previousFirstIndex)){

				}else{
					panic("unknown equation item FoilAllNeighboringParenthesis()")	
				}						
			
			}else if(IsMultiply(firstIndex, secondIndex)){
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsCP(previousFirstIndex, previousSecondIndex)){
					if(foilStateMachine == 4){
						foilStateMachine = 5
					}
				}else if(IsPlus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMinus(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsMultiply(previousFirstIndex, previousSecondIndex)){

				}else if(IsDivide(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsNumber(previousFirstIndex)){

				}else{
					panic("unknown equation item FoilAllNeighboringParenthesis()")	
				}						
			
			}else if(IsDivide(firstIndex, secondIndex)){
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsCP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsPlus(previousFirstIndex, previousSecondIndex)){
					if(foilStateMachine == 2){
						foilStateMachine = 3
					}
					if(foilStateMachine == 7){
						foilStateMachine = 8
					}
				}else if(IsMinus(previousFirstIndex, previousSecondIndex)){
					if(foilStateMachine == 3){
						foilStateMachine = 2
					}
					if(foilStateMachine == 7){
						foilStateMachine = 8
					}
				}else if(IsMultiply(previousFirstIndex, previousSecondIndex)){

				}else if(IsDivide(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsNumber(previousFirstIndex)){

				}else{
					panic("unknown equation item FoilAllNeighboringParenthesis()")	
				}						
			
			}else if(IsNumber(firstIndex, secondIndex)){

				if(IsOP(previousFirstIndex, previousSecondIndex)){
					if(foilStateMachine == 1){
						foilStateMachine = 2
						numbersToFoilFirstTerm = append(numbersToFoilFirstTerm, g(firstIndex, secondIndex))
					}

					if(foilStateMachine == 6){
						foilStateMachine = 7
						numbersToFoilFirstTerm = append(numbersToFoilFirstTerm, g(firstIndex, secondIndex))
					}
				}else if(IsCP(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsPlus(previousFirstIndex, previousSecondIndex)){
					if(foilStateMachine == 3){
						numbersToFoilFirstTerm = append(numbersToFoilFirstTerm, g(firstIndex, secondIndex))
						foilStateMachine = 2
					}
					if(foilStateMachine == 8){
						foilStateMachine = 7
						numbersToFoilFirstTerm = append(numbersToFoilFirstTerm, g(firstIndex, secondIndex))
					}
				}else if(IsMinus(previousFirstIndex, previousSecondIndex)){
					if(foilStateMachine == 3){
						numbersToFoilFirstTerm = append(numbersToFoilFirstTerm, g(firstIndex * - 1, secondIndex * -1))
						foilStateMachine = 2
					}
					if(foilStateMachine == 8){
						foilStateMachine = 7
						numbersToFoilFirstTerm = append(numbersToFoilFirstTerm, g(firstIndex * -1 , secondIndex * -1))
					}
				}else if(IsMultiply(previousFirstIndex, previousSecondIndex)){

				}else if(IsDivide(previousFirstIndex, previousSecondIndex)){
					
				}else if(IsNumber(previousFirstIndex)){

				}else{
					panic("unknown equation item FoilAllNeighboringParenthesis()")	
				}												
			}else{
				panic("unknown equation item FoilAllNeighboringParenthesis()")	
			}						

			if(foilStateMachinePrevious == foilStateMachine){
				foilStateMachine = 0
				foilStateMachinePrevious = 0
			}else if(foilStateMachine == 9){
				foundFoil = true
				break

			}else{
				foilStateMachinePrevious = foilStateMachine
			}

			previousTerm = g(firstIndex, secondIndex)

		}

	}


	if(foundFoil){

		newSlice := []float64{}

		for i := 0; i < len(numbersToFoilFirstTerm); i++ {

			firstNum := numbersToFoilFirstTerm[i]

			for j := 0; j < len(numbersToFoilSecondTerm); j++ {

				secondNum := numbersToFoilSecondTerm[j]

				newNum := []float64{firstNum[0]*secondNum[0], firstNum[1]*secondNum[1]}

				

				newSlice = append(newSlice



			}

		}



	}


}




func IsOP(num1 float64, num2 float64) bool {
	if(num1 == 0 && num2 == 0){
		return true
	}else{
		return false
	}
}
func IsCP(num1 float64, num2 float64) bool {
	if(num1 == 0 && num2 == 1){
		return true
	}else{
		return false
	}
}
func IsPlus(num1 float64, num2 float64) bool {
	if(num1 == 0 && num2 == 2){
		return true
	}else{
		return false
	}
}
func IsMinus(num1 float64, num2 float64) bool {
	if(num1 == 0 && num2 == 3){
		return true
	}else{
		return false
	}
}
func IsMultiply(num1 float64, num2 float64) bool {
	if(num1 == 0 && num2 == 4){
		return true
	}else{
		return false
	}
}
func IsDivide(num1 float64, num2 float64) bool {
	if(num1 == 0 && num2 == 5){
		return true
	}else{
		return false
	}
}
func IsNumber(num1 float64) bool {
	if(num1 != 0){
		return true
	}else{
		return false
	}
}


func CheckEquationForSyntaxErrors(equation [][]float64) {

	depthLevel := 0

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
			}else if(IsCP(firstIndex, secondIndex)){
				depthLevel--
			}else if(IsPlus(firstIndex, secondIndex)){
				
			}else if(IsMinus(firstIndex, secondIndex)){
				
			}else if(IsMultiply(firstIndex, secondIndex)){
				
			}else if(IsDivide(firstIndex, secondIndex)){
				
			}else{
				
			}

			//this would occur if there's more closing parenthesis than opening
			if(depthLevel == -1){
				panic("Syntax Error too many ) for the number of ( CheckEquationForSyntaxErrors()")
			}

		}
	}

	if(depthLevel != 0){
		panic("Syntax Error not all ( items were closed properly CheckEquationForSyntaxErrors()")
	}

}











