package main

import (
	"fmt"
	"math/rand"
	"time"
)

var taskChan = make(chan Task)
var resultChan = make(chan string)

type Task struct {
	name string
	a, b int
}

func processTask(task string) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func worker(id int, resultChan chan<- string){
	for {
		task, ok := <-taskChan
		if !ok {
			return
		}
		switch task.name {
			case "add":
				processTask(task.name)
				resultChan <- fmt.Sprintf("Worker %d hasil penjumlahan: %d", id, task.a+task.b)
			case "subtract":
				processTask(task.name)
				resultChan <- fmt.Sprintf("Worker %d hasil pengurangan: %d", id, task.a-task.b)
			case "multiply":
				processTask(task.name)
				resultChan <- fmt.Sprintf("Worker %d hasil perkalian: %d", id, task.a*task.b)
			case "divide":
				processTask(task.name)
				if task.b != 0 {
					resultChan <- fmt.Sprintf("Worker %d hasil pembagian: %.2f", id, float64(task.a)/float64(task.b))
				} else {
					resultChan <- fmt.Sprintf("Worker %d pembagian dengan nol tidak diperbolehkan", id)
				}
			default:
				resultChan <- fmt.Sprintf("Worker %d tugas tidak dikenal: %s", id, task.name)
			}
	}
}


func main() {

	for i := 1; i <=4; i++ {
		go worker(i, resultChan)
	}
	taskChan <- Task{"add", 10, 5}
	taskChan <- Task{"subtract", 20, 8}
	taskChan <- Task{"multiply", 7, 6}
	taskChan <- Task{"divide", 15, 3}
	close(taskChan)

	for i := 0; i < 4; i++ {
		result := <-resultChan
		fmt.Println(result)
	}
}