package after

type animal interface {
	eat()
	sleep()
}

type swimable interface {
	swim()
}

type climbable interface {
	climb()
}

type runable interface {
	run()
}
