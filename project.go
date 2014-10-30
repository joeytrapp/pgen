package main

import (
	"fmt"
	"strings"
)

type Project struct {
	Name    string
	Path    string
	Windows []*Window
}

func NewProject(name string) *Project {
	return &Project{Name: name}
}

func (p *Project) AddWindow(name string) *Window {
	win := *NewWindow(name, len(p.Windows)+1)
	p.Windows = append(p.Windows, &win)
	return &win
}

func (p *Project) SetPath(path string) {
	p.Path = strings.TrimSpace(path)
}

func (p *Project) String() string {
	return "Project name:" + p.Name + " windows:" + string(len(p.Windows))
}

func (p *Project) Render() string {
	windows := ""
	for _, window := range p.Windows {
		win := window.Render()
		c := ""
		for _, line := range strings.Split(win, "\n") {
			c = c + "\t" + line + "\n"
		}
		windows = windows + c
	}
	if len(windows) > 1 {
		runes := []rune(windows)
		windows = string(runes[0 : len(runes)-1])
	}
	return fmt.Sprintf(p.Template(), p.Name, p.Path, windows)
}

func (p *Project) Template() string {
	return `#!/usr/bin/env bash

PROJECT_NAME='%s'
PROJECT_PWD='%s'

_TMUX="$TMUX"
export TMUX=""

tmux has-session -t $PROJECT_NAME
if [ $? != 0 ]; then
%s
	tmux select-window -t $PROJECT_NAME:1
fi

if [ "$_TMUX" != "" ]; then
  tmux switch-client -t $PROJECT_NAME
else
  tmux attach -t $PROJECT_NAME
fi`
}
