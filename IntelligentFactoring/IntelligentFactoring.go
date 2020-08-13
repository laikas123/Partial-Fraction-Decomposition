package main

import (

	"fmt"
	"math/cmplx"
	"sort"
	"strings"
	"strconv"
)


//TODO, when multiple variables get involved a third index needs to be added to the float slice
//which will allow the third index to essentially span the alphabet 0-25 A-Z for variable names
//for now since this is only being used for inverse laplace transform of one variable, everything is
//assumed to be 's'	


func main() {
	

	equation := [][]complex128{gOP(), gNum(2, 3, 3, 0), gCP(3),  gOP(), gOP(), gNum(3, 2, 1, 1, 3, 0), gCP(1), gOP(), gNum(2, 1, 1, 0), gCP(2), gCP(1)}       


	equation = DetectAndFactorQuadratics(equation)

	equation = RemoveUnusedParenthesis(equation)

	fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))



	foiledEquation := FoilAllNeighboringParenthesis(equation)
	fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(foiledEquation), " ", ""))
	foiledEquation = RemoveUnusedParenthesis(foiledEquation)
	foiledEquation = FoilAllNeighboringParenthesis(equation)
	fmt.Println(foiledEquation)
	fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(foiledEquation), " ", ""))
	foiledEquation = RemoveUnusedParenthesis(foiledEquation)
	fmt.Println(foiledEquation)
	fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(foiledEquation), " ", ""))
	
	fmt.Println("makes it here")

fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(foiledEquation), " ", ""))
	foiledEquation = RemoveUnusedParenthesis(foiledEquation)
	fmt.Println("makes it here 1.5")
	fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(foiledEquation), " ", ""))
	fmt.Println()
	foiledEquation = FoilAllNeighboringParenthesis(equation)
	fmt.Println("makes it here 2")
	fmt.Println(foiledEquation)
	foiledEquation = RemoveUnusedParenthesis(foiledEquation)
	
	fmt.Println(foiledEquation)

	fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(foiledEquation), " ", ""))
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



func DecodeFloatSliceToEquation(equation [][]complex128 ) string {

//	CheckEquationForSyntaxErrors(equation)

	equationString := ""

	depthLevel := 0

	for i := 0; i < len(equation); i++ {
		
		currentItem := equation[i]

		for j := 0; j < len(currentItem); j = (j+2) {
			firstIndex := currentItem[j]
			secondIndex := currentItem[j+1]

			firstIndexString := strconv.FormatComplex(firstIndex, 'f', -1, 64)

			secondIndexString := strconv.FormatComplex(secondIndex, 'f', -1, 64)
			


			if(IsOP(firstIndex, secondIndex)){
				equationString += "( "
				depthLevel++ 
			}else if(IsCP(currentItem)){
				equationString += " )"
				if(currentItem[j+2] != 0 && currentItem[j+2] != 1){
					equationString += "^" + strconv.FormatComplex(currentItem[2], 'f', -1, 64) + " "
				}
				depthLevel--
			}else if(IsNumber(firstIndex)){
				if(real(firstIndex) != 1){
					if(real(secondIndex) == 0){
						equationString += firstIndexString + " + "
					}else if(real(secondIndex) == 1){
						equationString += firstIndexString + "S + "
					}else{
						equationString += firstIndexString + "S^" + secondIndexString + " + "
					}
				}else{
					if(real(secondIndex) == 0){
						equationString += firstIndexString + " + "
					}else if(real(secondIndex) == 1){
						equationString += "S + "
					}else{
						equationString += "S^" + secondIndexString + " +"
					}
				}
				
			}else{
				panic("unknown equation item DecodeFloatSliceToEquation()")	
			}
			// fmt.Println(equationString)
			// fmt.Println(depthLevel)

			if(IsCP(currentItem)){
				break	
			}

		}

	}

	okToRemovePlus := false

	plusIndicesToRemove := []int{}

	trackCurrent := -1

	for i := 0; i < len(equationString); i++ {
		if(rune(equationString[i]) == ' '){
			continue
		}else if(rune(equationString[i]) == ')' && !okToRemovePlus){
			okToRemovePlus= true
		}else if(rune(equationString[i]) == '+'){
			if(okToRemovePlus){
				trackCurrent = i
			}
		}else if(rune(equationString[i]) == ')' && okToRemovePlus){
			if(okToRemovePlus){
				plusIndicesToRemove = append(plusIndicesToRemove, trackCurrent)
			}
			okToRemovePlus = false
		}else{
			okToRemovePlus = false
		}

	}

	for i := 0; i < len(plusIndicesToRemove); i++ {

		equationString = equationString[0:plusIndicesToRemove[i]] + equationString[(plusIndicesToRemove[i] + 1):len(equationString)]

		for j := 0; j < len(plusIndicesToRemove); j++ {
			plusIndicesToRemove[j]--
		}
	}

	//strings.TrimSpace(equationString)


	return equationString

}

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


// // //for two parenthesis to be eligible to foil..
// // //they need to be at the same depth level
// // //and the operator between them needs to be multiplication '*'

func FoilAllNeighboringParenthesis(equation [][]complex128) [][]complex128 {

	fmt.Println(equation)
	// fmt.Println(DecodeFloatSliceToEquation(equation))

	equation = RemoveUnusedParenthesis(equation)

	fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))

	CheckEquationForSyntaxErrors(equation)

	fmt.Println("makes it here 3")

	equation = RemoveUnusedParenthesis(equation)

	depthLevel := 0

	foilStateMachine := 0

	foilStateMachinePrevious := 0

	numbersToFoilFirstTerm := []complex128{}

	numbersToFoilSecondTerm := []complex128{}

	previousTerm := []complex128{0, -1}

	foilStart := -1

	foilEnd := -1

	foundFoil := false

	indexOfLastOpeningParenthesis := -1

	timesToFoil := 1

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

			// fmt.Println("current item", currentItem)
			// fmt.Println("previous item", previousTerm)

			if(IsOP(firstIndex, secondIndex)){
				
				indexOfLastOpeningParenthesis = i

				// fmt.Println("is OP")
				depthLevel++ 
				if(foilStateMachine == 0){
					foilStateMachine = 1
					foilStart = i					
				}
				
				
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					// fmt.Println("previous is OP")
					if(foilStateMachine == 4){
						foilStateMachine = 1
						numbersToFoilFirstTerm = []complex128{}
						numbersToFoilSecondTerm = []complex128{}
						foilStart = i
					}
				}else if(IsCP(previousTerm)){
					// fmt.Println("previous is CP")
					if(foilStateMachine == 3){
						foilStateMachine = 4
					}
					
				}else if(IsNumber(previousFirstIndex)){
					// fmt.Println("previous is Numbers")
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

				// fmt.Println("is CP")



				depthLevel--

				if(real(currentItem[2]) > 1){
					foilStart = indexOfLastOpeningParenthesis
					foilEnd = i	
					//TODO: when fractional exponents are added this needs a better method
					timesToFoil = int(real(currentItem[2]) - 1)
					foundFoil = true
					if(foilStateMachine == 2){
						numbersToFoilSecondTerm = numbersToFoilFirstTerm
					}else if(foilStateMachine == 5){
						numbersToFoilFirstTerm = numbersToFoilSecondTerm
					}else{
						panic("should have been in state 2 or 5 FoilAllNeighboringParenthesis()")
					}
				}
				if(IsOP(previousFirstIndex, previousSecondIndex)){
					// fmt.Println("previous is OP")
				}else if(IsCP(previousTerm)){
					// fmt.Println("previous is CP")
				}else if(IsNumber(previousFirstIndex)){
					// fmt.Println("previous is Numbers")
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

				// fmt.Println("is Numbers")


				if(IsOP(previousFirstIndex, previousSecondIndex)){
					// fmt.Println("previous is OP")
					if(foilStateMachine == 1){
						foilStateMachine = 2
						numbersToFoilFirstTerm = append(numbersToFoilFirstTerm, gNum(firstIndex, secondIndex)...)
					}
					if(foilStateMachine == 4){
						foilStateMachine = 5
						numbersToFoilSecondTerm = append(numbersToFoilSecondTerm, gNum(firstIndex, secondIndex)...)
					}
				}else if(IsCP(previousTerm)){
					// fmt.Println("previous is CP")
				}else if(IsNumber(previousFirstIndex)){
					// fmt.Println("previous is Numbers")
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


			// fmt.Println("previous foil state machine", foilStateMachinePrevious)
			// fmt.Println("current foil state machine", foilStateMachine)

			if( (foilStateMachinePrevious == foilStateMachine) && foilStateMachine != 1 && foilStateMachine != 2 && foilStateMachine != 5 && foilStateMachine != 0 && foilStateMachine != 4){
				foilStateMachine = 0
				foilStateMachinePrevious = 0
				numbersToFoilFirstTerm = []complex128{}
				numbersToFoilSecondTerm = []complex128{}
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

		// fmt.Println(foilStart, foilEnd)

		newSlice := []complex128{}



		for timesToFoil > 0 {

			// fmt.Println("first terms and second terms")
			// fmt.Println(numbersToFoilFirstTerm)
			// fmt.Println(numbersToFoilSecondTerm)

			newSlice = []complex128{}

			for i := 0; i < len(numbersToFoilFirstTerm); i = (i + 2) {

				firstNumMultiplier := numbersToFoilFirstTerm[i]

				firstNumExponent := numbersToFoilFirstTerm[i+1]


				for j := 0; j < len(numbersToFoilSecondTerm); j = (j + 2) {

					secondNumMultiplier := numbersToFoilSecondTerm[j]

					secondNumExponent := numbersToFoilSecondTerm[j+1]

					newNum := []complex128{firstNumMultiplier*secondNumMultiplier, firstNumExponent+secondNumExponent}


					// fmt.Print("new num", newNum)

					newSlice = append(newSlice, newNum...)



				}

			}
			newSlice = SimplifyLikeTermsEquationSectionAndSortByDescendningExponent(newSlice)
			// fmt.Println("newSlice iteration", newSlice)
			numbersToFoilFirstTerm = make([]complex128, len(newSlice))
			itemsCopied := copy(numbersToFoilFirstTerm, newSlice)

			if(itemsCopied != len(newSlice)){
				panic("invalid copy FoilAllNeighboringParenthesis()")
			}
			timesToFoil--
		}

		// fmt.Println("first terms and second terms")
		// fmt.Println(numbersToFoilFirstTerm)
		// fmt.Println(numbersToFoilSecondTerm)

		// fmt.Println(newSlice)
		newSlice = SimplifyLikeTermsEquationSectionAndSortByDescendningExponent(newSlice)
		

		return Substitute1DSliceInto2DSliceStartAndEnd(foilStart, foilEnd, newSlice, equation)


	}else{
		return equation
	}


}


func SimplifyLikeTermsEquationSectionAndSortByDescendningExponent(equationSection []complex128) []complex128 {

	termsMap := make(map[complex128][]complex128)



	for i := 0; i < len(equationSection); i = (i + 2) {
		if(termsMap[equationSection[i+1]] == nil){
			termsMap[equationSection[i+1]] = []complex128{equationSection[i]}
		}else{
			termsMap[equationSection[i+1]] = append(termsMap[equationSection[i+1]], equationSection[i])
		}
	}

	exponentsSliceComplex := []complex128{}

	exponentsSliceFloat := []float64{}

	for exponents, _ := range termsMap {
		exponentsSliceFloat = append(exponentsSliceFloat, real(exponents) )
		exponentsSliceComplex = append(exponentsSliceComplex, exponents) 
	}



	// sort.complex128s(exponentsSlice)

	// sort.Reverse(sort.complex128Slice(exponentsSlice))

	copyToMatch := make([]float64, len(exponentsSliceFloat))

	itemsCopied := copy(copyToMatch, exponentsSliceFloat)

	if(itemsCopied != len(exponentsSliceFloat)){
		panic("invalid number of items copied SimplifyLikeTermsEquationSectionAndSortByDescendningExponent()")
	}

	sort.Sort(sort.Reverse(sort.Float64Slice((exponentsSliceFloat))))

	newIndices := []int{}

	for i := 0; i < len(copyToMatch); i++ {

		currentVal := copyToMatch[i]

		for j := 0; j < len(exponentsSliceFloat); j++ {
			currentValInner := exponentsSliceFloat[j]

			if(currentVal == currentValInner){
				newIndices = append(newIndices, j)
				break
			}
		}


	}

	sortedExponentsSliceComplex := make([]complex128, len(exponentsSliceComplex))

	for i := 0; i < len(exponentsSliceComplex); i++ {

		sortedExponentsSliceComplex[newIndices[i]] = exponentsSliceComplex[i]

	}


	returnSlice := []complex128{}


	for i := 0; i < len(sortedExponentsSliceComplex); i++ {

		currentMultipliers := termsMap[sortedExponentsSliceComplex[i]]

		// fmt.Println("exponent", exponentsSlice[i])
		// fmt.Println("multipliers", currentMultipliers)

		mutlipliersAdded := complex128(0)

		for j := 0; j < len(currentMultipliers); j++ {
			mutlipliersAdded += currentMultipliers[j]
		}

		returnSlice = append(returnSlice, mutlipliersAdded)
		returnSlice = append(returnSlice, sortedExponentsSliceComplex[i])


	}

	return returnSlice


}


func Substitute1DSliceInto2DSliceStartAndEnd(start int, end int, new1DSlice []complex128, equation [][]complex128) [][]complex128{

	// fmt.Println("presub ", equation)

	
	returnSlice := append(equation[0:start], []complex128{0, 0})	

	returnSlice = append(returnSlice, new1DSlice)

	returnSlice = append(returnSlice, []complex128{0, 1, 1})	

	returnSlice = append(returnSlice, equation[(end + 1): len(equation)]...)

	

	// fmt.Println("post sub", returnSlice)

	return returnSlice

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


func CheckEquationForSyntaxErrors(equation [][]complex128) {

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


func RemoveUnusedParenthesis(equation [][]complex128) [][]complex128 {

	for (IsOP(equation[len(equation)-1][0], equation[len(equation)-1][1])) {
		equation = equation[0:(len(equation)-1)]
	} 
	
	// fmt.Println("initial parentheis remove check", equation)
	// fmt.Println(DecodeFloatSliceToEquation(equation))	

	tooManyClosingParenthesis := true

	last3Counts := [][]int{}

	for tooManyClosingParenthesis {

		openingParenthesisCount := 0
		closingParenthesisCount := 0


		for i := 0; i < len(equation); i++ {
			if(IsOP(equation[i][0], equation[i][1])){
				openingParenthesisCount++
			}else if(IsCP(equation[i])){
				closingParenthesisCount++
			}

		}

		fmt.Println("parenthesis count ", openingParenthesisCount, closingParenthesisCount)

		if(closingParenthesisCount > openingParenthesisCount){
			if(IsCP(equation[len(equation) - 1])){
				equation = equation[0:len(equation)-1]				


			}
			last3Counts = append(last3Counts, []int{openingParenthesisCount, closingParenthesisCount})
				for len(last3Counts) > 3 {
					last3Counts = append(last3Counts[1:4])
				}

				fmt.Println(len(last3Counts))
				if(len(last3Counts) == 3){
					if(last3Counts[0][0] == last3Counts[1][0] && last3Counts[0][1] == last3Counts[1][1] && last3Counts[1][0] == last3Counts[2][0] && last3Counts[1][1] == last3Counts[2][1]){
						tooManyClosingParenthesis = false
					}
				}
		}else{
			tooManyClosingParenthesis = false
		}


	}



	depthLevel := 0

	//fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))	

	restart := false

	doneCheckingDepth := false

	for !doneCheckingDepth {

		restart = false

		for i := 0; i < len(equation); i++ {

			if(restart){
				break
			}

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
				}else if(IsNumber(firstIndex)){

				}else{
					fmt.Println(currentItem)
					panic("Syntax Error unknown item type RemoveUnusedParenthesis()")
				}

				//TODO FIGURE OUT HOW TO GET RID OF THE UNUSED PARENTHESIS

				//this would occur if there's more closing parenthesis than opening
				if(depthLevel == -1 && IsCP(currentItem)) {
					// panic("got here")
					fmt.Println("first half", strings.ReplaceAll(DecodeFloatSliceToEquation(equation[0:i]), " ", ""))	
					fmt.Println("second half", strings.ReplaceAll(DecodeFloatSliceToEquation(equation[(i+1):len(equation)]), " ", ""))
					fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))	
					equation = append(equation[0:i], equation[(i+1):len(equation)]...) 
					fmt.Println(strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))	
					panic("got here")
					restart = true
					break
				}

				if(len(currentItem) == 3){
					break	
				}
			}
		}

		doneCheckingDepth = true
	}



	CheckEquationForSyntaxErrors(equation)

	

	for i := 0; i < len(equation); i++ {

		currentItem := equation[i]

		for j := 0; j < len(currentItem); j = (j+2) {

			firstIndex := currentItem[j]
			secondIndex := currentItem[j+1]

			if(IsOP(firstIndex, secondIndex)){
				
				startIndexOpenParentParenthesis := i

				endIndexClosedParentParenthesis := -1

				cursor := i+1

				depthLevel := 1

				maxDepth := 1

				uniqueNumbers := true

				uniqueNumbersFound := 0

				for depthLevel > 0 {

					nextItem := equation[cursor] 

					firstIndexInner := nextItem[0]
					secondIndexInner := nextItem[1]

					if(IsOP(firstIndexInner, secondIndexInner)){
						depthLevel++ 
					}else if(IsCP(nextItem)){
						if(depthLevel - 1 == 0){
							endIndexClosedParentParenthesis = cursor
						}
						depthLevel--
						uniqueNumbers = true
					}else if(IsNumber(firstIndexInner)){
						if(uniqueNumbers){
							uniqueNumbersFound++
							uniqueNumbers = false
						}
					}else{
						panic("unknown item type RemoveUnusedParenthesis()")
					}

					if(depthLevel > maxDepth){
						maxDepth = depthLevel
					}

					cursor++

				}

				if(maxDepth > 1 || (maxDepth == 1 )){
					if(uniqueNumbersFound < maxDepth){
						returnEquation := append(equation[0:startIndexOpenParentParenthesis], equation[startIndexOpenParentParenthesis+1: endIndexClosedParentParenthesis]...)
						returnEquation = append(returnEquation, equation[endIndexClosedParentParenthesis+1:len(equation)]...)

						return RemoveUnusedParenthesis(returnEquation)
					}
				}


			}else if(IsCP(currentItem)){
				
			}else if(IsNumber(firstIndex)){

			}else{
				panic("unknown item type RemoveUnusedParenthesis()")
			}

			if(len(currentItem) == 3){
				break	
			}
		}
	}


	return equation
	




}




func DetectAndFactorQuadratics(equation [][]complex128) [][]complex128 {

	CheckEquationForSyntaxErrors(equation)


	for i := 0; i < len(equation); i++ {

		currentItem := equation[i]

		for j := 0; j < len(currentItem); j = (j+2) {

			firstIndex := currentItem[j]
			secondIndex := currentItem[j+1]

			if(IsOP(firstIndex, secondIndex)){
				
				startIndexQuadratic := i

				stopIndexQuadratic := i+2

				if(IsNumber(equation[i+1][0]) && IsCP(equation[i+2])){
					fmt.Println(i, j)
					
					numbers := equation[i+1]

					//which really means 2 or 3 terms
					if(len(numbers) == 4 || len(numbers) == 6){


						aTerm := complex(0, 0)
						bTerm := complex(0, 0)
						cTerm := complex(0, 0)

						

						for k := 0; k < len(numbers); k = (k+2){
							if(real(numbers[k + 1]) == 2){
								aTerm = numbers[k]
							}else if(real(numbers[k + 1]) == 1){
								bTerm = numbers[k]
							}else if(real(numbers[k + 1]) == 0){
								cTerm = numbers[k]
							}
						}
						//make sure the a term and at least one other number are there if 2 items
						if( (IsNumber(aTerm) && IsNumber(bTerm) || IsNumber(aTerm) && IsNumber(cTerm)) && len(numbers) == 4){

							//TODO MAKE THIS MATCH THE LOWER PORTION SO THAT IT WORKS LIKE THAT

							fmt.Println(numbers)
							leftTerm, rightTerm := QuadraticFormula(aTerm, bTerm, cTerm)

							leftTermSlice := []complex128{complex(1, 0), complex(0, 0), leftTerm}
							rightTermSlice := []complex128{complex(1, 0), complex(0, 0), rightTerm}


							bothTerms := [][]complex128{leftTermSlice, rightTermSlice}

							fmt.Println("first half", equation[0:startIndexQuadratic+1])
							fmt.Println("second half", equation[stopIndexQuadratic:len(equation)])

							returnEquation := append(equation[0:startIndexQuadratic+1], bothTerms...)
							returnEquation = append(returnEquation, equation[stopIndexQuadratic:len(equation)]...)
							fmt.Println("return equation", returnEquation)
							fmt.Println("return equation", DecodeFloatSliceToEquation(returnEquation))
							return DetectAndFactorQuadratics(returnEquation)

						//make sure all terms are there if length 3
						}else if(IsNumber(aTerm) && IsNumber(bTerm) && IsNumber(cTerm)){

							fmt.Println("original equation", DecodeFloatSliceToEquation(equation))

							fmt.Println(numbers)
							leftTerm, rightTerm := QuadraticFormula(aTerm, bTerm, cTerm)

							fmt.Println("left term ", leftTerm)
							fmt.Println("right term", rightTerm)

							leftTermSlice := []complex128{complex(1, 0), complex(1, 0), leftTerm, complex(0, 0),}
							rightTermSlice := []complex128{complex(1, 0), complex(1, 0), rightTerm, complex(0, 0),}

							fmt.Println("left term slice ", leftTermSlice)
							fmt.Println("right term slice", rightTermSlice)

							bothTerms := [][]complex128{leftTermSlice, rightTermSlice}

							fmt.Println("first half", equation[0:startIndexQuadratic+1])
							fmt.Println("middle", bothTerms)
							fmt.Println("second half", equation[stopIndexQuadratic:len(equation)])

							cleanCopyMainEquation := CleanCopyEntire2DComplex128Slice(equation)

							returnEquation := append(cleanCopyMainEquation[0:startIndexQuadratic+1], []complex128{complex(0, 0), complex(0, 0)})
							cleanCopyMainEquation = CleanCopyEntire2DComplex128Slice(equation)
							fmt.Println("return equation 1", DecodeFloatSliceToEquation(returnEquation))

							returnEquation = append(returnEquation, leftTermSlice)
							fmt.Println("return equation 2", DecodeFloatSliceToEquation(returnEquation))
							cleanCopyMainEquation = CleanCopyEntire2DComplex128Slice(equation)
							returnEquation = append(returnEquation, []complex128{complex(0, 0), complex(1, 0), complex(1, 0)})
							fmt.Println("return equation 3", DecodeFloatSliceToEquation(returnEquation))
							cleanCopyMainEquation = CleanCopyEntire2DComplex128Slice(equation)
							returnEquation = append(returnEquation, []complex128{complex(0, 0), complex(0, 0)})
							cleanCopyMainEquation = CleanCopyEntire2DComplex128Slice(equation)
							returnEquation = append(returnEquation, rightTermSlice)
							cleanCopyMainEquation = CleanCopyEntire2DComplex128Slice(equation)
							returnEquation = append(returnEquation, []complex128{complex(0, 0), complex(1, 0), complex(1, 0)})
							cleanCopyMainEquation = CleanCopyEntire2DComplex128Slice(equation)
							returnEquation = append(returnEquation, cleanCopyMainEquation[stopIndexQuadratic:len(cleanCopyMainEquation)]...)

							fmt.Println("return equation 4", DecodeFloatSliceToEquation(returnEquation))
							cleanCopyMainEquation = CleanCopyEntire2DComplex128Slice(equation)

							returnEquation = append(returnEquation, cleanCopyMainEquation[0])
							

							fmt.Println("return equation2", returnEquation)

							returnEquation = RemoveUnusedParenthesis(returnEquation)

							fmt.Println("return equation remove", returnEquation)

							
							fmt.Println("return equation", DecodeFloatSliceToEquation(returnEquation))


							return DetectAndFactorQuadratics(returnEquation)
						}

					}



				}


				

			}else if(IsCP(currentItem)){
				
			}else if(IsNumber(firstIndex)){

			}else{
				panic("unknown item type RemoveUnusedParenthesis()")
			}

			if(len(currentItem) == 3){
				break	
			}
		}
	}


	return equation
	



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




