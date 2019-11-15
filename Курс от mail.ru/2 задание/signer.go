package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func SingleHash(data string, Out chan string, quota chan struct{}) {
	start := time.Now()

	quota <- struct{}{}
	Md5 := DataSignerMd5(data)
	<-quota

	Crc32 := DataSignerCrc32(data)
	Crc32FromMd5 := DataSignerCrc32(Md5)
	result := Crc32 + "~" + Crc32FromMd5

	Out <- result
	fmt.Println("SingleHash", time.Since(start))

}

/*
func MultiHash(SingleHashOut chan string, Out chan string) {
	start := time.Now()
	SingleHashResult := <-SingleHashOut
	ArrayForMultiHash := [6]string{"0", "1", "2", "3", "4", "5"}
	result := ""
	for _, v := range ArrayForMultiHash {
		result += DataSignerCrc32(v + SingleHashResult)
	}
	Out <- result
	fmt.Println("MultiHash", time.Since(start))
}
*/

func MultiHash(SingleHashOut chan string, Out chan string) {
	start := time.Now()
	SingleHashResult := <-SingleHashOut

	MultiHashResultsChannel := make(chan string, 6)
	ArrayForMultiHash := [6]string{"0", "1", "2", "3", "4", "5"}
	for i := 0; i < 6; i++ {
		go MultiHashSingleStep(SingleHashResult, ArrayForMultiHash[i], MultiHashResultsChannel)
	}

	var result string
	for i := 0; i < 6; i++ {
		result += <-MultiHashResultsChannel
	}

	Out <- result
	fmt.Println("MultiHash", time.Since(start))
}

func MultiHashSingleStep(data string, v string, Out chan string) {
	Out <- DataSignerCrc32(v + data)
}

func worker(index int, data string, wg *sync.WaitGroup, quota chan struct{}) {
	defer wg.Done()
	Out1 := make(chan string)
	Out2 := make(chan string)

	go SingleHash(data, Out1, quota)
	runtime.Gosched()
	go MultiHash(Out1, Out2)
	result := <-Out2
	fmt.Printf("Worker %v finished with result: %v\n", index, result)
}

func ExecutePipeline(data []string) {
	wg := &sync.WaitGroup{}
	quota := make(chan struct{}, 1)
	start := time.Now()
	for i, v := range data {
		wg.Add(1)
		go worker(i, v, wg, quota)
	}
	wg.Wait()

	fmt.Println(time.Since(start))
}

func main() {
	data := []string{"0", "1"}
	ExecutePipeline(data)
}
