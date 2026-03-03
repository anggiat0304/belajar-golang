package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	for tick := range ticker.C {
		fmt.Println(tick)
	}
	ticker.Stop()
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	for tick := range channel {
		fmt.Println(tick)
	}
}
