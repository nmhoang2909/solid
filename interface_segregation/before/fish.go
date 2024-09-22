package before

var _ animal = fish{}

type fish struct{}

// climb implements animal.
func (f fish) climb() {
	panic("fish cannot climb")
}

// eat implements animal.
func (f fish) eat() {
}

// run implements animal.
func (f fish) run() {
	panic("fish cannot run")
}

// sleep implements animal.
func (f fish) sleep() {
}

// swim implements animal.
func (f fish) swim() {
}
