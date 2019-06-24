package main

import "fmt"

func main() {

	a := []int64{0, 1, 2, 3, 4, 5}
	b := a[:3]
	c := a[3:]
	fmt.Println(b, "<>", c)
}
