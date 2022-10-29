package main

import (
	"fmt"
)

func 意大利炮() {

	r := make(chan 操作, 10)
	f := make(chan 操作, 5)
	g := make(chan 操作, 3)
	r <- 操作{
		Name:  "装弹",
		Count: 10,
	}
	f <- 操作{
		Name:  "瞄准",
		Count: 5,
	}
	g <- 操作{
		Name:  "发射！",
		Count: 3,
	}
	if g == nil && f == nil {
		f <- 操作{
			Name:  "瞄准",
			Count: 5,
		}
		g <- 操作{
			Name:  "发射！",
			Count: 3,
		}
		fmt.Println("装弹->瞄准->发射！")

	} else if g == nil && f != nil {
		_ = <-f
		g <- 操作{
			Name:  "发射！",
			Count: 3,
		}
		fmt.Println("装弹->瞄准->发射！")
	} else if g != nil && f == nil {
		_ = <-g
		f <- 操作{
			Name:  "瞄准",
			Count: 5,
		}
		fmt.Println("装弹->瞄准->发射！")
	} else if g != nil && f != nil {
		_ = <-g
		_ = <-f
	}
	fmt.Println("装弹->瞄准->发射！")
}

func main() {
	for {
		go 意大利炮()
	}
}

type 操作 struct {
	Name  string
	Count int
}
