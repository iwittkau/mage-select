package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui/list"

	"github.com/magefile/mage/mage"
	"github.com/manifoldco/promptui"
)

const (
	maxSize = 10
)

var (
	mageVersion = "<unknown>"
	timestamp   = "<unknown>"
	hash        = "<unknown>"
	tag         = "<unknown>"
)

func searcher(targets []string) list.Searcher {
	return func(input string, index int) bool {
		if strings.Contains(strings.ToLower(targets[index]), input) {
			return true
		}
		return false
	}
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-version" {
			fmt.Printf("mage-select CLI frontend: %s\nBuild Date: %s\nCommit: %s\nMage: %s\n", tag, timestamp, hash, mageVersion)
			os.Exit(0)
		}
	}

	cmd := exec.Command("mage", "-l")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return
	}

	scan := bufio.NewScanner(bytes.NewBuffer(out))

	targets := []string{}
	for scan.Scan() {
		line := scan.Text()
		if strings.HasPrefix(line, "Targets:") {
			continue
		}
		line = strings.TrimSpace(line)
		targets = append(targets, line)
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{.}}",
		Active:   promptui.IconSelect + " {{.}}",
		Inactive: "  {{.|faint}}",
		Selected: promptui.IconGood + " {{.}}",
	}

	size := maxSize
	if len(targets) < size {
		size = len(targets)
	}

	prompt := promptui.Select{
		Label:             "Select a mage target:",
		Items:             targets,
		Templates:         templates,
		HideHelp:          true,
		Size:              size,
		Searcher:          searcher(targets),
		StartInSearchMode: true,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return
	}

	result = strings.Split(result, " ")[0]

	fmt.Printf("mage %s\n", result)

	os.Exit(mage.ParseAndRun(os.Stdout, os.Stderr, os.Stdin, []string{result}))
}
