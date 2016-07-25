/*****************************************************************************
 *  file name : HashTable.go
 *  author : Wu Yinghao
 *  email  : wyh817@gmail.com
 *
 *  file description : 哈希表结构，底层数据结构
 *
******************************************************************************/

package SparrowCache



type LinkList struct {
	linklen int
	node    *Node //NodeInterface
    alloc   func() *Node
    free    func(n *Node) error 

}

func (ll *LinkList) getLen() int {
	return ll.linklen
}

func (ll *LinkList) addkv(key, value string) error {
	if ll.linklen == 0 {
		ll.node = NewNode()//ll.alloc()
		ll.node.key = key
		ll.node.value = value
		ll.linklen = 1
		return nil
	}

	p := ll.node
	for p.Next != nil {
        
        if p.key==key {
            p.value = value
            return nil
        }

		p = p.Next
	}

	np := NewNode()//ll.alloc()
	np.key = key
	np.value = value
	p.Next = np
	ll.linklen++
	return nil

}

func (ll *LinkList) findKey(key string) (string, bool) {

	if ll.linklen == 1 && key == ll.node.key {
		return ll.node.value, true
	}

	if ll.linklen == 0 {
		return "", false
	}

	if ll.linklen > 1 {
		p := ll.node
		for p != nil {
			if key == p.key {
				return p.value, true
			}

			p = p.Next
		}

		return "", false
	}

	return "", false

}

// SHashTable 哈希表结构
type SHashTable struct {
	buketNum int64
	HBukets  []*LinkList
}

func NewHashTable() *SHashTable {

	this := &SHashTable{buketNum: 1403641}

	this.HBukets = make([]*LinkList, this.buketNum)

	return this

}

func (this *SHashTable) InitHashTable() error {
    
   // pool:=NewMemPool()
   // pool.InitMemPool()

	for i := range this.HBukets {

		this.HBukets[i] = new(LinkList)
		this.HBukets[i].linklen = 0
		this.HBukets[i].node = nil
     //   this.HBukets[i].alloc=pool.Alloc
     //   this.HBukets[i].free=pool.Free
	}

	return nil

}

func (this *SHashTable) Set(key, value string) error {

	index := ELFHash(key, this.buketNum)

	return this.HBukets[index].addkv(key, value)

}

func (this *SHashTable) Get(key string) (string, bool) {

	index := ELFHash(key, this.buketNum)

	return this.HBukets[index].findKey(key)

}

func ELFHash(str string, bukets int64) int64 {
	var hash int64
	var x int64
	for _, v := range str {

		hash = (hash << 4) + int64(v)
		x = hash
		if (x & 0xF0000000) != 0 {
			hash ^= (x >> 24)
			hash &= ^x
		}
	}
	return (hash & 0x7FFFFFFF) % bukets
}
