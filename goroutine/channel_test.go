package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Anggiat Pangaribuan"
		fmt.Println("Selesai kirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 & time.Second)
	channel <- "My name is Anggiat Pangaribuan"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// Mengirim
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "My name is Anggiat Pangaribuan. I hope we can friend"
}
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)
	time.Sleep(3 * time.Second)
}
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 4)
	defer close(channel)

	go func() {
		channel <- "Buku"
		channel <- "Beras"
		channel <- "Babi"
		channel <- "Bunga"
	}()
	go func() {
		for data := range channel {
			fmt.Println(data)
		}
	}()
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel2)
	go GiveMeResponse(channel1)
	counter := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		}

		if counter == 2 {
			break
		}

	}
	fmt.Println("Selesai")
}
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel2)
	go GiveMeResponse(channel1)
	counter := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}

		if counter == 2 {
			break
		}

	}
	fmt.Println("Selesai")
}
