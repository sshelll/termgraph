package component

import "github.com/gdamore/tcell/v2"

type Component interface {
	SetName(string)
	GetName() *string
	SetLayout(x1, y1, x2, y2 int)
	Draw(tcell.Screen, tcell.Style) error
}

type baseComponent struct {
	Name           *string
	x1, y1, x2, y2 int
}

func (c *baseComponent) SetName(name string) {
	c.Name = &name
}

func (c *baseComponent) GetName() *string {
	return c.Name
}

func (c *baseComponent) SetLayout(x1, y1, x2, y2 int) {
	c.x1 = x1
	c.x2 = x2
	c.y1 = y1
	c.y2 = y2
}
