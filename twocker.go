package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    switch os.Args[1] {
    case "run":
        parent()
    case "child":
        child()
    default:
        panic("wat should I do")
    }
}

func parent() {
    cmd := exec.Command("/twocker", append([]string{"child"}, os.Args[2:]...)...)
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
    }
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }
}

func child() {
    //(source string, target string, fstype string, flags uintptr, data string)
    //must(syscall.Mount("/root/wocker/rootfs/", "/", "rootfs", , ""))
    //must(os.MkdirAll("rootfs/oldrootfs", 0700))
    //must(syscall.PivotRoot("rootfs", "/root/wocker/rootfs/"))
    //must(os.Chdir("/"))
    
    must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
    must(os.MkdirAll("/root/wocker/rootfs/", 0700))
    must(syscall.PivotRoot("rootfs", "/root/wocker/rootfs/"))
    must(os.Chdir("/"))

    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Println("ERROR", err)
        os.Exit(1)
    }
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}