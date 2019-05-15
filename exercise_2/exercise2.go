package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
* Hacer un programa que reciba un parámetro p por consola y que calcule la
* suma de los valores v tales que 0 <= v <= p , siendo v un número divisible por 3 o por 5
 */
func main() {
	p, _ := strconv.Atoi(os.Args[1])
	fmt.Println(sumNumberDivisibleByThreeOrFive(p))
}

func sumNumberDivisibleByThreeOrFive(number int) int {
	var result int
	for i := 0; i <= number; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}
	return result
}
