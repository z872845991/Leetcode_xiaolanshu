/*
143. 重排链表
给定一个单链表L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
func reorderList(head *ListNode)  {
	s:=[]*ListNode{}
	temps:=head
	for temps!=nil{
		s=append(s,temps)
		temps=temps.Next
	}
	i:=0
	j:=len(s)-1
	//t2:=head
	for ;i<j;{
		/*
		t2.Next=s[i]
		t2=t2.Next
		if i==j{
			break
		}
		t2.Next=s[j]
		t2=t2.Next
		i++
		j--
		*/
		s[i].Next=s[j]
		i++
		if i==j{
			break
		}
		s[j].Next=s[i]
		j--
	}
	s[i].Next=nil
}
func main() {
	var head=new(ListNode)
	head.Next=nil
	p:=head
	for i:=0;i<4;i++{
		var tmps=new(ListNode)
		tmps.Next=nil
		tmps.Val=i+1
		p.Next=tmps
		p=p.Next
	}
	reorderList(head.Next)
	p1:=head.Next
	for p1!=nil{
		fmt.Println(p1.Val)
		p1=p1.Next
	}
}
