package main

import "fmt"

func main() {
	var fibo func(n uint) uint
	fibo = func(n uint) uint {
		switch n {
		case 0:
			return 0
		case 1:
			return 1
		default:
			return fibo(n-1) + fibo(n-2)
		}
	}

	fmt.Println("Olá gopher!")
	i := 0
	for i = 0; i <= 10; i++ {
		fmt.Println("Fibo(", i, ") =", fibo(uint(i)))
	}
}
