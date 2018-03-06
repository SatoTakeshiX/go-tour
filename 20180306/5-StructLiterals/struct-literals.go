package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1        = Vertex{1, 2}  // 通常の初期化
	v2        = Vertex{X: 1}  // Xラベルに値を入れる。Y:0 は暗黙的に決定
	v3        = Vertex{}      // X:0 and Y:0
	p         = &Vertex{1, 2} // has type *Vertex
	withLabel = Vertex{X: 111, Y: 200}
)

func main() {
	fmt.Println(v1, p, v2, v3, withLabel)
}
