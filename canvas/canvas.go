package canvas

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

type Canvas interface {
	Draw() error
	AddComponent(component.Component, *Position, *tcell.Style)
	SetSize(width, height int)
	SetStyle(style tcell.Style)
}
