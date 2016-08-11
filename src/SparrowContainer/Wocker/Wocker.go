

package main



import (
    
    "Container"
)



func main() {
    
    container := Container.NewContainerWithConfig("hello","/root/wocker/rootfs","/bin/bash",nil)
    
    container.RunContainer()
    
    
    
}