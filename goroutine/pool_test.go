package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "new"
		},
	}
	// var group sync.WaitGroup
	pool.Put("Anggiat")
	pool.Put("Pangaribuan")
	pool.Put("Yang")
	pool.Put("Ganteng")

	for i := 0; i < 10; i++ {
		// group.Add(1)
		go func() {
			// defer group.Done()
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	// group.Wait()
	time.Sleep(11 * time.Second)
	fmt.Println("Done")
}
