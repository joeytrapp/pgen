# pgen

A script for creating `tmux` project files in `~/.projects`. The project file is created with the name of the first arg to pgen and then there is a questionnaire for naming the windows and giving them a relative path.

## Installation

If you haven't yet, [install Go](https://golang.org/doc/install) and make sure to [setup $GOPATH](https://golang.org/doc/code.html#GOPATH).

    go get github.com/joeytrapp/pgen
    mkdir ~/.projects
    
## Usage

### Create a new project file

Run the command, replacing `<project>` with your desired project name and follow the instructions.

    pgen <project>
    
### Generate a script to automatically load the `tmux` project

    sudo pgen -s > /usr/local/bin/p
    chmod +x /usr/local/bin/p
    
### Start tmux with a project

Run the executable, replacing `<project>` with your previously created project.

    p <project>
