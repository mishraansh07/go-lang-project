package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sum int64

	fmt.Println("Hello, Welcome to main program")

	var arr_1 [5]int

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 5; i++ {
		fmt.Printf("Input the digits of array %d: ", i)
		fmt.Fscan(reader, &arr_1[i])
	}

	fmt.Println("The new Array is:", arr_1)

	for i := 0; i < 5; i++ {
		sum += int64(arr_1[i])
	}
	fmt.Println("Sum:", sum)

	avg := float64(sum) / 5

	fmt.Println("The Average of the array inputs are:\n", avg)
}
