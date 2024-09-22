package after

var _ swimable = fish{}
var _ animal = fish{}

type fish struct{}

// eat implements animal.
func (f fish) eat() {
}

// sleep implements animal.
func (f fish) sleep() {
}

// swim implements swimable.
func (f fish) swim() {
}
