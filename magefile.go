// +build mage

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/magefile/mage/mg"

	"github.com/magefile/mage/sh"
)

// this is a test target
func LogTest() {
	fmt.Println("Test")
}

// runs mages
func Run() {
	sh.RunV("go", "run", "main.go")
}

// fails with runtime error
func Fail() {
	fmt.Println(os.Args[math.MaxInt64])
}

// tries to panic
func Panic() {
	panic("mage panic")
}

// Runs "go install" for mage-select.
func Install() error {
	name := "mages"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}

	gocmd := mg.GoCmd()
	gopath, err := sh.Output(gocmd, "env", "GOPATH")
	if err != nil {
		return fmt.Errorf("can't determine GOPATH: %v", err)
	}
	paths := strings.Split(gopath, string([]rune{os.PathListSeparator}))
	bin := filepath.Join(paths[0], "bin")
	// specifically don't mkdirall, if you have an invalid gopath in the first
	// place, that's not on us to fix.
	if err := os.Mkdir(bin, 0700); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create %q: %v", bin, err)
	}
	path := filepath.Join(bin, name)

	// we use go build here because if someone built with go get, then `go
	// install` turns into a no-op, and `go install -a` fails on people's
	// machines that have go installed in a non-writeable directory (such as
	// normal OS installs in /usr/bin)
	return sh.RunV(gocmd, "build", "-o", path, "-ldflags="+flags(), "github.com/iwittkau/mage-select")
}

func flags() string {
	timestamp := time.Now().Format(time.RFC3339)
	hash := hash()
	tag := tag()
	if tag == "" {
		tag = "dev"
	}
	mageVersion := mageVersion()
	return fmt.Sprintf(`-X "main.timestamp=%s" -X "main.hash=%s" -X "main.tag=%s" -X "main.mageVersion=%s"`, timestamp, hash, tag, mageVersion)
}

// tag returns the git tag for the current branch or "" if none.
func tag() string {
	s, _ := sh.Output("git", "describe", "--tags")
	// s := "v0.0.0"
	return s
}

// hash returns the git hash for the current repo or "" if none.
func hash() string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	// hash := "0000000"
	return hash
}

// mage tag returns the git tag for the current branch or "" if none.
func mageVersion() string {
	s, _ := sh.Output("cat", "go.mod")
	scan := bufio.NewScanner(bytes.NewBufferString(s))
	ver := ""
	for scan.Scan() {
		line := scan.Text()
		if !strings.Contains(line, "github.com/magefile/mage") {
			continue
		}
		parts := strings.Split(line, " ")
		ver = parts[len(parts)-1]
		break
	}
	return ver
}
