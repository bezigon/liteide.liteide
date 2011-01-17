package main

import (
	"runtime"
	"os"
	"exec"
)

type GoBin struct {
	compiler string
	link     string
	pack     string
	cgo      string
	objext   string
	exeext   string
	pakext   string
	rm		 string
}

func newGoBin() (p *GoBin, err os.Error) {
	var goos string
	goos = os.Getenv("GOOS")
	if goos == "" {
		goos = runtime.GOOS
	}

	var exeext string
	var rm string = "rm"

	switch goos {
	case "windows":
		exeext = ".exe"
		rm ="del"
	}

	var goarch string
	goarch = os.Getenv("GOARCH")
	if goarch == "" {
		goarch = runtime.GOARCH
	}
	var o string
	switch goarch {
	case "amd64":
		o = "6"
	case "386":
		o = "8"
	case "arm":
		o = "5"
	default:
		err = os.NewError("Unsupported arch: " + goarch)
		return
	}

	p = new(GoBin)
	p.compiler = o + "g"
	p.link = o + "l"
	p.pack = "gopack"
	p.cgo = "cgo"
	p.objext = "." + o
	p.exeext = exeext
	p.pakext = ".a"
	p.rm = rm

	p.compiler, err = exec.LookPath(p.compiler)
	if err != nil {
		return
	}
	p.link, err = exec.LookPath(p.link)
	if err != nil {
		return
	}
	p.pack, err = exec.LookPath(p.pack)
	if err != nil {
		return
	}
	p.cgo, err = exec.LookPath(p.cgo)
	if err != nil {
		return
	}

	return
}