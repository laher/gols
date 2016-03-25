package gols

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"unicode"
)

func Main(execDefault string) {
	var ignoreDirsS = flag.String("ignore", "/vendor/", "ignore packages (comma-delimited)")
	var execs = flag.String("exec", "", "exec (e.g. 'go test')")
	flag.Parse()
	args := flag.Args()
	ignoreDirs := []string{}
	if *ignoreDirsS != "" {
		ignoreDirs = strings.Split(*ignoreDirsS, ",")
	}
	pkgs, err := Ls(args, ignoreDirs, *execs == "")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *execs != "" {
		//TODO handle quotes
		//execArr := strings.Split(*execs, " ")
		execArr := splitQuotedString(*execs)
		err = Exec(execArr, pkgs)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func splitQuotedString(s string) []string {
	lastQuote := rune(0)
	f := func(c rune) bool {
		switch {
		case c == lastQuote:
			lastQuote = rune(0)
			return false
		case lastQuote != rune(0):
			return false
		case unicode.In(c, unicode.Quotation_Mark):
			lastQuote = c
			return false
		default:
			return unicode.IsSpace(c)

		}
	}

	m := strings.FieldsFunc(s, f)
	return m
}

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
