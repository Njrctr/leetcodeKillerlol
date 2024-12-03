package skbkontur

import (
	"bufio"
	"fmt"
	"os"
)

func _() {
	reader := bufio.NewReader(os.Stdin)
	var countOfShots int
	fmt.Fscan(reader, &countOfShots)

	var predicttargets [][]int
	for x := countOfShots; x != 0; x-- {
		var x int
		var y int
		fmt.Fscan(reader, &x)
		fmt.Fscan(reader, &y)
		predicttargets = append(predicttargets, []int{x, y})
		x++
	}
	var facttargets [][]int
	for x := countOfShots; x != 0; x-- {
		var x int
		var y int
		fmt.Fscan(reader, &x)
		fmt.Fscan(reader, &y)
		facttargets = append(facttargets, []int{x, y})
		x++
	}
	fmt.Println(checkChangeTraectory(countOfShots, predicttargets, facttargets))

}

func checkChangeTraectory(countShot int, predicttargets, facttargets [][]int) (int, int) {
	totalPredict := []int{0, 0}
	totalFact := []int{0, 0}
	for i, shot := range predicttargets {
		totalPredict[0] += shot[0]
		totalPredict[1] += shot[1]
		totalFact[0] += facttargets[i][0]
		totalFact[1] += facttargets[i][1]
	}
	result := []int{0, 0}
	result[0], result[1] = (totalFact[0]-totalPredict[0])/countShot, (totalFact[1]-totalPredict[1])/countShot
	return result[0], result[1]
}
