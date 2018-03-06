package main

import "fmt"

//シグネチャ
//func append(slice []Type, elems ...Type) []Type

func main() {

	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	s = append(s, 5, 6)
	printSlice(s) //なんでcapが増えてるんだろう？
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
