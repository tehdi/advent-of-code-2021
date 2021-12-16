// Stack from https://github.com/golang-collections/collections/blob/604e922904d3/stack/stack.go
type (
    Stack struct {
        top *Node
        length int
    }
    Node struct {
        value *Cave
        prev *Node
    }
)
func NewStack() *Stack {
    return &Stack{nil,0}
}
func (this *Stack) Len() int {
    return this.length
}
func (this *Stack) Peek() *Cave {
    return this.top.value
}
func (this *Stack) Pop() *Cave {
    n := this.top
    this.top = n.prev
    this.length--
    return n.value
}
func (this *Stack) Push(cave *Cave) {
    n := &Node{cave,this.top}
    this.top = n
    this.length++
}
