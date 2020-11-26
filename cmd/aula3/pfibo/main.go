package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

const buffSize = 100
const maxN = 10_000
const entries = 50_000

func main() {
	store := newStore()
	transport := newTransport(buffSize)

	for i := 0; i < runtime.NumCPU(); i++ {
		go newWorker(store, transport).Work()
	}

	tasks := make(chan *task, buffSize)
	wg := sync.WaitGroup{}

	wg.Add(2)
	defer wg.Wait()

	go func() {
		defer wg.Done()
		for i := 0; i < entries; i++ {
			tasks <- transport.Put(newTask(rand.Uint64() % maxN))
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < entries; i++ {
			task := <-tasks
			fmt.Println("fibo(", task.n, "):", task.Result())
		}
	}()
}
