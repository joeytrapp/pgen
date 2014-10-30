# pgen

Create `tmux` project files for easy tmux session management. Project files are currently stored in `~/.projects`. The project file is created with the name of the first arg to `pgen` and then there is a questionnaire for naming the windows and giving them a relative path.

## Installation

### Direct Download

Download the [latest release](https://github.com/joeytrapp/pgen/releases/latest), put it in your path, and make it executable. Then create your projects folder.

    mkdir ~/.projects

### Using Go

You can optionally install `pgen` using Go. If you haven't yet, [install Go](https://golang.org/doc/install) and make sure to [setup $GOPATH](https://golang.org/doc/code.html#GOPATH).

    go get github.com/joeytrapp/pgen
    mkdir ~/.projects
    
## Usage

### Create a new project file

Run the command, replacing `<project>` with your desired project name and follow the instructions.

    pgen <project>
    
### Generate a script to easily load `tmux` project files

    sudo pgen -s > /usr/local/bin/p
    chmod +x /usr/local/bin/p
    
### Start tmux with a project

Run the executable, replacing `<project>` with your previously created project. You can even run this from inside existing `tmux` sessions.

    p <project>
