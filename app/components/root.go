package components

import (
	"github.com/Diogenesoftoronto/wasmdude/app/state"
	"github.com/vugu/vugu"
)

type Root struct {
	Body     vugu.Builder // main body content
	FullBody vugu.Builder // body section without any wrapping

	state.PageInfoRef
	vugu.Builder
	AutoReload bool // set to true during dev
}
