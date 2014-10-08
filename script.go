package main

func ScriptText() string {
	return `#!/usr/bin/env bash
PROJECTS_DIR="$HOME/.projects"
if [ "$1" == "" ]; then
	ls $PROJECTS_DIR
	exit 0
fi
if [ -f "$PROJECTS_DIR/$1" ]; then
	$PROJECTS_DIR/$1
else
	echo "$1 does not exist in $PROJECTS_DIR"
fi
`
}
