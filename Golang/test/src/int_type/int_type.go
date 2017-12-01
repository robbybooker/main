package int_type

type IntType int

func (this *IntType) Reset() {
	*this = 0;
}

func (this *IntType) Increment() {
	*this++;
}