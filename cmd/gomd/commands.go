package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var implementedCommands = []string{
	"copy",
	"cp",
	"move",
	"mv",
	"remove",
	"rm",
	"mkdir",
	"makedir",
	"q",
	"quit",
	"sync",
	"sw",
	"switch",
}

func (a *app) cmdAutocomplete(currentText string) (entries []string) {
	if len(currentText) <= 1 {
		return entries
	}
	for _, cmd := range a.commands {
		if strings.HasPrefix(strings.ToLower(cmd), strings.ToLower(currentText)) {
			entries = append(entries, cmd)
		}
	}
	return entries
}

func (a *app) executeCommand(command string) {
	if command == "" {
		return
	}
	c := strings.Split(command, " ")
	switch c[0] {
	case "copy", "cp":
		if err := a.cmdCopy(); err != nil {
			fmt.Fprintln(a.appOut, "error: ", err)
		}
	case "move", "mv":
		if err := a.cmdMove(); err != nil {
			fmt.Fprintln(a.appOut, "error: ", err)
		}
	case "mkdir":
		if err := a.cmdMkdir(); err != nil {
			fmt.Fprintln(a.appOut, "error: ", err)
		}
	case "remove", "rm":
		if err := os.RemoveAll(filepath.Join(a.left.Folder.Path, a.left.Folder.SelectedFile().Name())); err != nil {
			fmt.Fprintln(a.appOut, "error: ", err)
		}
	case "sync":
		if err := a.cmdSync(); err != nil {
			fmt.Fprintln(a.appOut, "error: ", err)
		}
	case "sw", "switch":
		if err := a.cmdSwitch(); err != nil {
			fmt.Fprintln(a.appOut, "error: ", err)
		}
	case "quit", "q":
		a.view.Stop()
	}
	a.refreshView()
	a.cmd.SetText("")
}

func (a *app) getPaths() (leftPath string, rightPath string) {
	leftPath = filepath.Join(
		a.left.Folder.Path,
		a.left.Folder.SelectedFile().Name())
	rightPath = filepath.Join(
		a.right.Folder.Path,
		a.left.Folder.SelectedFile().Name())
	return
}

func (a *app) cmdCopy() error {
	srcPath, dstPath := a.getPaths()
	sourceFileStat, err := os.Stat(srcPath)
	if err != nil {
		return err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", srcPath)
	}
	source, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer source.Close()
	destination, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	if nBytes != sourceFileStat.Size() {
		return fmt.Errorf("copy of file %s has not been completed", srcPath)
	}
	return err
}

func (a *app) cmdMove() error {
	oldpath, newpath := a.getPaths()
	if oldpath == newpath {
		a.Println("nothing to move here")
		return nil
	}
	return os.Rename(oldpath, newpath)
}

func (a *app) refreshView() {
	a.left.Folder.Update()
	a.left.makeTableView()
	a.right.Folder.Update()
	a.right.makeTableView()
}

func (a *app) cmdMkdir() error {
	dirName := strings.TrimSpace(strings.TrimLeft(a.cmd.GetText(), "mkdir"))
	return os.MkdirAll(filepath.Join(a.left.Folder.Path, dirName), 0755)
}

func (a * app) cmdSync() error {
	return a.right.Folder.SetDir(a.left.Folder.Path)
}

func (a * app) cmdSwitch() error {
	leftPath := a.left.Folder.Path
	rightPath := a.right.Folder.Path
	err1 := a.left.Folder.SetDir(rightPath)
	if(err1 != nil){
		return err1
	}
	err2 := a.right.Folder.SetDir(leftPath)
	if err2 != nil{
		return err2
	} 
	return nil	
}
