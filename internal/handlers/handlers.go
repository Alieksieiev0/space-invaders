package handlers

import "github.com/veandco/go-sdl2/sdl"

type KeyboardEventHandler interface {
	HandleKeyboardEvent(event *sdl.KeyboardEvent)
}

type AfterEventHandler interface {
	HandleAfterEvent()
}

type Handler interface {
	KeyboardEventHandler
	AfterEventHandler
}
