package main

import "fmt"

func main() {
	//make関数を使う
	//第一引数:配列の型?
	//第二引数:長さ（len)
	//第三引数:容量(cap)
	//https://blog.golang.org/go-slices-usage-and-internals
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
