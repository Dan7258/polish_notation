package convert

import (
	"polish_notation/Containers/queue"
	"polish_notation/Containers/stack"
	"strings"
)

func ConvertToPolish(formula *queue.Queuelist) *queue.Queuelist {
	stack := new(stack.Stacklist)
	polishFormula := new(queue.Queuelist)
	for !(formula.IsEmpty()) {
		op1 := formula.Dequeue()
		if IsOperator(op1) || IsMathFunc(op1) {
			op2 := stack.Pop()
			if op1 ==")" {
				for op2 != "" && op2 != "(" {
					polishFormula.Enqueue(op2)
					op2 = stack.Pop()
				}
				continue
			}
			if (priority(op1) > priority(op2)) || op1 == "(" {
				stack.Push(op2)
				stack.Push(op1)
			} else {
				for priority(op1) <= priority(op2) && op2 != "" {
					polishFormula.Enqueue(op2)
					op2 = stack.Pop()
				}
				stack.Push(op1)
			}
		} else {
			polishFormula.Enqueue(op1)
		}
	}
	
	for op2:= stack.Pop(); op2 != ""; {
		polishFormula.Enqueue(op2)
		op2 = stack.Pop()
	}
	return polishFormula
}

func IsOperator(data string) bool {
	operators := "+-*/()~"
	answer := strings.Contains(operators,data)
	return answer
}

func priority(operator string) int{
	list_priority := map[string]int{"": -1, "(": 1, ")": 1, "+": 2, "-": 2, "*": 3, "/": 3, "sin": 5, "cos": 5, "tan": 5, "ctg": 5, "ln": 5, "sqrt": 5, "~": 4, }
	prior := list_priority[operator]
	return prior
}
