package main

import "fmt"

func sum(sc []int, c chan int) {
	res := 0
	for _, val := range sc {
		res += val
	}
	c <- res
}

func main() {
	c := make(chan int)

	testArr := []int{12, 34, 5, 5, 1, 13, 5, 4, 123, 3, 321}
	go sum(testArr[:len(testArr)/2], c)
	go sum(testArr[len(testArr)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
