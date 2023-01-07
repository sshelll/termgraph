package canvas

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

type Canvas interface {
	Show() error
	AddComponent(comp component.Component, id string, pos *Position, style *tcell.Style)
	GetComponent(id string) (canvasComponent, bool)
	RemoveComponent(id string)
	SetLayout(width, height int)
	SetStyle(style tcell.Style)
	Fini()
}

type baseCanvas struct {
	screen tcell.Screen

	style tcell.Style
	pos   *Position

	components map[string]canvasComponent
}

func newBaseCanvas() (baseCanvas, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return baseCanvas{}, err
	}
	return baseCanvas{
		screen:     screen,
		style:      tcell.StyleDefault,
		components: make(map[string]canvasComponent),
	}, nil
}

func (c *baseCanvas) AddComponent(comp component.Component, id string, pos *Position, style *tcell.Style) {
	if pos == nil {
		panic("pos should not be nil")
	}
	if style == nil {
		style = &c.style
	}
	if _, ok := c.components[id]; ok {
		panic("id duplicated")
	}
	c.components[id] = newStableComponent(comp, pos, *style)
}

func (c *baseCanvas) GetComponent(id string) (canvasComponent, bool) {
	comp, ok := c.components[id]
	return comp, ok
}

func (c *baseCanvas) RemoveComponent(id string) {
	delete(c.components, id)
}

func (c *baseCanvas) SetLayout(width int, height int) {
	if c.pos != nil {
		c.pos.Set(0, height, width, height)
		return
	}
	c.pos = NewPosition(0, height, width, height)
}

func (c *baseCanvas) SetStyle(style tcell.Style) {
	c.style = style
}

func (c *baseCanvas) Fini() {
	c.screen.Fini()
}
