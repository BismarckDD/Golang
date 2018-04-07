package list

// Node is an node of a linked list.
type Node struct {
    // Next and previous pointers in the doubly-linked list of nodes.
    // To simplify the implementation, internally a list l is implemented
    // as a ring, such that &p_list.head is both the next node of the last
    // list node (p_list.Back()) and the previous node of the first list
    // node (p_list.Front()).
    next, prev *Node

    // The list to which this node belongs.
    list *List

    // The value stored with this node.
    Value interface{}
}

// Next returns the next list node or nil.
func (p_node *Node) Next() *Node {
    if ptr := p_node.next; p_node.list != nil && ptr != &p_node.list.head {
        return ptr
    }
    return nil
}

// Prev returns the previous list node or nil.
func (p_node *Node) Prev() *Node {
    if ptr := p_node.prev; p_node.list != nil && ptr != &p_node.list.head {
        return ptr
    }
    return nil
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
    head Node // sentinel list node, only &head, head.prev, and head.next are used
    len  int     // current list length excluding (this) sentinel node
}

// Init initializes or clears list l.
func (p_list *List) Init() *List {
    p_list.head.next = &p_list.head
    p_list.head.prev = &p_list.head
    p_list.len = 0
    return p_list
}

// New returns an initialized list.
func New() *List { return new(List).Init() }

// Len returns the number of nodes of list l.
// The complexity is O(1).
func (p_list *List) Len() int { return p_list.len }

// Front returns the first node of list l or nil if the list is empty.
func (p_list *List) Front() *Node {
    if p_list.len == 0 {
        return nil
    }
    return p_list.head.next
}

// Back returns the last node of list l or nil if the list is empty.
func (p_list *List) Back() *Node {
    if p_list.len == 0 {
        return nil
    }
    return p_list.head.prev
}

// lazyInit lazily initializes a zero List value.
func (p_list *List) lazyInit() {
    if p_list.head.next == nil {
        p_list.Init()
    }
}

// insert inserts p_node after p_curr, increments p_list.len, and returns p_node.
func (p_list *List) insert(p_node, p_curr *Node) *Node {
    temp := p_curr.next
    p_curr.next = p_node
    p_node.prev = p_curr
    p_node.next = temp
    temp.prev = p_node
    p_node.list = p_list
    p_list.len++
    return p_node
}

// insertValue is a convenience wrapper for insert(&Node{Value: p_value}, p_curr).
func (p_list *List) insertValue(p_value interface{}, p_curr *Node) *Node {
    return p_list.insert(&Node{Value: p_value}, p_curr)
}

// remove removes p_node from its list, decrements p_list.len, and returns p_node.
func (p_list *List) remove(p_node *Node) *Node {
    p_node.prev.next = p_node.next
    p_node.next.prev = p_node.prev
    p_node.next = nil // avoid memory leaks
    p_node.prev = nil // avoid memory leaks
    p_node.list = nil
    p_list.len--
    return p_node
}

// Remove removes e from l if e is an node of list l.
// It returns the node value e.Value.
// The node must not be nil.
func (p_list *List) Remove(p_node *Node) interface{} {
    if p_node.list == p_list {
        // if p_node.list == l, l must have been initialized when e was inserted
        // in l or l == nil (e is a zero Node) and l.remove will crash
        p_list.remove(p_node)
    }
    return p_node.Value
}

// PushFront inserts a new node e with value v at the front of list l and returns e.
func (p_list *List) PushFront(v interface{}) *Node {
    p_list.lazyInit()
    return p_list.insertValue(v, &p_list.head)
}

// PushBack inserts a new node e with value v at the back of list l and returns e.
func (p_list *List) PushBack(v interface{}) *Node {
    p_list.lazyInit()
    return p_list.insertValue(v, p_list.head.prev)
}

// InsertBefore inserts a new node p_node with value v immediately before p_node and returns e.
// If p_node is not an node of l, the list is not modified.
// The p_node must not be nil.
func (p_list *List) InsertBefore(p_value interface{}, p_node *Node) *Node {
    if p_node.list != p_list {
        return nil
    }
    // see comment in List.Remove about initialization of l
    return p_list.insertValue(p_value, p_node.prev)
}

// InsertAfter inserts a new node e with value v immediately after p_node and returns e.
// If p_node is not an node of l, the list is not modified.
// The p_node must not be nil.
func (p_list *List) InsertAfter(v interface{}, p_node *Node) *Node {
    if p_node.list != p_list {
        return nil
    }
    // see comment in List.Remove about initialization of l
    return p_list.insertValue(v, p_node)
}

// MoveToFront moves node e to the front of list l.
// If e is not an node of l, the list is not modified.
// The node must not be nil.
func (p_list *List) MoveToFront(p_node *Node) {
    if p_node.list != p_list || p_list.head.next == p_node {
        return
    }
    // see comment in List.Remove about initialization of l
    p_list.insert(p_list.remove(p_node), &p_list.head)
}

// MoveToBack moves node e to the back of list l.
// If e is not an node of l, the list is not modified.
// The node must not be nil.
func (p_list *List) MoveToBack(p_node *Node) {
    if p_node.list != p_list || p_list.head.prev == p_node {
        return
    }
    // see comment in List.Remove about initialization of l
    p_list.insert(p_list.remove(p_node), p_list.head.prev)
}

// MoveBefore moves node e to its new position before p_node.
// If e or p_node is not an node of l, or e == p_node, the list is not modified.
// The node and p_node must not be nil.
func (p_list *List) MoveBefore(p_node, p_curr *Node) {
    if p_curr.list != p_list || p_curr == p_node || p_node.list != p_list {
        return
    }
    p_list.insert(p_list.remove(p_node), p_curr.prev)
}

// MoveAfter moves node e to its new position after p_node.
// If e or p_node is not an node of l, or e == p_node, the list is not modified.
// The node and p_node must not be nil.
func (p_list *List) MoveAfter(p_node, p_curr *Node) {
    if p_curr.list != p_list || p_curr == p_node || p_node.list != p_list {
        return
    }
    p_list.insert(p_list.remove(p_node), p_curr)
}

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (p_list *List) PushBackList(other *List) {
    p_list.lazyInit()
    for i, e := other.Len(), other.Front(); i > 0; i, e = i - 1, e.Next() {
        p_list.insertValue(e.Value, p_list.head.prev)
    }
}

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (p_list *List) PushFrontList(other *List) {
    p_list.lazyInit()
    for i, e := other.Len(), other.Back(); i > 0; i, e = i - 1, e.Prev() {
        p_list.insertValue(e.Value, &p_list.head)
    }
}
