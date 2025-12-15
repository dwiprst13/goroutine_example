package main

import (
	"fmt"
)
	
type Number struct {
	numbers chan int
}

var angka = []int{5,7,2,9,4,6,3,8,1,0}
var taskChan = make(chan int)
var resultChan = make(chan string)

func numberProcessor(id int, taskChan <-chan int) {
	for {
		num, ok := <-taskChan
		if !ok {
			return
		}
		var hasil int
		if num % 2  == 0 {
			hasil = num *2
		} else {
			hasil = num * num
		}
		resultChan <- fmt.Sprintf("Goroutine %d memproses angka %d menjadi %d\n", id, num, hasil)
	}
}

func main() {
	go numberProcessor(1, taskChan)
	go numberProcessor(2, taskChan)
	go numberProcessor(3, taskChan)

	go func() {
		for _, n := range angka {
			taskChan <- n
		}
		close(taskChan)
	}()

	for i := 0; i < len(angka); i++ {
		fmt.Print(<-resultChan)
	}
}