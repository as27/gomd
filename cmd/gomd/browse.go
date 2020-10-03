package main

import (
	"fmt"

	"github.com/as27/gomd/internal/gocmd"
	"github.com/rivo/tview"
)

type browser struct {
	*tview.Table
	*gocmd.Folder
}

func newBrowser(f *gocmd.Folder) *browser {
	b := &browser{
		Table:  tview.NewTable(),
		Folder: f,
	}
	b.SetBorder(true)
	b.SetSelectable(true, false)
	b.makeTableView()
	return b
}

func (b *browser) makeTableView() error {
	b.Table.Clear()
	b.Table.SetTitle(b.Folder.Path)
	b.Table.SetCellSimple(0, 0, "..")
	b.Table.SetCellSimple(0, 1, "")
	b.Table.SetCellSimple(0, 2, "")
	for i, f := range b.Folder.Files() {
		b.Table.SetCell(i+1, 0, tview.NewTableCell(f.Name()).SetExpansion(10))
		b.Table.SetCell(i+1, 1, tview.NewTableCell(
			f.ModTime().Format("2006.01.02 15:04:05")).
			SetAlign(tview.AlignRight))
		b.Table.SetCell(i+1, 2, tview.NewTableCell(
			fmt.Sprintf("%v", f.Size())).SetAlign(tview.AlignRight))
	}
	b.Table.Select(b.Folder.Selected()+1, 0)
	return nil
}
