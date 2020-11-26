package main

import "fmt"

const buffSize = 1000

func main() {
	store := newStore()
	transport := newTransport(buffSize)

	go newWorker(store, transport).Work()
	go newWorker(store, transport).Work()
	go newWorker(store, transport).Work()
	go newWorker(store, transport).Work()
	go newWorker(store, transport).Work()
	go newWorker(store, transport).Work()

	task := transport.Put(newTask(1000))

	fmt.Println("fibo(1000):", task.Result())
}
