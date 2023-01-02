package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/canvas"
	"github.com/sshelll/termgraph/component"
	"github.com/sshelll/termgraph/screen/graph"
)

func main() {
	stableCanvas, err := canvas.NewStableCanvas()
	if err != nil {
		log.Fatalln(err)
	}
	stableCanvas.AddComponent(component.NewBox(nil), canvas.NewPosition(0, 0, 100, 10), nil)
	stableCanvas.AddComponent(component.NewBox(nil), canvas.NewPosition(0, 11, 100, 20), nil)
	stableCanvas.AddComponent(component.NewBox(nil), canvas.NewPosition(101, 0, 201, 10), nil)
	stableCanvas.AddComponent(component.NewBox(nil), canvas.NewPosition(101, 11, 201, 20), nil)
	if err := stableCanvas.Draw(); err != nil {
		log.Fatalln(err)
	}
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

	// Draw initial boxes
	graph.DrawBox(s, 0, 0, 42, 7, defStyle)
	graph.DrawBox(s, 0, 8, 32, 14, defStyle)
	s.Show()

loop:
	for {
		switch event := s.PollEvent().(type) {
		case *tcell.EventKey:
			key := event.Key()
			//drawBox(screen, 0, 0, 10, 10, defStyle, "dadasdasdadad")
			if key == tcell.KeyESC {
				break loop
			}
		}
	}

	s.Fini()
	log.Println("fini")

}
