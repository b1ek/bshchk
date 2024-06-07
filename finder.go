package main

import (
	"strings"

	"mvdan.cc/sh/v3/syntax"
)

func get_ignored_and_deps(code string) ([]string, []string) {
	var ignored []string = []string{"alias", "bind", "builtin", "caller", "command", "declare", "echo", "enable", "let", "local", "logout", "mapfile", "printf", "read", "readarray", "source", "type", "typeset", "ulimit", "unalias"}
	var deps []string

	for _, line := range strings.Split(code, "\n") {
		splitted := strings.Split(line, " ")
		if len(splitted) < 2 {
			continue
		}
		if splitted[0] == "#bshchk:ignore-cmd" {
			ignored = append(ignored, splitted[1:]...)
		}
		if splitted[0] == "#bshchk:add-cmd" {
			deps = append(deps, splitted[1:]...)
		}
	}

	return ignored, deps
}

func find(code string) ([]string, error) {
	r := strings.NewReader(code)
	f, err := syntax.NewParser().Parse(r, "")
	if err != nil {
		return make([]string, 0), err
	}

	ignored, deps := get_ignored_and_deps(code)

	// 1. find function declarations
	syntax.Walk(f, func(node syntax.Node) bool {
		switch x := node.(type) {
		case *syntax.FuncDecl:
			ignored = append(ignored, x.Name.Value)
		}

		return true
	})

	// 2. collect all commands
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
		for _, builtin := range ignored {
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
