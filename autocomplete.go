package main

func AutocompleteText() string {
	return `# Completion script for pgen
#
# To use, add the following to your .bashrc:
#
#    . pgen_bash_completion.sh
#
# This script assumes that project files are located in ~/.projects and that
# the "p" script is actually named "p".

__pgencomp_words_include ()
{
    local i=1
    while [[ $i -lt $COMP_CWORD ]]; do
        if [[ "${COMP_WORDS[i]}" = "$1" ]]; then
            return 0
        fi
        i=$((++i))
    done
    return 1
}

__pgencomp ()
{
    # break $1 on space, tab, and newline characters,
    # and turn it into a newline separated list of words
    local list s sep=$'\n' IFS=$' '$'\t'$'\n'
    local cur="${COMP_WORDS[COMP_CWORD]}"

    for s in $1; do
        __pgencomp_words_include "$s" && continue
        list="$list$s$sep"
    done

    IFS=$sep
    COMPREPLY=($(compgen -W "$list" -- "$cur"))
}

_pgen ()
{
    local i=1 cmd

    # find the subcommand
    while [[ $i -lt $COMP_CWORD ]]; do
        local s="${COMP_WORDS[i]}"
        case "$s" in
        --*)
            cmd="$s"
            break
            ;;
        -*)
            ;;
        *)
            cmd="$s"
            break
            ;;
        esac
        i=$((++i))
    done

    if [[ $i -eq $COMP_CWORD ]]; then
        local ext=$(\ls -p $HOME/.projects \
                2>/dev/null | sed -e "s/.*\///g")
        __pgencomp "
            -s
            -a
            "
        return
    fi

    # subcommands can have their own completion functions
    case "$cmd" in
    *)                          ;;
    esac
}

_pgen_p ()
{
    local i=1 cmd

    # find the subcommand
    while [[ $i -lt $COMP_CWORD ]]; do
        local s="${COMP_WORDS[i]}"
        case "$s" in
        --*)
            cmd="$s"
            break
            ;;
        -*)
            ;;
        *)
            cmd="$s"
            break
            ;;
        esac
        i=$((++i))
    done

    if [[ $i -eq $COMP_CWORD ]]; then
        local ext=$(\ls -p $HOME/.projects \
                2>/dev/null | sed -e "s/.*\///g")
        __pgencomp "
            $ext
            "
        return
    fi

    # subcommands can have their own completion functions
    case "$cmd" in
    *)                          ;;
    esac
}

complete -o bashdefault -o default -F _pgen pgen
complete -o bashdefault -o default -F _pgen_p p
`
}

