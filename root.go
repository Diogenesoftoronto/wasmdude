package main

import (
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

type Root struct {
	vgrouter.NavigatorRef

	Body vugu.Builder
}

func (c *Root) HandleClick(event vugu.DOMEvent) {

}
