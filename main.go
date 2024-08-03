package main

import "polish_notation/Web/server"

// "fmt"
// "log"
// "polish_notation/Opportunities/calculate"
// "polish_notation/Opportunities/convert"

func main() {
	
	// val := ""
	// fmt.Scanf("%s", &val)
	// formula, err := convert.SplitOnTokens(val)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	
	// polish_formula := convert.ConvertToPolish(formula)
	// fmt.Printf("%f", calculate.Calculate(polish_formula))
	// for !polish_formula.IsEmpty() {
	// 	fmt.Printf("%s ", polish_formula.Dequeue())
	// }

	server.RunServer()
}