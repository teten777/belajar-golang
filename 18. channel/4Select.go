package main

import (
	"fmt"
	"runtime"
)

func getAvg(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}

	}
	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{2, 1, 3, 5, 3, 2, 5, 7, 66, 3, 33, 6, 6, 4}
	fmt.Println("numbers ", numbers)

	var ch1 = make(chan float64)
	go getAvg(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg\t: %.2f\n", avg)
		case max := <-ch2:
			fmt.Printf("Max\t:%d\n", max)
		}
	}
}
