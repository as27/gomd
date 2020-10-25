package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/as27/gomd/internal/gocmd"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type app struct {
	conf  appConf
	files gocmd.Files
	view  *tview.Application
	root  tview.Primitive
	// head    tview.Primitive
	left     *browser
	right    *browser
	cmd      *tview.InputField
	bottom   tview.Primitive
	log      tview.Primitive
	cmdMode  bool
	commands []string
	appOut   io.Writer
}

type appConf struct {
	verbose bool
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
		view:     tview.NewApplication(),
		commands: implementedCommands,
	}
	a.view.SetInputCapture(a.inputEvents)
	// a.head = tview.NewTextView()
	a.left = newBrowser(a.files.Left)
	a.right = newBrowser(a.files.Right)
	a.cmd = tview.NewInputField().
		SetLabel("Cmd:").
		SetFieldBackgroundColor(tcell.ColorBlack)
	a.cmd.SetAutocompleteFunc(a.cmdAutocomplete)
	a.bottom = tview.NewTextView()
	a.log = tview.NewTextView()

	w, ok := a.bottom.(io.Writer)
	if ok {
		a.appOut = w
	} else {
		a.appOut = os.Stdout
	}
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
		//AddItem(a.head, 1, 1, false).
		AddItem(tview.NewFlex().
			AddItem(a.left, 0, 2, false).
			AddItem(a.right, 0, 2, false),
			0, 2, true).
		AddItem(a.cmd, 1, 1, false).
		AddItem(a.bottom, 2, 1, false)
	if a.conf.verbose {
		root = root.AddItem(a.log, 4, 1, false)
	}
	//root.SetBorder(true)
	root.SetTitle("gomd")
	a.root = root
	return a.view.SetRoot(a.root, true).Run()
}

func (a *app) update() {
	a.view.Draw()
}

func (a *app) enterFolder(b *browser) error {
	if b == nil || !b.HasFocus() {
		return nil
	}
	selectedFile := b.Folder.SelectedFile()
	if !selectedFile.IsDir() {
		return nil
	}
	return b.Folder.SetDir(
		filepath.Join(b.Folder.Path, selectedFile.Name()))
}

func (a *app) Println(i ...interface{}) {
	fmt.Fprintln(a.appOut, i...)
}
