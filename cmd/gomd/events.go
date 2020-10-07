package main

import (
	"log"
	"path/filepath"

	"github.com/gdamore/tcell"
)

func (a *app) inputEvents(event *tcell.EventKey) *tcell.EventKey {
	log.Printf("Key(): %v Rune():%v", event.Key(), event.Rune())
	// ensure that ctrl c always works
	/*if event.Key() == tcell.KeyCtrlC {
		os.Exit(0)
	}*/
	if a.cmdMode {
		if event.Key() == tcell.KeyEsc {
			a.cmd.SetText("")
			a.cmd.SetFieldBackgroundColor(tcell.ColorBlack)
			a.view.SetFocus(a.left)
			a.cmdMode = false
		}
		return event
	}
	switch event.Key() {
	case tcell.KeyEnter:
		if a.left.HasFocus() {
			a.enterFolder(a.left)
			a.left.makeTableView()
			return nil
		} else if a.right.HasFocus() {
			a.enterFolder(a.right)
			a.right.makeTableView()
			return nil
		}
	case tcell.KeyTab:
		if a.left.HasFocus() {
			a.view.SetFocus(a.right)
		} else if a.right.HasFocus() {
			a.view.SetFocus(a.left)
		}
	}
	switch event.Rune() {
	case ':': // switch to cmd mode
		a.view.SetFocus(a.cmd)
		a.cmdMode = true
		a.cmd.SetFieldBackgroundColor(tcell.ColorGray)
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
		a.right.Folder.SetDir(
			filepath.Join(a.right.Path, ".."))
		a.right.makeTableView()
		return nil
	case 'h': // enter folder right
		a.view.SetFocus(a.right)
		a.enterFolder(a.right)
		a.right.makeTableView()
		return nil
	case 's': // one folder up left
		a.view.SetFocus(a.left)
		a.left.Folder.SetDir(
			filepath.Join(a.left.Folder.Path, ".."))
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
		a.enterFolder(a.left)
		a.left.makeTableView()
		return nil
	}
	return event
}
