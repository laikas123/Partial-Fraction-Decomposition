package main

import (


	"testing"
	"fmt"
	"math"
	"os"
)


func TestFindingHighestDegree(t *testing.T) {		

	Init()	


	generalVariable1 := CreateGeneralVariable("A", 44, 2)

	VerbosePrint(generalVariable1)

	sVar1 := CreateSVar(43, 2)

	VerbosePrint(sVar1)


	genVar1 := CreateGenVar("A", 33)

	genVar2 := CreateGenVar("A", 22)

	VerbosePrint(genVar1)

	VerbosePrint(genVar2)

	VerbosePrint(TwoGenVarsAreEqual(genVar1, genVar2))

	VerbosePrint(TwoGenVarsAreSameVariable(genVar1, genVar2))



	generalVariableFinal1 := CreateGeneralVariable("A", 1, 1)
	generalVariableFinal2 := CreateGeneralVariable("B", 1, 0)
	
	leftNumerator := []GeneralVariable{generalVariableFinal1, generalVariableFinal2}

	generalVariableFinal3 := CreateGeneralVariable("C", 1, 1)
	generalVariableFinal4 := CreateGeneralVariable("D", 1, 0)

	rightNumerator := []GeneralVariable{generalVariableFinal3, generalVariableFinal4}

	rightSvar1 := CreateSVar(1, 2)
	rightSvar2 := CreateSVar(-6, 1)

	rightDenomS := []S_Var{rightSvar1, rightSvar2}

	rightConstant := float64(15)

	leftSvar1 := CreateSVar(1, 2)
	
	leftDenomS := []S_Var{leftSvar1}
	
	leftConstant := float64(9)


	originalNumeratorSVar := []S_Var{CreateSVar(-1, 3), CreateSVar(2,2), CreateSVar(-9, 1)}
	originalNumeratorConstant := float64(24)

	fmt.Println("num * denom")


	crossMultipliedResult := MultiplyNumeratorByOppositeDenominatorAndOrganizeTheData(leftNumerator, rightDenomS, rightConstant, rightNumerator, leftDenomS, leftConstant, originalNumeratorSVar, originalNumeratorConstant)


	VerbosePrintSlice(crossMultipliedResult)


	fmt.Println()

	fmt.Println("All Aliases")

	allAliases := ReturnAllPossibleAliases(crossMultipliedResult)

	VerbosePrintSlice(allAliases)


	fmt.Println()

	fmt.Println("Cleaned Up Aliases")

	cleanedUpAliases := CleanUpAliases(allAliases)

	VerbosePrintSlice(cleanedUpAliases)

	for i := 0; i < len(cleanedUpAliases); i++ {
		AddToAliasDatabase(cleanedUpAliases[i])
	}



	original := cleanedUpAliases[3]

	fmt.Println()

	fmt.Println("original selected")

	VerbosePrint(original)



	substitution := cleanedUpAliases[9]



	fmt.Println()

	fmt.Println("substitution selected")

	VerbosePrint(substitution)





	cleanCopySubstitution := CleanCopyAlias(substitution) 


	fmt.Println()

	fmt.Println("test clean copy of alias on substitution")

	VerbosePrint(cleanCopySubstitution)





	CheckLeftSideIsOnly1Long(cleanCopySubstitution.LGenVar, "TestFindingHighestDegree")

	scaleVal := cleanCopySubstitution.LGenVar[0].Multiplier

	cleanCopySubstitution.LGenVar = ScaleDownSliceGenVar(cleanCopySubstitution.LGenVar, scaleVal)
	cleanCopySubstitution.RGenVar = ScaleDownSliceGenVar(cleanCopySubstitution.RGenVar, scaleVal)

	cleanCopySubstitution.LNum = ScaleDownSliceFloat(cleanCopySubstitution.LNum, scaleVal)
	cleanCopySubstitution.RNum = ScaleDownSliceFloat(cleanCopySubstitution.RNum, scaleVal)

	fmt.Println()

	fmt.Println("scaled substitution selected")

	VerbosePrint(cleanCopySubstitution)


	testingSubstitution, dataValid := SubstituteAnAlias(original, substitution)

	if(!dataValid){
		fmt.Println("not data valid TestFindingHighestDegree")
		os.Exit(1)
	}

	cleanCopyTestingSubstitution := CleanCopyAlias(testingSubstitution)

	VerbosePrint(testingSubstitution)
	VerbosePrint(cleanCopyTestingSubstitution)



	fmt.Println()

	fmt.Println("test sub")

	VerbosePrint(testingSubstitution)


	cleanedSubstitution := SimplifyGenVarRightHandGenVarSlice(testingSubstitution)


	fmt.Println()

	fmt.Println("test clean sub gen var")

	VerbosePrint(cleanedSubstitution)


	cleanedSubstitution = SimplifyRightHandNumSlice(cleanedSubstitution)

	fmt.Println()

	fmt.Println("test clean sub num")

	VerbosePrint(cleanedSubstitution)


	cleanedSubstitution = MoveVarsEqualToLeftHandSideToLeftSide(cleanedSubstitution)

	fmt.Println()

	fmt.Println("test clean sub moved left")

	VerbosePrint(cleanedSubstitution)

	fmt.Println()

	fmt.Println("check full clean up matches")


	fullCleanUpTest, leftSideZero := FullCleanUp(cleanCopyTestingSubstitution)


	VerbosePrint(fullCleanUpTest)

	VerbosePrint(leftSideZero)


	//TESTING BOOLEAN RETURNS FOR SUBSTITUTIONS

	fmt.Println()
	fmt.Println()


	gVar1 := CreateGenVar("A", 3)
	gVar2 := CreateGenVar("B", 22)
	gVar3 := CreateGenVar("C", 1)




	newAliasOld := CreateAlias([]GenVar{gVar1}, []GenVar{gVar2, gVar3}, []float64{}, []float64{})



	gVar4 := CreateGenVar("C", 1)
	gVar5 := CreateGenVar("B", -22)
	gVar6 := CreateGenVar("D", -22)
	gVar7 := CreateGenVar("E", 12)
	gVar8 := CreateGenVar("A", -3)

	newAliasNew := CreateAlias([]GenVar{gVar4}, []GenVar{gVar5, gVar6, gVar7, gVar8}, []float64{}, []float64{})


	VerbosePrint(NewAliasEqualsLeftSideVariableNoIncrease(newAliasOld, newAliasNew))
	



	VerbosePrint(NewAliasReducesVariablesOnRightHandSide(newAliasOld, newAliasNew))


	
	testAliasEqual1 := CreateAlias([]GenVar{gVar4}, []GenVar{gVar5, gVar6, gVar7, gVar8}, []float64{}, []float64{})

	testAliasEqual2 := CreateAlias([]GenVar{gVar4}, []GenVar{ gVar6, gVar7, gVar8}, []float64{}, []float64{})


	fmt.Println()

	fmt.Println("two aliases are equal")

	VerbosePrint(TwoAliasesAreEqual(testAliasEqual1, testAliasEqual2))


	SolutionListener(4)










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















