package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	err := lsdirs(os.Args)
	if err != nil {
		os.Exit(1)
	}
}

func lsdirs(args []string) error {
	cmd := exec.Command("go", "list", "./...")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			txt := scanner.Text()
			if !strings.Contains(txt, "/vendor/") {
				fmt.Println(txt)
			}
		}
	}()
	err = cmd.Start()
	if err != nil {
		return err
	}
	//log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		return err
	}

	err = scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return err
}
