package main

import (
	"os"
	"os/exec"
	"syscall"
	"log"
)

func main() {
	switch os.Args[1] {
	case "run":
		img := os.Args[2]
		configProcess(img)
	case "actual":
		img := os.Args[2]
		actualProcess(img)
	default:
		panic("insufficient options")
	}
}

func configProcess(img string) {
	cmd := exec.Command("/proc/self/exe", "actual", img)
	attachIO(cmd)
	createNameSpace(cmd)
	cmd.Run()
}

func attachIO(c *exec.Cmd) {
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
}

func createNameSpace(c *exec.Cmd) {
	syscall.Sethostname([]byte("container"))

	c.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
		Unshareflags: syscall.CLONE_NEWNS,
	}
}

func createControlGroup(c *exec.Cmd) { }

func actualProcess(img string) {
	log.Println(img)
	cmd := exec.Command(img)
	attachIO(cmd)
	cmd.Run()
}
