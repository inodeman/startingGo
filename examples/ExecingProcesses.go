package examples

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func ExecingProcesses() {
	fmt.Println("**Examples Execing Processes**")

	binary, lookErr := exec.LookPath("ls")

	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}

}
