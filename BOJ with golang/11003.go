package main

import (
	"bufio"
	"fmt"
	"os"
)

func binarySearch(arr []int, x int) int {
	l := 0
	r := len(arr)
	for l < r {
		m := (l + r) / 2

		if arr[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}
func merge(arr []int, arr2 []int) []int {
	i := 0
	j := 0
	k := 0
	ans := make([]int, len(arr)+len(arr2))
	for i < len(arr) && j < len(arr2) {
		if arr[i] < arr2[j] {
			ans[k] = arr[i]
			i++
		} else {
			ans[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr) {
		ans[k] = arr[i]
		i++
		k++
	}
	for j < len(arr2) {
		ans[k] = arr2[j]
		j++
		k++
	}
	return ans
}

func initTree(start int, end int, node int, tree [][]int, arr []int) []int {

	if start == end {
		tree[node] = append(tree[node], arr[start])
		return tree[node]
	}

	mid := (start + end) / 2

	tree[node] = merge(initTree(start, mid, 2*node, tree, arr), initTree(mid+1, end, 2*node+1, tree, arr))
	return tree[node]
}

func query(start int, end int, node int, tree [][]int, arr []int, l, r int, x int) int {

	if start > r || end < l {
		return 0
	}
	if start >= l && end <= r {
		k := binarySearch(tree[node], x)
		//fmt.Println(k)
		return k
	}
	mid := (start + end) / 2
	a := query(start, mid, 2*node, tree, arr, l, r, x)
	b := query(mid+1, end, 2*node+1, tree, arr, l, r, x)
	//fmt.Println("a: ", a, " b: ", b)
	return a + b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscanf(reader, "%d %d\n", &n, &q)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d", &arr[i])
	}
	fmt.Fscanf(reader, "\n")
	tree := make([][]int, 4*n)
	initTree(0, n-1, 1, tree, arr)

	for i := 0; i < q; i++ {
		var x int
		var y int
		var z int
		fmt.Fscanf(reader, "%d %d %d\n", &x, &y, &z)
		left := int(-1e9)
		right := int(1e9)
		for left <= right {
			mid := (left + right) / 2

			k := query(0, n-1, 1, tree, arr, x-1, y-1, mid)
			if k >= z {
				right = mid - 1
			} else {

				left = mid + 1
			}
		}
		fmt.Println(left)
	}
	fmt.Print(tree)

}
