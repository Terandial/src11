package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxMum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxMum)
	fmt.Println("please input your guess")
	var a int
	var c int
	for {

		fmt.Scan(&a)
		fmt.Scanln()

		if a > secretNumber {
			fmt.Println("傻逼你是错的")
			fmt.Println(a)
			c = c + 1

		} else if a < secretNumber {
			fmt.Println("你太小了")
			fmt.Println(a)
			c = c + 1
		} else {
			fmt.Println("靠！居然被你猜中！")
			fmt.Println(a)
			c = c + 1
			break

		}

	}
	fmt.Println("你还记得你循环了几次码")
	fmt.Println("有种就输入啊!")
	var b int

	fmt.Scanf("%d", &b)
	if b == c {
		fmt.Println("哼！算你厉害！")

	} else {
		fmt.Println("不会吧不会吧！不会真的有人会被猜数游戏吸引注意力吧！这都记不住！")
	}
}
