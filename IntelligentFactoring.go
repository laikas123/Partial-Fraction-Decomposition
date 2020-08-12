package main

import (

	"fmt"
	"sort"
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
	

	equation := [][]float64{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(2),  gOP(), gOP(), gNum(3, 2, 1, 1, 3, 0), gCP(1), gOP(), gNum(2, 1, 1, 0), gCP(1), gCP(1), gCP(1)}       

	fmt.Println(DecodeFloatSliceToEquation(equation))

	foiledEquation := FoilAllNeighboringParenthesis(equation)

	fmt.Println(DecodeFloatSliceToEquation(foiledEquation))
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
			}else if(IsCP(currentItem)){
				equationString += " )"
				if(currentItem[j+2] != 0 ){
					equationString += "^" + strconv.FormatFloat(currentItem[2], 'f', -1, 64) + " "
				}
				depthLevel--
			}else if(IsNumber(firstIndex)){
				if(firstIndex != 1){
					if(secondIndex == 0){
						equationString += firstIndexString + " + "
					}else if(secondIndex == 1){
						equationString += firstIndexString + "S + "
					}else{
						equationString += firstIndexString + "S^" + secondIndexString + " + "
					}
				}else{
					if(secondIndex == 0){
						equationString += firstIndexString + " + "
					}else if(secondIndex == 1){
						equationString += "S + "
					}else{
						equationString += "S^" + secondIndexString + " +"
					}
				}
				
			}else{
				panic("unknown equation item DecodeFloatSliceToEquation()")	
			}
			fmt.Println(equationString)
			fmt.Println(depthLevel)

			if(IsCP(currentItem)){
				break	
			}

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

func gOP() []float64 {
	return []float64{0, 0}
}

func gCP(exponent float64) []float64 {
	return []float64{0, 1, exponent}
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

	previousTerm := []float64{0, -1}

	foilStart := -1

	foilEnd := -1

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

			fmt.Println("current item", currentItem)
			fmt.Println("previous item", previousTerm)

			if(IsOP(firstIndex, secondIndex)){
				

				fmt.Println("is OP")
				depthLevel++ 
				if(foilStateMachine == 0){
					foilStateMachine = 1
					foilStart = i					
				}
				
				
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					fmt.Println("previous is OP")
					if(foilStateMachine == 4){
						foilStateMachine = 1
						numbersToFoilFirstTerm = []float64{}
						numbersToFoilSecondTerm = []float64{}
						foilStart = i
					}
				}else if(IsCP(previousTerm)){
					fmt.Println("previous is CP")
					if(foilStateMachine == 3){
						foilStateMachine = 4
					}
					
				}else if(IsNumber(previousFirstIndex)){
					fmt.Println("previous is Numbers")
				}else{
					
					//-1 is given for the previous term second index of the first
					//cycle since there is no previous term 
					//this is the only place in all if else needed to check
					//since first term is always opening parenthesis
					if(previousFirstIndex == 0 && previousSecondIndex != -1){
						panic("unknown equation item FoilAllNeighboringParenthesis()")	
					}	
					
				}
			}else if(IsCP(currentItem)){

				fmt.Println("is CP")

				depthLevel--
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					fmt.Println("previous is OP")
				}else if(IsCP(previousTerm)){
					fmt.Println("previous is CP")
				}else if(IsNumber(previousFirstIndex)){
					fmt.Println("previous is Numbers")
					if(foilStateMachine == 2){
						foilStateMachine = 3
					}
					if(foilStateMachine == 5){
						foilStateMachine = 6
						foilEnd = i
						foundFoil = true
					}
				}else{
					fmt.Println(currentItem)
					panic("unknown equation item FoilAllNeighboringParenthesis()")	
				}						
			}else if(IsNumber(firstIndex)){

				fmt.Println("is Numbers")


				if(IsOP(previousFirstIndex, previousSecondIndex)){
					fmt.Println("previous is OP")
					if(foilStateMachine == 1){
						foilStateMachine = 2
						numbersToFoilFirstTerm = append(numbersToFoilFirstTerm, gNum(firstIndex, secondIndex)...)
					}
					if(foilStateMachine == 4){
						foilStateMachine = 5
						numbersToFoilSecondTerm = append(numbersToFoilSecondTerm, gNum(firstIndex, secondIndex)...)
					}
				}else if(IsCP(previousTerm)){
					fmt.Println("previous is CP")
				}else if(IsNumber(previousFirstIndex)){
					fmt.Println("previous is Numbers")
					if(foilStateMachine == 2){
						foilStateMachine = 2
						numbersToFoilFirstTerm = append(numbersToFoilFirstTerm, gNum(firstIndex, secondIndex)...)
					}
					if(foilStateMachine == 5){
						foilStateMachine = 5
						numbersToFoilSecondTerm = append(numbersToFoilSecondTerm, gNum(firstIndex, secondIndex)...)
					}
				}else{
					fmt.Println(currentItem)
					panic("unknown equation item FoilAllNeighboringParenthesis()")	
				}												
			}else{
				fmt.Println(currentItem)
				panic("unknown equation item FoilAllNeighboringParenthesis()")	
			}						


			fmt.Println("previous foil state machine", foilStateMachinePrevious)
			fmt.Println("current foil state machine", foilStateMachine)

			if( (foilStateMachinePrevious == foilStateMachine) && foilStateMachine != 1 && foilStateMachine != 2 && foilStateMachine != 5 && foilStateMachine != 0 && foilStateMachine != 4){
				foilStateMachine = 0
				foilStateMachinePrevious = 0
				numbersToFoilFirstTerm = []float64{}
				numbersToFoilSecondTerm = []float64{}
			}else if(foilStateMachine == 7){
				foundFoil = true
				break

			}else{
				foilStateMachinePrevious = foilStateMachine
			}

			previousTerm = currentItem

			if(IsCP(currentItem)){
				break
			}

		}

		
		


	}


	if(foundFoil){

		fmt.Println(foilStart, foilEnd)

		newSlice := []float64{}

		for i := 0; i < len(numbersToFoilFirstTerm); i = (i + 2) {

			firstNumMultiplier := numbersToFoilFirstTerm[i]

			firstNumExponent := numbersToFoilFirstTerm[i+1]

			for j := 0; j < len(numbersToFoilSecondTerm); j = (j + 2) {

				secondNumMultiplier := numbersToFoilSecondTerm[j]

				secondNumExponent := numbersToFoilSecondTerm[j+1]

				newNum := []float64{firstNumMultiplier*secondNumMultiplier, firstNumExponent+secondNumExponent}

				newSlice = append(newSlice, newNum...)



			}

		}

		fmt.Println("first terms and second terms")
		fmt.Println(numbersToFoilFirstTerm)
		fmt.Println(numbersToFoilSecondTerm)

		fmt.Println(newSlice)
		newSlice = SimplifyLikeTermsEquationSectionAndSortByDescendningExponent(newSlice)
		

		return Substitute1DSliceInto2DSliceStartAndEnd(foilStart, foilEnd, newSlice, equation)


	}else{
		return equation
	}


}


func SimplifyLikeTermsEquationSectionAndSortByDescendningExponent(equationSection []float64) []float64 {

	termsMap := make(map[float64][]float64)

	for i := 0; i < len(equationSection); i = (i + 2) {
		if(termsMap[equationSection[i+1]] == nil){
			termsMap[equationSection[i+1]] = []float64{equationSection[i]}
		}else{
			termsMap[equationSection[i+1]] = append(termsMap[equationSection[i+1]], equationSection[i])
		}
	}

	exponentsSlice := []float64{}

	for exponents, _ := range termsMap {
		exponentsSlice = append(exponentsSlice, exponents)
	}

	// sort.Float64s(exponentsSlice)

	// sort.Reverse(sort.Float64Slice(exponentsSlice))

	sort.Sort(sort.Reverse(sort.Float64Slice(exponentsSlice)))

	returnSlice := []float64{}

	for i := 0; i < len(exponentsSlice); i++ {

		currentMultipliers := termsMap[exponentsSlice[i]]

		mutlipliersAdded := float64(0)

		for j := 0; j < len(currentMultipliers); j++ {
			mutlipliersAdded += currentMultipliers[j]
		}

		returnSlice = append(returnSlice, mutlipliersAdded)
		returnSlice = append(returnSlice, exponentsSlice[i])


	}

	return returnSlice


}


func Substitute1DSliceInto2DSliceStartAndEnd(start int, end int, new1DSlice []float64, equation [][]float64) [][]float64{

	returnSlice := append(equation[0:start], new1DSlice)

	returnSlice = append(returnSlice, equation[(end + 1): len(equation)]...)

	return returnSlice

}

func IsOP(num1 float64, num2 float64) bool {
	if(num1 == 0 && num2 == 0){
		return true
	}else{
		return false
	}
}
func IsCP(nums []float64) bool {
	if(len(nums) == 3){
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











