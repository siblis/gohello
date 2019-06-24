package stack

var x []string

//Push places the str on the top of stack
func Push(str string) {

	x = append(x, str)
}

//Pop gets the str from the top of stack
func Pop() string {
	num := len(x)
	if num == 0 {
		return "[empty]"
	}
	popElem := x[num-1]
	x = x[:num-1]
	return popElem

}
