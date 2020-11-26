package main

type task struct {
	n, value uint64
	done     chan bool
}

func newTask(n uint64) *task {
	return &task{
		n:     n,
		value: 0,
		done:  make(chan bool, 1),
	}
}

func (t *task) Done() {
	t.done <- true
}

func (t *task) Result() uint64 {
	<-t.done
	return t.value
}
