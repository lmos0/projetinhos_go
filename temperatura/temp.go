package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {
	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32) * 5 / 9 //conversão de F para C
	fmt.Println("A temperatura em F é igual ", tempF)
	fmt.Println("A temperatura em C é igual ", tempC)

}
