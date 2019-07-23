package main

import (
	"fmt"
)

func main() {

        test := []int{9,3,9,3,9,7,9}
	//fmt.Println(len(test))
	fmt.Println(Solution(test))
}

func Solution(A []int)(oddOne int) {
    var histogram = make(map[int]int)

	for _, value := range A {
		histogram[value]++;
		fmt.Printf("%v,%v\n",value,histogram[value])
	}

	for k, value := range histogram {
		if isOdd(value) {
			oddOne = k
			return;
		}
	}

	return -1
}

func isOdd(n int) bool {
	return n % 2 == 1
}
