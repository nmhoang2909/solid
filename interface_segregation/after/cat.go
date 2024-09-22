package after

type cat struct{}

// climb implements climbable.
func (c cat) climb() {
}

var _ climbable = cat{}
