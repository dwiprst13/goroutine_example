package main
// import "fmt"

var daftar_angka = []int{1,2,3,4,5,6,7,8,9,10}
var inputChan = make(chan int)
var evenChan = make(chan int)
var squareChan = make(chan int)

func filter(inputChan <-chan int, evenChan chan<- int) {
	for v := range inputChan {
		if v % 2 == 0 {
			evenChan <- v
		}
		close(evenChan)
	}
}

func transform(evenChan <-chan int, squareChan chan<- int) {
	for v := range evenChan {
		squareChan <- v * v
	}
	close(squareChan)
}



func main() {
	go filter(inputChan, evenChan)

	go func() {
		for _, v := range daftar_angka {
			inputChan <- v
		}
		close(inputChan)
	}()
}