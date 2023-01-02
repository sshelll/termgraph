package content

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/util"
)

func SetContentH(s tcell.Screen, x, y int, content string, style tcell.Style) {
	for _, c := range content {
		r, w, comb := util.ReplaceRuneWithComb(c)
		s.SetContent(x, y, r, comb, style)
		x += w
	}
}

func SetContentV(s tcell.Screen, x, y int, content string, style tcell.Style) {
	for _, c := range content {
		r, _, comb := util.ReplaceRuneWithComb(c)
		s.SetContent(x, y, r, comb, style)
		y++
	}
}

func FillContentH(s tcell.Screen, x1, x2, y int, r rune, style tcell.Style) {
	ch, width, comb := util.ReplaceRuneWithComb(r)
	for col := x1; col <= x2; col += width {
		s.SetContent(col, y, ch, comb, style)
	}
}

func FillContentV(s tcell.Screen, y1, y2, x int, r rune, style tcell.Style) {
	ch, _, comb := util.ReplaceRuneWithComb(r)
	for ; y1 <= y2; y1++ {
		s.SetContent(x, y1, ch, comb, style)
	}
}
