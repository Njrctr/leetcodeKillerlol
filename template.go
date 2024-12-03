package main

import (
	"bufio"
	"fmt"
	"os"
)

func test() {
	// 	5 6 1
	// 7 2 7 4 7
	reader := bufio.NewReader(os.Stdin)
	var countOfTowers, maximalHeight, maxPhotoshop int
	var heightsOfTowers []int
	fmt.Fscan(reader, &countOfTowers)
	fmt.Fscan(reader, &maximalHeight)
	fmt.Fscan(reader, &maxPhotoshop)
	x := 0
	for x < countOfTowers {
		var now int
		fmt.Fscan(reader, &now)
		heightsOfTowers = append(heightsOfTowers, now)
	}

}
