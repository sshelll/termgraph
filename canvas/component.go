package canvas

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

type canvasComponent interface {
	Position() *Position
	Component() component.Component
	SetStyle(style tcell.Style)
	Draw(tcell.Screen) error
}

type baseCanvasComponent struct {
	style tcell.Style
	comp  component.Component
}

func (b *baseCanvasComponent) Component() component.Component {
	return b.comp
}

func (b *baseCanvasComponent) SetStyle(style tcell.Style) {
	b.style = style
}
