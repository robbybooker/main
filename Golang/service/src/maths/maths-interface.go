package maths

type MathsInterface interface {
	IncrementNumber(*int)
	IncrementCounter()
	GetCounter() int
	Error()
}
