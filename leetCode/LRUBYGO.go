package leetCode

/*
*  @author liqiqiorz
*  @data 2021/6/22 21:42
 */

type LRUCache struct {
	Size       int
	Cap        int
	Data       map[int]*DLL
	Head, Tail *DLL
}
type DLL struct {
	key, val   int
	prev, next *DLL
}

func initDLinkedNode(key, value int) *DLL {
	return &DLL{
		key: key,
		val: value,
	}
}
func Constructor(capcity int) LRUCache {
	l := LRUCache{
		Data: map[int]*DLL{},
		Head: initDLinkedNode(0, 0),
		Tail: initDLinkedNode(0, 0),
	}
	l.Head.next = l.Tail
	l.Tail.prev = l.Head
	return l
}
func (l *LRUCache) Get(key int) int {
	if node, ok := l.Data[key]; ok {
		l.moveToHead(node)
		return node.val
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if n, ok := l.Data[key]; ok {
		n.val = value
		l.moveToHead(n)
	} else {
		n := initDLinkedNode(key, value)
		l.Data[key] = n
		l.addToHead(n)
		l.Size++
		if l.Size > l.Cap {
			removed := l.removeTail()
			delete(l.Data, removed.key)
			l.Size--
		}
	}
}
func (l *LRUCache) addToHead(n *DLL) {
	n.prev, n.next = l.Head, l.Head.next
	l.Head.next.prev, l.Head.next = n, n
}
func (l *LRUCache) removeNode(n *DLL) {
	n.prev.next, n.next.prev = n.next, n.prev
}
func (l *LRUCache) moveToHead(n *DLL) {
	l.removeNode(n)
	l.addToHead(n)
}

func (l *LRUCache) removeTail() *DLL {
	n := l.Tail.prev
	l.removeNode(n)
	return n
}
