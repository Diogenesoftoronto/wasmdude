package components

import (
	"github.com/Diogenesoftoronto/wasmdude/app/state"
	"github.com/vugu/vugu"
)

type Root struct {
	Hero     vugu.Builder // below header but above Body
	Body     vugu.Builder // main body content
	Sidebar  vugu.Builder // sidebar to right of Body
	FullBody vugu.Builder // body section without any wrapping

	state.PageInfoRef
	vugu.Builder
	AutoReload bool // set to true during dev
}
