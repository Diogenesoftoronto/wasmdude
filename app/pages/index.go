package pages

import (
	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

type Index struct {
	vgrouter.NavigatorRef

	Body vugu.Builder
}

func (c *Index) HandleClick(event vugu.DOMEvent) {

}
