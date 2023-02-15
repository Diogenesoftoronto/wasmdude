package components

import (
	"github.com/Diogenesoftoronto/wasmdude/app/state"
	"github.com/vugu/vgrouter"
)

type Navigation struct {
	vgrouter.NavigatorRef
	state.PageInfoRef
}
