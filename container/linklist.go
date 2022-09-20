package container

type node struct {
	value any
	next  *node
}

type LinkList struct {
	len  int
	head *node
	tail *node
}

func NewLinkList() LinkList {
	return LinkList{len: 0, head: nil, tail: nil}
}

func (link *LinkList) PushHead(v any) {
	n := &node{value: v, next: nil}
	if link.len == 0 {
		link.tail = n
		link.head = n
	} else {
		n.next = link.head
		link.head = n
	}
	link.len++
}

func (link *LinkList) PushTail(v any) {
	n := &node{value: v, next: nil}
	if link.len == 0 {
		link.head = n
		link.tail = n
	} else {
		link.tail.next = n
		link.tail = n
	}
	link.len++
}

type FilterFunc interface {
	Filter(*node) bool
}

func (link *LinkList) Each(f func(any)) {
	cur := link.head
	for {
		f(cur.value)
		if cur.next != nil {
			cur = cur.next
		} else {
			break
		}
	}
}

func (link *LinkList) Filter(filter func(any) bool) {
	cur := link.head
	var prev *node
	for {
		if filter(cur.value) {
			link.len--
			//中间元素
			if prev != nil && cur.next != nil {
				//跳过自己
				prev.next = cur.next
			} else {
				//最后一个元素被删除
				if cur.next == nil {
					link.tail = prev
					link.tail.next = nil
					cur.next = nil
				} else {
					//第一个元素被删除
					if prev == nil {
						link.head = link.head.next
					}
				}
			}

		}
		if cur.next != nil {
			prev = cur
			cur = cur.next
		} else {
			break
		}
	}

}

func (link *LinkList) Len() int {
	return link.len
}
