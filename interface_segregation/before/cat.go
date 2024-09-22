package before

var _ animal = cat{}

type cat struct{}

// climb implements animal.
func (c cat) climb() {
}

// eat implements animal.
func (c cat) eat() {
}

// run implements animal.
func (c cat) run() {
}

// sleep implements animal.
func (c cat) sleep() {
}

// swim implements animal.
func (c cat) swim() {
	panic("cat cannot swim")
}
