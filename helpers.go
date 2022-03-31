package main

import (
	"os"
	"os/exec"
	"runtime"
)

/*
	function to check on the os
	where the app is running in.
*/
func clearScreen() {
	os := runtime.GOOS
	switch os {
	case "windows":
		clearStdOut("cls")
	case "darwin":
		clearStdOut("clear")
	default:
		clearStdOut("clear")
	}
}

func clearStdOut(command string) {
	cmd := exec.Command(command)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
