package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("insufficient args")
	}

	img, err := NewImage(os.Args[2])
	if err != nil {
		log.Fatalf("failed to load image: %v", err)
	}

	switch os.Args[1] {
	case "run":
		configProcess(img)
	case "actual":
		actualProcess(img)
	default:
		panic("invalid arg, must be (run, actual)")
	}
}

func configProcess(img *Image) {
	cmd := exec.Command("/proc/self/exe", "actual", img.Cmd)
	attachIO(cmd)
	createNameSpace(cmd)
	createControlGroup(cmd, img)

	if err := cmd.Run(); err != nil {
		log.Fatalf("from config: %v", err)
	}
}

func actualProcess(img *Image) {
	err := syscall.Sethostname([]byte(img.Hostname))
	if err != nil {
		log.Fatalf("can't set hostname: %v", err)
	}

	cmd := exec.Command(img.Cmd)
	attachIO(cmd)

	if err := cmd.Run(); err != nil {
		log.Fatalf("from actual: %v", err)
	}
}

func attachIO(c *exec.Cmd) {
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
}

func createNameSpace(c *exec.Cmd) {
	c.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
		Unshareflags: syscall.CLONE_NEWNS,
	}
}

func createControlGroup(c *exec.Cmd, img *Image) {
	cgroups := "/sys/fs/cgroup"

	opts := map[string]string{
		"pids":   img.MaxPids,
		"memory": img.MaxMem,
	}

	for opt, v := range opts {
		path := filepath.Join(cgroups, opt, "dev")
		os.MkdirAll(path, 0777)
		must(os.WriteFile(filepath.Join(path, opt+".max"), []byte(v), 0777))
		must(os.WriteFile(filepath.Join(path, "notify_on_release"), []byte("1"), 0777))
		must(os.WriteFile(filepath.Join(path, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0777))
	}
}

func must(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
