package graph

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/draw/content"
	"github.com/sshelll/termgraph/util"
)

type DrawBoxOpt struct {
	Title       *string
	TitleOffset int

	Text              *string
	AutoCut, AutoWrap bool
	MarginU, MarginB  int
	MarginL, MarginR  int

	TextStyle, TitleStyle *tcell.Style
}

func DrawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, opt *DrawBoxOpt) {
	drawer := &boxDrawer{}
	drawer.Draw(s, x1, y1, x2, y2, style, opt)
}

type boxDrawer struct {
}

func (b *boxDrawer) Draw(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, opt *DrawBoxOpt) {

	// param check
	x1, x2 = util.SwapByOrder(x1, x2)
	y1, y2 = util.SwapByOrder(y1, y2)
	b.drawBox(s, x1, y1, x2, y2, style, opt)

	if opt == nil {
		return
	}

	// draw title
	if opt.Title != nil {
		content.SetContentH(s, x1+opt.TitleOffset, y1, *opt.Title, style)
	}

	// draw content
	if opt.Text != nil {
		if opt.TextStyle == nil {
			opt.TextStyle = &style
		}
		b.drawText(s, x1+1, y1+1, x2-1, y2-1, style, opt)
	}

}

func (b *boxDrawer) drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, opt *DrawBoxOpt) {

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

func (b *boxDrawer) drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, opt *DrawBoxOpt) {

	if opt == nil {
		return
	}

	if opt.MarginL < 0 || opt.MarginR < 0 || opt.MarginU < 0 || opt.MarginB < 0 {
		panic("margin should not be negative")
	}

	textMaxWidth := x2 - x1 + 1 - opt.MarginL - opt.MarginR
	textMaxHeight := y2 - y1 + 1 - opt.MarginU - opt.MarginB

	if textMaxWidth <= 0 || textMaxHeight <= 0 {
		panic("margin too wide")
	}

	// single line
	if !opt.AutoWrap {
		cuttedContent := *opt.Text
		if opt.AutoCut { // cut content
			cuttedContent = b.cutBoxText(cuttedContent, textMaxWidth)
		}
		content.SetContentH(s, x1+opt.MarginL, y1+opt.MarginU, cuttedContent, *opt.TextStyle)
		return
	}

	cuttedContent := *opt.Text
	if opt.AutoCut {
		cuttedContent = b.cutBoxText(cuttedContent, textMaxHeight*textMaxWidth)
	}

	content.SetWrappedContentH(s, x1+opt.MarginL, y1+opt.MarginU, cuttedContent, textMaxWidth, *opt.TextStyle)

}

func (b *boxDrawer) cutBoxText(ori string, max int) string {
	cutted := util.CutString(ori, max-3)
	if max > 3 {
		cutted += "..."
	}
	return cutted
}
