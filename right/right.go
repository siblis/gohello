package right

import "fmt"

//Standard File Rights
const (
	Execute = 1
	Write   = 2
	Read    = 4
)

//Println - prints octal right for file
func Println(b uint) {
	fmt.Printf("%o\n", b)
}
