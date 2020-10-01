package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/rivo/tview"
)

type browser struct {
	*tview.Table
	fpath string
	files []string
}

func newBrowser(fpath string) *browser {
	b := &browser{
		Table: tview.NewTable(),
		fpath: fpath,
	}
	b.SetBorder(true)
	b.SetTitle(b.fpath)
	b.SetSelectable(true, false)
	b.dir(fpath)
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
	for i, f := range fs {
		name := f.Name()
		b.Table.SetCellSimple(i, 0, name)
		b.Table.SetCellSimple(i, 1, f.ModTime().Format("2006.01.02 15:04:05"))
		b.Table.SetCellSimple(i, 2, fmt.Sprintf("%v", f.Size()))
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
