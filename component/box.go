package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/draw/graph"
	"github.com/sshelll/termgraph/util"
)

type Box struct {
	baseComponent
	drawer *graph.BoxDrawer

	rawContent string

	curLn int

	scrollBarEnabled bool
}

func NewBox(name *string, enableScrollBar bool) *Box {
	box := &Box{
		baseComponent: baseComponent{
			name: name,
		},
		scrollBarEnabled: enableScrollBar,
	}
	box.initCtrlMap()
	return box
}

func (b *Box) Draw(s tcell.Screen) error {
	b.drawer = graph.NewBoxDrawer(s, b.x1, b.y1, b.x2, b.y2)
	b.sync()
	return nil
}

func (b *Box) EnableScrollBar(enabled bool) {
	b.scrollBarEnabled = enabled
	b.initCtrlMap()
}

func (b *Box) SetContent(content string) {
	b.rawContent = content
}

func (b *Box) sync() {
	if b.name != nil {
		b.drawer.SetTitle(*b.name, 1)
	}
	b.drawer.DrawBorder()
	b.drawContent()
	b.drawScrollBar()
}

func (b *Box) drawContent() {
	width, height := b.calLineWidthAndHeight()
	offset := b.curLn * width
	maxCnt := width * height

	curContent := b.rawContent[offset:util.Min(offset+maxCnt, len(b.rawContent))]

	b.drawer.SetText(curContent, true, true)
	b.drawer.DrawText()
}

func (b *Box) drawScrollBar() {
	if !b.scrollBarEnabled {
		return
	}

	_, height := b.calLineWidthAndHeight()
	if height < 1 {
		return
	}

	_, overCnt := b.calContentLineCnt()
	barHeight := util.Max(1, height-overCnt)
	scrollable := height - barHeight

	percent := float64(b.curLn) / float64(overCnt)
	offset := int(float64(scrollable) * percent)

	b.drawer.DrawScrollBar(offset, barHeight)
}

func (b *Box) calLineWidthAndHeight() (width, height int) {
	x1, y1, x2, y2 := b.x1+1, b.y1+1, b.x2-1, b.y2-1
	marginU, marginB, marginL, marginR := b.drawer.GetTextMargin()
	width = x2 - x1 + 1 - marginL - marginR
	height = y2 - y1 + 1 - marginU - marginB
	return
}

func (b *Box) calContentLineCnt() (lineCnt, overCnt int) {
	width, height := b.calLineWidthAndHeight()
	lineCnt = util.Max(1, len(b.rawContent)/width)
	overCnt = util.Max(0, lineCnt-height)
	return
}

func (b *Box) initCtrlMap() {
	// box ctrl is only designed for scroll bar.
	if !b.scrollBarEnabled {
		return
	}
	b.ctrlMap = make(map[tcell.Key]func(tcell.EventKey) error, 3)
	b.ctrlMap[tcell.KeyUp] = b.ctrlUp
	b.ctrlMap[tcell.KeyDown] = b.ctrlDown
}

func (b *Box) ctrlUp(_ tcell.EventKey) error {
	b.curLn = util.Max(0, b.curLn-1)
	b.sync()
	return nil
}

func (b *Box) ctrlDown(_ tcell.EventKey) error {
	_, overCnt := b.calContentLineCnt()
	b.curLn = util.Min(overCnt, b.curLn+1)
	b.sync()
	return nil
}
