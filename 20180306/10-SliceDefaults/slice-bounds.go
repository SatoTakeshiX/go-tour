package main

import "fmt"

//スライスのデファルと
func main() {

	s := []int{2, 3, 5, 7, 11, 13}

	//スライスのデファオルト。下限が0で上限がスライスの長さになる
	s = s[:]
	fmt.Println(s) //[2 3 5 7 11 13]

	s = s[1:4]     //2番目以降5番目未満を取る
	fmt.Println(s) //[3 5 7]

	s = s[:2]      //スライスの3番目未満(2番目まで)をとる
	fmt.Println(s) //[3 5]

	s = s[1:]      //二番目を取る
	fmt.Println(s) //[5]

	s = s[1:]
	fmt.Println(s)

	//スライスの範囲外を取得しようとするとエラー
	//panic: runtime error: slice bounds out of range

	// s = s[1:]
	// fmt.Println(s)

}
