package gofocused

import (
	"fmt"
	"runtime"
)

func ChanSelect() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("Numbers:", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}

func getAverage(num []int, ch chan float64) {
	var sum = 0
	for _, e := range num {
		sum += e
	}
	ch <- float64(sum) / float64(len(num))
}

func getMax(num []int, ch chan int) {
	var max = num[0]
	for _, e := range num {
		if max < e {
			max = e
		}
	}
	ch <- max
}
