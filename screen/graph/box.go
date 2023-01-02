package graph

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/screen/content"
	"github.com/sshelll/termgraph/util"
)

func DrawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style) {

	// param check
	x1, x2 = util.SwapByOrder(x1, x2)
	y1, y2 = util.SwapByOrder(y1, y2)

	// fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// draw horizonal borders
	content.FillContentH(s, x1, x2, y1, tcell.RuneHLine, style)
	content.FillContentH(s, x1, x2, y2, tcell.RuneHLine, style)

	// draw vertial borders
	// NOTE: set start y coordinate as 'y1+1' to remain the corner cell.
	content.FillContentV(s, y1+1, y2, x1, tcell.RuneVLine, style)
	content.FillContentV(s, y1+1, y2, x2, tcell.RuneVLine, style)

	// draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

}
