package main

import "gohello/right"

func main() {

	right.Println(right.Execute | right.Read)
	right.Println(right.Execute | right.Read | right.Write)
	right.Println(right.Execute | right.Write)
}
