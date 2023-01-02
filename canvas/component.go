package canvas

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

type canvasComponent interface {
	Position() *Position
	Component() component.Component
	Draw(tcell.Screen) error
}

type stableCanvasComponent struct {
	pos   *Position
	style tcell.Style
	comp  component.Component
}

func newStableCanvasComponent(comp component.Component, pos Position, style tcell.Style) *stableCanvasComponent {
	comp.SetSize(pos.x1, pos.y1, pos.x2, pos.y2)
	return &stableCanvasComponent{
		comp:  comp,
		pos:   &pos,
		style: style,
	}
}

func (c *stableCanvasComponent) Position() *Position {
	return c.pos
}

func (c *stableCanvasComponent) Component() component.Component {
	return c.Component()
}

func (c *stableCanvasComponent) Draw(s tcell.Screen) error {
	return c.comp.Draw(s, c.style)
}
