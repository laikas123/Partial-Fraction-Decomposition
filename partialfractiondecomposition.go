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

	LGenVar []GeneralVariable
	RGenVar []GeneralVariable

	LNum []float64
	RNum []float64

}


type AliasOneDEquationSimple struct {

	LGenVar []GenVar
	RGenVar []GenVar

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

var AliasDatabase []AliasOneDEquationSimple

var mutex *sync.Mutex 

var solvedMutex *sync.Mutex 


var Solutions []ConcreteSolution

var Solved bool

var SolvedCheck bool

func main() {
	
	//	Init()	

}










func MultiplyNumeratorByOppositeDenominatorAndOrganizeTheData(genVarSlice1 []GeneralVariable, sVarSlice1 []S_Var, constant1 float64, genVarSlice2 []GeneralVariable, sVarSlice2 []S_Var, constant2 float64, originalNumeratorSVarSlice []S_Var, originalNumeratorConstant float64) []OneDEquation {




	//fraction operations #1 operations 

	returnGeneralVariablesSlice1 := []GeneralVariable{}


	//this is a clean copy functionality no need to use function


	//multiply every general variable by the opposite denominator
	for i := 0; i < len(genVarSlice1); i++ {
		for j := 0; j < len(sVarSlice1); j++ {
			returnGeneralVariablesSlice1 = append(returnGeneralVariablesSlice1, genVarTimesSVar(genVarSlice1[i], sVarSlice1[j]))
		}
		
	}


	for i := 0; i < len(genVarSlice1); i++ {



		returnGeneralVariablesSlice1 = append(returnGeneralVariablesSlice1, GeneralVariable{genVarSlice1[i].Name, (genVarSlice1[i].Multiplier * constant1), genVarSlice1[i].DegreeToCompareToS})
	
	}


	//#2 operations	

	returnGeneralVariablesSlice2 := []GeneralVariable{}

	for i := 0; i < len(genVarSlice2); i++ {
		for j := 0; j < len(sVarSlice2); j++ {
			returnGeneralVariablesSlice2 = append(returnGeneralVariablesSlice2, genVarTimesSVar(genVarSlice2[i], sVarSlice2[j]))
		}
		
	}


	for i := 0; i < len(genVarSlice2); i++ {



		returnGeneralVariablesSlice2 = append(returnGeneralVariablesSlice2, GeneralVariable{genVarSlice2[i].Name, (genVarSlice2[i].Multiplier * constant2), genVarSlice2[i].DegreeToCompareToS})
	
	}






	combinedReturnSlices := append(returnGeneralVariablesSlice1, returnGeneralVariablesSlice2...)

	oneDEqtnSliceToReturn := []OneDEquation{}


	restrictedIndices := []int{}

	var powerToFocusOn int

	for i := 0; i < len(combinedReturnSlices); i++ {
		if(!(isRestrictedIndex(restrictedIndices, i))){
			
			restrictedIndices = append(restrictedIndices, i)

			oneDEqtn := OneDEquation{[]GeneralVariable{}, []GeneralVariable{}, []float64{}, []float64{}}

			powerToFocusOn = combinedReturnSlices[i].DegreeToCompareToS

			oneDEqtn.LGenVar = append(oneDEqtn.LGenVar, combinedReturnSlices[i])


			for j := 0; j < len(combinedReturnSlices); j++ {

				if(!(isRestrictedIndex(restrictedIndices, j))){
					if(combinedReturnSlices[j].DegreeToCompareToS == powerToFocusOn){
						
						restrictedIndices = append(restrictedIndices, j)

						oneDEqtn.LGenVar = append(oneDEqtn.LGenVar, combinedReturnSlices[j])

					}
				}

			}


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


//this function needs one more test
func CleanUpVars(oneDSlice []OneDEquation) []AliasOneDEquationSimple {


	returnOneDEqtnSlice := []AliasOneDEquationSimple{}

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


func WorkOnOneItem(singleOneD AliasOneDEquationSimple, solutionToSend chan ConcreteSolution, restrictedThisValue int) {


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



func AddToAliasDataBase(newAlias AliasOneDEquationSimple, canReadAgain chan bool) {




	mutex.Lock()
   
	AliasDatabase = append(AliasDatabase, newAlias)

    mutex.Unlock()




}

func PrintAliasDataBase() {
	
	mutex.Lock()
   
	fmt.Printf("%#v\n", AliasDatabase)

    mutex.Unlock()

}



func ReadItemFromAliasDataBase(index int) (AliasOneDEquationSimple, bool) {

	mutex.Lock()

		if(!(index >= len(AliasDatabase))){


			//when reading data from the data base clean copy is used for the slices because otherwise pointers to the underlying database itself would be
			//passed on... not good

			return CreateAlias(CleanCopySliceDataGenVar(AliasDatabase[index].LGenVar), CleanCopySliceDataGenVar(AliasDatabase[index].RGenVar), CleanCopySliceDataFloat(AliasDatabase[index].LNum), CleanCopySliceDataFloat(AliasDatabase[index].RNum)), true

			
		}else{
			return AliasOneDEquationSimple{}, false
		}

	mutex.Unlock()

	return AliasOneDEquationSimple{}, false

}



func FullCleanUp(genVarInput AliasOneDEquationSimple)  {

	// copyGenVar := genVarInput

	// copyGenVar = SimplifyGenVarRightHandGenVarSlice(copyGenVar)
	// copyGenVar = SimplifyRightHandNumSlice(copyGenVar)
	// copyGenVar = MoveVarsEqualToLeftHandSideToLeftSide(copyGenVar)

	// var leftSideNull bool

	// copyGenVar, leftSideNull = RemoveZerosWarnIfLeftHandSideZero(copyGenVar)

	// if(leftSideNull){
	// 	fmt.Println("left side eqtn is null")
	// 	os.Exit(1)
	// }

	// fmt.Println(copyGenVar)


} 



//this function works well
func SimplifyGenVarRightHandGenVarSlice(genVarInput AliasOneDEquationSimple) AliasOneDEquationSimple{


	//if there's more variables on the left than one something went wrong earlier in the program
	if(len(genVarInput.LGenVar) > 1){
		fmt.Println("equation doesn't only have one variable on the left side")
		os.Exit(1)
	}



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



func SimplifyRightHandNumSlice(genVarInput AliasOneDEquationSimple) AliasOneDEquationSimple{



	//don't modify the underlying slice instead copy the data
	copyRNumData := CleanCopySliceDataGenVar(genVarInput.RNum)


	//if there's more variables on the left than one something went wrong earlier in the program
	if(len(genVarInput.LGenVar) > 1){
		fmt.Println("equation doesn't only have one variable on the left side")
		os.Exit(1)
	}

	leftSidedVar := genVarInput.LGenVar[0]


	//these will be indices that already found a combination
	//hence we don't want to double combine and corrupt the data
	restrictedIndices := []int{}


	//this will be the output slice of the combined right hand side variables
	outputRNumData := []float64{}



	//keep track of indices that matched,
	//this makes sure items that didnt match wont get lost
	indicesThatMatched := []int{}


	//since this function is called recursively 
	//once no simplifications were made 
	//the result can be returned
	oneSimplificationWasMade :=	false

	for i := 0; i < len(copyRNumData); i++ {
		if(!isRestrictedIndex(restrictedIndices, i)){

			restrictedIndices = append(restrictedIndices, i)

			checkVal1 := copyRNumData[i]

			for j := 0; j < len(copyRNumData); j++ {

				if(!isRestrictedIndex(restrictedIndices, j)){
				//restrictedIndices = append(restrictedIndices, j)
					checkVal2 := copyRNumData[j]


					//if the two variables are of the same name
					if(checkVal1 == checkVal2){


						//add them together to the output slice
						outputRNumData = append(outputRNumData, (checkVal1 + checkVal2))
					
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




	for i := 0; i < len(copyRNumData); i++ {

		//the restricted index function can be used
		//also here, dont let the name play tricks
		if(!isRestrictedIndex(indicesThatMatched, i)){
				outputRNumData = append(outputRNumData, copyRNumData[i])
		}


	}

	//if no simplifications were made return the data as is
	if(!oneSimplificationWasMade){
		return CreateAlias(CleanCopySliceDataGenVar(genVarInput.LGenVar), CleanCopySliceDataGenVar(genVarInput.RGenVar), CleanCopySliceDataFloat(genVarInput.LNum), outputRNumData))

	//else recursively call until there can be no more simplifications
	}else{
		return SimplifyGenVarRightHandGenVarSlice(CreateAlias(CleanCopySliceDataGenVar(genVarInput.LGenVar), CleanCopySliceDataGenVar(genVarInput.RGenVar), CleanCopySliceDataFloat(genVarInput.LNum), outputRNumData))
	}



}



// func SimplifyRightHandNumSlice(genVarInput AliasOneDEquationSimple) AliasOneDEquationSimple {

// 	copyRNumData := []GenVar

// 	for i := 0; i < len(genVarInput.RNum); i++ {

// 		copyRNumData = append(copyRNumData, genVarInput.RNum[i])

// 	}

// 	restrictedIndices := []int{}

// 	outputRNumData := []int{}

// 	indicesThatMatched := []int{}

// 	oneSimplificationWasMade :=	false

// 	for i := 0; i < len(copyRNumData); j++ {
// 		if(!isRestrictedIndex(restrictedIndices, i)){

// 			restrictedIndices = append(restrictedIndices, i)

// 			checkVal1 := copyRNumData[i]

// 			for j := 0; j < len(copyRNumData); j++ {

// 				if(!isRestrictedIndex(restrictedIndices, j)){
// 				//restrictedIndices = append(restrictedIndices, j)
// 					checkVal2 := copyRNumData[j]

// 					if(checkVal1.Name == checkVal2.Name){

// 						outputRNumData = append(outputRNumData, (checkVal1+checkVal2) )
					
// 						restrictedIndices = append(restrictedIndices, j)

// 						indicesThatMatched = append(indicesThatMatched, i)

// 						indicesThatMatched = append(indicesThatMatched, j)

// 						oneSimplificationWasMade = true

// 						break

// 					}

// 				}	
// 			}

// 		}

// 	}




// 	for i := 0; i < len(copyRNumData); i++ {

// 		if(!isRestrictedIndex(indicesThatMatched)){
// 			outputRNumData = append(outputRNumData, copyRGenVarData[i])
// 		}


// 	}

// 	if(!oneSimplificationWasMade){
// 		return CreateAlias(genVarInput.LGenVar, genVarInput.RGenVar, genVarInput.LNum, outputRNumData)
// 	}else{
// 		return SimplifyGenVarRightHandGenVarSlice(CreateAlias(genVarInput.LGenVar, genVarInput.RGenVar, genVarInput.LNum, outputRNumData))
// 	}



// }


// func MoveVarsEqualToLeftHandSideToLeftSide(genVarInput AliasOneDEquationSimple) AliasOneDEquationSimple {


// 	if(len(genVarInput.LGenVar) > 0){
// 		fmt.Println("equation doesn't only have one variable on the left side")
// 		os.Exit(1)
// 	}

// 	leftSidedVarName := genVarInput.LGenVar[0].Name

// 	indicesToRemove := []int{}

// 	for i := 0; i < len(genVarInput.RGenVar); i++ {
// 		if(genVarInput.RGenVar[i].Name == leftSidedVarName){
// 			indicesToRemove = append(indicesToRemove, i)
// 			genVarInput.LGenVar[0] = []GeneralVariable{GeneralVariable{leftSidedVarName, (genVarInput.RGenVar[i].Multiplier + genVarInput.LGenVar[0].Multiplier) } }
// 		}
// 	}

// 	newRGenVar := []GeneralVariable{}

// 	for i := 0; i < len(genVarInput.RGenVar); i++ {

// 		okToAdd := false

// 		for j := 0; j < len(indicesToRemove); j++ {
// 			if(i == indicesToRemove[j]){
// 				okToAdd = true
// 				break
// 			}
// 		}

// 		if(okToAdd){
// 			newRGenVar = append(newRGenVar, genVarInput.RGenVar[i])
// 		}


// 	}


// 	return CreateAlias(genVarInput.LGenVar, newRGenVar, genVarInput.LNum, genVarInput.RNum)


// }



// //the boolean return value is set true only if the left hand side is 0
// //as this means we have nulled out the variable we need to know the value of
// func RemoveZerosWarnIfLeftHandSideZero(genVarInput AliasOneDEquationSimple) (AliasOneDEquationSimple, bool) {


// 	if(len(genVarInput.LGenVar) > 0){
// 		fmt.Println("equation doesn't only have one variable on the left side")
// 		os.Exit(1)
// 	}
	
// 	warnThatLeftVarIsNull := false

// 	if(genVarInput.LGenVar[0].Multiplier == 0){
// 			warnThatLeftVarIsNull = true
// 	}

// 	indicesToRemove := []int{}

// 	for i := 0; i < len(genVarInput.RGenVar); i++ {
// 		if(genVarInput.RGenVar[i].Multiplier == 0){
// 			indicesToRemove = append(indicesToRemove, i)
// 		}
// 	}

// 	newRGenVarSlice := []GenVar{}

// 	for i := 0; i < len(genVarInput.RGenVar); i++ {
		
// 		okToAdd := false

// 		for j := 0; j < len(indicesToRemove); j++ {
// 			if(i == indicesToRemove[j]){
// 				okToAdd = true
// 				break
// 			}
// 		}

// 		if(okToAdd){
// 			newRGenVarSlice = append(newRGenVarSlice, genVarInput.RGenVar[i])
// 		}

// 	}


// 	return CreateAlias(genVarInput.LGenVar, newRGenVarSlice, genVarInput.LNum, genVarInput.RNum), warnThatLeftVarIsNull


// }






func NewAliasTransmuteToExistingVariableAndReducesTheVariablesOnTheRight(oldAlias AliasOneDEquationSimple, newAlias AliasOneDEquationSimple) bool {
	return false
}

func NewAliasHasVariableAlreadyOnLeft(oldAlias AliasOneDEquationSimple, newAlias AliasOneDEquationSimple) bool {
	return false
}






















func Init() {
	mutex = &sync.Mutex{}
	solvedMutex = &sync.Mutex{}
	AliasDatabase = []AliasOneDEquationSimple{}
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




func CreateAlias(lGenVar []GenVar, rGenVar []GenVar, lNum []float64, rNum []float64) AliasOneDEquationSimple {
	return AliasOneDEquationSimple{lGenVar, rGenVar, lNum, rNum}
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


func CleanCopySliceGeneralVariable(input []GeneralVariable) []GeneralVariable {

	outputSlice := []GeneralVariable{}

	for i := 0; i < len(input); i++ {
		outputSlice = append(outputSlice, input[i])
	}

	return outputSlice

}






