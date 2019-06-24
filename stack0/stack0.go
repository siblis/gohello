package stack0

var x []string

//Push - push string to stack
func Push(s string) {
	x = append(x, s)
}

//Pop - extract string from stack
func Pop() string {
	num := len(x)
	if num == 0 {
		return "[Empty]"
	}
	popEl := x[num-1]
	x = x[:num-1]
	return popEl
}
