package main

import (
	algo "Algorithm/Algo"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	s := "rabbbit"
	t := "rabbit"
	fmt.Printf("S: %v\tT: %v\nANSWER: %v\n", len(s), len(t), algo.NumDistinct(s, t))
	fmt.Println(time.Since(start))
}

/*

 */
