package main

import "fmt"

func main() {
	// 声明了长度为5的数组，数组中的每一个元素都是int类型
	var a [5]int
	// 给数组a的第4位元素赋值为100
	a[4] = 100
	fmt.Println("get:", a[2])
	fmt.Println("len:", len(a))

	// 在给数组声明的同时赋值
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 声明二位数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
