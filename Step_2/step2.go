package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run <container> command args
// go run main.go run command args

func main() {
	// Key off of command
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what?")
	}
}

func run() {
	// Print out what we plan to run from arg 2 on.
	fmt.Print("running %v\n", os.Args[2:])

	// Execute what ever we passed in.
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	// Connect Standard I/O so we see whats going on
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Unixs Timeshare System... Give me new namespace.
	cmd.SysProcAttr = &syscall.SysProcAttr {
		Cloneflags: syscall.CLONE_NEWUTS
	}

	// Run the command
	must(cmd.Run())
}

// Error catcher
func must (err error) {
	if err != nil {
		panic(err)
	}
}