package canvas

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

type Canvas interface {
	Show() error
	AddComponent(component.Component, *Position, *tcell.Style)
	SetLayout(width, height int)
	SetStyle(style tcell.Style)
}
