package main // import "github.com/sachaos/lsh-go"

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var buildInCmd = map[string]func([]string) error{
	"cd":   lshChdir,
	"exit": lshExit,
}

func lshChdir(args []string) error {
	return os.Chdir(args[1])
}

func lshExit(args []string) error {
	os.Exit(0)
	return nil
}

func lshReadLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func lshSplitLine(line string) []string {
	return strings.Split(line, " ")
}

func lshLaunch(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func lshExecute(args []string) error {
	f, ok := buildInCmd[args[0]]
	if ok {
		return f(args)
	}
	return lshLaunch(args)
}

func lshLoop() {
	for {
		fmt.Printf("> ")
		line := lshReadLine()
		args := lshSplitLine(line)
		err := lshExecute(args)

		if err != nil {
			fmt.Printf("lsh: %v\n", err)
		}
	}
}

func main() {
	lshLoop()
}
