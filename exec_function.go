package main

import "fmt"

func main() {
	fmt.Println("Hello everybody!")
	foo()
	print_even_numbers()
	fmt.Println("this is after func foo", 42, "hello everybody")

}

func foo() {
	n,err := fmt.Println("this is func foo")
	fmt.Println(n)
	fmt.Println(err)
}

func print_even_numbers() {
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}
}
