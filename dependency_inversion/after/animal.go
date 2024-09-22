package before

type animal struct {
	r runable
}

func newAnimal(r runable) animal {
	return animal{r: r}
}

type runable interface {
	run()
}

var _ runable = cat{}

type cat struct{}

// run implements runable.
func (c cat) run() {
}

func test() {
	c := cat{}
	a := newAnimal(c)
	a.r.run()
}
