package component

import "github.com/gdamore/tcell/v2"

type Component interface {
	SetName(string)
	SetLayout(x1, y1, x2, y2 int)
	Draw(s tcell.Screen, style tcell.Style, drawName bool) error
	// TODO: how to update component properly if I wanna design a live chart component?
}

type baseComponent struct {
	name           *string
	x1, y1, x2, y2 int
}

func (c *baseComponent) SetName(name string) {
	c.name = &name
}

func (c *baseComponent) SetLayout(x1, y1, x2, y2 int) {
	c.x1 = x1
	c.x2 = x2
	c.y1 = y1
	c.y2 = y2
}
