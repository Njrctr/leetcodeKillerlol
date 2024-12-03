package skbkontur

import (
	"bufio"
	"fmt"
	"os"
)

var runeShip = []rune("X")[0]
var emptyRune = []rune(".")[0]
var fire = []rune("f")[0]

func _() {
	reader := bufio.NewReader(os.Stdin)
	var heiht, width, countofshots int
	fmt.Fscan(reader, &heiht)
	fmt.Fscan(reader, &width)
	fmt.Fscan(reader, &countofshots)
	var pole [][]rune
	for x := 0; x < heiht; x++ {
		var now string
		fmt.Fscan(reader, &now)
		pole = append(pole, []rune(now))
	}
	var shots [][]int
	for x := 0; x < countofshots; x++ {
		var nowx int
		var nowy int
		fmt.Fscan(reader, &nowx)
		fmt.Fscan(reader, &nowy)
		shots = append(shots, []int{nowx - 1, nowy - 1})
	}
	spaceWar(pole, shots)
}

func spaceWar(place [][]rune, shots [][]int) {
	for _, x := range shots {
		// place[x[0]] - строка в которую идёт выстрел
		var result rune
		if place[x[0]][x[1]] == emptyRune {
			fmt.Println(" MISS")
			continue
		}
		if len(place[x[0]]) == 1 {
			if place[x[0]][x[1]] == runeShip {
				result = emptyRune
			}
		} else if x[1] == 0 {
			result = check(place[x[0]], x[1]+1, true)
			place[x[0]][x[1]] = result
		} else if x[1] == len(place[x[0]])-1 {
			result = check(place[x[0]], x[1]-1, false)
			place[x[0]][x[1]] = result
		} else {
			resultleft := check(place[x[0]], x[1]-1, false)
			resultright := check(place[x[0]], x[1]+1, true)
			if resultleft == resultright {
				result = resultleft
				place[x[0]][x[1]] = result
			} else {
				result = fire
				place[x[0]][x[1]] = result
			}
		}
		switch result {
		case fire:
			fmt.Println("HIT")
		case emptyRune:
			fmt.Println("DESTROY")
		}
	}
}

func check(row []rune, point int, toright bool) rune {
	switch row[point] {
	case runeShip:
		return fire
	case fire:
		if toright {
			if point == len(row)-1 {
				return emptyRune
			}
			return check(row, point+1, toright)
		}
		if point == 0 {
			return emptyRune
		}
		return check(row, point-1, toright)
	default:
		return emptyRune
	}
}
