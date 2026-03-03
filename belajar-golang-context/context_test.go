package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextC, "g", "G")
	contextH := context.WithValue(contextG, "h", "H")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)
	fmt.Println(contextH)

	fmt.Println(contextA.Value("a"))
	fmt.Println(contextB.Value("b"))
	fmt.Println(contextC.Value("c"))
	fmt.Println(contextD.Value("d"))
	fmt.Println(contextE.Value("e"))
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextG.Value("g"))
	fmt.Println(contextH.Value("c"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return destination
}
func TestWithCancel(t *testing.T) {
	fmt.Println("Total goroutine :", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	destination := CreateCounter(ctx)

	fmt.Println("Total goroutine :", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter ", n)
		if n == 10 {
			break
		}
	}
	cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("Total goroutine :", runtime.NumGoroutine())

}
func TestWithTimeout(t *testing.T) {
	fmt.Println("Total goroutine :", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total goroutine :", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter ", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total goroutine :", runtime.NumGoroutine())

}
func TestWithDeadline(t *testing.T) {
	fmt.Println("Total goroutine :", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)

	fmt.Println("Total goroutine :", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter ", n)
		if n == 10 {
			break
		}
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total goroutine :", runtime.NumGoroutine())

}
