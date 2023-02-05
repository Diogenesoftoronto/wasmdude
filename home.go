package main

import (
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

type Home struct {
	vgrouter.NavigatorRef
	Body vugu.Builder
}
