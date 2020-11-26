package main

type transport struct {
	channel chan *task
}

func newTransport(memory uint64) *transport {
	return &transport{channel: make(chan *task, memory)}
}

func (t *transport) Put(task *task) *task {
	t.channel <- task

	return task
}

func (t *transport) Get() *task {
	return <-t.channel
}
