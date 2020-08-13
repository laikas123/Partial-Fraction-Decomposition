package main

import (
	"fmt"
	"strings"
	"testing"

)


func TestAreIdentical(t *testing.T){


	equation1 := [][]complex128{gOP(), gNum(2, 3, 3, 1), gCP(3), gCP(3), gCP(3)}       
	equation2 := [][]complex128{gOP(), gNum(2, 3, 3, 1), gCP(3), gCP(3), gCP(3)}       

	fmt.Println("equation1", strings.ReplaceAll(DecodeFloatSliceToEquation(equation1), " ", ""))
	fmt.Println("equation2", strings.ReplaceAll(DecodeFloatSliceToEquation(equation2), " ", ""))

	fmt.Println("are identical", TwoEquationsAreExactlyIdentical(equation2, equation1))



	if(false){
		t.Errorf("failure")
	}
}



func TestCleanCopy(t *testing.T){


	equation := [][]complex128{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3), gCP(3), gOP(), gOP(), gOP(), gCP(3), gCP(3), gOP(), gOP(), gOP()}       

	cleanCopy := CleanCopyEntire2DComplex128Slice(equation)

	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))



	fmt.Println("   copy", strings.ReplaceAll(DecodeFloatSliceToEquation(cleanCopy), " ", ""))

	

	if(false){
		t.Errorf("failure")
	}
}

func TestRemoveTrailingOpeners(t *testing.T){


	equation := [][]complex128{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3), gCP(3), gOP(), gOP(), gOP()}       

	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
	equation = RemoveLastItemIfItIsOpeningParenthesis(equation)

	fmt.Println("after", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))

	

	if(false){
		t.Errorf("failure")
	}
}


func TestDepthCheckRemove(t *testing.T){


	equation := [][]complex128{gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3), gCP(3), gCP(3)}       

	fmt.Println("initial", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))
	
	equation = RemoveExcessParenthesisViaDepthCheck(equation)

	fmt.Println(" after", strings.ReplaceAll(DecodeFloatSliceToEquation(equation), " ", ""))

	

	if(false){
		t.Errorf("failure")
	}
}


