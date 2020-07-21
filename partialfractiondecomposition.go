package main

import (

	"fmt"
	"os"
)

type EquationItem struct {
	Items []interface{}
}


type S_Var struct {
	Multiplier float64
	Exponent int

}

type GeneralVariable struct {
	Name string
	Multiplier float64
	DegreeToCompareToS int
	
}



func main() {
	fmt.Println("hello")
}




func MultiplyNumeratorByOppositeDenominator(genVarSlice []GeneralVariable, sVarSlice []S_Var, constant float64) []GeneralVariable {

	returnGeneralVariablesSlice := []GeneralVariable{}

	for i := 0; i < len(genVarSlice); i++ {
		for j := 0; j < len(sVarSlice); j++ {
			returnGeneralVariablesSlice = append(returnGeneralVariablesSlice, genVarTimesSVar(genVarSlice[i], sVarSlice[j]))
		}
		
	}


	for i := 0; i < len(genVarSlice); i++ {



		returnGeneralVariablesSlice = append(returnGeneralVariablesSlice, GeneralVariable{genVarSlice[i].Name, (genVarSlice[i].Multiplier * constant), genVarSlice[i].DegreeToCompareToS})
	
	}



	return returnGeneralVariablesSlice


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













