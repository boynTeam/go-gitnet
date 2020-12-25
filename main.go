package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ListEffect func(*tview.List, *tview.Application)

var effects = []ListEffect{
	addQuitEventListener,
}

func showSelectList(uris []string) {
	app := tview.NewApplication()
	list := tview.NewList()
	for _, e := range effects {
		e(list, app)
	}
	for _, uri := range uris {
		localURI := uri
		list.AddItem(localURI, "", 0, func() {
			Open(localURI)
			app.Stop()
		})
	}
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		fmt.Println(err)
		return
	}
}

func addQuitEventListener(l *tview.List, app *tview.Application) {
	l.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
			return tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModNone)
		}
		return event
	})
}

func main() {
	wd, _ := os.Getwd()
	uris, err := LoadRemoteURI(wd)
	if err != nil {
		fmt.Println(err)
		return
	}
	showSelectList(uris)
}
