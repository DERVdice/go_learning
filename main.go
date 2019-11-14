package main

import (
	"fmt"
	"io/ioutil"
	//"math"
	"os"
	"strings"
	"time"
)

// Нахождение квадратного корня числа методом последовательного приближения Ньютона
func NewtonSqrt(x float64) float64 {
	result := x
	resultLast := 0.0
	for {
		resultLast = result
		result = result - (result*result-x)/(2*result)
		if resultLast == result || resultLast-result < 0.0000000000000000000000000000000000000000000000000000000000000001 {
			return result
		}
	}
}

// Возвращает map с количеством повторений каждого слова строки s
func WordCount(s string) map[string]int {
	result := make(map[string]int)

	for _, v := range strings.Fields(s) {
		val, ok := result[v]
		if ok {
			result[v] = val + 1
		} else {
			result[v] = 1
		}

		fmt.Println(v)
	}
	return result
}

func Clock() func() int {
	x := -1
	return func() int {
		if x > 23 {
			x = 0
		} else {
			x++
		}
		return x
	}
}

// минимальное число массива
func MinArrayValue(array []int) int {
	min := 9223372036854775807;
	for i := 0; i < len(array); i++ {
		if array[i] <= min {
			min = array[i]
		}
	}
	return min
}

// максимальное число массива
func MaxArrayValue(array []int) int {
	max := -9223372036854775808;
	for i := 0; i < len(array); i++ {
		if array[i] >= max {
			max = array[i]
		}
	}
	return max
}

// n-ое число Фибоначчи
func Fib(n int) int {
	if n > 1 {
		return Fib(n-1) + Fib(n-2)
	} else {
		switch n {
		case 0:
			return 0
		case 1:
			return 1
		}
	}

	return -1
}

// Меняет местами значения x и y
func SwapValues(x *int, y *int) {
	temp := new(int)
	*temp = *x
	*x = *y
	*y = *temp
}

func faa() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		amt := time.Duration(100)
		time.Sleep(time.Millisecond * amt)
	}
}

func Chantest1(a int, ch chan int) {
	a = a * 2
	ch <- a
}

func Chantest2(b int, ch chan int) {
	b = b * 3
	ch <- b
}

func dirTree(path string, space string) {
	if space == "└──" {
		space = space + " "
	} else {
		space = "     " + space
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, f := range files {
		if f.IsDir() {
			fmt.Printf("%v%v\n", space, f.Name())
			dirTree(path+"/"+f.Name(), space)
		} else {
			if f.Size() != 0 {
				fmt.Printf("%v%v (%vb)\n", space, f.Name(), f.Size())
			} else {
				fmt.Printf("%v%v (%v)\n", space, f.Name(), "empty")
			}
		}
	}
}

func backgroundWorkerOne()  {
	for{
		select {

		}
	}
}

func backgroundWorkerTwo()  {

}

func main() {
/*
	//	1
	x := float64(256)
	fmt.Println("x = ", x)
	fmt.Println("NewtonSqrt : ", NewtonSqrt(x))
	fmt.Println("math.Sqrt : ", math.Sqrt(x))
	fmt.Println()

	//	2
	s := "a b b c c c d d d d "
	fmt.Println("WordCounter : ", WordCount(s))
	fmt.Println()

	//	3
	f := Clock()
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
	fmt.Println()

	//	4
	z := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(MinArrayValue(z))
	fmt.Println(MaxArrayValue(z))
	fmt.Println()

	//	5
	fmt.Println(Fib(12))
	fmt.Println()

	//	6
	e := 1
	y := 2
	fmt.Println(e, "	", y)
	SwapValues(&e, &y)
	fmt.Println(e, "	", y)
	fmt.Println()

	//go faa()

	ch := make(chan int, 10)
	go Chantest2(1, ch)
	go Chantest1(1, ch)
	a, b := <-ch, <-ch
	println(a, b)

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dirTree(pwd, "└──")
	*/
	crc32InUint32 := crc32.ChecksumIEEE([]byte("Checksum returns uint32"))
	crc32InString := strconv.FormatUint(uint64(crc32InUint32), 16)
	fmt.Println(crc32InString)
}
