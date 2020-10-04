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
	//b.Table.SetCellSimple(0, 0, "..")
	//b.Table.SetCellSimple(0, 1, "")
	//b.Table.SetCellSimple(0, 2, "")
	for i, f := range b.Folder.Files() {
		dir := ""
		if f.IsDir() {
			dir = "d"
		}
		b.Table.SetCell(i, 0, tview.NewTableCell(dir))
		b.Table.SetCell(i, 1, tview.NewTableCell(f.Name()).SetExpansion(10))
		b.Table.SetCell(i, 2, tview.NewTableCell(
			f.ModTime().Format("2006.01.02 15:04:05")).
			SetAlign(tview.AlignRight))
		b.Table.SetCell(i, 3, tview.NewTableCell(
			fmt.Sprintf("%v", f.Size())).SetAlign(tview.AlignRight))
	}
	b.Table.Select(b.Folder.Selected(), 0)
	return nil
}
