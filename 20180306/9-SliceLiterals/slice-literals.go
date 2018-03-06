package main

import "fmt"

func main() {
	//配列リテラル
	array := [3]bool{true, false, false}
	fmt.Println(array)

	//スライスリテラル。Int編
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	//スライスリテラル。Bool編
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	//スライスリテラル。struct編
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
