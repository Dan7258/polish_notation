package main

import (
	"fmt"
	"log"
	//"polish_notation/Containers/queue"
	"polish_notation/Opportunities/convert"
	// "polish_notation/Containers/stack"
)

func main() {
	
	val := ""
	fmt.Scanf("%s", &val)
	formula, err := convert.SplitOnTokens(val)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	polish_formula := convert.ConvertToPolish(formula)
	for !polish_formula.IsEmpty() {
		fmt.Printf("%s ", polish_formula.Dequeue())
	}
}