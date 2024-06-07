# bshchk (https://git.blek.codes/blek/bshchk)
{{.DepsVar}}=({{.Deps}})
non_ok=()

for d in ${{.DepsVar}}
do
    if ! command -v $d > /dev/null 2>&1; then
        non_ok+=$d
    fi
done

if (( ${#non_ok[@]} != 0 )); then
    >&2 echo "RDC Failed!"
    >&2 echo "  This program requires these commands:"
    >&2 echo "  > ${{.DepsVar}}"
    >&2 echo "    --- "
    >&2 echo "  From which, these are missing:"
    >&2 echo "  > $non_ok"
    >&2 echo "Make sure that those are installed and are present in \$PATH."
fi

unset non_ok{{if .UnsetDeps}}
unset {{.DepsVar}}
{{end}}# Dependencies are OK at this point