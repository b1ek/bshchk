package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
)

var version string // to be set by compiler

var args struct {
	File          string `arg:"positional" help:"if not specified, will read from stdin" default:""`
	Outfile       string `arg:"positional" help:"if not specified, will emit to stdout" default:""`
	Version       bool   `arg:"-v" help:"print version and exit"`
	YieldDepsOnly bool   `arg:"--yield-deps-only" help:"print dependencies as a JSON array and exit" default:false`
	ExposeDeps    bool   `arg:"--expose-deps" help:"expose dependencies to program" default:false`
	DepsVarName   string `arg:"--deps-var-name" help:"override deps variable name" default:"deps"`
	IgnoreShebang bool   `arg:"--ignore-shebang" help:"ignore shebang requirement" default:false`
}

func main() {
	arg.MustParse(&args)
	if args.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	var file []byte
	if args.File != "" {
		f, err := os.ReadFile(args.File)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "File %s does not exist!", args.File)
				os.Exit(1)
			}
			panic(err)
		}
		file = f
	} else {
		data, err := io.ReadAll(os.Stdin)

		if err != nil {
			fmt.Printf("Coudln't read from stdin: %s", err.Error())
			os.Exit(1)
		}
		file = data
	}

	code := string(file)
	found, err := find(code)
	if err != nil {
		panic(err)
	}

	codelines := strings.Split(code, "\n")

	if len(codelines) < 2 {
		fmt.Fprintf(os.Stderr, "The code must have at least two lines!\n")
		os.Exit(3)
	}

	shebang := codelines[0]

	if !((shebang == "#!/bin/bash") || (shebang == "#!/usr/bin/env bash")) && (!args.IgnoreShebang) {
		fmt.Fprintf(os.Stderr, "The code must start with one of those shebangs: #!/bin/bash OR #!/usr/bin/env bash\n")
		os.Exit(2)
	}

	gen := shebang + "\n\n" + gencode(found) + "\n\n" + strings.Join(codelines[1:], "\n")

	if args.Outfile == "" {
		fmt.Printf("%s", gen)
	} else {
		os.WriteFile(args.Outfile, []byte(gen), os.FileMode(0o755))
	}
}
