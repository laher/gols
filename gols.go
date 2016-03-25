package gols

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func Ls(args []string, ignoreDirs []string, print bool) ([]string, error) {
	cmd := exec.Command("go", "list")
	cmd.Args = append(cmd.Args, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
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
				if print {
					fmt.Println(txt)
				}
				lines = append(lines, txt)
			}
		}
		l.Unlock()
	}()
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	l.Lock()
	defer l.Unlock()
	return lines, err
}

func Exec(execArr []string, pkgs []string) error {

	cmd2 := exec.Command(execArr[0])
	cmd2.Args = execArr
	cmd2.Args = append(cmd2.Args, pkgs...)
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	err := cmd2.Start()
	if err != nil {
		return err
	}
	err = cmd2.Wait()
	if err != nil {
		return err
	}
	return err
}
