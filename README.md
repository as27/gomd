# gomd

_Important: This tool is a prototype and not production ready! Use it play with it and send some feedback. Just open an issue at github for this._

* A prototype of file browser for command line. 

* Main Motive:
    * To use tview and create something useful.
    * Get more people to contribute via hacktoberfest. 

The logic of gomd is to have split file browsers. Not a new idea though (say for example, the midnight commander).


![gomd screenshot](gomd_screenshot.png)


The navigation within the file browsers is slightly similar to vim. In order to make selection and move down you should use the key under your index finger. 


At present the navigation logics are quite simple:

* index-finger navigates down
    * f: navigates down to the left side
    * j: navigates down to the right side
* middle-finger navigates up
    * d: navigates up to the left side
    * k: navigates up to the right side
* ring-finger goes one folder up
    * s: one folder up left
    * l: one folder up right
* the keys in the middle let you navigate inside a folder
    * g: for the left side
    * h: for the right side

For copying or other commands there is a command-mode. To activate that mode, just press `:`. To go back into normal mode use `esc`

At the moment there are following commands implemented:

* `copy`, `cp`: copies the file from the left to the right side
* `move`, `mv`: moves the file from the left to the right side
* `remove`,	`rm`: removes the file from the right side
* `mkdir [dirname]`, `makedir [dirname]`: creates that folder on the left side
* `q`, `quit`: exits the app

## Installation

This tool is written in Go. For a fast installation just run 

`go get github.com/as27/gomd/cmd/gomd`

To update just type

`go get -u github.com/as27/gomd/cmd/gomd`