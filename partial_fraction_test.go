package main

import (


	"testing"
	"fmt"
	"math"

)


func TestFindingHighestDegree(t *testing.T) {		

	Init()	



	


	if (result != 2) {
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

















