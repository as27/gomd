package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type browser struct {
	*tview.Table
	fpath string
	files []string
}

func newBrowser(fpath string, nextFocus chan struct{}) *browser {
	b := &browser{
		Table: tview.NewTable(),
		fpath: fpath,
	}
	b.SetBorder(true)
	b.SetTitle(b.fpath)
	b.SetSelectable(true, false)
	b.dir(fpath)
	b.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		log.Println(event.Key())
		switch event.Key() {
		case 9: // tab key
			nextFocus <- struct{}{}
		case 13: // enter key
			r, _ := b.GetSelection()
			v := b.GetCell(r, 0).Text
			b.dir(filepath.Join(fpath, v))
		}
		//r, _ := b.GetSelection()
		//log.Println("Selected row: ", r)
		/*if event.Key() == tcell.KeyCtrlJ || event.Key() == 9 {
			a.view.SetFocus(a.navView)
		}*/
		return event
	})
	return b
}

func (b *browser) dir(fpath string) error {
	fpath, _ = filepath.Abs(fpath)
	fs, err := ioutil.ReadDir(fpath)
	if err != nil {
		log.Fatal("Cannot scan dir:", err)
	}
	b.files = []string{}
	b.Table.Clear()
	b.Table.SetTitle(fpath)
	b.Table.SetCellSimple(0, 0, "..")
	b.Table.SetCellSimple(0, 1, "")
	b.Table.SetCellSimple(0, 2, "")
	for i, f := range fs {
		b.Table.SetCell(i+1, 0, tview.NewTableCell(f.Name()).SetExpansion(10))
		b.Table.SetCell(i+1, 1, tview.NewTableCell(
			f.ModTime().Format("2006.01.02 15:04:05")).
			SetAlign(tview.AlignRight))
		b.Table.SetCell(i+1, 2, tview.NewTableCell(
			fmt.Sprintf("%v", f.Size())).SetAlign(tview.AlignRight))
	}

	return nil
	/*
		a.navView.AddItem("..", "", 0, func() {
			newPath := filepath.Join(fpath, "..")
			a.dir(newPath)
		})
		for _, f := range fs {
			name := f.Name()
			if f.IsDir() {
				a.navView.AddItem(fmt.Sprintf("+ %s", name), "", 0, func() {
					newPath := filepath.Join(fpath, name)
					log.Println("New Folder: ", newPath)
					a.dir(newPath)
				})
				continue
			}
			a.navView.AddItem(name, "", 0, func() {
				a.txtView.SetTitle(name)
				a.setFile(filepath.Join(fpath, name))
				a.view.SetFocus(a.txtView)
			})
			a.files = append(a.files, f.Name())
		}
		return nil
	*/
}
