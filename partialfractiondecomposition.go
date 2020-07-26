package main

import (

	"fmt"
	"os"
	"sync"
	"time"
)


type SolutionItem struct {
	BinaryCursor string
	PseudoNamesChosenCursor []int
	HighestNetChange int
}


type VarPseudoNames struct {
	PseudoNames [][]string
	//vals for RNum for each variable
	//0 if nil
	LoneNumberVals []float64
	ScaledDownMultipliers []float64
	ParentVar string
}


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
	Name string

	Value float64

}

var AliasDatabase []Alias

var mutex *sync.Mutex 

var solvedMutex *sync.Mutex 

var seedMutex *sync.Mutex 

var printAliasMutex *sync.Mutex

var Solutions []ConcreteSolution

var Solved bool

var SolvedCheck bool

var SeedsTested []Alias

var OldGlobal Alias

var SubGlobal Alias

var NetGlobal Alias


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


//this is essentially the same as the "ReturnAllAliases" method
//that method is called for when the data is still in OneDEquation format
//whereas this is called when the data is already in an Alias format

//this function passes all tests
func AllAliasPermutationsAndAddToDatabase(alias Alias)  {


	CheckLeftSideIsOnly1Long(alias.LGenVar, "AllAliasPermutationsAndAddToDatabase")


	copyAlias := CleanCopyAlias(alias)

	leftSideVal := copyAlias.LGenVar[0]


	AddToAliasDatabase(copyAlias)

	//this value is being subtracted from the left so adjust the sign
	leftSideVal.Multiplier = (leftSideVal.Multiplier * -1)

	for i := 0; i < len(copyAlias.RGenVar); i++ {

		newLeftSideVal := copyAlias.RGenVar[i]


		//we are subtracting this item from the right so the sign needs to be changed
		newLeftSideVal.Multiplier = (newLeftSideVal.Multiplier * -1)



		newRightSideSlice := []GenVar{}

		for j := 0; j < len(copyAlias.RGenVar); j++{
			if(j != i){

				newRightSideSlice = append(newRightSideSlice, copyAlias.RGenVar[j])
			}
		}


		newRightSideSlice = append(newRightSideSlice, leftSideVal)


		aliasToAdd := CreateAlias([]GenVar{newLeftSideVal}, newRightSideSlice, copyAlias.LNum, copyAlias.RNum)

		AddToAliasDatabase(aliasToAdd)


	}

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
func SubstituteAnAlias(originalAlias Alias, substituteAlias Alias) (Alias, bool){



	

	CheckLeftSideIsOnly1Long(originalAlias.LGenVar, "SubstituteAnAlias")
	CheckLeftSideIsOnly1Long(substituteAlias.LGenVar, "SubstituteAnAlias")

	dataValid := true

	if(originalAlias.LGenVar[0].Name == substituteAlias.LGenVar[0].Name){
		
	//	fmt.Println("invalid substitution")
		// VerbosePrintln(originalAlias)
		// VerbosePrintln(substituteAlias)

		dataValid = false

		return Alias{}, dataValid


	}



	cleanCopySubstituteAlias := CleanCopyAlias(substituteAlias)


	scaleValSub := substituteAlias.LGenVar[0].Multiplier


	scaledCleanCopySubstituteAlias := ScaleDownEntireAlias(cleanCopySubstituteAlias, scaleValSub)


	//its ok to index 0 since above its checked that there is only one element
	// leftSideMultiplierSub := substituteAlias.LGenVar[0].Multiplier

	cleanCopyRGenVarSubScaled := scaledCleanCopySubstituteAlias.RGenVar
	cleanCopyRNumSubScaled := scaledCleanCopySubstituteAlias.RNum


	cleanCopyRNumOriginal := CleanCopySliceDataFloat(originalAlias.RNum)





	//this is the variable slice of the original alias without the variable to remove
	//and the multiplier of that variable since it will mutliply the newly added substitute values
	originalAliasSubstituteVariableRemoved, multiplierForSubstitute, validRemoval := RemoveExistingGenVarReturnMultiplier(originalAlias.RGenVar, substituteAlias.LGenVar[0].Name)

	VerbosePrintln(multiplierForSubstitute)

	if(!validRemoval){
		dataValid = false
		return Alias{}, dataValid
	}


	for i := 0; i < len(cleanCopyRGenVarSubScaled); i++ {
		cleanCopyRGenVarSubScaled[i].Multiplier = cleanCopyRGenVarSubScaled[i].Multiplier * multiplierForSubstitute
	}

	for i := 0; i < len(cleanCopyRNumSubScaled); i++ {
		cleanCopyRNumSubScaled[i] = cleanCopyRNumSubScaled[i] * multiplierForSubstitute
	}

	originalAliasSubstituteVariableRemoved = append(originalAliasSubstituteVariableRemoved, cleanCopyRGenVarSubScaled...)

	cleanCopyRNumOriginal = append(cleanCopyRNumOriginal, cleanCopyRNumSubScaled...)




	returnAlias := CreateAlias(CleanCopySliceDataGenVar(originalAlias.LGenVar), originalAliasSubstituteVariableRemoved, CleanCopySliceDataFloat(originalAlias.LNum), cleanCopyRNumOriginal)

	var leftSideZero bool

	returnAlias, leftSideZero = FullCleanUp(returnAlias)

	if(leftSideZero){
	//	fmt.Println("error left side 0 in SubstituteAnAlias")	
		// os.Exit(1)
		dataValid = false
	}

	return returnAlias, dataValid


}





func SolutionListener(numberOfSolutionsNeeded int ) []ConcreteSolution{


	soltnsChan := make(chan ConcreteSolution)

	go WorkerSpawnAndAliasListener(soltnsChan)


	solNeededCount := numberOfSolutionsNeeded


	returnSolutionsSlice := []ConcreteSolution{}

	for (solNeededCount > 0) {

		newSolution := <- soltnsChan

		if(!(IsDuplicateConcreteSolution(returnSolutionsSlice, newSolution))){

			returnSolutionsSlice = append(returnSolutionsSlice, newSolution)

			solNeededCount-- 

			fmt.Println("hello!!!!")

			VerbosePrintln(newSolution)

		}

	}

	Solved = true

	for i := 0; i < len(returnSolutionsSlice); i++ {
		VerbosePrintln(returnSolutionsSlice[i])
	}
	
	os.Exit(1)

	return returnSolutionsSlice




}



func IsDuplicateConcreteSolution(soltns []ConcreteSolution, checkVal ConcreteSolution) bool {

	isDuplicate := false 

	for i := 0; i < len(soltns); i++ {
		if(soltns[i].Name == checkVal.Name){
			isDuplicate = true
		}
	}

	return isDuplicate


}


func WorkerSpawnAndAliasListener(soltnsChan chan ConcreteSolution)   {

	if(len(AliasDatabase) == 0){
		fmt.Println("WorkerSpawnAndAliasListener called with an empty AliasDatabase")
		os.Exit(1)
	}

		

	cursor := 0


	for !Solved {


		//alias to send to work
		aliasToSend, canSend := ReadItemFromAliasDataBase(cursor)

		// fmt.Println("Alias to send")
		// VerbosePrintln(aliasToSend)

		if(len(aliasToSend.RGenVar) == 0){
			canSend = false
		}

		if(len(aliasToSend.LGenVar) == 0){
			canSend = false
		}

		if(canSend){

			go WorkOnOneItem(aliasToSend, soltnsChan)

			cursor++ 
		}else{
			if(CursorIsLongAsOrLongerThanDatabase(cursor)){
				time.Sleep(time.Duration(1) * time.Second)
			}else{
				cursor++ 
				time.Sleep(time.Duration(1) * time.Second)
			}
			
		}


		//Checks if main go routine found its solutions

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


func SeedIsDuplicateAlias(seed Alias) bool {

	isDuplicate := false

	if(len(seed.LGenVar) == 0){
		return true
	}

	seedMutex.Lock()
	for i := 0; i < len(SeedsTested); i++ {
		if(TwoAliasesAreEqual(SeedsTested[i], seed, "SeedIsDuplicateAlias")){
			isDuplicate = true
		}
	}
	seedMutex.Unlock()


	return isDuplicate

}

func AddSeedToSeedsTested(seed Alias) {
	seedMutex.Lock()

	SeedsTested = append(SeedsTested, seed)

	seedMutex.Unlock()
}

func PrintSeedsTested() {

	seedMutex.Lock()

	fmt.Println("START---- Printing Seeds Tested")

	for i := 0; i < len(SeedsTested); i++ {
		VerbosePrintln(SeedsTested[i])
	}


	fmt.Println("END---- Printing Seeds Tested")

	seedMutex.Unlock()


}




func WorkOnOneItem(seed Alias, solutionToSend chan ConcreteSolution) {

	if(len(seed.LGenVar) == 0){
		return
	}

	cleanCopySeed := CleanCopyAlias(seed)

	bestSubChoice, dataValid1 := BestCandidateFromDataBase(cleanCopySeed)

	VerbosePrintln("WorkOnOneItem() called")

	VerbosePrintln(cleanCopySeed)

	if(dataValid1){

		bestSub, dataValid2 := SubstituteAnAlias(cleanCopySeed, bestSubChoice)

		

		if(dataValid2){

			fmt.Println("net sub effect")

			VerbosePrintln(bestSub)

			clnBestSub1 := CleanCopyAlias(bestSub)
			clnBestSub2 := CleanCopyAlias(bestSub)
			clnBestSub3 := CleanCopyAlias(bestSub)

			AddToAliasDatabase(clnBestSub3)

			AllAliasPermutationsAndAddToDatabase(clnBestSub2)

			if(len(clnBestSub1.RGenVar) == 0){
				if(IsConcreteSolution(clnBestSub1)){

					solutionFound := ConcreteSolution{clnBestSub1.LGenVar[0].Name,  clnBestSub1.RNum[0]}

					solutionToSend <- solutionFound
				}else{
					fmt.Println("this should be a concrete solution... WorkOnOneItem()")
					os.Exit(1)
				}
			}else{
				WorkOnOneItem(clnBestSub1, solutionToSend)
			}



		}

	}else{
		time.Sleep(time.Duration(1) * time.Second)
		
		//PrintAliasDataBase()
		WorkOnOneItem(cleanCopySeed, solutionToSend)

	}


	// cursor := 0

	// doneWorking := false

	// cleanCopySeed := CleanCopyAlias(seed)

	// isDup := SeedIsDuplicateAlias(cleanCopySeed)
	
	// if(!isDup){
	// 	AddSeedToSeedsTested(cleanCopySeed)
	// 	//PrintSeedsTested()
	// 	//VerbosePrintln(cleanCopySeed)

	// }else{
	// 	return 
	// 	doneWorking = true
	// }


	// for !doneWorking {

		
	// 		//read a new item from the database
	// 		valToWorkWith, dataValid := ReadItemFromAliasDataBase(cursor)

	// 	//	fmt.Println("Val To Work With")
	// 	//	VerbosePrintln(valToWorkWith)

	// 		cleanCopyValToWorkWith := CleanCopyAlias(valToWorkWith)


	// 		PrintOldAliasSubAliasAndNetChange(cleanCopySeed, cleanCopyValToWorkWith, Alias{})


	// 		dontCheckIfEqual := false

	// 		if(len(cleanCopySeed.LGenVar) == 0){
	// 			dataValid = false
	// 			dontCheckIfEqual = true
				
	// 		}


	// 		if(len(cleanCopyValToWorkWith.LGenVar)  == 0 ){
	// 			dataValid = false
	// 			dontCheckIfEqual = true
	// 		}



	// 		if(!dontCheckIfEqual){
	// 			if(TwoAliasesAreEqual(cleanCopySeed, cleanCopyValToWorkWith, "WorkOneItem")){
	// 				dataValid = false
					 
	// 			}


	// 			if(TwoAliasesAreVaritaionsOfEachOther(cleanCopySeed, cleanCopyValToWorkWith)){
	// 				dataValid = false
					
	// 			}

	// 		}






	// 		if(IsImpossibleSubstitution(cleanCopySeed,cleanCopyValToWorkWith) && dataValid){
	// 			dataValid = false
				
	// 		}


	// 		if(dataValid){


	// 			cursor++ 

	// 			//if that new item is helpful
	// 			if(NewAliasEqualsLeftSideVariableNoIncrease(cleanCopySeed, cleanCopyValToWorkWith)){
					


	// 				testSub, subValid := SubstituteAnAlias(CleanCopyAlias(cleanCopySeed), CleanCopyAlias(cleanCopyValToWorkWith) )

	// 				// fmt.Println()
	// 				// fmt.Println()
	// 				// fmt.Println("Test Sub")
	// 				// VerbosePrintln(testSub)
	// 				// fmt.Println()
	// 				// fmt.Println()

					

	// 				if(subValid){

	// 					PrintOldAliasSubAliasAndNetChange(cleanCopySeed, cleanCopyValToWorkWith, testSub)


	// 					if(AliasOnlyHasOneVariableOnTheRight(testSub)){

	// 						cleanCopyTestSub1 := CleanCopyAlias(testSub)

	// 						cleanCopyTestSub2 := CleanCopyAlias(testSub)

	// 						AddToAliasDatabase(cleanCopyTestSub1)

	// 						AllAliasPermutationsAndAddToDatabase(cleanCopyTestSub2)

	// 						go OnlyOneVarLeftOnRightSideWorker(testSub, solutionToSend)

	// 						doneWorking = true
	// 					}else{





	// 						cleanCopyTestSub1 := CleanCopyAlias(testSub)

	// 						cleanCopyTestSub2 := CleanCopyAlias(testSub)

	// 						cleanCopyTestSub3 := CleanCopyAlias(testSub)

	// 						go WorkOnOneItem(cleanCopyTestSub1, solutionToSend)

							
	// 						AddToAliasDatabase(cleanCopyTestSub2)

	// 						AllAliasPermutationsAndAddToDatabase(cleanCopyTestSub3)


							

	// 						if(IsConcreteSolution(testSub)){

	// 							//Full clean up gets called via substitution above so it is 
	// 							//known there is only 1 variable on the left hand side
	// 							//and concrete solution check function checks there is only
	// 							//one constant on the right
	// 							solutionFound := ConcreteSolution{testSub.LGenVar[0].Name,  testSub.RNum[0]}

	// 							solutionToSend <- solutionFound

	// 							doneWorking = true

	// 						}	
	// 					}

	// 				}else{
	// 					doneWorking = true
	// 				}
	// 			}else if(NewAliasReducesVariablesOnRightHandSide(cleanCopySeed, cleanCopyValToWorkWith)){



	// 				testSub, subValid := SubstituteAnAlias(CleanCopyAlias(cleanCopySeed), CleanCopyAlias(cleanCopyValToWorkWith))


	// 				if(subValid){


	// 					PrintOldAliasSubAliasAndNetChange(cleanCopySeed, cleanCopyValToWorkWith, testSub)
					
	// 					if(AliasOnlyHasOneVariableOnTheRight(testSub)){

	// 						cleanCopyTestSub1 := CleanCopyAlias(testSub)

	// 						cleanCopyTestSub2 := CleanCopyAlias(testSub)

	// 						AddToAliasDatabase(cleanCopyTestSub1)

	// 						AllAliasPermutationsAndAddToDatabase(cleanCopyTestSub2)

	// 						go OnlyOneVarLeftOnRightSideWorker(testSub, solutionToSend)

	// 						doneWorking = true
	// 					}else{

	// 						cleanCopyTestSub1 := CleanCopyAlias(testSub)

	// 						cleanCopyTestSub2 := CleanCopyAlias(testSub)

	// 						cleanCopyTestSub3 := CleanCopyAlias(testSub)

	// 						go WorkOnOneItem(cleanCopyTestSub1, solutionToSend)

							
	// 						AddToAliasDatabase(cleanCopyTestSub2)

	// 						AllAliasPermutationsAndAddToDatabase(cleanCopyTestSub3)


							

	// 						if(IsConcreteSolution(testSub)){

	// 							//Full clean up gets called via substitution above so it is 
	// 							//known there is only 1 variable on the left hand side
	// 							//and concrete solution check function checks there is only
	// 							//one constant on the right
	// 							solutionFound := ConcreteSolution{testSub.LGenVar[0].Name,  testSub.RNum[0]}

	// 							solutionToSend <- solutionFound

	// 							doneWorking = true

	// 						}

	// 					}


	// 				}else{
	// 						doneWorking = true
	// 					}
	// 			}


	// 		}else{

	// 			if(CursorIsLongAsOrLongerThanDatabase(cursor)){
	// 				time.Sleep(time.Duration(1) * time.Second)		
	// 			}else{
	// 				cursor++
	// 				time.Sleep(time.Duration(1) * time.Second)		
	// 			}

				
	// 		}


	// 	}


	



}



func OnlyOneVarLeftOnRightSideWorker(alias Alias, solutionToSend chan ConcreteSolution) {


//	cursor := 0

//	testedAllVals := false

	if(len(alias.RGenVar) == 0 || len(alias.LGenVar) == 0){
		fmt.Println("invalid input to OnlyOneVarLeftOnRightSideWorker")
		os.Exit(1)
	}

	cleanCopyInputAlias := CleanCopyAlias(alias)

//	nameToMatchRight := alias.RGenVar[0].Name

//	nameToMatchLeft := alias.LGenVar[0].Name


	solutionAlias, dataValid1 := GetValidCanidateForOneVarLeftCase(cleanCopyInputAlias)

	fmt.Println("SOLUTION ALIAS PRE SUB")

	PrintOldAliasSubAliasAndNetChange(cleanCopyInputAlias, solutionAlias, Alias{})


	if(dataValid1){

		netSolutionAlias, dataValid2 := SubstituteAnAlias(cleanCopyInputAlias, solutionAlias)

		fmt.Println("SOLUTION ALIAS")

		PrintOldAliasSubAliasAndNetChange(cleanCopyInputAlias, solutionAlias, netSolutionAlias)


		if(dataValid2){
		
			if(IsConcreteSolution(netSolutionAlias)){

								
				solutionFound := ConcreteSolution{netSolutionAlias.LGenVar[0].Name,  netSolutionAlias.RNum[0]}

				solutionToSend <- solutionFound

	//			testedAllVals = true

								

			}else{
				fmt.Println("this must be a concrete solution...")
				os.Exit(1)
			}

		}
	}



// 	for !testedAllVals {







// 		compareAlias, dataValid :=  ReadItemFromAliasDataBase(cursor)

		

// 		cleanCopyCompareAlias := CleanCopyAlias(compareAlias)


// 		fmt.Println("ONE LEFT ATTEMPT")
// 		PrintOldAliasSubAliasAndNetChange(cleanCopyInputAlias, cleanCopyCompareAlias, Alias{})



// 		dontCheckIfEqual := false

// 		if(len(cleanCopyCompareAlias.LGenVar)  == 0){
// 				dataValid = false
// 				dontCheckIfEqual = true		
// 		}


// 		if(!dontCheckIfEqual){


// 			if(TwoAliasesAreEqual(cleanCopyInputAlias, cleanCopyCompareAlias, "OnlyOneVarLeftOnRightSideWorker")){
// 					dataValid = false
					
// 			}


// 			if(TwoAliasesAreVaritaionsOfEachOther(cleanCopyInputAlias, cleanCopyCompareAlias) && dataValid){
// 				dataValid = false
				

// 			}	
// 		}



// 		if(IsImpossibleSubstitution(cleanCopyInputAlias, cleanCopyCompareAlias) && dataValid){
// 			dataValid = false
			
// 		}


// 		var compareAliasHasRightSideVars bool

// 		if(len(cleanCopyCompareAlias.RGenVar) == 0){
// 			compareAliasHasRightSideVars = false
// 		}else{
// 			compareAliasHasRightSideVars = true
// 		}


// 		if(dataValid && compareAliasHasRightSideVars){
// 			cursor++



// 			CheckLeftSideIsOnly1Long(compareAlias.LGenVar, "OnlyOneVarLeftOnRightSideWorker")

// 			compareAliasLeftVarName := cleanCopyCompareAlias.LGenVar[0].Name

// 			compareAliasRightVarName := cleanCopyCompareAlias.RGenVar[0].Name


// 			fmt.Println(compareAliasRightVarName, nameToMatchLeft, compareAliasLeftVarName, nameToMatchRight)


// 			if(compareAliasRightVarName == nameToMatchLeft && compareAliasLeftVarName == nameToMatchRight){

					

// 					testSub, subValid := SubstituteAnAlias(cleanCopyInputAlias, cleanCopyCompareAlias)


// 					if(subValid){

// 						PrintOldAliasSubAliasAndNetChange(cleanCopyInputAlias, cleanCopyCompareAlias, testSub)


// 						cleanCopyTestSub1 := CleanCopyAlias(testSub)

						

// 						AddToAliasDatabase(cleanCopyTestSub1)

// 						if(IsConcreteSolution(testSub)){

							

// 							//Full clean up gets called via substitution above so it is 
// 							//known there is only 1 variable on the left hand side
// 							//and concrete solution check function checks there is only
// 							//one constant on the right
// 							solutionFound := ConcreteSolution{testSub.LGenVar[0].Name,  testSub.RNum[0]}

// 							solutionToSend <- solutionFound

// 							testedAllVals = true

							

// 						}else{
// 							fmt.Println("this must be a concrete solution...")
// 							os.Exit(1)
// 						}

// 			} 



// 		}

// 	}else if(dataValid && !compareAliasHasRightSideVars){

					
				
					

// 					testSub, subValid := SubstituteAnAlias(cleanCopyInputAlias, cleanCopyCompareAlias)

					

// 					if(subValid){

// 						cleanCopyTestSub1 := CleanCopyAlias(testSub)

						
// 						PrintOldAliasSubAliasAndNetChange(cleanCopyInputAlias, cleanCopyCompareAlias, testSub)
						

// 						AddToAliasDatabase(cleanCopyTestSub1)

// 						if(IsConcreteSolution(testSub)){

							

// 							//Full clean up gets called via substitution above so it is 
// 							//known there is only 1 variable on the left hand side
// 							//and concrete solution check function checks there is only
// 							//one constant on the right
// 							solutionFound := ConcreteSolution{testSub.LGenVar[0].Name,  testSub.RNum[0]}

// 							solutionToSend <- solutionFound

// 							testedAllVals = true

							

// 						}else{
// 							fmt.Println("this must be a concrete solution...")
// 							os.Exit(1)
// 						}

// 					}
// 	}else{
// 		if(CursorIsLongAsOrLongerThanDatabase(cursor)){
// 			time.Sleep(time.Duration(1) * time.Second)		
// 		}else{
// 			cursor++
// 			time.Sleep(time.Duration(1) * time.Second)		
// 		}	
// 	}


// }

}


func TwoAliasesAreVaritaionsOfEachOther(a1 Alias, a2 Alias) bool {

	if(len(a1.LGenVar) != len(a2.LGenVar) || len(a1.RGenVar) != len(a2.RGenVar)) {
		

		return false
	}


	if(len(a1.LGenVar) == 0 || len(a2.LGenVar) == 0) {
		return false
	}

	CheckLeftSideIsOnly1Long(a1.LGenVar, "TwoAliasesAreNotInverseOfEachOther")
	CheckLeftSideIsOnly1Long(a2.LGenVar, "TwoAliasesAreNotInverseOfEachOther")


	a1CleanCopy := CleanCopyAlias(a1)
	a2CleanCopy := CleanCopyAlias(a2)


	fmt.Println()

	


	a1LeftSideVarName := a1CleanCopy.LGenVar[0].Name

	a2RightSideHasCorrectVariable := false


	fmt.Println()

	for i := 0; i < len(a2CleanCopy.RGenVar); i++ {
		if(a2CleanCopy.RGenVar[i].Name == a1LeftSideVarName){
			a2RightSideHasCorrectVariable = true

			a2LVar := a2CleanCopy.LGenVar[0]

			a2LVar.Multiplier = (a2LVar.Multiplier * (-1) )

			currentVar := a2CleanCopy.RGenVar[i]

			currentVar.Multiplier = (currentVar.Multiplier * (-1))


			a2CleanCopy.LGenVar = []GenVar{currentVar}

			a2CleanCopy.RGenVar[i] = a2LVar

		}
	} 




	if(!a2RightSideHasCorrectVariable) {


		return false
	}

	a1CleanCopy = ScaleDownEntireAlias(a1CleanCopy, a1CleanCopy.LGenVar[0].Multiplier)
	a2CleanCopy = ScaleDownEntireAlias(a2CleanCopy, a2CleanCopy.LGenVar[0].Multiplier)

	

	if(TwoAliasesAreEqual(a1CleanCopy, a2CleanCopy, "TwoAliasesAreVaritaionsOfEachOther")){

		return true
	}else{



		return false
	}





}



func AliasOnlyHasOneVariableOnTheRight(aliasInput Alias) bool {

	if( len(aliasInput.RGenVar) == 1 ){
		return true
	}else{
		return false
	}


}



func CursorIsLongAsOrLongerThanDatabase(cursor int) bool {
	
	isLongOrLonger := false

	mutex.Lock()

		if(cursor >= len(AliasDatabase)){
			isLongOrLonger = true
		}

	mutex.Unlock()

	return isLongOrLonger

}



func IsConcreteSolution(checkAlias Alias) bool {
	if(len(checkAlias.LGenVar)  == 1 && len(checkAlias.RGenVar) == 0){
		
		if(len(checkAlias.LNum) == 1){
			if(checkAlias.LNum[0] != 0){
				return false
			}
		}

		return true
	}else{
		return false
	}
} 





func IsImpossibleSubstitution(oldAlias Alias, subAlias Alias) bool {


	if( (len(oldAlias.LGenVar) == 0)  || (len(subAlias.LGenVar) == 0)){
		return true
	}

	CheckLeftSideIsOnly1Long(oldAlias.LGenVar, "IsImpossibleSubstitution")
	CheckLeftSideIsOnly1Long(subAlias.LGenVar, "IsImpossibleSubstitution")

	subAliasLeftVar := subAlias.LGenVar[0]

	oldAliasLeftVar	:= oldAlias.LGenVar[0]


	//if they are both referring to the same variable after a full cleanup
	//then there's no use, the substitution cannot occur
	if(oldAliasLeftVar.Name == subAliasLeftVar.Name){
		return true
	}

	isImpossible := true

	for i := 0; i < len(oldAlias.RGenVar); i++ {
		if(oldAlias.RGenVar[i].Name == subAliasLeftVar.Name){
			isImpossible = false
		}
	}


	return isImpossible



}


func AddToAliasDatabase(newAlias Alias) {




	mutex.Lock()

	notDuplicateValue := true

	

	if(len(newAlias.LGenVar) == 0){
		return
	}

	CheckLeftSideIsOnly1Long(newAlias.LGenVar, "AddToAliasDatabase")

	for i := 0; i < len(AliasDatabase); i++ {
		if(TwoAliasesAreEqual(CleanCopyAlias(newAlias), CleanCopyAlias(AliasDatabase[i]), "AddToAliasDatabase")){
			notDuplicateValue = false
		}

	}	




   	if(notDuplicateValue){
   		//fmt.Println("DATABASE VAL ADDED")
   		//VerbosePrintln(newAlias)
   		fmt.Println()
   		fmt.Println("NEW VAL ADDED TO DATABASE")
   		VerbosePrintln(newAlias)

		AliasDatabase = append(AliasDatabase, newAlias)
	}
    mutex.Unlock()

    if(notDuplicateValue){
    	PrintAliasDataBase()
    }



}



func PrintAliasDataBase() {
	
	mutex.Lock()
   	

	fmt.Println("START---- Printing Alias Database")

	for i := 0; i < len(AliasDatabase); i++ {
		fmt.Print(i)
		VerbosePrintln(AliasDatabase[i])
	}

	fmt.Println("END---- Printing Alias Database")

    mutex.Unlock()

}



func ReadItemFromAliasDataBase(index int) (Alias, bool) {

	returnDataValid := false

	var returnData Alias



	mutex.Lock()

		if(!(index >= len(AliasDatabase))){


			returnDataValid = true

			//when reading data from the data base clean copy is used for the slices because otherwise pointers to the underlying database itself would be
			//passed on... not good

			returnData = CreateAlias(CleanCopySliceDataGenVar(AliasDatabase[index].LGenVar), CleanCopySliceDataGenVar(AliasDatabase[index].RGenVar), CleanCopySliceDataFloat(AliasDatabase[index].LNum), CleanCopySliceDataFloat(AliasDatabase[index].RNum))

			
		}

	mutex.Unlock()

	if(returnDataValid){


		//PrintAliasDataBase()

		return returnData, returnDataValid
	}else{
		return Alias{}, false
	}



}




// this function passes all tests
func FullCleanUp(genVarInput Alias)  (Alias, bool) {

	cleanCopy := CleanCopyAlias(genVarInput)

	cleanCopy = SimplifyGenVarRightHandGenVarSlice(cleanCopy)


	cleanCopy = SimplifyRightHandNumSlice(cleanCopy)
	

	cleanCopy = MoveVarsEqualToLeftHandSideToLeftSide(cleanCopy)
	

	var leftSideZero bool

	cleanCopy, leftSideZero = RemoveZerosWarnIfLeftHandSideZero(cleanCopy)


	// if(leftSideZero){
	// 	fmt.Println("left hand side 0")
	// 	os.Exit(1)
	// }


	CheckLeftSideIsOnly1Long(cleanCopy.LGenVar, "FullCleanUp")


	return cleanCopy, leftSideZero

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

	if(variablesMoved == 0){
		return CleanCopyAlias(genVarInput)
	}

	cleanCopyRightHandSideForRemoval := CleanCopySliceDataGenVar(genVarInput.RGenVar)


	//discard the multiplier variable returned since here the only goal is to remove the element
	cleanCopyRightHandSideForRemoval, _, _ =  RemoveExistingGenVarReturnMultiplier(cleanCopyRightHandSideForRemoval, leftHandVarName)




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


//this function passes all tests
func NewAliasEqualsLeftSideVariableNoIncrease(oldAlias Alias, newAlias Alias) bool {


	// the net result is what matters 

	oldAliasRightHandCountInitial := len(oldAlias.RGenVar)




	// a clean copy of old and new is taken

	cleanCopyOldAlias := CleanCopyAlias(oldAlias)
	cleanCopyNewAlias := CleanCopyAlias(newAlias)


	CheckLeftSideIsOnly1Long(cleanCopyOldAlias.LGenVar, "NewAliasEqualsLeftSideVariableNoIncrease")
	CheckLeftSideIsOnly1Long(cleanCopyNewAlias.LGenVar, "NewAliasEqualsLeftSideVariableNoIncrease")



	newAliasRightHandSideContainsOldAliasLeftHandSideVariable := false

	oldAliasLeftHandSideVarName := cleanCopyOldAlias.LGenVar[0].Name

	for i := 0; i < len(cleanCopyNewAlias.RGenVar); i++ {
		if(cleanCopyNewAlias.RGenVar[i].Name == oldAliasLeftHandSideVarName){
			newAliasRightHandSideContainsOldAliasLeftHandSideVariable = true
		}
	}

	if(!newAliasRightHandSideContainsOldAliasLeftHandSideVariable){
		return false
	}


	cleanSubstitute, dataValid := SubstituteAnAlias(cleanCopyOldAlias, cleanCopyNewAlias)

	if(!dataValid){
		return false
	}


	var leftSideZero bool

	cleanSubstitute, leftSideZero = FullCleanUp(cleanSubstitute)


	if(leftSideZero){
		fmt.Println("left side zero NewAliasHasVariableAlreadyOnLeft")
		os.Exit(1)
	}


	cleanSubstituteRightHandSideCount := len(cleanSubstitute.RGenVar)

	if(oldAliasRightHandCountInitial > cleanSubstituteRightHandSideCount){
		return true
	}



	return false


}

//this function passes all tests 
func NewAliasReducesVariablesOnRightHandSide(oldAlias Alias, newAlias Alias) bool {


	// the net result is what matters 

	oldAliasRightHandCountInitial := len(oldAlias.RGenVar)


	// a clean copy of old and new is taken

	cleanCopyOldAlias := CleanCopyAlias(oldAlias)
	cleanCopyNewAlias := CleanCopyAlias(newAlias)


	CheckLeftSideIsOnly1Long(cleanCopyOldAlias.LGenVar, "NewAliasEqualsLeftSideVariableNoIncrease")
	CheckLeftSideIsOnly1Long(cleanCopyNewAlias.LGenVar, "NewAliasEqualsLeftSideVariableNoIncrease")



	cleanSubstitute, dataValid := SubstituteAnAlias(cleanCopyOldAlias, cleanCopyNewAlias)

	if(!dataValid){
		return false
	}

	var leftSideZero bool

	cleanSubstitute, leftSideZero = FullCleanUp(cleanSubstitute)


	if(leftSideZero){
		fmt.Println("left side zero NewAliasReducesVariablesOnRightHandSide")
		os.Exit(1)
	}


	cleanSubstituteRightHandSideCount := len(cleanSubstitute.RGenVar)

	if(oldAliasRightHandCountInitial > cleanSubstituteRightHandSideCount){
		return true
	}



	return false


}




















func Init() {
	mutex = &sync.Mutex{}
	solvedMutex = &sync.Mutex{}
	seedMutex = &sync.Mutex{}
	printAliasMutex = &sync.Mutex{}
	AliasDatabase = []Alias{}
	SeedsTested = []Alias{}
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


	leftNumCleanCopy := CleanCopySliceDataFloat(leftNum)

	rightNumCleanCopy := CleanCopySliceDataFloat(rightNum)

	if(len(leftNumCleanCopy)  == 0){
		leftNumCleanCopy = []float64{0}
	}

	if(len(rightNumCleanCopy)  == 0){
		rightNumCleanCopy = []float64{0}
	}

	return Alias{leftGenVar, rightGenVar, leftNumCleanCopy, rightNumCleanCopy}
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


func CleanCopySliceDataInt(input []int) []int {

	outputSlice := []int{}

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



func RemoveExistingGenVarReturnMultiplier(genVarSlice []GenVar, removeName string) ([]GenVar, float64, bool) {

	dataValid := true

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
		//fmt.Println("couldn't remove value from slice or couldn't find multiplier")

		dataValid = false		

		// os.Exit(1)

	}

	if(removedItems > 1){
		fmt.Println("more than one item removed")
		os.Exit(1)
	}

	return returnSlice, returnMultiplier, dataValid


}



func CheckLeftSideIsOnly1Long(genVarInput []GenVar, functionCaller string){
	//if there's more variables on the left than one something went wrong earlier in the program
	if(len(genVarInput) != 1){
		fmt.Println("equation doesn't only have one variable on the left side from function:", functionCaller)
		VerbosePrintln(genVarInput)
		PrintAliasDataBase()
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

//this function passes all tests
func ScaleDownEntireAlias(aliasInput Alias, scaleVal float64) Alias{

	cleanAliasInputCopy := CleanCopyAlias(aliasInput)

	cleanAliasInputCopy.LGenVar = ScaleDownSliceGenVar(cleanAliasInputCopy.LGenVar, scaleVal)

	cleanAliasInputCopy.RGenVar = ScaleDownSliceGenVar(cleanAliasInputCopy.RGenVar, scaleVal)

	cleanAliasInputCopy.LNum = ScaleDownSliceFloat(cleanAliasInputCopy.LNum, scaleVal)

	cleanAliasInputCopy.RNum = ScaleDownSliceFloat(cleanAliasInputCopy.RNum, scaleVal)

	return cleanAliasInputCopy

}


//this function passes all tests
func TwoAliasesAreEqual(alias1 Alias, alias2 Alias, parent string) bool {


	CheckLeftSideIsOnly1Long(alias1.LGenVar, "TwoAliasesAreEqual")
	CheckLeftSideIsOnly1Long(alias2.LGenVar, "TwoAliasesAreEqual")


	alias1LGenVar := alias1.LGenVar

	alias2LGenVar := alias2.LGenVar

	if((alias1LGenVar[0].Name != alias2LGenVar[0].Name) && (alias1LGenVar[0].Multiplier != alias2LGenVar[0].Multiplier) ){
		return false
	}
	
	scaleValAlias1 := alias1.LGenVar[0].Multiplier

	scaledAlias1 := ScaleDownEntireAlias(alias1, scaleValAlias1)


	scaleValAlias2 := alias2.LGenVar[0].Multiplier

	scaledAlias2 := ScaleDownEntireAlias(alias2, scaleValAlias2)



	scaledAlias1, _ = FullCleanUp(scaledAlias1)
	scaledAlias2, _ = FullCleanUp(scaledAlias2)


	if(len(scaledAlias1.RGenVar) != len(scaledAlias2.RGenVar)){
		return false
	}

	matchesNeededToConfirmEqual := len(scaledAlias1.RGenVar)

	matchesMade := 0

	for i := 0; i < len(scaledAlias1.RGenVar); i++ {

		currentVal := scaledAlias1.RGenVar[i]

		for j := 0; j < len(scaledAlias2.RGenVar); j++ {
			compareVal := scaledAlias2.RGenVar[j]

			if((currentVal.Name == compareVal.Name) && (currentVal.Multiplier == compareVal.Multiplier)){
				matchesMade++
			}

		}


	}

	if(matchesMade != matchesNeededToConfirmEqual){
		return false
	}

	summationLAlias1 := float64(0)

	for i := 0; i < len(scaledAlias1.LNum); i++ {
		summationLAlias1 = summationLAlias1 + scaledAlias1.LNum[i]
	}

	summationLAlias2 := float64(0)

	for i := 0; i < len(scaledAlias2.LNum); i++ {
		summationLAlias2 = summationLAlias2 + scaledAlias2.LNum[i]
	}

	if(summationLAlias1 != summationLAlias2){
		return false
	}	



	summationRAlias1 := float64(0)

	for i := 0; i < len(scaledAlias1.RNum); i++ {
		summationRAlias1 = summationRAlias1 + scaledAlias1.RNum[i]
	}

	summationRAlias2 := float64(0)

	for i := 0; i < len(scaledAlias2.RNum); i++ {
		summationRAlias2 = summationRAlias2 + scaledAlias2.RNum[i]
	}

	if(summationRAlias1 != summationRAlias2){
		return false
	}	


	return true



}




func PrintOldAliasSubAliasAndNetChange(old Alias, sub Alias, net Alias) {

	printAliasMutex.Lock()


	OldGlobal = old

	SubGlobal = sub

	NetGlobal = net

	fmt.Println("OLD ALIAS")
	VerbosePrintln(OldGlobal)
	fmt.Println()


	fmt.Println("SUB ALIAS")
	VerbosePrintln(SubGlobal)
	fmt.Println()

	fmt.Println("NET ALIAS")
	VerbosePrintln(NetGlobal)
	fmt.Println()

	printAliasMutex.Unlock()


}





















//returns a value from the data base that is a valid choice
func GetValidCanidateForOneVarLeftCase(aliasOneVarLeft Alias) (Alias, bool) {

	cleanCopyOfAlias := CleanCopyAlias(aliasOneVarLeft)

	if(len(cleanCopyOfAlias.LGenVar) != 1 && len(cleanCopyOfAlias.LGenVar) != 1){
		fmt.Println("invalid input to GetValidCanidateForOneVarLeftCase")
		os.Exit(1)
	}

	CheckLeftSideIsOnly1Long(cleanCopyOfAlias.LGenVar, "GetValidCanidateForOneVarLeftCase")

	leftHandName := cleanCopyOfAlias.LGenVar[0].Name

	rightHandName := cleanCopyOfAlias.RGenVar[0].Name

	mutex.Lock()

		for i := 0; i < len(AliasDatabase); i++ {

			cleanCopyPosition := CleanCopyAlias(AliasDatabase[i])



			if(len(cleanCopyPosition.RGenVar) != 1 && len(cleanCopyPosition.RGenVar) != 0){
				continue
			}

			CheckLeftSideIsOnly1Long(cleanCopyPosition.LGenVar, "GetValidCanidateForOneVarLeftCase")

			isLength1 := false

			isLength0 := false

			if(len(cleanCopyPosition.RGenVar) == 1){
				isLength1 = true
			}

			if(len(cleanCopyPosition.RGenVar) == 1){
				isLength0 = true
			}



			if(cleanCopyPosition.LGenVar[0].Name == rightHandName && isLength1){

				if(cleanCopyPosition.RGenVar[0].Name == leftHandName){

					dontCheckIfEqual := false

					if(len(cleanCopyPosition.LGenVar) == 0){
						
						dontCheckIfEqual = true
						
					}


					if(!dontCheckIfEqual){

						dontReturn := false

						if(TwoAliasesAreEqual(cleanCopyOfAlias, cleanCopyPosition, "GetValidCanidateForOneVarLeftCase")){
							dontReturn = true
							 
						}


						if(TwoAliasesAreVaritaionsOfEachOther(cleanCopyOfAlias, cleanCopyPosition)){
							dontReturn = true
							
						}

						if(!dontReturn){
							return cleanCopyPosition, true
						}

					}					

					
				}

			}else if(cleanCopyPosition.LGenVar[0].Name == rightHandName && isLength0){

					dontCheckIfEqual := false

					if(len(cleanCopyPosition.LGenVar) == 0){
						
						dontCheckIfEqual = true
						
					}


					if(!dontCheckIfEqual){

						dontReturn := false

						if(TwoAliasesAreEqual(cleanCopyOfAlias, cleanCopyPosition, "GetValidCanidateForOneVarLeftCase")){
							dontReturn = true
							 
						}


						if(TwoAliasesAreVaritaionsOfEachOther(cleanCopyOfAlias, cleanCopyPosition)){
							dontReturn = true
							
						}

						if(!dontReturn){
							return cleanCopyPosition, true
						}

					}									

			}


		}

	mutex.Unlock()


	return Alias{}, false


}





func BestCandidateFromDataBase(inputAlias Alias) (Alias, bool) {


	CheckLeftSideIsOnly1Long(inputAlias.LGenVar, "BestCandidateFromDataBase")

	if(len(inputAlias.RGenVar) == 0){
		fmt.Println("Interesting case, a check in worker go routine should not have sent this type")
		os.Exit(1)

	}

	clnInput := CleanCopyAlias(inputAlias)


	validSub := false
	reducesVars := false
	numberVarsReduced := 0

	bestChoice := Alias{}


	mutex.Lock()


		for i := 0; i < len(AliasDatabase); i++ {
			dblCleanInput := CleanCopyAlias(clnInput)
			dtbsClean := CleanCopyAlias(AliasDatabase[i])

			if(isValidCompareAlias(dtbsClean)){
				if(!TwoAliasesDiffer(clnInput, dtbsClean)){
					continue
				}
			}

			validSubNew, reducesVarsNew, numberVarsReducedNew := NetEffectOfSubstitution(dblCleanInput, dtbsClean)

			if( (!validSub) && validSubNew){
				validSub = validSubNew
				bestChoice = dtbsClean
			}else if(validSub && !(reducesVars) && reducesVarsNew){
				reducesVars = reducesVarsNew
				bestChoice = dtbsClean
			}else if(validSub && reducesVars && (numberVarsReducedNew > numberVarsReduced)){
				numberVarsReduced = numberVarsReducedNew
				bestChoice = dtbsClean
			}



		}


	mutex.Unlock()


	if(!validSub || !reducesVars){
		return Alias{}, false
	}else{
		return bestChoice, true
	}

}


//Assume clean copy is passed
func NetEffectOfSubstitution(oldA Alias, subA Alias) (bool, bool, int) {


	validSubstitution := false
	reducesVariablesOnRight := false
	numberVariablesReduced := 0

	initialVariableCount := len(oldA.RGenVar)


	netA, dataValid := SubstituteAnAlias(oldA, subA)

	if(!dataValid){
		return validSubstitution, reducesVariablesOnRight, numberVariablesReduced
	}else{
		validSubstitution = true
	}


	newVariableCount := len(netA.RGenVar)


	if(!(initialVariableCount > newVariableCount)){
		return validSubstitution, reducesVariablesOnRight, numberVariablesReduced
	}else{
		reducesVariablesOnRight = true
		numberVariablesReduced = initialVariableCount - newVariableCount
		return validSubstitution, reducesVariablesOnRight, numberVariablesReduced
	}



}


//this function is read only
//no clean copy required
func TwoAliasesDiffer(alias1 Alias, alias2 Alias) bool {


			equal := TwoAliasesAreEqual(alias1, alias2, "OnlyOneVarLeftOnRightSideWorker")
					
			if(equal){
				return false
			}

			variations := TwoAliasesAreVaritaionsOfEachOther(alias1, alias2) 

			if(variations){
				return false
			}

			return true

}


func isValidCompareAlias(aliasInput Alias) bool {
	if(len(aliasInput.LGenVar) != 0 && len(aliasInput.RGenVar) != 0){
		return true
	}else{
		return false
	}
}


//the program may not initially see some substitutions
//for instance take take the equation a = b + c
//if it is also known that c = d
//and b = 2d
//the program won't see this
//however if we look for PseudoNames
//which are essentially names a variable could equal
//if a substitution was made, then more solutions can be found

//this function passes all tests
func GetPseudoNamesForRGenVar(varName string) VarPseudoNames {

	valsSlice := []float64{}

	pseudoNamesSlice := [][]string{}

	//the variable is can always be referred to as just itself
	//so add the variables own name to pseudonames
	pseudoNamesSlice = append(pseudoNamesSlice, []string{varName})

	scaledDownMultipliersToVars := []float64{}

	valsSlice = append(valsSlice, 0)

	mutex.Lock()

		for i := 0; i < len(AliasDatabase); i++ {

			clnDB := CleanCopyAlias(AliasDatabase[i])

			if(isValidCompareAlias(clnDB)) {
				if(clnDB.LGenVar[0].Name == varName){
					pseudoNameInner := []string{}

					leftSideScaleVal := clnDB.LGenVar[0].Multiplier

					for i := 0; i < len(clnDB.RGenVar); i++ {
						pseudoNameInner = append(pseudoNameInner, clnDB.RGenVar[i].Name)
						scaledDownMultipliersToVars = append(scaledDownMultipliersToVars, (clnDB.RGenVar[i].Multiplier)/(leftSideScaleVal))
					}
					if(len(clnDB.RNum) == 0){
						valsSlice = append(valsSlice, 0)
					}else if(len(clnDB.RNum) == 1){
						valsSlice = append(valsSlice, (clnDB.RNum[0]/leftSideScaleVal) )
					}else{
						fmt.Println("nums not cleaned up GetPseudoNamesForRGenVar")
						os.Exit(1)
					}

					if(!(IsDuplicatePseudoName(pseudoNamesSlice, pseudoNameInner))){

							pseudoNamesSlice = append(pseudoNamesSlice, pseudoNameInner)

					}

				}
			}

		}

	mutex.Unlock()


	return VarPseudoNames{pseudoNamesSlice, valsSlice, varName}

}



func IsDuplicatePseudoName(pseudoNamesSlice [][]string, pseudoName []string) bool {

	totalMatchesNeeded := len(pseudoName)

	for i := 0; i < len(pseudoNamesSlice); i++ {
		currentPseudoName := pseudoNamesSlice[i]


		//if they are the same length they could be identical
		if(len(currentPseudoName) == len(pseudoName)){

			if(TwoStringsSlicesContainsSameVars(currentPseudoName, pseudoName, totalMatchesNeeded)){
				return true
			}
		}

	}

	return false


}

func TwoStringsSlicesContainsSameVars(slice1 []string, slice2 []string, matchesNeeded int) bool {



	totalMatches := 0

	restrictedIndicesSecondSlice := []int{}

	for i := 0; i < len(slice1); i++ {

		currentVal := slice1[i]

		for j := 0; j < len(slice2); j++ {

			if(!isRestrictedIndex(restrictedIndicesSecondSlice, j)){
				checkVal := slice2[j]

				if(currentVal == checkVal){
					totalMatches++
					restrictedIndicesSecondSlice = append(restrictedIndicesSecondSlice, j)
					break
				}

			}
		}


	}	


	if(totalMatches == matchesNeeded){
		return true
	}else{
		return false
	}

}


//say A has pseudo name [x y]
//B has pseudo name[y]
//C has pseudo name[z]
//D has pseudo name[y z]
//E has pseduo name[x z]

//the summation is x y z aka 3 unique vars when combining A B C D E

//to note, parent left var is the left side variable that the parent equation has
//so for instance.. if  
// x = A + B + C + D + E
// x is the parent variable, and upon cleaning up the equation, any x substitute 
//will be moved to the left thereby reducing the right hand side by one further,
//thus the parentLeftVar string helps account for this
func SumOfPseudoNamesNetChangeIsGood(pseudoNamesSlice [][]string, parentLeftVar string) (int, bool) {

	seenMap := make(map[string]bool)

	preSumVarCount := len(pseudoNamesSlice)

	parentVarPresent := false

	for i := 0; i < len(pseudoNamesSlice); i++ {
		currentPseudoName := pseudoNamesSlice[i]

		for j := 0; j < len(currentPseudoName); j++ {
			seenMap[currentPseudoName[j]] = true	
			if(currentPseudoName[j] == parentLeftVar){
				parentVarPresent = true
			}
		}
	}

	

	postSumVarCount := len(seenMap)

	if(parentVarPresent){
		postSumVarCount = postSumVarCount - 1
	}


	if(preSumVarCount > postSumVarCount){

		return (preSumVarCount - postSumVarCount), true
	}else{
		return (preSumVarCount - postSumVarCount), false
	}

	
}


func ConvertGoodNetChangePseudoNamesAndAddToDataBase(){
	//TODO 
	//even if during the search for the best pseudonames aka substituions,
	//if one substitution wasn't the best it doesn't mean it couldn't be added to 
	//the database and further help out other substitutions

	//implement converting pseudonames to substitutions and adding to the data base

	//do the following....

	//in the method SumOfPseudoNamesNetChangeIsGood() change the input to 
	//be VarPseudoNames that way each Pseudo name has a parent, then use 
	//GetAliasFromDatabase via pseudoname, figure out a way to also get the equation
	//being referred to
	//and then perform the substitutions and add it to the database

	//this is a good function so best to implement well, when more awake...





}




func GetDataBaseItemByPseudoName(leftVarName string, pseudoName []string) (Alias, bool) {

	returnAlias := Alias{}

	dataValid := false

	mutex.Lock()

		for i := 0; i < len(AliasDatabase); i++ {

			clnDB := CleanCopyAlias(AliasDatabase[i])

			if(len(clnDB.RGenVar) != len(pseudoName)){
				continue
			}


			if(clnDB.LGenVar[0].Name == leftVarName){
				if(TwoStringsSlicesContainsSameVars(GenVarSliceToVarNameStringSlice(clnDB.RGenVar), pseudoName, len(pseudoName))) {
					returnAlias = clnDB
					dataValid = true
					break
				}
			}

		}

	mutex.Unlock()


	return returnAlias, dataValid


}


func GenVarSliceToVarNameStringSlice(gvSlice []GenVar) []string {

	returnSlice := []string{}

	for i := 0; i < len(gvSlice); i++ {
		returnSlice = append(returnSlice, gvSlice[i].Name)
	}

	return returnSlice


}


func CircularSolutionSolver(solutionChan chan ConcreteSolution) {

	doneSolving = false

	cursor := 0

	for !doneSolving {

		val, dataValid := ReadItemFromAliasDataBase(cursor)



		if(dataValid){

			clnVal := CleanCopyAlias(val)

			FindBestSubstitutionForAlias(clnVal, solutionChan)

			cursor++ 
		}else{
			cursor = 0
		}


	}


}



func FindBestSubstitutionForAlias(aliasInput Alias, solutionChan chan ConcreteSolution) {

	CheckLeftSideIsOnly1Long(aliasInput.LGenVar,"FindBestSubstitutionForAlias")


	parentLeftVar := aliasInput.LGenVar[0]

	if(len(aliasInput.RGenVar) == 0){
		fmt.Println("Invalid Alias to test FindBestSubstitutionForAlias")
		os.Exit(1)
	}

	clnInputalias := CleanCopyAlias(aliasInput)

	varPseudoNamesSlice := []VarPseudoNames{}

	for i := 0; i < len(clnInputalias.RGenVar); i++ {

		currentVar := clnInputalias.RGenVar[i]

		currentVarName := currentVar.Name

		pseudoNamesCurrentVar := GetPseudoNamesForRGenVar(currentVarName)

		varPseudoNamesSlice = append(varPseudoNamesSlice, pseudoNamesCurrentVar)

	}

	BestAliasSliceForSubstitution(varPseudoNamesSlice, parentLeftVar, clnInputalias)


}




func BestAliasSliceForSubstitution(varsWithPseudoNames []VarPseudoNames, parentLeftVar string, inputAlias Alias) SolutionItem {

	//max combos is (2^n) - 1 where n is the number of different items 

	maxNumberCombos := 2


	//minus 1 since 2 is already present above
	for i := 0; i < (len(varsWithPseudoNames) - 1); i++ {
		maxNumberCombos = maxNumberCombos * 2
	}
	
	maxNumberCombos = maxNumberCombos - 1









	// highestNetChange := 0

	// binaryCursorValForHighestNetChange := ""

	// cursorForHighestNetChange := []int{}


	cursor := 1

	bestSolutions := []SolutionItem{}

	for (cursor <= maxNumberCombos){


		//the binary cursor ensures all possible
		//variable combinations get chosen
		binaryCursor := fmt.Sprintf("%b", cursor)

		for len(binaryCursor) < len(varsWithPseudoNames){
			binaryCursor = "0" + binaryCursor
		}

		if(len(binaryCursor) != len(varsWithPseudoNames)){
			fmt.Println("binary cursor wrong size")
			os.Exit(1)
		}

		activeVars := []VarPseudoNames{}

		for i := 0; i < len(binaryCursor); i++ {

			bit := rune(binaryCursor[i])

			if(bit == rune('1')){
				activeVars = append(activeVars, varsWithPseudoNames[i])
			}


		}

		cursorSolution, highestVal := AllDifferentPseudoNamesTested(activeVars, parentLeftVar, inputAlias)

		bestSolutions = append(bestSolutions, SolutionItem{binaryCursor, cursorSolution, highestVal})


		cursor++ 

	}

	returnSolution := SolutionItem{"0", []int{}, 0}

	for i := 0; i < len(bestSolutions); i++ {
		if(bestSolutions[i].HighestNetChange  > returnSolution.HighestNetChange){
			returnSolution = bestSolutions[i]
		}
	}

	return returnSolution
	
}



//this function returns the pseudo name combination that reduces the 
//number of variables on the right hand side the most

//more specifically this function returns the cursor combination that yielded the best
//reduction

//it also returns the amount of variables that were reduced

//this function needs to track the best possible reduction for the given variable combination
func AllDifferentPseudoNamesTested(chosenVars []VarPseudoNames, parentLeftVar string, inputAlias Alias) ([]int, int) {


	//this tracks the selected pseudoname per variable

	//the cursor slice ensures that for the current variable combination
	//all different pseudoname combinations are chosen, which means every possible 
	//pseudoname for each variable is matched with every combination of pseudonames of all 
	//the other variables 
	cursorSlice := []int{}

	//start column cursor least significant column

	cursorIsMaxedOut := false

	for i := 0; i < len(chosenVars); i++ {
		cursorSlice = append(cursorSlice, 0)
	}

	

	maxVals := []int{}

	for i := 0; i < len(chosenVars); i++ {
		maxVals = append(maxVals, len(chosenVars[i].PseudoNames))
	}

	VerbosePrintln(maxVals)

	
	doneTesting := false

	highestNetChange := 0

	cursorForHighestNetChange := []int{}

	
	for !doneTesting {

		pseudoNames := ReturnPseudoNamesForCursor(cursorSlice, maxVals, chosenVars)



		netChange, goodNetChange := SumOfPseudoNamesNetChangeIsGood(pseudoNames, parentLeftVar)

		if(goodNetChange){
			if(netChange > highestNetChange){
				highestNetChange = netChange
				cursorForHighestNetChange = CleanCopySliceDataInt(cursorSlice)
			}

			AddPseudoNameSubToDatabase(chosenVars, cursorSlice, inputAlias)

		}


		cursorSlice, cursorIsMaxedOut = IncrementCursorObject(CleanCopySliceDataInt(cursorSlice), maxVals)



		VerbosePrintln(cursorSlice)


		VerbosePrintln(cursorIsMaxedOut)



		if(cursorIsMaxedOut){

			doneTesting = true
		}


	}

	return cursorForHighestNetChange, highestNetChange

}

//this function returns the pseudonames for the cursor slice
func ReturnPseudoNamesForCursor(cursorSlice []int, maxVals []int, chosenVars []VarPseudoNames) [][]string {

	activePseudoNames := [][]string{}

	
	if(len(cursorSlice) != len(chosenVars)){
		fmt.Println("error cursor slice needs same length as chosen vars")
		os.Exit(1)
	}

	for i := 0; i < len(chosenVars); i++ {
		
		activePseudoNames = append(activePseudoNames, chosenVars[i].PseudoNames[cursorSlice[i]])
			
	}

	return activePseudoNames

} 

func Add



//returns incremented cursor, warns if 
func IncrementCursorObject(cursorSlice []int, maxVals []int) ([]int, bool) {

	if(len(cursorSlice) != len(maxVals)){
		fmt.Println("error cursor slice not same length as max vals slice")
		os.Exit(1)
	}

	incrementNext := true

	columnCursor := len(cursorSlice) - 1

	for i := columnCursor; i > -1; i-- {

		if(incrementNext){



			currentVal := cursorSlice[i]

			currentMaxVal := maxVals[i]


			if((currentVal + 1) == currentMaxVal){ 
				cursorSlice[i] = 0
				incrementNext = true
				
				if(columnCursor == 0){
					columnCursor = (len(cursorSlice) - 1)
				}else{
					columnCursor = columnCursor - 1
				}


				//if in the most significant column
				//and incremented to the last column,
				//the cursor is maxed out
				if(i == 0){
					return []int{}, true
				}
			}else{
				cursorSlice[i]++
				incrementNext = false
	

				return cursorSlice, false
				
			}
		}

	}

	return cursorSlice,  false

}

func AddPseudoNameSubToDatabase(chosenVars []VarPseudoNames, cursorSlice []int, inputAlias Alias) {

	if(len(chosenVars) != len(cursorSlice)){
		fmt.Println("error chosen vars slice must be same length as cursor slice AddPseudoNameSubToDatabase")
		os.Exit(1)
	}

	restrictedIndicesAliasRGenVar := []int{}

	aliasesToSub := []Alias{}

	
	for i := 0; i < len(chosenVars); i++ {
		currentVar := chosenVars[i]

		parentVarName := currentVar.ParentVar

		currentVarIndividualNumbers := currentVar.LoneNumberVals[cursorSlice[i]]

		currentVarPseudoNames := currentVar.PseudoNames[cursorSlice[i]]

		currentVarFloatMultiplierVals := currentVar.ScaledDownMultipliers[cursorSlice[i]]


		if(len(currentVarPseudoNames) != len(currentVarFloatVals)){
			fmt.Println("something went wrong when creating this VarPseudoNames AddPseudoNameSubToDatabase")
		}

		for j := 0; j < len(inputAlias.RGenVar); j++ {
			if(!(isRestrictedIndex(restrictedIndicesAliasRGenVar, j))){
				if(inputAlias.RGenVar[j].Name == parentVarName){

					varToSubMultiplier := inputAlias.RGenVar[j].Multiplier

					rightGenVarsForSubAlias := []GenVar{}

					//length check done above so ok to index both at k
					for k := 0; k < len(currentVarPseudoNames); k++ {
						rightGenVarsForSubAlias = append(rightGenVarsForSubAlias, CreateGenVar(currentVarPseudoNames[k], varToSubMultiplier * currentVarFloatMultiplierVals[k]))
					}

					currentVarIndividualNumber = currentVarIndividualNumber * varToSubMultiplier

					aliasToAdd := CreateAlias(CreateGenVar(parentVarName, 1), rightGenVarsForSubAlias, inputAlias.LNum, []float64{currentVarIndividualNumber})

					aliasesToSub = append(aliasesToSub, aliasToAdd)

					restrictedIndicesAliasRGenVar = append(restrictedIndicesAliasRGenVar, j)
				}
			}
		}

	}



	if(len(restrictedIndicesAliasRGenVar) != len(chosenVars)){
		fmt.Println("not all substitutions were made it appears... AddPseudoNameSubToDatabase")
		os.Exit(1)
	}

	cleanCopyInput := CleanCopyAlias(aliasInput)

	var outPutAliasToAdd Alias 

	previousOutput := cleanCopyInput

	for l := 0; l < len(aliasesToSub); l++ {

		outPutAliasToAdd = SubstituteAnAlias(previousOutput, aliasesToSub[i])

		previousOutput = CleanCopyAlias(outPutAliasToAdd)

	}

	AddToAliasDatabase(outPutAliasToAdd)


}








