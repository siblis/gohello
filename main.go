package main

import (
	"fmt"
	"gohello/right"
)

func main() {
	fmt.Println("My first homework")
	right.Println(right.Execute | right.Read)
	right.Println(right.Execute | right.Read | right.Write)
	right.Println(right.Execute | right.Write)
}
