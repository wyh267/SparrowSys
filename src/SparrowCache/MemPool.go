/*****************************************************************************
 *  file name : MemPool.go
 *  author : Wu Yinghao
 *  email  : wyh817@gmail.com
 *
 *  file description : 内存池
 *
******************************************************************************/

package SparrowCache

import (
	"fmt"
	"time"
)

var MAX_NODE_COUNT int = 1000

type Node struct {
	key   string
	value string
	hash1 uint64
	hash2 uint64
	Next  *Node
}

func NewNode() *Node {
	node := &Node{hash1: 0, hash2: 0, Next: nil}
	return node
}

// SMemPool 内存池对象
type SMemPool struct {
	nodechanGet   chan *Node
	nodechanGiven chan *Node
	nodeList      []Node
	freeList      []*Node
}

func (pool *SMemPool) makeNodeList() error {

	pool.nodeList = make([]Node, MAX_NODE_COUNT)

	return nil
}

func NewMemPool() *SMemPool {

	this := &SMemPool{nodechanGet: make(chan *Node, MAX_NODE_COUNT),
		nodechanGiven: make(chan *Node, MAX_NODE_COUNT),
		nodeList:      make([]Node, MAX_NODE_COUNT),
		freeList:      make([]*Node, 0)}

	return this
}

func (this *SMemPool) Alloc() *Node {

	return <-this.nodechanGet

}

func (this *SMemPool) Free(node *Node) error {

	this.nodechanGiven <- node

	return nil

}

func (this *SMemPool) InitMemPool() error {

	//初始化node

	go func() {
		//q := new(list.List)
		q := make([]Node, MAX_NODE_COUNT)
		for {

			if len(q) == 0 {
				q = append(q, make([]Node, MAX_NODE_COUNT)...)
			}
			e := q[0]
			timeout := time.NewTimer(time.Second)
			select {
			case b := <-this.nodechanGiven:
				timeout.Stop()
				fmt.Printf("Free Buffer...\n")
				//b=b[:MAX_DOCID_LEN]
				q = append(q, *b)
			case this.nodechanGet <- &e:
				timeout.Stop()
				q = q[1:]
				//fmt.Printf("Alloc Buffer...\n")
				//q.Remove(e)

			case <-timeout.C:
				fmt.Printf("remove Buffer...\n")

			}

		}

	}()

	return nil

}
