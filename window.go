package main

import (
	"fmt"
	"strings"
)

type Window struct {
	Name  string
	Path  string
	Index int
}

func NewWindow(name string, index int) *Window {
	return &Window{Name: name, Index: index}
}

func (w *Window) SetPath(path string) {
	w.Path = strings.TrimSpace(path)
}

func (w *Window) String() string {
	return "Window name:" + w.Name + " index:" + string(w.Index)
}

func (w *Window) Render() string {
	cmd := "new-window -t"
	if w.Index == 1 {
		cmd = "new-session -s"
	}
	return fmt.Sprintf(w.Template(), w.Index, cmd, w.Name, w.Path)
}

func (w *Window) Template() string {
	return `WINDOW_NUM=%d
tmux %s $PROJECT_NAME -n %s -d
tmux send-keys -t $PROJECT_NAME:$WINDOW_NUM "cd $PROJECT_PWD/%s" C-m
tmux send-keys -t $PROJECT_NAME:$WINDOW_NUM "clear" C-m
`
}
