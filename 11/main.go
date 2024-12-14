package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	stoneStrings := strings.Split(string(data), " ")
	stones := make(map[int]int)
	for _, s := range stoneStrings {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		stones[v]++
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}
	var count int
	for _, c := range stones {
		count += c
	}
	fmt.Println(count)

	for i := 0; i < 50; i++ {
		stones = blink(stones)
	}
	count = 0
	for _, c := range stones {
		count += c
	}
	fmt.Println(count)
}

func blink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)
	for stone, count := range stones {
		for _, newStone := range replaceStone(stone) {
			newStones[newStone] += count
		}
	}
	return newStones
}

func replaceStone(v int) []int {
	if v == 0 {
		return []int{1}
	}
	l := digitLen(v)
	if l%2 == 1 {
		return []int{v * 2024}
	}
	p := 1
	for i := 0; i < l/2; i++ {
		p *= 10
	}
	return []int{v / p, v % p}
}

func digitLen(n int) int {
	i := 10
	l := 1
	for n >= i {
		i *= 10
		l += 1
	}
	return l
}
