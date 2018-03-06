package main

import "fmt"

//sliceはarrayのポインターの・ようなもの。取得した配列要素を書き換えると元の配列データも書き換わる。
func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"       //bのスライスPaulを書き換え。
	fmt.Println(a, b)  //2インデックス目の要素がXXXになった
	fmt.Println(names) //もとの配列も書き換わってる
}
