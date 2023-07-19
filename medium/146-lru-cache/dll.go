package LRUCache

type QNode struct {
	Key  int
	Val  int
	Prev *QNode
	Next *QNode
}

type LRUCache struct {
	collection map[int]*QNode
	capacity   int
	freqRoot   *QNode
	freqTail   *QNode
}

func Constructor(capacity int) LRUCache {
	c := make(map[int]*QNode)
	dr := &QNode{
		Key:  -1,
		Val:  -1,
		Prev: nil,
	}
	dt := &QNode{
		Key:  -1,
		Val:  -1,
		Prev: dr,
		Next: nil,
	}
	dr.Next = dt
	return LRUCache{
		collection: c,
		capacity:   capacity,
		freqRoot:   dr,
		freqTail:   dt,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.collection[key]

	if !ok {
		return -1
	}

	removeNode(v)
	newNode := &QNode{
		Key: v.Key,
		Val: v.Val,
	}
	this.addNodeAtFront(newNode)
	this.collection[key] = newNode

	return v.Val
}

func removeNode(n *QNode) {
	tFront := n.Prev
	tNext := n.Next

	tFront.Next = tNext
	tNext.Prev = tFront

	n.Prev = nil
	n.Next = nil
}

func (this *LRUCache) addNodeAtFront(n *QNode) {
	tNext := this.freqRoot.Next

	this.freqRoot.Next = n
	n.Prev = this.freqRoot

	n.Next = tNext
	tNext.Prev = n
}

func (this *LRUCache) Put(key int, value int) {
	newRoot := &QNode{
		Key: key,
		Val: value,
	}
	v, exists := this.collection[key]
	if exists {
		removeNode(v)
		this.addNodeAtFront(newRoot)
		this.collection[key] = newRoot
	} else {
		if len(this.collection) == this.capacity {
			tbd := this.freqTail.Prev
			removeNode(this.freqTail.Prev)
			delete(this.collection, tbd.Key)

			this.addNodeAtFront(newRoot)
			this.collection[key] = newRoot
		} else {
			this.addNodeAtFront(newRoot)
			this.collection[key] = newRoot
		}
	}
}

/**
* Your LRUCache object will be instantiated and called as such:
* obj := Constructor(capacity);
* param_1 := obj.Get(key);
* obj.Put(key,value);
 */
