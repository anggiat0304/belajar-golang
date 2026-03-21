package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock User 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock User 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Anggiat",
		Balance: 10,
	}
	user2 := UserBalance{
		Name:    "Jaka",
		Balance: 10,
	}

	go Transfer(&user1, &user2, 1)
	time.Sleep(3 * time.Second)
	go Transfer(&user2, &user1, 2)
	time.Sleep(3 * time.Second)
	fmt.Println("User  ", user1.Name, "have balance ", user1.Balance)
	fmt.Println("User  ", user2.Name, "have balance ", user2.Balance)

}

func RunAsync(group *sync.WaitGroup, i int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Bilangan ke : ", i)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 9999; i++ {
		go RunAsync(group, i)
	}
	group.Wait()

	fmt.Println("Complete")
}
