package main

import "fmt"

type vertex struct {
	x int
	y int
}

func main() {
	// v := vertex{1, 2}
	// v.x = 4
	// fmt.Println(v.x)

	v := vertex{1, 2}
	notPoint := v
	notPoint.y = 122
	p := &v //vのpointを取る
	p.x = 1e9

	fmt.Println(v)        //ポイントpを変更したからvも変わった！
	fmt.Println(p)        //変更したとおりの値になった！
	fmt.Println(notPoint) //ポイントpの影響は受けていない！
}
