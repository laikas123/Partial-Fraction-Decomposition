package main

import (
	"fmt"
	
	"testing"

)

// func TestGetFirstIndexOfEquation(t *testing.T){

// 	numberSlice := gNum(1, 2, 1, 6, 0, 1, 3, 3)

// 	numberSlice2 := gNum(3, 2, 1, 6, 0, 1, 3, 3)

// 	equation := Create2DEquationFromSliceInputs(gOP(), numberSlice2, gCP(2, 1), gOP(), numberSlice, gCP(2, 1))


// 	factor, endIndex := GetFirstFactorFromEquation(equation)

// 	fmt.Println("equation", DecodeFloatSliceToEquation(equation), "\n", "first factor", DecodeFloatSliceToEquation(factor), "\n", "index first factor", endIndex)

// 	panic("done testing")

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestFactorRestrictedIndices(t *testing.T){


// 	restrictedIndices := []int{3, 4, 6}


// 	indexToCheck := 8

// 	fmt.Println("is restricted", IsRestrictedIndex(indexToCheck, restrictedIndices))

// 	panic("done testing")

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }


// func TestFoilingExponentParenthesis(t *testing.T){

	

// 	numberSlice := gNum(1, 7, 1, 2, 2, 1, 2,2)

// 	// numberSlice2 := gNum(1, 7, 1, 2, 2, 3, 2,2)

// 	equation := Create2DEquationFromSliceInputs(gOP(), numberSlice, gCP(2, 1))



// 	// equation := Create2DEquationFromSliceInputs(gOP(), gNum(3, 0, 3), gNum(1, 1, 2), gNum(3, 0, 5), gCP(1, 3))

// 	fmt.Println("initial equation", DecodeFloatSliceToEquation(equation))

// 	result := FoilOutParenthesisRaisedToExponent(equation)
// 	fmt.Println("initial equation", DecodeFloatSliceToEquation(equation))


// 	fmt.Println("resutl", DecodeFloatSliceToEquation(result))

// 	panic("exit")
	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }


// func TestGenerateNumbersSliceAndSimplifyInnerParenthesis(t *testing.T){

	

// 	numberSlice := gNum(1, 7, 3, 2, 2, 3, 2,2)

	

// 	equation := Create2DEquationFromSliceInputs(gOP(), numberSlice, gCP(1, 1))



// 	// equation := Create2DEquationFromSliceInputs(gOP(), gNum(3, 0, 3), gNum(1, 1, 2), gNum(3, 0, 5), gCP(1, 3))

// 	// fmt.Println(DecodeFloatSliceToEquation(equation))

// 	result := SimplifyInnerParenthesis(equation)

// 	fmt.Println("result result", DecodeFloatSliceToEquation(result))
	

// 	panic("exit")

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }



//this function gathers terms in series that multiply or divide each other
//it also performs foils for any factor raised to an exponent greater than 1
//if two factors adjacent have powers >= 1 they get foiled
func TestGatherFactorsIntoSeriesThatMultiplyOrDivideEachOtherSimplify(t *testing.T){

	
	numberSlice1 := gNum(14, 1, 2, 28, 0)

	numberSlice2 := gNum(14, 1, 1, 28, 0, 1, 15, 1, 1, -15, 2)

	// equation1 := Create2DEquationFromSliceInputs(gOP(), numberSlice2, gCP(1, 1))

	// numbersHolder, operatorsHolder := SortParenthesisContainingOnlyPlusAndMinusBySExponent([][]complex128{[]complex128{complex(14, 0), complex(1, 0)}, []complex128{complex(28, 0), complex(0, 0)}, []complex128{complex(15, 0), complex(1, 0)}, []complex128{complex(15, 0), complex(2, 0)}},  [][]complex128{[]complex128{complex(0, 0), complex(1, 0)}, []complex128{complex(0, 0), complex(1, 0)}, []complex128{complex(0, 0), complex(2, 0)}})

	// fmt.Println("numbers", DecodeFloatSliceToEquation(numbersHolder), "operators", DecodeFloatSliceToEquation(operatorsHolder))

	// panic("error")

	// fmt.Println("simplified further", DecodeFloatSliceToEquation(SimplifyInnerParenthesis(equation1)))

	// panic("end test")

	numberSlice3 := gNum(2, 1, 1, 28, 0)

	numberSlice4 := gNum(222, 1, 1, 28, 0)

	equation := Create2DEquationFromSliceInputs(gOP(), numberSlice1, gCP(2, 3), gOP(), numberSlice2, gCP(1, 4), gOP(), numberSlice3, gCP(1, 1), gOP(), numberSlice4, gCP(1, 3),)

	factors2dslice := GatherFactorsInSeriesThatMultiplyOrDivideEachOther(equation)

	for i := 0; i < len(factors2dslice); i++ {

		currentFactorsInSeries := factors2dslice[i]

		fmt.Println("New Factors In Series ", i)

		for j := 0; j < len(currentFactorsInSeries); j++ {
			fmt.Println("factor ", j, " ", DecodeFloatSliceToEquation(currentFactorsInSeries[j].Data))
		}

	}





	panic("done testing")

	if(false){
		t.Errorf("failure")
	}
}



//returns the exact same data when there are no factor mathces,
//returns the factors properly removed when there are matches
//test is setup to show this right now
func TestFactorNumeratorDenomiantor(t *testing.T){

	// numberSlice := gNum(14, 2, 1, 28, 0)

	// numberSlice2 := gNum(7, 2, 1, 14, 0)

	numbersSlice1 := gNum(14, 1, 2, 28, 0)

	numbersSlice2 := gNum(14, 1, 1, 28, 0)

	numbersSlice3 := gNum(222, 1, 1, 28, 0)

	numbersSlice4 := gNum(222, 1, 1, 28, 0)


	numerator := Create2DEquationFromSliceInputs(gOP(), numbersSlice1, gCP(3, 1), gOP(), numbersSlice2, gCP(1, 1), gOP(), numbersSlice3, gCP(1, 1))

	denominator := Create2DEquationFromSliceInputs(gOP(), numbersSlice1, gCP(1, 1), gOP(), numbersSlice2, gCP(3, 1), gOP(), numbersSlice4, gCP(1, 1))

	newNumeratorFactors, newDenominatorFactors := FactorNumeratorAndDenonminatorRemoveLikeFactors(numerator, denominator)

	for i := 0; i < len(newNumeratorFactors); i++ {
		fmt.Println("New Numerator factor ", i, " ", DecodeFloatSliceToEquation(newNumeratorFactors[i].Data) )
	}

	for i := 0; i < len(newDenominatorFactors); i++ {
		fmt.Println("New Denominator factor ", i, " ", DecodeFloatSliceToEquation(newDenominatorFactors[i].Data) )
	}

	panic("done testing")

	if(false){
		t.Errorf("failure")
	}
}

func TestMakeSurePuttingAnEquationThroughFactorFunctionsDoesNotChangeIt(t *testing.T){

	numberSlice := gNum(1, 9, 1, 6, 22)


	equation := Create2DEquationFromSliceInputs(gOP(), numberSlice, gCP(2, 1))

	copyOriginal := CleanCopyEntire2Dcomplex128Slice(equation)

	fmt.Println("initial equation", DecodeFloatSliceToEquation(equation))

	equation = FactorQuadraticsWithABCAllPresent(equation)
	equation = FactorQuadraticsWithABOnlyPresent(equation)
	equation = FactorQuadraticsWithACOnlyPresent(equation)


	fmt.Println("post quadratic functions equation", DecodeFloatSliceToEquation(equation))
	

	fmt.Println("two equations identical", TwoEquationsAreExactlyIdentical(equation, copyOriginal))

	panic("done testing")

	if(false){
		t.Errorf("failure")
	}
}





// func TestFoilingNeighborParenthesis(t *testing.T){

	

// 	numberSlice := gNum(1, 7, 1, 2, 2, 1, 2,2)

// 	numberSlice2 := gNum(1, 7, 1, 2, 2, 3, 2,2)

// 	equation := Create2DEquationFromSliceInputs(numberSlice, numberSlice2)



// 	// equation := Create2DEquationFromSliceInputs(gOP(), gNum(3, 0, 3), gNum(1, 1, 2), gNum(3, 0, 5), gCP(1, 3))

// 	fmt.Println("initial equation", DecodeFloatSliceToEquation(equation))

// 	result := FoilNeighborParenthesis(equation)

// 	fmt.Println(DecodeFloatSliceToEquation(result))
	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestFoilingExponentParenthesis(t *testing.T){

	

// 	numberSlice := gNum(1, 7, 1, 2, 2, 1, 2,2)

// 	// numberSlice2 := gNum(1, 7, 1, 2, 2, 3, 2,2)

// 	equation := Create2DEquationFromSliceInputs(gOP(), numberSlice, gCP(2, 1))



// 	// equation := Create2DEquationFromSliceInputs(gOP(), gNum(3, 0, 3), gNum(1, 1, 2), gNum(3, 0, 5), gCP(1, 3))

// 	fmt.Println("initial equation", DecodeFloatSliceToEquation(equation))

// 	result := FoilOutParenthesisRaisedToExponent(equation)

// 	fmt.Println(DecodeFloatSliceToEquation(result))
	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }


// func TestSimplifyingInnerParenthesis(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(2, 3), gNum(3, 1), gCP(1)}       

// 	fmt.Println("original equation", DecodeFloatSliceToEquation(equation))

// 	SimplifyInnerParenthesis(equation)

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestAreIdentical(t *testing.T){


// 	equation1 := [][]complex128{gOP(), gNum(2, 3, 3, 1), gCP(3, 3), gCP(3, 3), gCP(3, 3)}       
// 	equation2 := [][]complex128{gOP(), gNum(2, 3, 3, 1), gCP(3, 3), gCP(3, 3), gCP(3, 3)}       

// 	fmt.Println("equation1", DecodeFloatSliceToEquation(equation1))
// 	fmt.Println("equation2", DecodeFloatSliceToEquation(equation2))

// 	fmt.Println("are identical", TwoEquationsAreExactlyIdentical(equation2, equation1))



// 	if(false){
// 		t.Errorf("failure")
// 	}
// }



// func TestCleanCopy(t *testing.T){


// 	equation := [][]complex128{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3, 3), gCP(3, 3), gOP(), gOP(), gOP(), gCP(3, 3), gCP(3, 3), gOP(), gOP(), gOP()}       

// 	cleanCopy := CleanCopyEntire2DComplex128Slice(equation)

// 	fmt.Println("initial", DecodeFloatSliceToEquation(equation))



// 	fmt.Println("   copy", DecodeFloatSliceToEquation(cleanCopy))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestRemoveTrailingOpeners(t *testing.T){


// 	equation := [][]complex128{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3, 3), gCP(3, 3), gOP(), gOP(), gOP()}       

// 	fmt.Println("initial", DecodeFloatSliceToEquation(equation))
	
// 	equation = RemoveLastItemIfItIsOpeningParenthesis(equation)

// 	fmt.Println("after", DecodeFloatSliceToEquation(equation))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }


// func TestDepthCheckRemove(t *testing.T){


// 	equation := [][]complex128{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3, 3), gCP(3, 3), gCP(3, 3)}       

// 	fmt.Println("initial", DecodeFloatSliceToEquation(equation))
	
// 	equation = RemoveExcessParenthesisViaDepthCheck(equation)

// 	fmt.Println(" after", DecodeFloatSliceToEquation(equation))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }



// func TestFoilingExponentStyle(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(2, 3, 3, 0), gCP(3, 3), gOP(), gNum(2, 3, 3, 0), gCP(3, 3), gOP(), gNum(2, 3, 3, 0), gCP(2, 3)}       

// 	fmt.Println("initial", DecodeFloatSliceToEquation(equation))
	
// 	equation = FoilOutParenthesisRaisedToExponent(equation)

// 	fmt.Println(" after", DecodeFloatSliceToEquation(equation))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestFoilingNeighborStyle(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(2, 3, 3, 0), gCP(1, 3), gOP(), gNum(2, 3, 3, 0), gCP(1, 3), gOP(), gNum(2, 3, 3, 0), gCP(1, 3), gOP(), gNum(2, 3, 3, 0), gCP(1, 3)}       

// 	fmt.Println("initial", DecodeFloatSliceToEquation(equation))
	
// 	equation = FoilNeighborParenthesis(equation)

// 	fmt.Println(" after", DecodeFloatSliceToEquation(equation))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestQuadraticFactoringABCPresent(t *testing.T){



// 	numberSlice := gNum(1, 2, 1, 1, 1, 1, -6, 0)

// 	// numberSlice2 := gNum(1, 7, 1, 2, 2, 3, 2,2)

// 	equation := Create2DEquationFromSliceInputs(gOP(), numberSlice, gCP(2, 1))



// 	// equation := Create2DEquationFromSliceInputs(gOP(), gNum(3, 0, 3), gNum(1, 1, 2), gNum(3, 0, 5), gCP(1, 3))

// 	fmt.Println("initial equation", DecodeFloatSliceToEquation(equation))

// 	result := FactorQuadraticsWithABCAllPresent(equation)

// 	fmt.Println(DecodeFloatSliceToEquation(result))
	

// 	if(false){
// 		t.Errorf("failure")
// 	}

// }

// func TestQuadraticFactoringACOnlyPresent(t *testing.T){



// 	numberSlice := gNum(1, 2, 1, 6, 0)

// 	// numberSlice2 := gNum(1, 7, 1, 2, 2, 3, 2,2)

// 	equation := Create2DEquationFromSliceInputs(gOP(), numberSlice, gCP(2, 1))



// 	// equation := Create2DEquationFromSliceInputs(gOP(), gNum(3, 0, 3), gNum(1, 1, 2), gNum(3, 0, 5), gCP(1, 3))

// 	fmt.Println("initial equation", DecodeFloatSliceToEquation(equation))

// 	result := FactorQuadraticsWithACOnlyPresent(equation)

// 	fmt.Println(DecodeFloatSliceToEquation(result))
	

// 	if(false){
// 		t.Errorf("failure")
// 	}

// }


func TestQuadraticFactoringABOnlyPresent(t *testing.T){



	numberSlice := gNum(1, 2, 1, 6, 1)


	equation := Create2DEquationFromSliceInputs(gOP(), numberSlice, gCP(2, 1))



	// equation := Create2DEquationFromSliceInputs(gOP(), gNum(3, 0, 3), gNum(1, 1, 2), gNum(3, 0, 5), gCP(1, 3))

	fmt.Println("initial equation", DecodeFloatSliceToEquation(equation))

	result := FactorQuadraticsWithABOnlyPresent(equation)

	fmt.Println(DecodeFloatSliceToEquation(result))
	

	if(false){
		t.Errorf("failure")
	}

}

func TestCreatingATreeMap(t *testing.T){



	numberSlice0 := gNum(4, 2, 1, 4, 0)
	numberSlice1 := gNum(2, 2, 1, 3, 1)
	numberSlice2 := gNum(4, 9, 1, 3, 0)


	equation := Create2DEquationFromSliceInputs(gOP(), gOP(), numberSlice0, gCP(1, 1), gOP(), gOP(), numberSlice1, gCP(1, 1), gOP(), numberSlice2, gCP(1, 1), gCP(1, 1), gCP(1, 1))

	fmt.Println("initial equation", DecodeFloatSliceToEquation(equation))

	treeSlice := CreateEntireTreeForEquation(equation)

	IntelligentlyPrintTree(treeSlice)




	if(false){
		t.Errorf("failure")
	}



}


// func TestQuadraticFactoringABOnlyPresent(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(3, 2, 9, 1), gCP(1, 3)}       

// 	fmt.Println("initial", DecodeFloatSliceToEquation(equation))
	
// 	equation = FactorQuadraticsWithABOnlyPresent(equation)

// 	// equation = RemoveParenthesisWith0DirectChildren(equation)

// 	fmt.Println(" after", DecodeFloatSliceToEquation(equation))



// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestQuadraticFactoringACOnlyPresent(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(3, 2, 9, 0), gCP(1, 3)}       

// 	fmt.Println("initial", DecodeFloatSliceToEquation(equation))
	
// 	equation = FactorQuadraticsWithACOnlyPresent(equation)

// 	// equation = RemoveParenthesisWith0DirectChildren(equation)

// 	fmt.Println(" after", DecodeFloatSliceToEquation(equation))



// 	if(false){
// 		t.Errorf("failure")
// 	}
// }


// func TestCreatingATreeMap(t *testing.T){


// 	// equation := [][]complex128{gOP(), gOP(), gNum(3, 2), gCP(1, 3), gOP(), gNum(1, 1, 2, 0), gCP(1, 3), gOP(), gNum(7, 2), gCP(1, 3), gCP(1, 3), gOP(), gOP(), gNum(4, 3, 3, 2, 9, 1), gCP(1, 3), gOP(), gNum(4, 1), gCP(1, 3), gCP(1, 3)}       

// 	equation := [][]complex128{gOP(), gOP(), gOP(), gNum(3, 2), gCP(1, 3), gCP(1, 3), gOP(), gNum(2, 2), gOP(), gNum(2, 1), gCP(1, 3), gCP(1, 3), gCP(1, 3), gCP(1, 3)}       

// 	fmt.Println("initial", DecodeFloatSliceToEquation(equation))
	
// 	treeSlice := CreateEntireTreeForEquation(equation)

// 	IntelligentlyPrintTree(treeSlice)
	
// 	// for i := 0; i < len(treeSlice); i++ {

// 	// 	fmt.Println("parent at depth 0", DecodeFloatSliceToEquation(*treeSlice[i].Parent))


// 	// }	

// 	// childrenDepth0 := treeSlice[0].Children

// 	// fmt.Println(childrenDepth0)

// 	// for i := 0; i < len(childrenDepth0); i++ {

// 	// 	fmt.Println("parents at depth 1", DecodeFloatSliceToEquation(*childrenDepth0[i].Parent))

// 	// }

// 	// childrenDepth1 := treeSlice[0].Children[0].Children

// 	// // childrenDepth0 = treeSlice[0].Children

// 	// for i := 0; i < len(childrenDepth1); i++ {

// 	// 	fmt.Println("parents at depth 2", DecodeFloatSliceToEquation(*childrenDepth1[i].Parent))

// 	// }


// 	// childrenDepth1 = treeSlice[0].Children[1].Children

// 	// // childrenDepth0 = treeSlice[0].Children

// 	// for i := 0; i < len(childrenDepth1); i++ {

// 	// 	fmt.Println("parents at depth 2", DecodeFloatSliceToEquation(*childrenDepth1[i].Parent))

// 	// }



// 	//fmt.Println(treeSlice)
	

// 	if(false){
// 		t.Errorf("failure")
// 	}

	
// }