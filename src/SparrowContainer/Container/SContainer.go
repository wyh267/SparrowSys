/*****************************************************************************
 *  file name : SContainer.go
 *  author : Wu Yinghao
 *  email  : wyh817@gmail.com
 *
 *  file description : 容器
 *
******************************************************************************/

package Container

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

var CONTAINER_PROCESS string = "./SparrowContainer"

type NetInfo struct {
	Dev     string
	Mac     string
	Ip      string
	Netmask string
}

type SContainer struct {
	Name        string
	ID          string
	Rootfs      string
	Net         []NetInfo
	InitProcess string
}

// NewContainer function description : 新建容器
// params :
// return :
func NewContainer(name string) *SContainer {

	this := &SContainer{Name: name, Net: make([]NetInfo, 0)}

	return this
}

// NewContainerWithConfig function description : 根据配置新建容器
// params :
// return :
func NewContainerWithConfig(name, rootfs, initprocess string, netinfo []NetInfo) *SContainer {

	this := &SContainer{Name: name, Rootfs: rootfs, InitProcess: initprocess, Net: netinfo}

	return this
}

func (this *SContainer) SetRootFS(rootfs string) error {

	this.Rootfs = rootfs
	return nil

}

// RunContainer function description : 运行容器
// params :
// return :
func (this *SContainer) RunContainer() error {

	args := make([]string, 0)

	//配置hostname
	hostname := fmt.Sprintf("-h=%v", this.Name)
	args = append(args, hostname)

	//配置根文件系统
	rootfsArg := fmt.Sprintf("-f=%v", this.Rootfs)
	args = append(args, rootfsArg)

	//配置启动进程
	startProcess := fmt.Sprintf("-p=%v", this.InitProcess)
	args = append(args, startProcess)
	//配置网络

	cmd := exec.Command(CONTAINER_PROCESS, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS, // | syscall.CLONE_NEWNET,
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
