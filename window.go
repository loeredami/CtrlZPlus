package main

type Window interface {
	run()
}

type CtrlZWindow struct {
	front *Frontend
	back  *Backend
}

func (window *CtrlZWindow) run() {
	(*window.front).run()
}