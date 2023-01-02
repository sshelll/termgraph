package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/screen/graph"
)

type Box struct {
	baseComponent
}

func NewBox(name *string) *Box {
	return &Box{
		baseComponent{
			Name: name,
		},
	}
}

func (b *Box) Draw(s tcell.Screen, style tcell.Style) error {
	graph.DrawBox(s, b.x1, b.y1, b.x2, b.y2, style)
	return nil
}
