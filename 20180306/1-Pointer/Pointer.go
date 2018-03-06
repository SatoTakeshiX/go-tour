package main

import "fmt"

//http://cuto.unirita.co.jp/gostudy/post/pointer/
//C言語と同じくGOにもポインタがある。
//変数 T のポインタは、 *T 型で、ゼロ値は nil です。
//& オペレータは、そのオペランド( operand )へのポインタを引き出します。
//* オペレータで、ポインタの指す先の変数を示します。
//C言語と異なりポイント演算はない。あるポインタを足すなどして別のポインタにすることはできない。
//invalid operation: p + p (operator + not defined on pointer)

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	//anotherPoint := p + p
	//fmt.Println(anotherPoint)
	fmt.Println(p) // iのポインタを示す。
	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
