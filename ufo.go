package main

import (
	"fmt"
)

func main() {
	var Team1 string
	var UFO string
	fmt.Scanf("%s%s", &Team1, &UFO)
	var T int = 1
	for i := 0; i < len(Team1); i++ {
		T = T * (int(Team1[i]) - 64)
	}
	T1 := T % 47
	var U int = 1
	for j := 0; j < len(UFO); j++ {
		U = U * (int(UFO[j]) - 64)
	}
	U1 := U % 47
	if T1 == U1 {
		fmt.Println("GO")
	}
	if T1 != U1 {
		fmt.Println("STAY")
	}

}
