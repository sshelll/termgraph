package main

import (
	"log"

	"github.com/AlekSi/pointer"
	"github.com/gdamore/tcell/v2"
	"github.com/sshelll/termgraph/canvas"
	"github.com/sshelll/termgraph/component"
	"github.com/sshelll/termgraph/draw/graph"
)

func main() {
	//testStableCanvas()
	testDrawBox()
}

func testStableCanvas() {
	stableCanvas, err := canvas.NewStableCanvas()
	if err != nil {
		log.Fatalln(err)
	}
	stableCanvas.AddComponent(component.NewBox(nil), "box1", canvas.NewPosition(0, 0, 100, 10), nil)
	stableCanvas.AddComponent(component.NewBox(nil), "box2", canvas.NewPosition(0, 11, 100, 20), nil)
	stableCanvas.AddComponent(component.NewBox(nil), "box3", canvas.NewPosition(101, 0, 201, 10), nil)
	stableCanvas.AddComponent(component.NewBox(nil), "box4", canvas.NewPosition(101, 11, 201, 20), nil)
	if err := stableCanvas.Show(); err != nil {
		log.Fatalln(err)
	}
	stableCanvas.RemoveComponent("box1")
	stableCanvas.RemoveComponent("box4")
	if err := stableCanvas.Show(); err != nil {
		log.Fatalln(err)
	}
	stableCanvas.Fini()
}

func testInitAgain() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	screen.Init()
	screen.Init()
	screen.Fini()
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
	//graph.DrawBox(s, 0, 0, 42, 7, defStyle)
	content := "hello world box, this is a box with text, now it's too long!"
	for i := 0; i < 5; i++ {
		content += content
	}
	graph.DrawBox(s, 0, 0, 42, 7, defStyle, &graph.DrawBoxOpt{
		Title:       pointer.ToString("Test"),
		TitleOffset: 1,
		Text:        pointer.ToString(content),
		AutoCut:     true,
		AutoWrap:    true,
	})
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
