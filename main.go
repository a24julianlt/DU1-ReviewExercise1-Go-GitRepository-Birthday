package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Introduce el número de casos de prueba: ")
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Println("Introduce la edad en segundos: ")
		var edad int
		fmt.Scan(&edad)

		binario := strconv.FormatInt(int64(edad), 2)

		var unos int
		for _, bit := range binario {
			if bit == '1' {
				unos++
			}
		}
		fmt.Println(unos)
	}
}
