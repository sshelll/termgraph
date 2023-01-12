package main

import (
	"log"

	"github.com/AlekSi/pointer"
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/component"
)

func main() {
	//testStableCanvas()
	testDrawBox()
}

func testDrawBox() {

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	defer func() {
		r := recover()
		s.Fini()
		if r != nil {
			log.Println("panic info =", r)
		}
	}()

	// Draw initial boxes
	content := "hello world box, this is a box with text, now it's too long, and and and.. !"
	for i := 0; i < 3; i++ {
		content += content
	}

	box := component.NewBox(pointer.ToString("Box"), true)
	box.SetLayout(0, 0, 42, 7)
	box.SetContent(content)
	box.Draw(s)

loop:
	for {
		s.Show()
		switch event := s.PollEvent().(type) {
		case *tcell.EventKey:
			key := event.Key()
			//drawBox(screen, 0, 0, 10, 10, defStyle, "dadasdasdadad")
			if key == tcell.KeyESC {
				break loop
			} else {
				box.Ctrl(*event)
			}
		}
	}

	log.Println("fini")

}
