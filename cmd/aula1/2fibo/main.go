package main

import "fmt"

func Fibo(n uint) uint {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fibo(n-1) + Fibo(n-2)
	}
}

func main() {
	fmt.Println("Olá gopher!")
	i := 0
	for i = 0; i <= 10; i++ {
		fmt.Println("Fibo(", i, ") =", Fibo(uint(i)))
	}
}
