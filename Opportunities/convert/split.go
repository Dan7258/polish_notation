package convert

import (
	"polish_notation/Containers/queue"
	"strconv"
	"unicode/utf8"
)

func SplitOnTokens(data string, x string) (*queue.Queuelist, error) {
	parse_data := new(queue.Queuelist)
	size := utf8.RuneCountInString(data)
	var err error
	for i := 0; i < size;  {
		if IsOperator(string(data[i])) {			
			if string(data[i]) == "-" && (i == 0 || string(data[i-1]) == "(") {
				parse_data.Enqueue("~")
			} else {
				parse_data.Enqueue(string(data[i]))
			}
			i++
		} else {
			var val string
			for i < size && !IsOperator(string(data[i])) {
				val += string(data[i])
				i++
			}
			if !IsMathFunc(val) && val != "x" {
				_, err = strconv.ParseFloat(val, 64)
				if err != nil {
					i = size
				}
			}
			if val == "x" {
				val = x
			}
			parse_data.Enqueue(val)
		}
	}
	return parse_data, err
}

func IsMathFunc(data string) bool {
	list_math_func := map[string]struct{}{"sin": {}, "cos": {}, "tan": {}, "ctg": {}, "ln": {}, "sqrt": {}, }
	_, ok := list_math_func[data]
	return ok
}