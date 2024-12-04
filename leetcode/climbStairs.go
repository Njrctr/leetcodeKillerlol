package leetcode

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	massive := make(map[int]int)
	massive[2], massive[3] = 2, 3
	ok := massive[n]
	if ok != 0 {
		return ok
	}
	for x := 4; x <= n; x++ {
		massive[x] = massive[x-1] + massive[x-2]
	}
	return massive[n]
}
