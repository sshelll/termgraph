package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/draw/graph"
)

type Box struct {
	baseComponent
}

func NewBox(name *string) *Box {
	return &Box{
		baseComponent{
			name: name,
		},
	}
}

func (b *Box) Draw(s tcell.Screen, style tcell.Style, drawName bool) error {
	graph.DrawBox(s, b.x1, b.y1, b.x2, b.y2, style, nil)
	return nil
}
