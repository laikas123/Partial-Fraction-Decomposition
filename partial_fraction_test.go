package main

import (


	"testing"
	"fmt"
	"math"

)

func TestFindingHighestDegree(t *testing.T) {		



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

	newGenVar1 := MultiplyNumeratorByOppositeDenominator(sliceGenVars1, sVarSlice2, constant2)

	fmt.Println("NEW Gen Vars 1")


	for i := 0; i < len(newGenVar1); i++ {
		fmt.Printf("%#v\n", newGenVar1[i])
	}	


	newGenVar2 := MultiplyNumeratorByOppositeDenominator(sliceGenVars2, sVarSlice1, constant1)

	fmt.Println("NEW Gen Vars 2")


	for i := 0; i < len(newGenVar2); i++ {
		fmt.Printf("%#v\n", newGenVar2[i])
	}	





	if (result != 2) {
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





