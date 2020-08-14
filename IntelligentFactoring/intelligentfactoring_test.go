package main

import (
	"fmt"
	"strings"
	"testing"

)


// func TestAreIdentical(t *testing.T){


// 	equation1 := [][]complex128{gOP(), gNum(2, 3, 3, 1), gCP(3), gCP(3), gCP(3)}       
// 	equation2 := [][]complex128{gOP(), gNum(2, 3, 3, 1), gCP(3), gCP(3), gCP(3)}       

// 	fmt.Println("equation1", strings.ReplaceAll(DecodeFloatSliceToEquation(equation1), " ", ""))
// 	fmt.Println("equation2", strings.ReplaceAll(DecodeFloatSliceToEquation(equation2), " ", ""))

// 	fmt.Println("are identical", TwoEquationsAreExactlyIdentical(equation2, equation1))



// 	if(false){
// 		t.Errorf("failure")
// 	}
// }



// func TestCleanCopy(t *testing.T){


// 	equation := [][]complex128{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3), gCP(3), gOP(), gOP(), gOP(), gCP(3), gCP(3), gOP(), gOP(), gOP()}       

// 	cleanCopy := CleanCopyEntire2DComplex128Slice(equation)

// 	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))



// 	fmt.Println("   copy", strings.ReplaceAll(DecodeFloatSliceToEquation(cleanCopy), " ", ""))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestRemoveTrailingOpeners(t *testing.T){


// 	equation := [][]complex128{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3), gCP(3), gOP(), gOP(), gOP()}       

// 	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
// 	equation = RemoveLastItemIfItIsOpeningParenthesis(equation)

// 	fmt.Println("after", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }


// func TestDepthCheckRemove(t *testing.T){


// 	equation := [][]complex128{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3), gCP(3), gCP(3)}       

// 	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
// 	equation = RemoveExcessParenthesisViaDepthCheck(equation)

// 	fmt.Println(" after", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }



// func TestFoilingExponentStyle(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(2, 3, 3, 0), gCP(3), gOP(), gNum(2, 3, 3, 0), gCP(3), gOP(), gNum(2, 3, 3, 0), gCP(2)}       

// 	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
// 	equation = FoilOutParenthesisRaisedToExponent(equation)

// 	fmt.Println(" after", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestFoilingNeighborStyle(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(2, 3, 3, 0), gCP(1), gOP(), gNum(2, 3, 3, 0), gCP(1), gOP(), gNum(2, 3, 3, 0), gCP(1), gOP(), gNum(2, 3, 3, 0), gCP(1)}       

// 	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
// 	equation = FoilNeighborParenthesis(equation)

// 	fmt.Println(" after", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))

	

// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestQuadraticFactoringABCPresent(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(2, 2, -3, 1, 5, 0), gCP(1), gOP(), gCP(1), gOP(), gCP(1), gOP(), gCP(1), gOP(), gCP(1), gOP(), gCP(1), gOP(), gNum(2, 2, -3, 1, 5, 0), gCP(1), gOP(), gOP(), gCP(1), gCP(1)}       

// 	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
// 	equation = FactorQuadraticsWithABCAllPresent(equation)

// 	equation = RemoveParenthesisWith0DirectChildren(equation)

// 	fmt.Println(" after", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))



// 	if(false){
// 		t.Errorf("failure")
// 	}
// }


// func TestQuadraticFactoringABOnlyPresent(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(3, 2, 9, 1), gCP(1)}       

// 	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
// 	equation = FactorQuadraticsWithABOnlyPresent(equation)

// 	// equation = RemoveParenthesisWith0DirectChildren(equation)

// 	fmt.Println(" after", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))



// 	if(false){
// 		t.Errorf("failure")
// 	}
// }

// func TestQuadraticFactoringACOnlyPresent(t *testing.T){


// 	equation := [][]complex128{gOP(), gNum(3, 2, 9, 0), gCP(1)}       

// 	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
// 	equation = FactorQuadraticsWithACOnlyPresent(equation)

// 	// equation = RemoveParenthesisWith0DirectChildren(equation)

// 	fmt.Println(" after", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))



// 	if(false){
// 		t.Errorf("failure")
// 	}
// }


func TestCreatingATreeMap(t *testing.T){


	// equation := [][]complex128{gOP(), gOP(), gNum(3, 2), gCP(1), gOP(), gNum(1, 1, 2, 0), gCP(1), gOP(), gNum(7, 2), gCP(1), gCP(1), gOP(), gOP(), gNum(4, 3, 3, 2, 9, 1), gCP(1), gOP(), gNum(4, 1), gCP(1), gCP(1)}       

	equation := [][]complex128{gOP(), gOP(), gOP(), gNum(3, 2), gCP(1), gCP(1), gOP(), gNum(2, 2), gOP(), gNum(2, 1), gCP(1), gCP(1), gCP(1), gCP(1)}       

	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
	treeSlice := CreateEntireTreeForEquation(equation)

	IntelligentlyPrintTree(treeSlice)
	
	// for i := 0; i < len(treeSlice); i++ {

	// 	fmt.Println("parent at depth 0", strings.ReplaceAll(DecodeFloatSliceToEquation(*treeSlice[i].Parent), " ", ""))


	// }	

	// childrenDepth0 := treeSlice[0].Children

	// fmt.Println(childrenDepth0)

	// for i := 0; i < len(childrenDepth0); i++ {

	// 	fmt.Println("parents at depth 1", strings.ReplaceAll(DecodeFloatSliceToEquation(*childrenDepth0[i].Parent), " ", ""))

	// }

	// childrenDepth1 := treeSlice[0].Children[0].Children

	// // childrenDepth0 = treeSlice[0].Children

	// for i := 0; i < len(childrenDepth1); i++ {

	// 	fmt.Println("parents at depth 2", strings.ReplaceAll(DecodeFloatSliceToEquation(*childrenDepth1[i].Parent), " ", ""))

	// }


	// childrenDepth1 = treeSlice[0].Children[1].Children

	// // childrenDepth0 = treeSlice[0].Children

	// for i := 0; i < len(childrenDepth1); i++ {

	// 	fmt.Println("parents at depth 2", strings.ReplaceAll(DecodeFloatSliceToEquation(*childrenDepth1[i].Parent), " ", ""))

	// }



	//fmt.Println(treeSlice)
	

	if(false){
		t.Errorf("failure")
	}

	
}