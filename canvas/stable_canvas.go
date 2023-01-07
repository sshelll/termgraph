package canvas

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

type stableCanvas struct {
	baseCanvas
	initOnce sync.Once
}

func NewStableCanvas() (canvas Canvas, err error) {
	base, err := newBaseCanvas()
	if err != nil {
		return
	}
	canvas = &stableCanvas{
		initOnce:   sync.Once{},
		baseCanvas: base,
	}
	return
}

func (c *stableCanvas) Show() (err error) {
	c.initOnce.Do(func() {
		err = c.screen.Init()
	})
	if err != nil {
		return
	}
	if err = c.redraw(); err != nil {
		return err
	}
	c.screen.Show()
	c.startCtrl()
	return nil
}

func (c *stableCanvas) startCtrl() {
loop:
	for {
		switch event := c.screen.PollEvent().(type) {
		case *tcell.EventKey:
			// TODO: @shaojiale maybe allow user to customize some key mappings.
			key := event.Key()
			if key == tcell.KeyESC {
				break loop
			}
		}
	}
}

func (c *stableCanvas) redraw() error {
	c.screen.Clear()
	for _, comp := range c.components {
		if err := comp.Draw(c.screen); err != nil {
			return err
		}
	}
	return nil
}
