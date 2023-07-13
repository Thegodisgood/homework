package main

import "fmt"

func main() {
	var n, l, r int
	fmt.Scan(&n, &l, &r)
	var a [1000]int
	for i := 0; i <= n; i++ {
		fmt.Scanf("%d", &a[i])

	}
	for i := l + 1; i <= r; i++ {
		if a[i] > a[i+1] {
			a[i], a[i+1] = a[i+1], a[i]
		}
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
