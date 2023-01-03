package canvas

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

type stableCanvas struct {
	screen tcell.Screen

	style *tcell.Style
	pos   *Position

	components []canvasComponent
}

func NewStableCanvas() (canvas Canvas, err error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return
	}
	if err = screen.Init(); err != nil {
		return
	}
	canvas = &stableCanvas{
		screen: screen,
		style:  &tcell.StyleDefault,
	}
	return
}

func (c *stableCanvas) Show() error {
	if err := c.drawComponents(); err != nil {
		return err
	}
	c.startCtrl()
	return nil
}

func (c *stableCanvas) AddComponent(comp component.Component, pos *Position, style *tcell.Style) {
	if pos == nil {
		panic("pos should not be nil")
	}
	if style == nil {
		style = c.style
	}
	c.components = append(c.components, newStableComponent(comp, pos, *style))
}

func (c *stableCanvas) SetLayout(width int, height int) {
	if c.pos != nil {
		c.pos.Set(0, height, width, height)
		return
	}
	c.pos = NewPosition(0, height, width, height)
}

func (c *stableCanvas) SetStyle(style tcell.Style) {
	c.style = &style
}

func (c *stableCanvas) startCtrl() {

	defer c.screen.Fini()
	c.screen.Show()

loop:
	for {
		switch event := c.screen.PollEvent().(type) {
		case *tcell.EventKey:
			// TODO: @shaojiale maybe allow user to customize some key mappings.
			key := event.Key()
			if key == tcell.KeyESC {
				break loop
			}
		}
	}

	log.Println("exit...")

}

func (c *stableCanvas) drawComponents() error {
	for _, comp := range c.components {
		if err := comp.Draw(c.screen); err != nil {
			return err
		}
	}
	return nil
}
