package before

type animal struct {
	c cat
}

func newAnimal(c cat) animal {
	return animal{c: c}
}

type cat struct{}

func (c cat) run() {
}

func test() {
	c := cat{}
	a := newAnimal(c)
	a.c.run()
}
