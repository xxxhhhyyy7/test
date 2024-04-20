package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var num [100]int

func main() {
	var group sync.WaitGroup
	group.Add(2)
	go fill(num[:50], &group)
	go fill(num[50:], &group)
	group.Wait()
	fmt.Println(num)
	ch1 := make(chan int)
	go sum(num[:33], ch1)
	ch2 := make(chan int)
	go sum(num[33:66], ch2)
	ch3 := make(chan int)
	go sum(num[66:], ch3)
	sum1 := <-ch1
	sum2 := <-ch2
	sum3 := <-ch3
	res := (sum1 + sum2 + sum3) / len(num)
	fmt.Println(res)
}

func fill(num []int, group *sync.WaitGroup) {
	defer group.Done()
	for i := range num {
		num[i] = rand.Intn(101)
	}
}
func sum(num []int, result chan<- int) {
	var sum int
	for _, n := range num {
		sum += n
	}
	result <- sum
}
