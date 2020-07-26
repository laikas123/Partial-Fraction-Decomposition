package main

import (


	"testing"
	"fmt"
	"math"
	

)


func TestFindingHighestDegree(t *testing.T) {		

	Init()	

// 	gVar1 := CreateGenVar("A", 1)
// 	gVar2 := CreateGenVar("B", 1)
// 	gVar3 := CreateGenVar("C", 1)
// 	gVar4 := CreateGenVar("D", 1)
// 	gVar5 := CreateGenVar("E", 1)
	
// 	x := CreateGenVar("X", 1)
// 	y := CreateGenVar("Y", 1)
// 	z := CreateGenVar("Z", 1)
	

// //say A has pseudo name [x y]
// //B has pseudo name[y]
// //C has pseudo name[z]
// //D has pseudo name[y z]
// //E has pseduo name[x z]


// 	alias1 := CreateAlias([]GenVar{gVar1}, []GenVar{x, y}, []float64{}, []float64{2})
	
// 	alias2 := CreateAlias([]GenVar{gVar2}, []GenVar{y}, []float64{}, []float64{2})

// 	alias3 := CreateAlias([]GenVar{gVar3}, []GenVar{z}, []float64{}, []float64{2})

// 	alias4 := CreateAlias([]GenVar{gVar4}, []GenVar{y, z}, []float64{}, []float64{2})

// 	alias5 := CreateAlias([]GenVar{gVar5}, []GenVar{x, z}, []float64{}, []float64{2})

// 	AllAliasPermutationsAndAddToDatabase(alias1)
// 	AllAliasPermutationsAndAddToDatabase(alias2)
// 	AllAliasPermutationsAndAddToDatabase(alias3)
// 	AllAliasPermutationsAndAddToDatabase(alias4)
// 	AllAliasPermutationsAndAddToDatabase(alias5)


// 	PrintAliasDataBase()

// 	pseudoNamesA := GetPseudoNamesForRGenVar("A")

// 	pseudoNamesB := GetPseudoNamesForRGenVar("B")

// 	pseudoNamesC := GetPseudoNamesForRGenVar("C")

// 	pseudoNamesD := GetPseudoNamesForRGenVar("D")

// 	pseudoNamesE := GetPseudoNamesForRGenVar("E")

// 	listPseudoNames := [][]string{}

// 	listPseudoNames = append(listPseudoNames, pseudoNamesA[1])
// 	listPseudoNames = append(listPseudoNames, pseudoNamesB[1])
// 	listPseudoNames = append(listPseudoNames, pseudoNamesC[1])
// 	listPseudoNames = append(listPseudoNames, pseudoNamesD[1])
// 	listPseudoNames = append(listPseudoNames, pseudoNamesE[1])

// 	reductionAmount, dataValid := SumOfPseudoNamesNetChangeIsGood(listPseudoNames, "X")

// 	VerbosePrint(dataValid)
// 	VerbosePrint(reductionAmount)


	// parentAlias := CreateAlias([]GenVar{CreateGenVar("L", 10)}, []GenVar{CreateGenVar("A", 2), CreateGenVar("B", 2), CreateGenVar("C", 2)}, []float64{0}, []float64{44})


	// varswpname1 := VarPseudoNames{[][]string{[]string{"x, y, z"}, []string{"j"}, []string{"y"}}, []float64{1, 2, 4}, [][]float64{[]float64{3, 4, 5}, []float64{3}, []float64{3}}, "A"}

	// varswpname2 := VarPseudoNames{[][]string{[]string{"l"}, []string{"k"}, []string{"j"}}, []float64{1, 2, 4}, [][]float64{[]float64{3}, []float64{3}, []float64{3}}, "B"}

	// varswpname3 := VarPseudoNames{[][]string{[]string{"j"}, []string{"y"}}, []float64{1, 2}, [][]float64{[]float64{3}, []float64{3}}, "C"}

	// chosenVars := []VarPseudoNames{varswpname1, varswpname2, varswpname3}

	// cursorSlice := []int{1, 2, 0}

	// AddPseudoNameSubToDatabase(chosenVars, cursorSlice, parentAlias)	


	// varswpname4 := VarPseudoNames{[][]string{[]string{"k"}, []string{"s"}, []string{"j"}}, "D"}
	

	// varsChosen := []VarPseudoNames{varswpname1, varswpname2, varswpname3, varswpname4}




	// solution := BestAliasSliceForSubstitution(varsChosen, "j")	

	// VerbosePrint(solution)
	



//CURSOR TEST

/*

	
	cursorSlice := []int{0, 0, 0, 0}

	maxVals := []int{3, 3, 2, 3}

	columnCursor := 3


	VerbosePrint(cursorSlice)
	VerbosePrint(maxVals)
	VerbosePrint(columnCursor)

	fmt.Println()

	newCurs := cursorSlice

	cursorMaxed := false


	for(!cursorMaxed){

		newCurs,  cursorMaxed = IncrementCursorObject(cursorSlice, maxVals)

		VerbosePrint(newCurs)

	}
*/

//CURSOR TEST

//PSEUDO NAME DATA 
/*
pseudoNamesReturned := ReturnPseudoNamesForCursor(cursorSlice, maxVals, varsChosen)


	for i := 0; i < len(pseudoNamesReturned); i++ {
		VerbosePrint(pseudoNamesReturned[i])
	}

*/
//PSEUDO NAME DATA




//MAIN TEST


	gVar1 := CreateGenVar("A", 2)
	gVar2 := CreateGenVar("D", 1)

	alias1 := CreateAlias([]GenVar{gVar1}, []GenVar{gVar2}, []float64{}, []float64{2})


	gVar3 := CreateGenVar("B", 2)
	gVar4 := CreateGenVar("A", 1)
	
	alias2 := CreateAlias([]GenVar{gVar3}, []GenVar{gVar4}, []float64{}, []float64{})
	
	
	gVar5 := CreateGenVar("C", 1)

	gVar6 := CreateGenVar("A", 1)
	
	gVar7 := CreateGenVar("B", 1)



	alias3 := CreateAlias([]GenVar{gVar5}, []GenVar{gVar6, gVar7}, []float64{}, []float64{})
	

	gVar8 := CreateGenVar("D", 1)

	gVar9 := CreateGenVar("C", 1)
	
	gVar10 := CreateGenVar("B", 1)

	
	alias4 := CreateAlias([]GenVar{gVar8}, []GenVar{gVar9, gVar10}, []float64{}, []float64{-2})
	




	AddToAliasDatabase(alias1)
	AddToAliasDatabase(alias2)
	AddToAliasDatabase(alias3)
	AddToAliasDatabase(alias4)
	
	AllAliasPermutationsAndAddToDatabase(alias1)
	AllAliasPermutationsAndAddToDatabase(alias2)
	AllAliasPermutationsAndAddToDatabase(alias3)
	AllAliasPermutationsAndAddToDatabase(alias4)



	SolutionListener(4)




//MAIN TEST











	
	


	// gVar1 := CreateGenVar("A", -5)

	// gVar1Neg := CreateGenVar("A", 5)


	// gVar2 := CreateGenVar("B", 2)

	// gVar2Neg := CreateGenVar("B", -2)


	// gVar3 := CreateGenVar("C", 12)

	// alias1 := CreateAlias([]GenVar{gVar1}, []GenVar{gVar2Neg, gVar3}, []float64{}, []float64{})


	// alias2 := CreateAlias([]GenVar{gVar2}, []GenVar{gVar3, gVar1Neg}, []float64{}, []float64{})


	// VerbosePrint(TwoAliasesAreVaritaionsOfEachOther(alias1, alias2))




//GOOOD

	// gVar1 := CreateGenVar("C", 1)

	// gVar2 := CreateGenVar("A", 1)

	// gVar3 := CreateGenVar("B", 1)


	// alias1 := CreateAlias([]GenVar{gVar1}, []GenVar{gVar2, gVar3}, []float64{}, []float64{})

	// gVar4 := CreateGenVar("B", 2)
	// gVar5 := CreateGenVar("A", 1)


	// alias2 := CreateAlias([]GenVar{gVar4}, []GenVar{gVar5}, []float64{}, []float64{})


	// testSub, _ := SubstituteAnAlias(alias1, alias2)

	// VerbosePrint(ScaleDownEntireAlias(CleanCopyAlias(alias2), alias2.LGenVar[0].Multiplier))

	// VerbosePrint(alias1)
	// VerbosePrint(alias2)

	// VerbosePrint(testSub)

//GOOOD



















	// generalVariable1 := CreateGeneralVariable("A", 44, 2)

	// VerbosePrint(generalVariable1)

	// sVar1 := CreateSVar(43, 2)

	// VerbosePrint(sVar1)


	// genVar1 := CreateGenVar("A", 33)

	// genVar2 := CreateGenVar("A", 22)

	// VerbosePrint(genVar1)

	// VerbosePrint(genVar2)

	// VerbosePrint(TwoGenVarsAreEqual(genVar1, genVar2))

	// VerbosePrint(TwoGenVarsAreSameVariable(genVar1, genVar2))



	// generalVariableFinal1 := CreateGeneralVariable("A", 1, 1)
	// generalVariableFinal2 := CreateGeneralVariable("B", 1, 0)
	
	// leftNumerator := []GeneralVariable{generalVariableFinal1, generalVariableFinal2}

	// generalVariableFinal3 := CreateGeneralVariable("C", 1, 1)
	// generalVariableFinal4 := CreateGeneralVariable("D", 1, 0)

	// rightNumerator := []GeneralVariable{generalVariableFinal3, generalVariableFinal4}

	// rightSvar1 := CreateSVar(1, 2)
	// rightSvar2 := CreateSVar(-6, 1)

	// rightDenomS := []S_Var{rightSvar1, rightSvar2}

	// rightConstant := float64(15)

	// leftSvar1 := CreateSVar(1, 2)
	
	// leftDenomS := []S_Var{leftSvar1}
	
	// leftConstant := float64(9)


	// originalNumeratorSVar := []S_Var{CreateSVar(-1, 3), CreateSVar(2,2), CreateSVar(-9, 1)}
	// originalNumeratorConstant := float64(24)

	// fmt.Println("num * denom")


	// crossMultipliedResult := MultiplyNumeratorByOppositeDenominatorAndOrganizeTheData(leftNumerator, rightDenomS, rightConstant, rightNumerator, leftDenomS, leftConstant, originalNumeratorSVar, originalNumeratorConstant)


	// VerbosePrintSlice(crossMultipliedResult)


	// fmt.Println()

	// fmt.Println("All Aliases")

	// allAliases := ReturnAllPossibleAliases(crossMultipliedResult)

	// VerbosePrintSlice(allAliases)


	// fmt.Println()

	// fmt.Println("Cleaned Up Aliases")

	// cleanedUpAliases := CleanUpAliases(allAliases)

	// VerbosePrintSlice(cleanedUpAliases)

	// for i := 0; i < len(cleanedUpAliases); i++ {
	// 	AddToAliasDatabase(cleanedUpAliases[i])
	// }



	// original := cleanedUpAliases[3]

	// fmt.Println()

	// fmt.Println("original selected")

	// VerbosePrint(original)



	// substitution := cleanedUpAliases[9]



	// fmt.Println()

	// fmt.Println("substitution selected")

	// VerbosePrint(substitution)





	// cleanCopySubstitution := CleanCopyAlias(substitution) 


	// fmt.Println()

	// fmt.Println("test clean copy of alias on substitution")

	// VerbosePrint(cleanCopySubstitution)





	// CheckLeftSideIsOnly1Long(cleanCopySubstitution.LGenVar, "TestFindingHighestDegree")

	// scaleVal := cleanCopySubstitution.LGenVar[0].Multiplier

	// cleanCopySubstitution.LGenVar = ScaleDownSliceGenVar(cleanCopySubstitution.LGenVar, scaleVal)
	// cleanCopySubstitution.RGenVar = ScaleDownSliceGenVar(cleanCopySubstitution.RGenVar, scaleVal)

	// cleanCopySubstitution.LNum = ScaleDownSliceFloat(cleanCopySubstitution.LNum, scaleVal)
	// cleanCopySubstitution.RNum = ScaleDownSliceFloat(cleanCopySubstitution.RNum, scaleVal)

	// fmt.Println()

	// fmt.Println("scaled substitution selected")

	// VerbosePrint(cleanCopySubstitution)


	// testingSubstitution, dataValid := SubstituteAnAlias(original, substitution)

	// if(!dataValid){
	// 	fmt.Println("not data valid TestFindingHighestDegree")
	// 	os.Exit(1)
	// }

	// cleanCopyTestingSubstitution := CleanCopyAlias(testingSubstitution)

	// VerbosePrint(testingSubstitution)
	// VerbosePrint(cleanCopyTestingSubstitution)



	// fmt.Println()

	// fmt.Println("test sub")

	// VerbosePrint(testingSubstitution)


	// cleanedSubstitution := SimplifyGenVarRightHandGenVarSlice(testingSubstitution)


	// fmt.Println()

	// fmt.Println("test clean sub gen var")

	// VerbosePrint(cleanedSubstitution)


	// cleanedSubstitution = SimplifyRightHandNumSlice(cleanedSubstitution)

	// fmt.Println()

	// fmt.Println("test clean sub num")

	// VerbosePrint(cleanedSubstitution)


	// cleanedSubstitution = MoveVarsEqualToLeftHandSideToLeftSide(cleanedSubstitution)

	// fmt.Println()

	// fmt.Println("test clean sub moved left")

	// VerbosePrint(cleanedSubstitution)

	// fmt.Println()

	// fmt.Println("check full clean up matches")


	// fullCleanUpTest, leftSideZero := FullCleanUp(cleanCopyTestingSubstitution)


	// VerbosePrint(fullCleanUpTest)

	// VerbosePrint(leftSideZero)


	// //TESTING BOOLEAN RETURNS FOR SUBSTITUTIONS

	// fmt.Println()
	// fmt.Println()


	// // gVar1 := CreateGenVar("A", 3)
	// // gVar2 := CreateGenVar("B", 22)
	// // gVar3 := CreateGenVar("C", 1)




	// // newAliasOld := CreateAlias([]GenVar{gVar1}, []GenVar{gVar2, gVar3}, []float64{}, []float64{})



	


	// // VerbosePrint(NewAliasEqualsLeftSideVariableNoIncrease(newAliasOld, newAliasNew))
	



	// // VerbosePrint(NewAliasReducesVariablesOnRightHandSide(newAliasOld, newAliasNew))


	// gVar4 := CreateGenVar("C", 1)
	// gVar5 := CreateGenVar("B", -22)
	// gVar6 := CreateGenVar("D", -22)
	// gVar7 := CreateGenVar("E", 12)
	// gVar8 := CreateGenVar("A", -3)

	

	
	// testAliasEqual1 := CreateAlias([]GenVar{gVar4}, []GenVar{gVar5, gVar6, gVar7, gVar8}, []float64{221}, []float64{1})

	// VerbosePrint(testAliasEqual1)

	// VerbosePrint(ScaleDownEntireAlias(testAliasEqual1, 2))


	// testAliasEqual2 := CreateAlias([]GenVar{gVar4}, []GenVar{ gVar6, gVar7, gVar8}, []float64{}, []float64{})


	// fmt.Println()

	// fmt.Println("two aliases are equal")

	// VerbosePrint(TwoAliasesAreEqual(testAliasEqual1, testAliasEqual2))


	// SolutionListener(4)






	if (false) {
       t.Errorf("failure")
    
    }

}

func TestCleanUpForGenVar(t *testing.T) {



	



	if (false) {
       t.Errorf("failure")
    
    }

}



func aboutEquals(checkVal float64, result float64) bool {
	
	difference := math.Abs(checkVal - result)


	if(difference < math.Abs(0.03) ) {
		return true
	}else{
		return false
	}
}


func VerbosePrint(input interface{}) {
	fmt.Printf("%#v\n", input)
}


func VerbosePrintSlice(item interface{}){
	

	switch item.(type){
		case []OneDEquation:
			
			value, ok := item.([]OneDEquation)
			if(ok){
				for i := 0; i < len(value); i++ {
					fmt.Print(i)
					fmt.Printf("%#v\n", value[i])		
				}
			}
		case []Alias:
			
			value, ok := item.([]Alias)
			if(ok){
				for i := 0; i < len(value); i++ {
					fmt.Print(i)
					fmt.Printf("%#v\n", value[i])		
				}
			}	
	}



	
}























