package gocmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Folder strores the relevant information
type Folder struct {
	Path     string
	files    []os.FileInfo
	selected int
}

// NewFolder creates a folder type and loads the FileInfos
// of a directory and sets the index 0 as selected.
func NewFolder(fpath string) (*Folder, error) {
	if !filepath.IsAbs(fpath) {
		fpath, err := filepath.Abs(fpath)
		if err != nil {
			return nil, fmt.Errorf("gocmd.NewFolder(%s) Abs(): %w", fpath, err)
		}
	}
	f := &Folder{
		Path: fpath,
	}
	files, err := ioutil.ReadDir(fpath)
	if err != nil {
		return f, fmt.Errorf("gocmd.NewFolder(%s) ReadDir: %w", fpath, err)
	}
	f.files = files
	f.selected = 0
	return f, nil
}

// Next selects the next file inside the folder
func (f *Folder) Next() int {
	f.selected++
	if f.selected > len(f.files)-1 {
		f.selected = len(f.files) - 1
	}
	return f.selected
}

// Prev selects the privious file inside the folder
func (f *Folder) Prev() int {
	f.selected--
	if f.selected < 0 {
		f.selected = 0
	}
	return f.selected
}

func (f *Folder) Files() []os.FileInfo {
	return f.files
}

func (f *Folder) Selected() int {
	return f.selected
}
