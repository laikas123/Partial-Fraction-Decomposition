package main

import (
	"fmt"
	"testing"
)





func TestRemoveUnusedParenthesis(t *testing.T){


	equation := [][]float64{gOP(), gOP(), gOP(), gNum(2, 3, 3, 0), gCP(3), gCP(2), gCP(1), gOP(), gOP(), gOP(), gCP(3), gCP(2), gCP(1)}       
	
	equation = RemoveUnusedParenthesis(equation)

	// equation = RemoveUnusedParenthesis(equation)

	fmt.Println(DecodeFloatSliceToEquation(equation))

	if(false){
		t.Errorf("failure")
	}
}