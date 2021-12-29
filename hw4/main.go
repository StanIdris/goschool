package main

import (
	"fmt"
	"sort"
)

func intScanln(n int) ([]int, error) {
	fmt.Println("Введите 6 чисел через пробле для примера 5 4 1 3 6 2")
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	sort.Ints(x)
	return x, err
}

func main() {
	x, err := intScanln(6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("После сортировки введенные вами данные выглядят так: %v\n", x)
}
