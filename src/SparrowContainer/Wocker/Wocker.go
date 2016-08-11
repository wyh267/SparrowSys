package main

import (
	"SparrowContainer/Container"
)

func main() {

	container := Container.NewContainerWithConfig("hello", "/root/wocker/rootfs/", "/root/start_container.sh", nil)

	container.RunContainer()

}
