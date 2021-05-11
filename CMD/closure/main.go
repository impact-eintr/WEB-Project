package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	x := 3
	f1 := foo1(&x)
	f2 := foo2(x)
	f1()
	f1()
	f2()
	f2()

	foo3()

	foo5()

	wg.Wait()

	f9 := foo9()
	f9()
	f9()
	f9()
}

func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Println("foo1 val =", *x)
	}

}

func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Println("foo2 val =", x)
	}
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fmt.Println("foo3 val =", val)
	}
}

func show(v interface{}) {
	fmt.Println("foo4 val =", v)
	wg.Done()
}

func foo4() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		wg.Add(1)
		go show(val)
	}
}

func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		wg.Add(1)
		go func() {
			fmt.Println("foo5 val =", val)
			wg.Done()
		}()
	}
}

var foo6Chan = make(chan int, 10)

func foo6() {
	for val := range foo6Chan {
		go func() {
			fmt.Printf("foo6 val = %d\n", val)

		}()

	}

}

func foo7(x int) []func() {
	var fs []func()
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		fs = append(fs, func() {
			fmt.Printf("foo7 val = %d\n", x+val)
		})

	}
	return fs

}

func foo8() func() {
	num := 10
	return func() {
		num++
		fmt.Println("foo8 val =", num)
	}
}

func foo9() func() {
	return foo8()
}
