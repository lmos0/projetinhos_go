package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := multiplca(10, 10)
	z := subtrai(12, 2, 2)
	a := dividi(4, 2)
	fmt.Println(x, y, z, a)

}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}

	return total

}

func multiplca(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total

}

func subtrai(i ...int) int {
	total := 0
	for _, v := range i[1:] {
		total -= v
	}

	return total

}

func dividi(i ...int) int {
	total := i[0]
	for _, v := range i[1:] {
		total /= v
	}
	return total

}
