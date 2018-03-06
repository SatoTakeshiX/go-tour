package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//型 []T は 型 T のスライスを表します。
	//a[low : high](low以上、high未満)
	var s []int = primes[1:4]
	fmt.Println(s)

	var slice2 []int = primes[0:6]
	fmt.Println(slice2) //[2 3 5 7 11 13]

	//配列の範囲を超えるとだめ
	//invalid slice index 8 (out of bounds for 6-element array) primes [6]int
	//slice3 := primes[8:10]

}
