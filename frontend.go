package main

import "cogentcore.org/core/core"

type Frontend interface {
	makeApp() *core.Body
	run()

	linkBackend(backend *Backend)
}

type CtrlZFront struct {
	body *core.Body
	backend *Backend
}

func makeCtrlZFront(back *Backend) Frontend {
	front := &(CtrlZFront{nil, nil})
	front.makeApp()
	front.linkBackend(back)
	return front
}

func (front *CtrlZFront) run() {
	front.body.RunMainWindow()
}

func (front *CtrlZFront) makeApp() *core.Body {
	front.body = core.NewBody()
	return front.body
}

func (front *CtrlZFront) linkBackend(backend *Backend) {
	front.backend = backend
}