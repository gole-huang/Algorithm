package main

import (
	"Algorithm/Algo"
	"fmt"
)

func main() {
	//655
	root := &Algo.TreeNode{Val: 1, Left: &Algo.TreeNode{Val: 2}}
	fmt.Printf("Root: %v\nAnswer: %v\n", root, Algo.PrintTree(root))
}

/*
{{'1','0','1','1','1'},{'0','1','0','1','0'},{'1','1','0','1','1'},{'1','1','0','1','1'},{'0','1','1','1','1'}}
[][]byte{{'1','0','1','0','1','1','1','0'},{'1','1','0','1','1','0','0','0'},{'1','1','1','0','0','1','0','1'},{'1','0','1','1','1','1','1','0'},{'0','0','0','1','1','1','1','0'}}
{{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','0','0','1','1','1','1','1','1','1','1','0','0','1','1','1','0','1','1','1','1','1','1','1','1'},{'1','1','1','1','0','1','1','0','1','1','1','1','1','1','1','1','1','0','1','1','0','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'0','1','1','1','1','0','1','0','1','1','1','1','1','1','0','1','1','0','1','1','0','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'0','1','0','1','1','0','1','0','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','0','1','0','1','1','0','1','1','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','0','1','1','0','0','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','0','1','1','0','1','0','1','1','1','1','1','1','1','1','1','1','1','0','1','0','1','1','1','1','1','1','0','1','1','1','1'},{'0','1','1','0','1','1','0','1','0','1','1','1','0','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','0','1','0','1'},{'0','1','1','1','1','1','1','1','1','1','1','1','1','1','0','0','1','1','1','1','1','1','1','0','0','1','1','0','0','1','1','0','1','1','0','1','0','1','0','1'},{'1','1','1','1','0','1','1','1','1','0','1','1','1','1','1','1','1','1','1','0','1','1','0','1','1','0','1','1','1','1','0','1','0','1','1','0','1','0','1','1'},{'1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','0','1','1','0','1','1','0','1','1','1','0','1','1','1','1','0','1','1','1','1'},{'1','1','1','0','1','1','0','0','1','1','1','1','1','1','1','1','1','1','1','1','0','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','0','1','1','1','1','1','1','1','0','1','1','1','1','0','1','1','1','1','0','0','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'0','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','0','1','1','1','0','1','1','1','1','1','0','1','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','0','1','1'},{'1','1','1','1','1','0','0','1','1','1','1','1','1','1','1','0','1','0','1','1','0','0','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1'},{'1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1'},{'1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','0','0','1','1','1','1','1','1','0','0','1','1','1','1','1'},{'1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','0','1','1','1'},{'1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','0','1','1','1','1','1','0','0','1','0','1','1','1','1','1','0','1','1','1','1','1','1'},{'1','1','1','1','1','1','0','0','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','0','1','1','1','1','1','0','1','1','1','1','1','0','1','1','0','1','1'},{'1','1','0','0','0','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','0','1','0','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','0','0','1','0','1','1','1','0','0','1','1','1','1','1','1','1','1'},{'1','1','1','0','0','1','0','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','0','1','1','1','1','0','1','1','1','1','0','1','1','1','1','1','0','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','0','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','1','1','0','1','1','1','1','1','1','0','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1'},{'1','1','1','0','0','1','1','1','1','1','1','1','1','1','1','0','1','1','1','0','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','0','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','0','1','1','1','1','1','1','0','1','0','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','0','0','1','1','1','1','1','1','1','1','1','0','1','1','1','0','1'},{'1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','0','0','1','1','1','0','1','1','0','1','1'},{'1','1','1','1','0','1','1','0','1','1','1','1','1','1','0','1','1','0','1','1','0','1','1','1','1','1','1','0','1','1','1','1','1','1','1','0','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'1','1','0','0','0','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','0','1','1'},{'1','1','1','1','1','0','1','1','1','1','1','1','1','1','0','1','1','1','1','0','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1'},{'0','1','1','1','1','1','1','1','1','1','1','1','0','0','1','1','1','1','1','1','1','1','1','1','0','1','0','1','0','1','1','0','1','1','1','1','1','1','1','1'},{'1','0','1','1','0','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','0','0','1','1'},{'1','0','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','0','1','1','1','1','1'},{'0','1','1','1','1','0','1','1','1','1','0','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','1'},{'0','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','0','1','1','1','0','1','1','1','1','0','1','1','1','0','1','1','1','1','1','1','1','1','1','1'},{'0','1','1','1','1','1','1','1','1','1','1','1','0','1','0','1','1','1','1','0','1','1','1','1','1','1','0','1','0','1','1','0','0','1','1','1','1','0','1','1'},{'1','1','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','0','1','1','1','1','1','1','1','0','1','1','1','1','1','1','1','1','0','1','1','1','0'},{'1','1','1','1','1','0','1','1','1','1','1','1','1','1','0','0','1','1','1','1','1','1','1','1','1','1','1','0','1','1','1','1','1','1','0','0','1','1','1','1'},{'1','1','0','1','1','0','1','1','1','1','1','1','0','1','0','1','1','1','1','1','0','1','1','1','1','1','1','1','1','0','0','1','1','1','0','1','0','1','0','0'},{'0','1','1','0','1','1','1','1','1','1','1','0','0','1','1','1','1','1','0','0','1','0','1','1','1','1','1','0','1','1','1','0','1','1','0','1','1','1','0','1'}}
*/
