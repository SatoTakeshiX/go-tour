package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	// マルバツゲームをします。
	//[]T{value1, value2}でスライスを作る。
	//Tがスライスの時、ここでは[]stringの時、スライスのスライスが作れる。
	//他の言語では２次元配列に当たる。
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	fmt.Println(board)

	// The players take turns.
	// 
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
