package convert

import (
	//"fmt"
	"polish_notation/Containers/queue"
	"strconv"
	"unicode/utf8"
)

func SplitOnTokens(data string) (*queue.Queuelist, error) {
	parse_data := new(queue.Queuelist)
	size := utf8.RuneCountInString(data)
	var err error
	for i := 0; i < size;  {
		if isOperator(string(data[i])) {			
			if string(data[i]) == "-" && (i == 0 || string(data[i-1]) == "(") {
				parse_data.Enqueue("~")
			} else {
				parse_data.Enqueue(string(data[i]))
			}
			
			i++
		} else {
			var val string
			for i < size && !isOperator(string(data[i])) {
				val += string(data[i])
				i++
			}
			_, err = strconv.ParseFloat(val, 64)
			if err != nil {
				i = size
			}
			parse_data.Enqueue(val)
		}
	}
	return parse_data, err
}