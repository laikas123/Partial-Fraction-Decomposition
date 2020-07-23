package main

import (

	"fmt"
	"os"
	"sync"
	"time"
)






type EquationItem struct {
	Items []interface{}
}


type S_Var struct {
	Multiplier float64
	Exponent int

}


//for before we organize the system of eqtns 
//when the SVar matters 
type GeneralVariable struct {
	Name string
	Multiplier float64
	DegreeToCompareToS int
	
}

//for after we have the system of eqtns
//and the S var matters
type GenVar struct {
	Name string
	Multiplier float64	
}

type OneDEquation struct {

	//this represents variables on either side of equation
	LGenVar []GeneralVariable
	RGenVar []GeneralVariable

	//this represents constants on either side of equations
	LNum []float64
	RNum []float64

}


type Alias struct {

	//this represents variables on either side of equations
	LGenVar []GenVar
	RGenVar []GenVar

	//this represents constants on either side of equations
	LNum []float64
	RNum []float64
}


type SolvedTracker struct {

	AliasesFound []OneDEquation
	ConcreteValuesFound []OneDEquation

}



type ConcreteSolution struct {
	varName string

	value float64

}

var AliasDatabase []Alias

var mutex *sync.Mutex 

var solvedMutex *sync.Mutex 


var Solutions []ConcreteSolution

var Solved bool

var SolvedCheck bool

func main() {
	
	//	Init()	

}









// this function passes all tests
func MultiplyNumeratorByOppositeDenominatorAndOrganizeTheData(leftNumerator []GeneralVariable, rightDenomS []S_Var, rightDenomConstant float64, rightNumerator []GeneralVariable, leftDenomS []S_Var, leftDenomConstant float64, originalNumeratorSVarSlice []S_Var, originalNumeratorConstant float64) []OneDEquation {




	//fraction operations #1 operations 

	returnGeneralVariablesSlice1 := []GeneralVariable{}


	//this is a clean copy functionality no need to use function


	//multiply every general variable by the opposite denominator
	for i := 0; i < len(leftNumerator); i++ {
		for j := 0; j < len(rightDenomS); j++ {
			returnGeneralVariablesSlice1 = append(returnGeneralVariablesSlice1, genVarTimesSVar(leftNumerator[i], rightDenomS[j]))
		}
	}


	for i := 0; i < len(leftNumerator); i++ {
		returnGeneralVariablesSlice1 = append(returnGeneralVariablesSlice1, CreateGeneralVariable(leftNumerator[i].Name, (leftNumerator[i].Multiplier * rightDenomConstant), leftNumerator[i].DegreeToCompareToS))	
	}


	//fraction #2 operations	

	returnGeneralVariablesSlice2 := []GeneralVariable{}

	for i := 0; i < len(rightNumerator); i++ {
		for j := 0; j < len(leftDenomS); j++ {
			returnGeneralVariablesSlice2 = append(returnGeneralVariablesSlice2, genVarTimesSVar(rightNumerator[i], leftDenomS[j]))
		}
		
	}


	for i := 0; i < len(rightNumerator); i++ {
		
		returnGeneralVariablesSlice2 = append(returnGeneralVariablesSlice2, CreateGeneralVariable(rightNumerator[i].Name, (rightNumerator[i].Multiplier * leftDenomConstant), rightNumerator[i].DegreeToCompareToS))
	
	}





	//combined return slices is a slice containing all the values 
	//from both numerators beings multiplied by their opposite denominators
	//each value holds its own sign
	combinedReturnSlices := append(returnGeneralVariablesSlice1, returnGeneralVariablesSlice2...)


	//this will be the one dimensional equation to return
	oneDEqtnSliceToReturn := []OneDEquation{}


	restrictedIndices := []int{}

	var powerToFocusOn int


	for i := 0; i < len(combinedReturnSlices); i++ {
		
		//for indices that havent already been added to the output
		if(!(isRestrictedIndex(restrictedIndices, i))){
			
			restrictedIndices = append(restrictedIndices, i)


			oneDEqtn := OneDEquation{[]GeneralVariable{}, []GeneralVariable{}, []float64{}, []float64{}}

			//what power of S did this variable get multiplied against
			powerToFocusOn = combinedReturnSlices[i].DegreeToCompareToS


			//plug in all variables on the left hand side initially
			oneDEqtn.LGenVar = append(oneDEqtn.LGenVar, combinedReturnSlices[i])


			for j := 0; j < len(combinedReturnSlices); j++ {

				if(!(isRestrictedIndex(restrictedIndices, j))){
					if(combinedReturnSlices[j].DegreeToCompareToS == powerToFocusOn){
						
						//gather all other variables that are of this Power of S
						//and append them to the slice

						restrictedIndices = append(restrictedIndices, j)

						oneDEqtn.LGenVar = append(oneDEqtn.LGenVar, combinedReturnSlices[j])

					}
				}

			}


			//when the "S Power" is 0 
			//this really means a constant 
			if(powerToFocusOn != 0){
			for k := 0; k < len(originalNumeratorSVarSlice); k++ {
				if(originalNumeratorSVarSlice[k].Exponent == powerToFocusOn){
				
					oneDEqtn.RNum = append(oneDEqtn.RNum, originalNumeratorSVarSlice[k].Multiplier)

				}


			}

			oneDEqtnSliceToReturn = append(oneDEqtnSliceToReturn, oneDEqtn)

		}else{
			oneDEqtn.RNum = append(oneDEqtn.RNum, originalNumeratorConstant)


			oneDEqtnSliceToReturn = append(oneDEqtnSliceToReturn, oneDEqtn)

		}

		}

		

	}



	return oneDEqtnSliceToReturn


}


//this function takes the system of linear equations generated 
//and returns all possible rearrangements such that there is only one 
//variable on the left hand side
//all possible permutations of this are returned

//this function passes all tests
func ReturnAllPossibleAliases(oneDEqtnSlice []OneDEquation) []OneDEquation {

	returnOneDEqtnSlice := []OneDEquation{}

	for i := 0; i < len(oneDEqtnSlice); i++ {


		currentOneDEqtn := oneDEqtnSlice[i]

		for j := 0; j < len(currentOneDEqtn.LGenVar); j++ {

			valToRemainLeft := []GeneralVariable{currentOneDEqtn.LGenVar[j]}


			newRightSide := []GeneralVariable{}

			for k := 0; k < len(currentOneDEqtn.LGenVar); k++ {
				if(k != j){
					appendVal := GeneralVariable{currentOneDEqtn.LGenVar[k].Name, (currentOneDEqtn.LGenVar[k].Multiplier* (-1)), currentOneDEqtn.LGenVar[k].DegreeToCompareToS}
					newRightSide = append(newRightSide, appendVal)
				}
			}

			alias := OneDEquation{valToRemainLeft, newRightSide, currentOneDEqtn.LNum, currentOneDEqtn.RNum}


			returnOneDEqtnSlice = append(returnOneDEqtnSlice, alias)

		}
	}


	return returnOneDEqtnSlice

}


//this function passes all tests
func CleanUpAliases(oneDSlice []OneDEquation) []Alias {


	returnOneDEqtnSlice := []Alias{}

	for i := 0; i < len(oneDSlice); i++ {

		currentOneDEqtn := oneDSlice[i]

		oneDEqtn :=  CreateAlias([]GenVar{}, []GenVar{}, currentOneDEqtn.LNum, currentOneDEqtn.RNum)


		//this is a clean copy method itself so no need to worry

		for j := 0; j < len(currentOneDEqtn.LGenVar); j++ {

			simpleGenVar := GenVar{currentOneDEqtn.LGenVar[j].Name, currentOneDEqtn.LGenVar[j].Multiplier}
			oneDEqtn.LGenVar = append(oneDEqtn.LGenVar, simpleGenVar)

		}


		for j := 0; j < len(currentOneDEqtn.RGenVar); j++ {

			simpleGenVar := GenVar{currentOneDEqtn.RGenVar[j].Name, currentOneDEqtn.RGenVar[j].Multiplier}
			oneDEqtn.RGenVar = append(oneDEqtn.RGenVar, simpleGenVar)

		}


		returnOneDEqtnSlice = append(returnOneDEqtnSlice, oneDEqtn)

	}


	return returnOneDEqtnSlice

}


//this method passes all tests
func SubstituteAnAlias(originalAlias Alias, substituteAlias Alias) Alias{

	

	CheckLeftSideIsOnly1Long(originalAlias.LGenVar, "SubstituteAnAlias")
	CheckLeftSideIsOnly1Long(substituteAlias.LGenVar, "SubstituteAnAlias")


	//its ok to index 0 since above its checked that there is only one element
	// leftSideMultiplierSub := substituteAlias.LGenVar[0].Multiplier

	cleanCopyRGenVarSub := CleanCopySliceDataGenVar(substituteAlias.RGenVar)
	cleanCopyRNumSub := CleanCopySliceDataFloat(substituteAlias.RNum)

	scaleValSub := substituteAlias.LGenVar[0].Multiplier



	cleanCopyRGenVarSubScaled := ScaleDownSliceGenVar(cleanCopyRGenVarSub, scaleValSub)
	cleanCopyRNumSubScaled := ScaleDownSliceFloat(cleanCopyRNumSub, scaleValSub)




	cleanCopyRNumOriginal := CleanCopySliceDataFloat(originalAlias.RNum)





	//this is the variable slice of the original alias without the variable to remove
	//and the multiplier of that variable since it will mutliply the newly added substitute values
	originalAliasSubstituteVariableRemoved, multiplierForSubstitute := RemoveExisitngGenVarReturnMultiplier(originalAlias.RGenVar, substituteAlias.LGenVar[0].Name)


	for i := 0; i < len(cleanCopyRGenVarSubScaled); i++ {
		cleanCopyRGenVarSubScaled[i].Multiplier = cleanCopyRGenVarSubScaled[i].Multiplier * multiplierForSubstitute
	}

	for i := 0; i < len(cleanCopyRNumSub); i++ {
		cleanCopyRNumSubScaled[i] = cleanCopyRNumSubScaled[i] * multiplierForSubstitute
	}

	originalAliasSubstituteVariableRemoved = append(originalAliasSubstituteVariableRemoved, cleanCopyRGenVarSubScaled...)

	cleanCopyRNumOriginal = append(cleanCopyRNumOriginal, cleanCopyRNumSubScaled...)



	return CreateAlias(CleanCopySliceDataGenVar(originalAlias.LGenVar), originalAliasSubstituteVariableRemoved, CleanCopySliceDataFloat(originalAlias.LNum), cleanCopyRNumOriginal)


}





func SolutionListener(numberOfSolutionsNeeded int ) []ConcreteSolution{


	soltnsChan := make(chan ConcreteSolution)

	go WorkerSpawnAndAliasListener(soltnsChan)


	solNeededCount := numberOfSolutionsNeeded


	returnSolutionsSlice := []ConcreteSolution{}

	for (solNeededCount > 0) {

		newSolution := <- soltnsChan

		returnSolutionsSlice = append(returnSolutionsSlice, newSolution)

		solNeededCount-- 


	}


	return returnSolutionsSlice




}



func WorkerSpawnAndAliasListener(soltnsChan chan ConcreteSolution)   {

	if(len(AliasDatabase) == 0){
		fmt.Println("SmartSubstitution() called with an empty AliasDatabase")
		os.Exit(1)
	}

		

	cursor := 0


	for !Solved {

		aliasToSend, canSend := ReadItemFromAliasDataBase(cursor)

		if(canSend){

			go WorkOnOneItem(aliasToSend, soltnsChan, cursor)

			cursor++ 
		}else{
			time.Sleep(time.Duration(1) * time.Second)
		}

		canFlipSolveToTrue := false

		solvedMutex.Lock()

			if(SolvedCheck){
				canFlipSolveToTrue = true
			}

		solvedMutex.Unlock()

		if(canFlipSolveToTrue){
			Solved = true
		}




	}





}


func WorkOnOneItem(singleOneD Alias, solutionToSend chan ConcreteSolution, restrictedThisValue int) {


	cursor := 0

	doneWorking := false


	for !doneWorking {

		if(cursor != restrictedThisValue){

			valToWorkWith, dataValid := ReadItemFromAliasDataBase(cursor)

			if(dataValid){

				fmt.Println(valToWorkWith)


			}else{
				time.Sleep(time.Duration(1) * time.Second)		
			}


		}else{
			cursor++ 
		}


	}



}



func AddToAliasDataBase(newAlias Alias, canReadAgain chan bool) {




	mutex.Lock()
   
	AliasDatabase = append(AliasDatabase, newAlias)

    mutex.Unlock()




}

func PrintAliasDataBase() {
	
	mutex.Lock()
   
	fmt.Printf("%#v\n", AliasDatabase)

    mutex.Unlock()

}



func ReadItemFromAliasDataBase(index int) (Alias, bool) {

	mutex.Lock()

		if(!(index >= len(AliasDatabase))){


			//when reading data from the data base clean copy is used for the slices because otherwise pointers to the underlying database itself would be
			//passed on... not good

			return CreateAlias(CleanCopySliceDataGenVar(AliasDatabase[index].LGenVar), CleanCopySliceDataGenVar(AliasDatabase[index].RGenVar), CleanCopySliceDataFloat(AliasDatabase[index].LNum), CleanCopySliceDataFloat(AliasDatabase[index].RNum)), true

			
		}else{
			return Alias{}, false
		}

	mutex.Unlock()

	return Alias{}, false

}




// this function passes all tests
func FullCleanUp(genVarInput Alias)  Alias {

	cleanCopy := CleanCopyAlias(genVarInput)

	fmt.Println("main function clean copy")
	VerbosePrintln(cleanCopy)

	cleanCopy = SimplifyGenVarRightHandGenVarSlice(cleanCopy)

	fmt.Println("main function clean copy")
	VerbosePrintln(cleanCopy)

	cleanCopy = SimplifyRightHandNumSlice(cleanCopy)
	
	fmt.Println("main function clean copy")
	VerbosePrintln(cleanCopy)

	cleanCopy = MoveVarsEqualToLeftHandSideToLeftSide(cleanCopy)
	
	fmt.Println("main function clean copy special ")
	VerbosePrintln(cleanCopy)

	var leftSideZero bool

	cleanCopy, leftSideZero = RemoveZerosWarnIfLeftHandSideZero(cleanCopy)

	fmt.Println("main function clean copy")
	VerbosePrintln(cleanCopy)

	if(leftSideZero){
		fmt.Println("left hand side 0")
		os.Exit(1)
	}


	return cleanCopy

} 



//this function works well
func SimplifyGenVarRightHandGenVarSlice(genVarInput Alias) Alias{


	CheckLeftSideIsOnly1Long(genVarInput.LGenVar, "SimplifyGenVarRightHandGenVarSlice")


	//don't modify the underlying slice instead copy the data
	copyRGenVarData := CleanCopySliceDataGenVar(genVarInput.RGenVar)



	//these will be indices that already found a combination
	//hence we don't want to double combine and corrupt the data
	restrictedIndices := []int{}


	//this will be the output slice of the combined right hand side variables
	outputRGenVarData := []GenVar{}



	//keep track of indices that matched,
	//this makes sure items that didnt match wont get lost
	indicesThatMatched := []int{}


	//since this function is called recursively 
	//once no simplifications were made 
	//the result can be returned
	oneSimplificationWasMade :=	false

	for i := 0; i < len(copyRGenVarData); i++ {
		if(!isRestrictedIndex(restrictedIndices, i)){

			restrictedIndices = append(restrictedIndices, i)

			checkVal1 := copyRGenVarData[i]

			for j := 0; j < len(copyRGenVarData); j++ {

				if(!isRestrictedIndex(restrictedIndices, j)){
				//restrictedIndices = append(restrictedIndices, j)
					checkVal2 := copyRGenVarData[j]


					//if the two variables are of the same name
					if(checkVal1.Name == checkVal2.Name){


						//add them together to the output slice
						outputRGenVarData = append(outputRGenVarData, GenVar{checkVal1.Name, (checkVal1.Multiplier + checkVal2.Multiplier)})
					
						//add this j value to restricted indices so it is not double added
						restrictedIndices = append(restrictedIndices, j)


						//add the matching indices
						indicesThatMatched = append(indicesThatMatched, i)

						indicesThatMatched = append(indicesThatMatched, j)

						oneSimplificationWasMade = true

						break

					}

				}	
			}

		}

	}




	for i := 0; i < len(copyRGenVarData); i++ {

		//the restricted index function can be used
		//also here, dont let the name play tricks
		if(!isRestrictedIndex(indicesThatMatched, i)){
				outputRGenVarData = append(outputRGenVarData, copyRGenVarData[i])
		}


	}

	//if no simplifications were made return the data as is
	if(!oneSimplificationWasMade){
		return CreateAlias(CleanCopySliceDataGenVar(genVarInput.LGenVar), outputRGenVarData, CleanCopySliceDataFloat(genVarInput.LNum), CleanCopySliceDataFloat(genVarInput.RNum))

	//else recursively call until there can be no more simplifications
	}else{
		return SimplifyGenVarRightHandGenVarSlice(CreateAlias(CleanCopySliceDataGenVar(genVarInput.LGenVar), outputRGenVarData, CleanCopySliceDataFloat(genVarInput.LNum), CleanCopySliceDataFloat(genVarInput.RNum)))
	}



}



func SimplifyRightHandNumSlice(genVarInput Alias) Alias{


	CheckLeftSideIsOnly1Long(genVarInput.LGenVar, "SimplifyRightHandNumSlice")

	//don't modify the underlying slice instead copy the data
	copyRNumData := CleanCopySliceDataFloat(genVarInput.RNum)

	summation := float64(0)

	for i := 0; i < len(copyRNumData); i++ {
	
		summation = summation + copyRNumData[i]

	}

	return CreateAlias(CleanCopySliceDataGenVar(genVarInput.LGenVar), CleanCopySliceDataGenVar(genVarInput.RGenVar), CleanCopySliceDataFloat(genVarInput.LNum), []float64{summation})

}



func MoveVarsEqualToLeftHandSideToLeftSide(genVarInput Alias) Alias {

	CheckLeftSideIsOnly1Long(genVarInput.LGenVar, "MoveVarsEqualToLeftHandSideToLeftSide")

	leftHandVarName := genVarInput.LGenVar[0].Name

	variablesMoved := 0

	var variableToMove GenVar

	for i := 0; i < len(genVarInput.RGenVar); i++ {
		if(genVarInput.RGenVar[i].Name == leftHandVarName){
			variablesMoved++
			variableToMove = genVarInput.RGenVar[i]
		}
	}

	//since right before this function 
	//like terms should have been combined
	//there never should be a case where two elements are 
	//moved to the left
	if(variablesMoved > 1){
		fmt.Println("too many variables moved")
		os.Exit(1)
	}

	cleanCopyRightHandSideForRemoval := CleanCopySliceDataGenVar(genVarInput.RGenVar)


	//discard the multiplier variable returned since here the only goal is to remove the element
	cleanCopyRightHandSideForRemoval, _ =  RemoveExisitngGenVarReturnMultiplier(cleanCopyRightHandSideForRemoval, leftHandVarName)


	leftHandGenVar := genVarInput.LGenVar[0]

	newCombinedVariable := GenVar{leftHandGenVar.Name, leftHandGenVar.Multiplier - variableToMove.Multiplier}


	return CreateAlias([]GenVar{newCombinedVariable}, cleanCopyRightHandSideForRemoval, CleanCopySliceDataFloat(genVarInput.LNum), CleanCopySliceDataFloat(genVarInput.RNum))


}



//the boolean return value is set true only if the left hand side is 0
//as this means we have nulled out the variable we need to know the value of
func RemoveZerosWarnIfLeftHandSideZero(genVarInput Alias) (Alias, bool) {

	CheckLeftSideIsOnly1Long(genVarInput.LGenVar, "RemoveZerosWarnIfLeftHandSideZero")


	warnThatLeftVarIsNull := false

	if(genVarInput.LGenVar[0].Multiplier == 0){
			warnThatLeftVarIsNull = true
	}

	indicesToRemove := []int{}

	for i := 0; i < len(genVarInput.RGenVar); i++ {
		if(genVarInput.RGenVar[i].Multiplier == 0){
			VerbosePrintln(genVarInput.RGenVar)
			indicesToRemove = append(indicesToRemove, i)
		}
	}

	newRGenVarSlice := []GenVar{}

	if(len(indicesToRemove) != 0){
		for i := 0; i < len(genVarInput.RGenVar); i++ {
		
			okToAdd := false

			for j := 0; j < len(indicesToRemove); j++ {
				if(i != indicesToRemove[j]){
					okToAdd = true
					break
				}
			}

			//this is a clean form of adding so no need to call clean function
			if(okToAdd){
				newRGenVarSlice = append(newRGenVarSlice, genVarInput.RGenVar[i])
			}

		}

		return CreateAlias(CleanCopySliceDataGenVar(genVarInput.LGenVar), CleanCopySliceDataGenVar(newRGenVarSlice), CleanCopySliceDataFloat(genVarInput.LNum), CleanCopySliceDataFloat(genVarInput.RNum)), warnThatLeftVarIsNull
	}else{
		return CleanCopyAlias(genVarInput), warnThatLeftVarIsNull
	}






}



//this "new alias" should be the scaled down version 
//scaled down being that its left side variable equals 1 and the right side is scaled
//accordingly
func NewAliasEqualsLeftSideVariableNoIncrease(oldAlias Alias, newAlias Alias) bool {


	//the net result is what matters 
	//a clean copy of old and new is taken

	// cleanCopyOldAlias := CleanCopyAlias(oldAlias)
	// cleanCopyNewAlias := CleanCopyAlias(newAlias)

	return false


}




func NewAliasTransmuteToExistingVariableAndReducesTheVariablesOnTheRight(oldAlias Alias, newAlias Alias) bool {
	return false
}

func NewAliasHasVariableAlreadyOnLeft(oldAlias Alias, newAlias Alias) bool {
	return false
}






















func Init() {
	mutex = &sync.Mutex{}
	solvedMutex = &sync.Mutex{}
	AliasDatabase = []Alias{}
	Solutions = []ConcreteSolution{}
	Solved = false
	SolvedCheck = false
}








func isRestrictedIndex(restrictedIntSlice []int, checkIndex int) bool {
	for i := 0; i < len(restrictedIntSlice); i++ {
		if(restrictedIntSlice[i] == checkIndex){
			return true
		}
	}

	return false
}

func genVarTimesSVar(genVar GeneralVariable, sVar S_Var) GeneralVariable {

	sExpGenVar := genVar.DegreeToCompareToS

	sExpSVar := sVar.Exponent

	newExponent := sExpGenVar + sExpSVar

	multiplierGenVar := genVar.Multiplier
	
	multiplierSVar := sVar.Multiplier

	newMultiplier := multiplierSVar*multiplierGenVar

	return GeneralVariable{genVar.Name, newMultiplier, newExponent}	

}



func PartialFractionDecompositionSolver() {
	
}



func PartialFractionDecomposition(numer []EquationItem, denom1 []EquationItem, denom2 []EquationItem) []EquationItem {



	return []EquationItem{}

}


func VariableNameAlphabetIndex(index int) string {

	if(index < 0){
		fmt.Println("impossible index for alphabet")
		os.Exit(1)
	}


	switch index {

		case 0:
			return "A"
		case 1:
			return "B"
		case 2:
			return "C"
		case 3:
			return "D"
		case 4:
			return "E"
		case 5:
			return "F"
		case 6:
			return "G"
		case 7:
			return "H"
		case 8:
			return "I"
		case 9:
			return "J"
		case 10:
			return "K"
		case 11:
			return "L"
		case 12:
			return "M"
		case 13:
			return "N"
		case 14:
			return "O"
		case 15:
			return "P"
		case 16:
			return "Q"
		case 17:
			return "R"
		case 18:
			return "S"
		case 19:
			return "T"
		case 20:
			return "U"
		case 21:
			return "V"
		case 22:
			return "W"
		case 23:
			return "X"
		case 24:
			return "Y"
		case 25:
			return "Z"
		default:


			newNumber := index - 25

			return "A"+ VariableNameAlphabetIndex(newNumber)
			
			
	}

	return "-1*z*error"



}



func returnHighestDegree(terms []EquationItem) int {



	var highestDegree int

	foundAtLeastOneS := false

	for i := 0; i < len(terms); i++ {

		for j := 0; j < len(terms[i].Items); j++ {

			value, ok := terms[i].Items[j].(S_Var)

			if(ok){
				if(!foundAtLeastOneS){
					highestDegree = value.Exponent
					foundAtLeastOneS = true
				}else if(value.Exponent > highestDegree){
					highestDegree = value.Exponent
				}
		}


		}

	}


	if(!foundAtLeastOneS){
		fmt.Println("error no s variables when searching for highest degree")
		os.Exit(1)
	}


	return highestDegree


}



func GEQI(items ...interface{}) EquationItem {


	returnSlice := []interface{}{}

	for i := 0; i < len(items); i++ {


		item := items[i]

		switch item.(type){
			case S_Var:

				value, ok := item.(S_Var)

				if(ok){
					returnSlice = append(returnSlice, value)
				}else{
					fmt.Println("couldn't assert type for")
					fmt.Printf("%#v\n", item)
					os.Exit(1)
				}
			case float64:
				value, ok := item.(float64)

				if(ok){
					returnSlice = append(returnSlice, value)
				}else{
					fmt.Println("couldn't assert type for")
					fmt.Printf("%#v\n", item)
					os.Exit(1)
				}
			case int:

				

				value, ok := item.(int)

				if(ok){
					intToFloat := float64(value)
					
					returnSlice = append(returnSlice, intToFloat)
				}else{
					fmt.Println("couldn't assert type for")
					fmt.Printf("%#v\n", item)
					os.Exit(1)
				}
			default:
				fmt.Println("unkown type for")
				fmt.Printf("%#v\n", item)
				os.Exit(1)
		}


     
    }

    return EquationItem{returnSlice}

}




//for now only whole number exponents will be considered
//in the future look up methods for partial fraction decomposition 
//with fractional exponents
func ReturnGeneralVariablesForDegree(degree int, startIndex int) ([]GeneralVariable, int) {




	generalVariableSlice := []GeneralVariable{}



	for i := 0; i < degree; i++ {
		generalVariableSlice = append(generalVariableSlice, GeneralVariable{VariableNameAlphabetIndex(i+startIndex), 1, degree-1-i})
	}

	return generalVariableSlice, (degree)


}




func CreateGeneralVariable(name string, multiplier float64, degreeCompareS int) GeneralVariable{
	return GeneralVariable{name, multiplier, degreeCompareS}
}

func CreateGenVar(name string, multiplier float64) GenVar{
	return GenVar{name, multiplier}
}


func CreateSVar(multiplier float64, exponent int) S_Var {
	return S_Var{multiplier, exponent}
}

func CreateAlias(leftGenVar []GenVar, rightGenVar []GenVar, leftNum []float64, rightNum []float64) Alias {
	return Alias{leftGenVar, rightGenVar, leftNum, rightNum}
} 




//since slices are a pointer to an underlying value and we 
//are constantly calling to the alias database which has values that we re use
//it's best to clean copy the data aka create a new pointer
func CleanCopySliceDataGenVar(input []GenVar) []GenVar {

	outputSlice := []GenVar{}

	for i := 0; i < len(input); i++ {
		outputSlice = append(outputSlice, input[i])
	}

	return outputSlice

}



func CleanCopySliceDataFloat(input []float64) []float64 {

	outputSlice := []float64{}

	for i := 0; i < len(input); i++ {
		outputSlice = append(outputSlice, input[i])
	}

	return outputSlice

}


func CleanCopySliceDataGeneralVariable(input []GeneralVariable) []GeneralVariable {

	outputSlice := []GeneralVariable{}

	for i := 0; i < len(input); i++ {
		outputSlice = append(outputSlice, input[i])
	}

	return outputSlice

}



func TwoGenVarsAreEqual(genvar1 GenVar, genvar2 GenVar) bool {


	if((genvar1.Name == genvar2.Name) && (genvar1.Multiplier == genvar2.Multiplier) ){
		return true
	}else{
		return false
	}


}


func TwoGenVarsAreSameVariable(genvar1 GenVar, genvar2 GenVar) bool {


	if(genvar1.Name == genvar2.Name){
		return true
	}else{
		return false
	}


}



func RemoveExisitngGenVarReturnMultiplier(genVarSlice []GenVar, removeName string) ([]GenVar, float64) {

	returnSlice := []GenVar{}

	var returnMultiplier float64

	valRemovedMultiplierFound := false

	removedItems := 0

	for i := 0; i < len(genVarSlice); i++ {
		currentVarName := genVarSlice[i].Name
		if(currentVarName != removeName){
			returnSlice = append(returnSlice, genVarSlice[i])
		}else{
			returnMultiplier = genVarSlice[i].Multiplier
			valRemovedMultiplierFound = true
			removedItems++
		}
	}

	if(!valRemovedMultiplierFound){
		fmt.Println("couldn't remove value from slice or couldn't find multiplier")
		os.Exit(1)

	}

	if(removedItems > 1){
		fmt.Println("more than one item removed")
		os.Exit(1)
	}

	return returnSlice, returnMultiplier


}



func CheckLeftSideIsOnly1Long(genVarInput []GenVar, functionCaller string){
	//if there's more variables on the left than one something went wrong earlier in the program
	if(len(genVarInput) > 1){
		fmt.Println("equation doesn't only have one variable on the left side from function:", functionCaller)
		VerbosePrintln(genVarInput)
		os.Exit(1)
	}


}

func VerbosePrintln(input interface{}) {
	fmt.Printf("%#v\n", input)
}








func ScaleDownSliceGenVar(genVarInput []GenVar, scaleVal float64) []GenVar {

	copiedData := CleanCopySliceDataGenVar(genVarInput)

	for i := 0; i < len(copiedData); i++ {
		copiedData[i].Multiplier = (copiedData[i].Multiplier)/(scaleVal)
	}

	return copiedData

}

func ScaleDownSliceFloat(floatSlice []float64, scaleVal float64) []float64 {

	copiedData := CleanCopySliceDataFloat(floatSlice)

	for i := 0; i < len(copiedData); i++ {
		copiedData[i] = (copiedData[i])/(scaleVal)
	}

	return copiedData

}


//this function passes all tests
func CleanCopyAlias(oldAlias Alias) Alias{


	return CreateAlias(CleanCopySliceDataGenVar(oldAlias.LGenVar), CleanCopySliceDataGenVar(oldAlias.RGenVar), CleanCopySliceDataFloat(oldAlias.LNum), CleanCopySliceDataFloat(oldAlias.RNum))


}














