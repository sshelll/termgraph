package canvas

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

type stableCanvas struct {
	screen tcell.Screen

	style *tcell.Style
	pos   *Position

	components []canvasComponent
}

func NewStableCanvas() (canvas *stableCanvas, err error) {
	canvas = &stableCanvas{}
	screen, err := tcell.NewScreen()
	if err != nil {
		return
	}
	if err = screen.Init(); err != nil {
		return
	}
	canvas.screen = screen
	canvas.style = &tcell.StyleDefault
	return
}

func (c *stableCanvas) Draw() error {
	for _, comp := range c.components {
		if err := comp.Draw(c.screen); err != nil {
			return err
		}
	}
	c.screen.Show()
loop:
	for {
		switch event := c.screen.PollEvent().(type) {
		case *tcell.EventKey:
			key := event.Key()
			if key == tcell.KeyESC {
				break loop
			}
		}
	}
	c.screen.Fini()
	return nil
}

func (c *stableCanvas) AddComponent(comp component.Component, pos *Position, style *tcell.Style) {
	if pos == nil {
		panic("pos should not be nil")
	}
	if style == nil {
		style = c.style
	}
	c.components = append(c.components, newStableCanvasComponent(comp, *pos, *style))
}

func (c *stableCanvas) SetSize(width int, height int) {
	if c.pos != nil {
		c.pos.Set(0, height, width, height)
		return
	}
	c.pos = NewPosition(0, height, width, height)
}

func (c *stableCanvas) SetStyle(style tcell.Style) {
	c.style = &style
}
