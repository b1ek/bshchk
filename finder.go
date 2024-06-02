package main

import (
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

func find(code string) ([]string, error) {
	r := strings.NewReader(code)
	f, err := syntax.NewParser().Parse(r, "")
	if err != nil {
		return make([]string, 0), err
	}

	var builtins = [...]string{"alias", "bind", "builtin", "caller", "command", "declare", "echo", "enable", "help", "let", "local", "logout", "mapfile", "printf", "read", "readarray", "source", "type", "typeset", "ulimit", "unalias"}
	var deps []string
	syntax.Walk(f, func(node syntax.Node) bool {
		switch x := node.(type) {
		case *syntax.CallExpr:
			for i := range x.Args {
				for _, part := range x.Args[i].Parts {
					switch xx := part.(type) {
					case *syntax.Lit:
						deps = append(deps, xx.Value)
					}
					return true
				}
			}
		}

		return true
	})

	var finished_deps []string

	for _, dep := range deps {
		is_builtin := false
		for _, builtin := range builtins {
			if dep == builtin {
				is_builtin = true
			}
		}

		if !is_builtin {
			finished_deps = append(finished_deps, dep)
		}
	}

	return finished_deps, nil
}
