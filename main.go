package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	a       = flag.Bool("a", false, "echo the autocomplete script content")
	h       = flag.Bool("h", false, "show command usage")
	v       = flag.Bool("v", false, "show command version")
	s       = flag.Bool("s", false, "echo the p script content")
	version = "0.3.0"
)

func main() {
	flag.Parse()

	if *v {
		fmt.Println(version)
		return
	}

	if *s {
		fmt.Print(ScriptText())
		return
	}

	if *a {
		fmt.Print(AutocompleteText())
		return
	}

	if *h || flag.Arg(0) == "" {
		fmt.Println(HelpText())
		return
	}

	path := "~/.projects"
	name := flag.Arg(0)
	fullpath := filepath.Join(path, name)
	fullpath = strings.Replace(fullpath, "~", os.Getenv("HOME"), 1)
	if _, err := os.Stat(fullpath); err == nil {
		fmt.Println("Tmux project already exists at " + fullpath)
		return
	}
	proj := NewProject(name)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Base path for the project: ")
	res, _ := reader.ReadString('\n')
	proj.SetPath(strings.TrimSpace(res))
	fmt.Print("Name of default window: ")
	res, _ = reader.ReadString('\n')
	win1 := proj.AddWindow(strings.TrimSpace(res))
	fmt.Print("Relative path of default window: ")
	res, _ = reader.ReadString('\n')
	win1.SetPath(res)

	done := false
	for !done {
		fmt.Print("Add another window [Yn]: ")
		res, _ := reader.ReadString('\n')
		if strings.TrimSpace(res) != "" {
			d, err := strToBool(strings.TrimSpace(res))
			if err != nil {
				fmt.Println(err)
				continue
			}
			done = !d
		}
		if !done {
			fmt.Print("Name of new window: ")
			res, _ := reader.ReadString('\n')
			win := proj.AddWindow(strings.TrimSpace(res))
			fmt.Print("Relative path of new window: ")
			res, _ = reader.ReadString('\n')
			win.SetPath(strings.TrimSpace(res))
		}
	}

	contents := proj.Render()
	if err := ioutil.WriteFile(fullpath, []byte(contents), 0744); err != nil {
		fmt.Println("Failed writing tmux project file")
		return
	}
	fmt.Println("Wrote tmux project file to " + fullpath)
}

func strToBool(str string) (bool, error) {
	if str == "Y" || str == "y" {
		return true, nil
	}
	if str == "N" || str == "n" {
		return false, nil
	}
	return false, errors.New("Invalid response")
}
