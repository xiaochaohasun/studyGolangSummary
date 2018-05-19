package main

func reverseList(p *Node) *Node {
	var pre *Node
	for p != nil {
		//暂存p的下一个节点
		next := p.Next
		//让当前节点更新为pre
		p.Next = pre
		//更新pre为当前节点
		pre=p
		//继续迭代，变更p为原来的next节点
		p=next
	}
	return pre
}
