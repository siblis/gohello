//#!/usr/bin/go run
package main

import "fmt"

func addr() {
	var x map[string][]int
	x = make(map[string][]int)
	x["A"] = append(x["A"], 8999999)
	x["B"] = []int{1111}
	x["A"] = append(x["A"], 911)
	x["D"] = []int{222}
	delete(x, "D")
	for i, nums := range x { //[]string{"A", "B", "C", "D"} {
		//if nums, ok := x[i]; ok {
		fmt.Print(i)
		for _, num := range nums {
			fmt.Print(" ", num)
		}
		fmt.Println("found!")
		//} else {
		//	fmt.Println(i, "not found")
		//}
	}
}
