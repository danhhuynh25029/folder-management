package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	err := InitLayout()
	if err != nil {
		panic(err)
	}
}
func InitLayout() error {
	box := tview.NewBox().SetBorder(true).SetTitle("Folder Manage")
	app := tview.NewApplication()
	EventHandle(app)
	if err := app.SetRoot(box, true).EnableMouse(true).Run(); err != nil {
		return err
	}
	return nil
}

func EventHandle(app *tview.Application) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		}
		return event
	})
}
