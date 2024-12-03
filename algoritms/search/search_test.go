package search

import (
	"alg/src"
	"alg/tracking"
	"math/rand"
	"testing"
)

func TestLazySearch(t *testing.T) {
	defer tracking.Duration(tracking.Track("Время выполнения Ленивого поиска: "))
	randomLen := rand.Intn(1000)
	myList := src.GetRange(randomLen)
	Lazy(randomLen, myList)
}

func TestBynarySearch(t *testing.T) {
	defer tracking.Duration(tracking.Track("Время выполнения Бинарного поиска: "))
	randomLen := rand.Intn(1000)
	myList := src.GetRange(randomLen)
	Binary(randomLen, myList)
}
