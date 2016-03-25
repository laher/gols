package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/laher/gols"
)

func main() {
	var ignoreDirsS = flag.String("ignore", "/vendor/", "ignore packages (comma-delimited)")
	var execs = flag.String("exec", "", "exec (e.g. 'go test')")
	flag.Parse()
	args := flag.Args()
	ignoreDirs := []string{}
	if *ignoreDirsS != "" {
		ignoreDirs = strings.Split(*ignoreDirsS, ",")
	}
	pkgs, err := gols.Ls(args, ignoreDirs, *execs == "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *execs != "" {
		//TODO handle quotes
		execArr := strings.Split(*execs, " ")
		err = gols.Exec(execArr, pkgs)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
