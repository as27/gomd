package main

import "github.com/rivo/tview"

type app struct {
	head   tview.Primitive
	left   tview.Primitive
	right  tview.Primitive
	cmd    tview.Primitive
	bottom tview.Primitive
}

func newApp() *app {
	return &app{}
}

func (a *app) run() error {
	return nil
}
