package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	err := lsdirs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func lsdirs(args []string) error {
	var ignoreDirsS = flag.String("ignore", "/vendor/", "ignore packages (comma-delimited)")
	var execs = flag.String("exec", "", "exec (e.g. 'go test')")
	flag.Parse()
	cmd := exec.Command("go", "list")
	args = flag.Args()
	cmd.Args = append(cmd.Args, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	ignoreDirs := []string{}
	if *ignoreDirsS != "" {
		ignoreDirs = strings.Split(*ignoreDirsS, ",")
	}
	scanner := bufio.NewScanner(stdout)
	lines := []string{}
	l := sync.Mutex{}
	go func() {
		l.Lock()
		for scanner.Scan() {
			txt := scanner.Text()
			ignore := false
			for _, i := range ignoreDirs {
				if strings.Contains(txt, i) {
					ignore = true
					break
				}
			}
			if !ignore {
				if *execs == "" {
					fmt.Println(txt)
				} else {
					lines = append(lines, txt)
				}
			}
		}
		l.Unlock()
	}()
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	err = scanner.Err()
	if err != nil {
		return err
	}
	l.Lock()
	defer l.Unlock()
	if *execs != "" {

		//TODO quotes
		execArr := strings.Split(*execs, " ")
		cmd2 := exec.Command(execArr[0])
		cmd2.Args = execArr
		cmd2.Args = append(cmd2.Args, lines...)
		cmd2.Stdout = os.Stdout
		cmd2.Stderr = os.Stderr
		err = cmd2.Start()
		if err != nil {
			return err
		}
		err = cmd2.Wait()
		if err != nil {
			return err
		}
	}
	return err
}
