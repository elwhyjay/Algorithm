package main

import "fmt"

func multiple(a [][]int64, b [][]int64) [][]int64 {
	_size := len(a)
	ans := make([][]int64, _size)
	for i := 0; i < _size; i++ {
		ans[i] = make([]int64, _size)
	}
	for i := 0; i < _size; i++ {
		for j := 0; j < _size; j++ {
			for k := 0; k < _size; k++ {
				ans[i][j] += a[i][k] * b[k][j]
				ans[i][j] %= 1000
			}
		}
	}

	return ans
}

func square(a [][]int64, b int64) [][]int64 {
	if b == 1 {
		return a
	}

	ans := square(a, b/2)

	if b%2 == 1 {
		return multiple(multiple(ans, ans), a)
	} else {
		return multiple(ans, ans)
	}
}

func main() {
	var n int
	var s int64
	fmt.Scan(&n, &s)
	a := make([][]int64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	ans := square(a, s)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", ans[i][j]%1000)
		}
		fmt.Println()
	}

}
