package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	/**
	100000
	2023/08/03 11:06:41 Total sum of square 333338333350000
	2023/08/03 11:06:41 Thời gian xử lý 1 ::: 3.138ms

	*/
	startTime := time.Now()
	randomNumbers := createArrayNumber(100)

	sum1 := 0
	for i := 0; i < len(randomNumbers); i++ {
		time.Sleep(100)
		sum1 += randomNumbers[i] * randomNumbers[i]
	}

	//numberChan := readNumberChan(randomNumbers)
	//c1 := square(numberChan)
	//c2 := square(numberChan)
	//c3 := square(numberChan)
	//c4 := square(numberChan)
	//c5 := square(numberChan)
	//c6 := square(numberChan)
	//c7 := square(numberChan)
	//c8 := square(numberChan)
	//c9 := square(numberChan)
	//c10 := square(numberChan)
	//for c := range merge(c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) {
	//	sum1 += c
	//	//log.Println("c1 ::: ", c)
	//}
	//wg := sync.WaitGroup{}
	//lock := sync.RWMutex{}
	//wg.Add(3)
	//go func() {
	//	for c := range merge(c1) {
	//		lock.Lock()
	//		sum1 += c
	//		lock.Unlock()
	//		//log.Println("c1 ::: ", c)
	//	}
	//	wg.Done()
	//}()
	//go func() {
	//	for c := range merge(c2) {
	//		lock.Lock()
	//		sum1 += c
	//		lock.Unlock()
	//		//log.Println("c2 ::: ", c)
	//	}
	//	wg.Done()
	//}()
	//go func() {
	//	for c := range merge(c3) {
	//		lock.Lock()
	//		sum1 += c
	//		lock.Unlock()
	//		//log.Println("c3 ::: ", c)
	//	}
	//	wg.Done()
	//}()
	//wg.Wait()
	log.Printf("Total sum of square %d", sum1)
	endTime := time.Since(startTime)
	log.Printf("Thời gian xử lý 1 ::: %s\n", endTime)
}

func readNumberChan(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, number := range numbers {
			out <- number
		}
		close(out)
	}()
	return out
}

func createArrayNumber(number int) (randomNumbers []int) {
	for i := 1; i <= number; i++ {
		randomNumbers = append(randomNumbers, i)

	}
	return
}
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			time.Sleep(100)
			out <- num * num
		}
	}()
	return out
}
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(ch <-chan int) {
			defer wg.Done()
			for n := range ch {
				out <- n
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
