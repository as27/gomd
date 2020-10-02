package gocmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func createTmpFiles(n int) (fpath string, fileInfos []os.FileInfo) {
	fpath, err := ioutil.TempDir(os.TempDir(), "*-gocmd")
	if err != nil {
		fmt.Println("Error ioutilTempDir: ", err)
		panic("Cannot create TmpDir")
	}
	fileInfos = []os.FileInfo{}
	for i := 0; i <= n; i++ {
		fd, err := ioutil.TempFile(fpath, "*-gocmd.txt")
		if err != nil {
			fmt.Println("Error ioutilTempFile: ", err)
			panic("Cannot create TmpFile")
		}
		if err = fd.Close(); err != nil {
			fmt.Println("error closing tmp file:", err)
		}
		fi, err := os.Stat(fd.Name())
		if err != nil {
			fmt.Println("error os.Stat", err)
		}
		fileInfos = append(fileInfos, fi)
	}
	return fpath, fileInfos
}

func TestFolder_Next(t *testing.T) {
	tmpFolder, _ := createTmpFiles(2)
	defer os.RemoveAll(tmpFolder)
	f, err := NewFolder(tmpFolder)
	if err != nil {
		t.Errorf("no error expected! Got: %v", err)
	}
	tests := []struct {
		name string
		f    *Folder
		want int
	}{
		{
			"index 1",
			f,
			1,
		},
		{
			"index 2",
			f,
			2,
		},
		{
			"index 2",
			f,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Next(); got != tt.want {
				t.Errorf("Folder.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFolder_Prev(t *testing.T) {
	tmpFolder, _ := createTmpFiles(3)
	defer os.RemoveAll(tmpFolder)
	f, err := NewFolder(tmpFolder)
	if err != nil {
		t.Errorf("no error expected! Got: %v", err)
	}
	f.Next() // set index to 1
	f.Next() // set index to 2
	tests := []struct {
		name string
		f    *Folder
		want int
	}{
		{
			"index 1",
			f,
			1,
		},
		{
			"index 0",
			f,
			0,
		},
		{
			"index 0",
			f,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Prev(); got != tt.want {
				t.Errorf("Folder.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}
