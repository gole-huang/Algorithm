package main

import (
	"Algorithm/Algo"
	"fmt"
)

func main() {
	/*	//31
		nums := [...]int{1, 3, 2}
		numsSlice := nums[:]
		Algo.NextPermutation(numsSlice)
		fmt.Printf("Result: %v\n", numsSlice)
	*/
	/*	//41
		nums := []int{1, 2, 0}
		fmt.Printf("%v\n", Algo.FirstMissingPositive(nums))
	*/
	/*//1200
	arr := []int{1, 3, 6, 10, 15}
	fmt.Printf("Result: %v\n", Algo.MinimumAbsDifference(arr))
	*/
	//33
	/*
		nums, target := []int{4, 5, 6, 7, 8, 1, 2, 3}, 8
		fmt.Printf("%v", Algo.Search(nums, target))
	*/
	//22
	/*
		str := Algo.GenerateParenthesis(8)
		for _, s := range str {
			fmt.Printf("%v\n", s)
		}
	*/
	//38
	/*
		fmt.Printf("%v", Algo.CountAndSay(20))
	*/
	//34
	/*
		nums := []int{1, 4}
		target := 4
		fmt.Printf("%v", Algo.SearchRange(nums, target))
	*/
	//36
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	//fmt.Printf("%v\n", board)
	fmt.Printf("%v", Algo.IsValidSudoku(board))
}

/*
board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
*/
