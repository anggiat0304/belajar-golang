package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	var x = 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter :", x)
}

type BackAccount struct {
	RwmMutex sync.RWMutex
	Balance  int
}

func (account *BackAccount) AddBalance(amount int) {
	account.RwmMutex.Lock()
	account.Balance = account.Balance + amount
	account.RwmMutex.Unlock()
}
func (account *BackAccount) GetBalance() int {
	account.RwmMutex.RLock()
	balance := account.Balance
	account.RwmMutex.RUnlock()
	return balance
}
func TestReadWriteMutex(t *testing.T) {
	account := BackAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("maka total uangnya adalah :", account.GetBalance())
}
