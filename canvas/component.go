package canvas

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

type canvasComponent interface {
	ID() string
	Position() *Position
	Component() component.Component
	SetStyle(style tcell.Style)
	ToggleName(bool)
	Draw(tcell.Screen) error
}

type baseCanvasComponent struct {
	id       string
	style    tcell.Style
	comp     component.Component
	showName bool
}

func (b *baseCanvasComponent) Component() component.Component {
	return b.comp
}

func (b *baseCanvasComponent) SetStyle(style tcell.Style) {
	b.style = style
}

func (b *baseCanvasComponent) ToggleName(show bool) {
	b.showName = show
}

func (b *baseCanvasComponent) ID() string {
	return b.id
}
