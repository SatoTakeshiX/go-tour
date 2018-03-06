package main

import "fmt"

//https://blog.golang.org/go-slices-usage-and-internals
/**
長さは、スライスによって参照される要素の数です。容量は、基本となる配列の要素の数です（スライスポインタで参照される要素から始まります）。
*/
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	//長さは0で容量は6
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	//panic: runtime error: slice bounds out of range
	s = s[2:]
	printSlice(s)

}

//スライスを引数に長さと容量を出力する関数
func printSlice(s []int) {
	//len関数とcap関数を実行する。
	//インスタンスに紐付いているわけではない？ s.len or cap.s　じゃない
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
