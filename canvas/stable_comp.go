package canvas

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

type stableComponent struct {
	baseCanvasComponent
	pos *Position
}

func (c *stableComponent) Position() *Position {
	return c.pos
}

func (c *stableComponent) Draw(screen tcell.Screen) error {
	return c.comp.Draw(screen, c.style)
}

func (c *stableComponent) setLayout() {
	pos := c.pos
	if pos == nil {
		return
	}
	c.comp.SetLayout(pos.x1, pos.y1, pos.x2, pos.y2)
}

func newStableComponent(comp component.Component, pos *Position, style tcell.Style) canvasComponent {
	ccomp := &stableComponent{
		pos: pos,
		baseCanvasComponent: baseCanvasComponent{
			comp:  comp,
			style: style,
		},
	}
	// set layout immediately
	ccomp.setLayout()
	return ccomp
}
