package calculate

import (
	"fmt"
	"math"
	"polish_notation/Containers/queue"
	"polish_notation/Containers/stack"
	"polish_notation/Opportunities/convert"
	"strconv"
)

func Calculate(polish_formula *queue.Queuelist,) float64 {
	stack := new(stack.Stacklist)

	for !polish_formula.IsEmpty() {
		token := polish_formula.Dequeue()
		if convert.IsOperator(token) || convert.IsMathFunc(token) {
			choseOperation(stack, token)
		} else {
			stack.Push(token)
		}
	}
	res, _ := strconv.ParseFloat(stack.Pop(), 64)
	return res
}

func choseOperation(stack *stack.Stacklist, token string) {
	switch token {
	case "+":
		op2, _ := strconv.ParseFloat(stack.Pop(), 64)
		op1, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", op1+op2))
		
	case "-":
		op2, _ := strconv.ParseFloat(stack.Pop(), 64)
		op1, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", op1-op2))
		
	case "*":
		op2, _ := strconv.ParseFloat(stack.Pop(), 64)
		op1, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", op1*op2))
		
	case "/":
		op2, _ := strconv.ParseFloat(stack.Pop(), 64)
		op1, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", op1/op2))
		
	case "~":
		op, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", op * -1))
		
	case "sin":
		op, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", math.Sin(op * math.Pi / 180.0)))
		
	case "cos":
		op, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", math.Cos(op * math.Pi / 180.0)) )
		
	case "tan":
		op, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", math.Tan(op * math.Pi / 180.0)))
		
	case "ctg":
		op, _ := strconv.ParseFloat(stack.Pop(), 64)
		op = op * math.Pi / 180.0
		stack.Push(fmt.Sprintf("%f", math.Cos(op) / math.Sin(op)))
		
	case "ln":
		op, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", math.Log(op)))
		
	case "sqrt":
		op, _ := strconv.ParseFloat(stack.Pop(), 64)
		stack.Push(fmt.Sprintf("%f", math.Sqrt(op)))
		
	}
}