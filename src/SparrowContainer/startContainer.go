package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {

	var rootfs string
	var hostname string
	var process string
	flag.StringVar(&rootfs, "f", "", "文件系统")
	flag.StringVar(&hostname, "h", "hello_world", "主机名称")
	flag.StringVar(&process, "p", "/bin/sh", "运行进程")
	flag.Parse()

	if err := runProcess(process, rootfs, hostname); err != nil {

		fmt.Printf("run process error : %v", err)
	}

}

func runProcess(processname, rootfs, hostname string) error {

	syscall.Sethostname([]byte(hostname))

	//挂载根
	//must(syscall.Mount("/root/wocker/rootfs", "rootfs", "", syscall.MS_BIND, ""))
	pivotBaseDir := "/"
	tmpMountPoint := "/tmp/"
	if err := syscall.Mount(rootfs, tmpMountPoint, "", syscall.MS_BIND, ""); err != nil {
		return err
	}

	procpath := filepath.Join(rootfs, "proc")
	tmpMountPointProc := filepath.Join(tmpMountPoint, "proc")
	if err := syscall.Mount(procpath, tmpMountPointProc, "proc", 0, ""); err != nil {
		return err
	}

	tmpDir := filepath.Join(tmpMountPoint, pivotBaseDir)
	os.MkdirAll(tmpDir, 0755)
	pivotDir, err := ioutil.TempDir(tmpDir, ".tmpmount")
	if err != nil {
		return err
	}

	if err := syscall.PivotRoot(tmpMountPoint, pivotDir); err != nil {

		fmt.Printf("root err :%v\n", err)
		return err
	}

	if err := os.Chdir("/"); err != nil {
		return err
	}

	fmt.Printf("pivotDir : %v \n ", pivotDir)
	pivotDir = filepath.Join(pivotBaseDir, filepath.Base(pivotDir))

	fmt.Printf("pivotDir : %v \n ", pivotDir)
	if err := syscall.Mount("", pivotDir, "", syscall.MS_PRIVATE|syscall.MS_REC, ""); err != nil {
		fmt.Printf("mout err : %v", err)
		return err
	}

	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		fmt.Printf("unmount pivot_root dir %s", err)
		return err
	}
	if err := os.Remove(pivotDir); err != nil {
		return err
	}
	//配置网络

	//fmt.Printf("%v  \n", os.Args[1])

	fmt.Printf("command : %v ", processname)

	processCmd := strings.Split(processname, " ")

	var cmd *exec.Cmd
	if len(processCmd) < 2 {
		cmd = exec.Command(processname)
	} else {
		cmd = exec.Command(processCmd[0], processCmd[1:]...)
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		return err
	}
	return nil
}
