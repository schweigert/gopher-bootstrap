package services

// Fibo service
type Fibo struct{}

// NewFibo service constructor
func NewFibo() *Fibo {
	return &Fibo{}
}

// Fibo number
func (service *Fibo) Fibo(value uint) uint {
	switch value {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return service.Fibo(value-1) + service.Fibo(value-2)
	}
}
