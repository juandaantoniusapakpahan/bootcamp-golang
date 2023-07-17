package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	sync.Mutex
	Name    string
	Balance int
}

// Dealock
// mutex.lock akan saling menunggu
func Transfer(user1 *Account, user2 *Account, amount int) {

	user1.Mutex.Lock()
	user1.Balance = user1.Balance - amount
	fmt.Println("Pengirim:", user1.Name)

	user2.Mutex.Lock()
	user2.Balance = user2.Balance + amount
	fmt.Println("Penerima:", user2.Name)

	user1.Mutex.Unlock()
	user2.Mutex.Unlock()
}

func TransferSolutionWithWaitGroup(user1 *Account, user2 *Account, amount int, wg *sync.WaitGroup) {
	wg.Add(1)

	user1.Balance = user1.Balance - amount

	user2.Balance = user2.Balance + amount
	wg.Done()

}
func TransferSolutionWithWaitMutex(user1 *Account, user2 *Account, amount int, mx *sync.Mutex) {
	mx.Lock()
	user1.Balance = user1.Balance - amount

	user2.Balance = user2.Balance + amount
	mx.Unlock()
}

func main() {
	user1 := Account{
		Name:    "Juanda",
		Balance: 1000000,
	}

	user2 := Account{
		Name:    "Mely",
		Balance: 2000000,
	}

	// Deadlock
	// go Transfer(&user1, &user2, 100000)
	// go Transfer(&user2, &user1, 300000)

	// Solution with waithGroup
	// wg := new(sync.WaitGroup)
	// go TransferSolutionWithWaitGroup(&user1, &user2, 100000, wg)
	// go TransferSolutionWithWaitGroup(&user2, &user1, 400000, wg)
	// go TransferSolutionWithWaitGroup(&user1, &user2, 200000, wg)
	// go TransferSolutionWithWaitGroup(&user2, &user1, 300000, wg)
	// wg.Wait()

	// Solution with Mutex
	mx := new(sync.Mutex)

	go TransferSolutionWithWaitMutex(&user1, &user2, 100000, mx)
	go TransferSolutionWithWaitMutex(&user2, &user1, 500000, mx)
	go TransferSolutionWithWaitMutex(&user1, &user2, 100000, mx)
	go TransferSolutionWithWaitMutex(&user2, &user1, 200000, mx)
	go TransferSolutionWithWaitMutex(&user1, &user2, 100000, mx)
	go TransferSolutionWithWaitMutex(&user2, &user1, 300000, mx)

	time.Sleep(time.Second * 5)
	fmt.Println("Saldo ", user1.Name, ":", user1.Balance)
	fmt.Println("Saldo ", user2.Name, ":", user2.Balance)

}
