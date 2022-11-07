package main

import (
	"fmt"
)

func main() {

	var n, i, r int
	_, err := fmt.Scanf("%d %d %d\n", &n, &i, &r)
	if err != nil {
		fmt.Println("输入的数有错误")
	}
	a := make([]int, 10)
	num := 0
	b := 0
	for num = 0; num < n; num++ {
		fmt.Scanf("%d", &b)

		a[num] = b
	}
	sort(a, i, r)
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
func sort(a []int, i, r int) []int {
	var num1, num2, num int
	for num1 = i; num1 <= r; num1++ {
		for num2 = num1; num2 <= r; num2++ {
			if a[num1] > a[num2] {
				num = a[num1]
				a[num1] = a[num2]
				a[num2] = num
			}
		}
	}
	return a
}
