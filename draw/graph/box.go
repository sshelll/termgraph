package graph

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/draw/content"
	"github.com/sshelll/termgraph/util"
)

type BoxDrawer struct {
	s              tcell.Screen
	x1, y1, x2, y2 int

	title       *string
	titleOffset int

	text              *string
	autoCut, autoWrap bool
	marginU, marginB  int
	marginL, marginR  int

	style, textStyle, titleStyle *tcell.Style
}

func NewBoxDrawer(s tcell.Screen, x1, y1, x2, y2 int) *BoxDrawer {
	return &BoxDrawer{
		s:  s,
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
}

func (b *BoxDrawer) SetTitle(t string, offset int) *BoxDrawer {
	b.title = &t
	b.titleOffset = offset
	return b
}

func (b *BoxDrawer) SetText(text string, autoCut, autoWrap bool) *BoxDrawer {
	b.text = &text
	b.autoCut = autoCut
	b.autoWrap = autoWrap
	return b
}

func (b *BoxDrawer) SetTextMargin(up, bottom, left, right int) *BoxDrawer {
	b.marginU = up
	b.marginB = bottom
	b.marginL = left
	b.marginR = right
	return b
}

func (b *BoxDrawer) GetTextMargin() (up, bottom, left, right int) {
	return b.marginU, b.marginB, b.marginL, b.marginR
}

func (b *BoxDrawer) SetBoxStyle(s tcell.Style) *BoxDrawer {
	b.style = &s
	return b
}

func (b *BoxDrawer) SetTitleStyle(s tcell.Style) *BoxDrawer {
	b.titleStyle = &s
	return b
}

func (b *BoxDrawer) SetTextStyle(s tcell.Style) *BoxDrawer {
	b.textStyle = &s
	return b
}

func (b *BoxDrawer) Draw() {

	b.init()
	b.x1, b.x2 = util.SwapByOrder(b.x1, b.x2)
	b.y1, b.y2 = util.SwapByOrder(b.y1, b.y2)

	b.DrawBorder()

	if b.title != nil {
		content.SetContentH(b.s, b.x1+b.titleOffset, b.y1, *b.title, *b.style)
	}

	if b.text != nil {
		b.DrawText()
	}

}

func (b *BoxDrawer) DrawBorder() {

	b.init()
	// fill background
	for row := b.y1; row <= b.y2; row++ {
		for col := b.x1; col <= b.x2; col++ {
			b.s.SetContent(col, row, ' ', nil, *b.style)
		}
	}

	// draw horizonal borders
	content.FillContentH(b.s, b.x1, b.x2, b.y1, tcell.RuneHLine, *b.style)
	content.FillContentH(b.s, b.x1, b.x2, b.y2, tcell.RuneHLine, *b.style)

	// draw vertial borders
	content.FillContentV(b.s, b.y1, b.y2, b.x1, tcell.RuneVLine, *b.style)
	content.FillContentV(b.s, b.y1, b.y2, b.x2, tcell.RuneVLine, *b.style)

	// draw corners if necessary
	if b.y1 != b.y2 && b.x1 != b.x2 {
		b.s.SetContent(b.x1, b.y1, tcell.RuneULCorner, nil, *b.style)
		b.s.SetContent(b.x2, b.y1, tcell.RuneURCorner, nil, *b.style)
		b.s.SetContent(b.x1, b.y2, tcell.RuneLLCorner, nil, *b.style)
		b.s.SetContent(b.x2, b.y2, tcell.RuneLRCorner, nil, *b.style)
	}

}

func (b *BoxDrawer) DrawScrollBar(offset, barHeight int) {
	b.init()
	y1 := util.Min(b.y1+1+offset, b.y2-barHeight)
	content.FillContentV(b.s, y1, y1+barHeight-1, b.x2, tcell.RuneBlock, *b.style)
}

func (b *BoxDrawer) DrawText() {

	b.init()
	// text pos should be 'inner'
	x1, y1, x2, y2 := b.x1+1, b.y1+1, b.x2-1, b.y2-1

	if b.marginL < 0 || b.marginR < 0 || b.marginU < 0 || b.marginB < 0 {
		panic("margin should not be negative")
	}

	textMaxWidth := x2 - x1 + 1 - b.marginL - b.marginR
	textMaxHeight := y2 - y1 + 1 - b.marginU - b.marginB

	if textMaxWidth <= 0 || textMaxHeight <= 0 {
		panic("margin too wide")
	}

	// single line
	if !b.autoWrap {
		cuttedContent := *b.text
		if b.autoCut { // cut content
			cuttedContent = util.CutString(cuttedContent, textMaxWidth)
		}
		content.SetContentH(b.s, x1+b.marginL, y1+b.marginU, cuttedContent, *b.textStyle)
		return
	}

	cutted := *b.text
	if b.autoCut {
		cutted = util.CutString(cutted, textMaxHeight*textMaxWidth)
	}

	content.SetWrappedContentH(b.s, x1+b.marginL, y1+b.marginU, cutted, textMaxWidth, *b.textStyle)

}

func (b *BoxDrawer) init() {
	if b.style == nil {
		b.style = &tcell.StyleDefault
	}
	if b.textStyle == nil {
		b.textStyle = b.style
	}
	if b.titleStyle == nil {
		b.titleStyle = b.style
	}
}
