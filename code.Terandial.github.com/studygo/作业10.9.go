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
	for {

		fmt.Scanf("%d", &a)
		if a > secretNumber {
			fmt.Println("傻逼你是错的")
			fmt.Println(a)

		} else if a < secretNumber {
			fmt.Println("你太小了")
			fmt.Println(a)

		} else {
			fmt.Println("靠！居然被你猜中！")
			fmt.Println(a)
			break
		}

	}

}
