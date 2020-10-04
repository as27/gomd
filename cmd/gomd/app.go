package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/as27/gomd/internal/gocmd"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type app struct {
	files   gocmd.Files
	view    *tview.Application
	root    tview.Primitive
	head    tview.Primitive
	left    *browser
	right   *browser
	cmd     tview.Primitive
	bottom  tview.Primitive
	cmdMode bool
}

func (a *app) inputEvents(event *tcell.EventKey) *tcell.EventKey {
	log.Printf("Key(): %v Rune():%v", event.Key(), event.Rune())
	// ensure that ctrl c always works
	if event.Key() == tcell.KeyCtrlC {
		os.Exit(0)
	}
	if a.cmdMode {
		if event.Key() == tcell.KeyEsc {
			a.cmdMode = false
		}
		return event
	}
	switch event.Rune() {
	case ':': // switch to cmd mode
		a.view.SetFocus(a.cmd)
		a.cmdMode = true
		return nil
	case 'j': // j next right
		a.view.SetFocus(a.right)
		a.files.Right.Next()
		a.right.makeTableView()
		return nil
	case 'k':
		a.view.SetFocus(a.right)
		a.files.Right.Prev()
		a.right.makeTableView()
		return nil
	case 'l': // one folder up right
		a.view.SetFocus(a.right)
		a.files.Right, _ = gocmd.NewFolder(
			filepath.Join(a.files.Right.Path, ".."))
		a.right.Folder = a.files.Right
		a.right.makeTableView()
		return nil
	case 'h': // enter folder right
		a.view.SetFocus(a.right)
		selectedFile := a.right.Folder.SelectedFile()
		if !selectedFile.IsDir() {
			return nil
		}
		a.files.Right, _ = gocmd.NewFolder(
			filepath.Join(a.files.Right.Path, selectedFile.Name()))
		a.right.Folder = a.files.Right
		a.right.makeTableView()
		return nil
	case 's': // one folder up left
		a.view.SetFocus(a.left)
		a.files.Left, _ = gocmd.NewFolder(
			filepath.Join(a.files.Left.Path, ".."))
		a.left.Folder = a.files.Left
		a.left.makeTableView()
		return nil
	case 'f': // f next left
		a.view.SetFocus(a.left)
		a.files.Left.Next()
		a.left.makeTableView()
		return nil
	case 'd': // f prev left
		a.view.SetFocus(a.left)
		a.files.Left.Prev()
		a.left.makeTableView()
		return nil
	case 'g': // enter folder left
		a.view.SetFocus(a.left)
		selectedFile := a.left.Folder.SelectedFile()
		if !selectedFile.IsDir() {
			return nil
		}
		a.files.Left, _ = gocmd.NewFolder(
			filepath.Join(a.files.Left.Path, selectedFile.Name()))
		a.left.Folder = a.files.Left
		a.left.makeTableView()
		return nil
	}
	return event
}

// newApp contains the initial configuration
func newApp() *app {
	left, err := gocmd.NewFolder("./")
	if err != nil {
		fmt.Println("cannot initialize folder: ", err)
		panic("cannot read folder")
	}
	right, err := gocmd.NewFolder("./")
	if err != nil {
		fmt.Println("cannot initialize folder: ", err)
		panic("cannot read folder")
	}
	a := app{
		files: gocmd.Files{
			Left:  left,
			Right: right,
		},
		view: tview.NewApplication(),
	}
	a.view.SetInputCapture(a.inputEvents)
	a.head = newTView("t1", "text1")
	a.left = newBrowser(a.files.Left)
	a.right = newBrowser(a.files.Right)
	a.cmd = tview.NewInputField().
		SetLabel("Cmd:")
	a.bottom = tview.NewTextView()
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
		AddItem(a.head, 1, 1, false).
		AddItem(tview.NewFlex().
			AddItem(a.left, 0, 2, false).
			AddItem(a.right, 0, 2, false),
			0, 2, true).
		AddItem(a.cmd, 1, 1, false).
		AddItem(a.bottom, 4, 1, false)
	//root.SetBorder(true)
	root.SetTitle("gomd")
	a.root = root
	return a.view.SetRoot(a.root, true).Run()
}

func (a *app) update() {
	a.view.Draw()
}
