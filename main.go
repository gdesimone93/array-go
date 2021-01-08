package main

import "fmt"

type Node struct {
	value int
	node  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) insert(v int) {
	n := Node{v, nil}
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	pr :=l.head
	for i := 0; i < l.len; i++ {
		if pr.node == nil{
			pr.node = &n
			l.len++
			return
		}
		pr = pr.node
	}

}



func main() {
	l := LinkedList{}
	l.insert(2)
	l.insert(4)
	l.insert(5)
	fmt.Println(l.head.value)
	fmt.Println(l.head.node.value)
	fmt.Println("ACA ELIMINO")

	l.delete(4)
	fmt.Println(l.head.value)
	fmt.Println(l.head.node.value)
	//fmt.Println(n.node)
}


func (l *LinkedList) findNode(v int) *Node{

	pr :=l.head
	for i := 0; i < l.len; i++ {
		if pr.value == v{
			return pr
		}
		pr = pr.node
	}
	return nil

}

func (l *LinkedList) size() int {
	return l.len
}

func (l *LinkedList) delete(v int) {

	if(l.head.value == v){
		if(l.len==1){
			l.head = nil
		}else{
			l.head = l.head.node.node
		}
		return
	}
	first :=l.head

	for i := 0; i < l.len; i++ {
		if first.node.value == v{
			n := l.findNode(first.node.node.value)
			first.node = n
			l.len--
			return
		}

	}

}