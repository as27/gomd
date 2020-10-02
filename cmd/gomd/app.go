package main

import (
	"github.com/rivo/tview"
)

type app struct {
	view   *tview.Application
	root   tview.Primitive
	head   tview.Primitive
	left   tview.Primitive
	right  tview.Primitive
	cmd    tview.Primitive
	bottom tview.Primitive
}

// newApp contains the initial configuration
func newApp() *app {
	a := app{
		view: tview.NewApplication(),
	}
	nextFocus := make(chan struct{})
	go func() {
		for {
			select {
			case <-nextFocus:
				if a.left.GetFocusable().HasFocus() {
					a.view.SetFocus(a.right)
				} else {
					a.view.SetFocus(a.left)
				}
				a.view.Draw()
			}
		}
	}()
	//head := tview.NewTextView()
	//head.SetTitle("head")
	a.head = newTView("t1", "text1")
	a.left = newBrowser("./", nextFocus)
	a.right = newBrowser("./../../", nextFocus)
	a.cmd = newTView("cmd", "text1 cmd")
	a.bottom = newTView("bottom", "text1 bottom")
	return &a
}

// small helper for defining the initial layout
func newTView(t, s string) *tview.TextView {
	v := tview.NewTextView()
	v.SetTitle(t)
	//v.SetBorder(true)
	v.SetText(s)
	return v
}

// run defines the place of each element inside the
// app and starts it.
func (a *app) run() error {
	root := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.head, 3, 1, false).
		AddItem(tview.NewFlex().
			AddItem(a.left, 0, 2, true).
			AddItem(a.right, 0, 2, false),
			0, 2, true).
		AddItem(a.bottom, 3, 1, false).
		AddItem(a.cmd, 3, 1, false)
	root.SetBorder(true)
	root.SetTitle("gomd")
	a.root = root
	return a.view.SetRoot(a.root, true).Run()
}
