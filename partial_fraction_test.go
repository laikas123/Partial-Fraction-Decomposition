package main

import (


	"testing"
	"fmt"
	"math"

)


func TestFindingHighestDegree(t *testing.T) {		

	Init()	

	originalNumerator := []S_Var{S_Var{-1, 3}, S_Var{2, 2}, S_Var{-9, 1}}

	originalNumeratorConstant := float64(24)




	testItem := GEQI(S_Var{1, 2}, 3, S_Var{1, 2})

	result := returnHighestDegree([]EquationItem{testItem})

	fmt.Printf("%#v\n", result)

	sliceGenVars1, lastIndex1 := ReturnGeneralVariablesForDegree(result, 0)

	fmt.Println("Gen Vars 1")


	for i := 0; i < len(sliceGenVars1); i++ {
		fmt.Printf("%#v\n", sliceGenVars1[i])
	}

	fmt.Println("last index 1")

	fmt.Printf("%#v\n", lastIndex1)

	sliceGenVars2, lastIndex2 := ReturnGeneralVariablesForDegree(result, lastIndex1)

	fmt.Println("Gen Vars 2")


	for i := 0; i < len(sliceGenVars2); i++ {
		fmt.Printf("%#v\n", sliceGenVars2[i])
	}

	fmt.Println("last index 2")

	fmt.Printf("%#v\n", lastIndex2)


	sVarSlice1 := []S_Var{S_Var{1,2}}

	constant1 := float64(9)

	
	sVarSlice2 := []S_Var{S_Var{1,2}, S_Var{-6, 1}}

	constant2 := float64(15)

	oneDEqtnSlice := MultiplyNumeratorByOppositeDenominatorAndOrganizeTheData(sliceGenVars1, sVarSlice2, constant2, sliceGenVars2, sVarSlice1, constant1, originalNumerator, originalNumeratorConstant)

	fmt.Println("eqtns")

	for i := 0; i < len(oneDEqtnSlice); i++ {
		fmt.Printf("%#v\n", oneDEqtnSlice[i])
	}


	allAliases :=  ReturnAllPossibleAliases(oneDEqtnSlice)


	fmt.Println("aliases")

	for i := 0; i < len(allAliases); i++ {
		fmt.Printf("%#v\n", allAliases[i])
	}


	cleanedUpAliases := CleanUpVars(allAliases)

	fmt.Println("cleaned up aliases")

	AliasDatabase = []AliasOneDEquationSimple{}

	

	newAlias := AliasOneDEquationSimple{[]GenVar{GenVar{"A", -1000000}}, []GenVar{}, []float64{}, []float64{}}


	AddToAliasDataBase(newAlias)

	PrintAliasDataBase()














	for i := 0; i < len(cleanedUpAliases); i++ {
		fmt.Printf("%#v\n", cleanedUpAliases[i])
	}


	if (result != 2) {
       t.Errorf("failure")
    
    }

}

func TestCleanUpForGenVar(t *testing.T) {



	testGenVar := GenVar{[]}



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





