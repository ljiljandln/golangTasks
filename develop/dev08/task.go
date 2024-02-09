package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Shell: ")
	for scanner.Scan() {
		line := scanner.Text()
		doShell(line)
		fmt.Print("Shell: ")
	}
}

func doShell(line string) {
	line = strings.TrimSpace(line)
	line = strings.TrimSuffix(line, "\n")
	cmd := strings.Fields(line)

	if len(cmd) == 0 {
		return
	}

	switch cmd[0] {
	case "cd":
		if len(cmd) != 2 {
			return
		}
		os.Chdir(cmd[1])
	case "quit":
		os.Exit(0)
	}

	command := exec.Command("bash", "-c", line)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	command.Run()
}
