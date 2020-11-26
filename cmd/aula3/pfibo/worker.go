package main

type worker struct {
	store     *store
	transport *transport
}

func newWorker(store *store, transport *transport) *worker {
	return &worker{
		store:     store,
		transport: transport,
	}
}

func (w *worker) Work() {
	for {
		task := w.transport.Get()
		task.value = w.Fibo(task.n)
		task.Done()
	}
}

func (w *worker) Fibo(n uint64) uint64 {
	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	case w.store.Has(n):
		return w.store.Get(n)
	default:
		n1 := w.Fibo(n - 1)
		n2 := w.Fibo(n - 2)

		w.store.Set(n, n1+n2)

		return n1 + n2
	}
}
