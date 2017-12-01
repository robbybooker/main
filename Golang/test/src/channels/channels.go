package channels

type channels struct {

}

func New() *channels {
	return new(channels);
}

func (this *channels) Moo(channel chan int) {
	channel <- 100;
}
