package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(3 * time.Second)
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU :", totalCPU)
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine :", totalGoRoutine)
}
