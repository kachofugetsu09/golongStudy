package main

import "container/list"

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack = list.New()
	res := []int{}
	stack.PushBack(root)
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		if e.Value == nil {
			e = stack.Back()
			res = append(res, e.Value.(*TreeNode).Val)
			continue
		}
		if e.Value.(*TreeNode).Right != nil {
			stack.PushBack(e.Value.(*TreeNode).Right)
		}
		if e.Value.(*TreeNode).Left != nil {
			stack.PushBack(e.Value.(*TreeNode).Left)
		}
		stack.PushBack(e.Value.(*TreeNode))
		stack.PushBack(nil)
	}
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := list.New() //栈
	res := []int{}      //结果集
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			e = stack.Back() //弹出元素（即中间节点）
			stack.Remove(e)  //删除中间节点
			node = e.Value.(*TreeNode)
			res = append(res, node.Val) //将中间节点加入到结果集中
			continue                    //继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		//压栈顺序：右中左
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return res
}

func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack = list.New() //栈
	res := []int{}         //结果集
	stack.PushBack(root)
	var node *TreeNode
	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		if e.Value == nil { // 如果为空，则表明是需要处理中间节点
			e = stack.Back() //弹出元素（即中间节点）
			stack.Remove(e)  //删除中间节点
			node = e.Value.(*TreeNode)
			res = append(res, node.Val) //将中间节点加入到结果集中
			continue                    //继续弹出栈中下一个节点
		}
		node = e.Value.(*TreeNode)
		//压栈顺序：中右左
		stack.PushBack(node) //中间节点压栈后再压入nil作为中间节点的标志符
		stack.PushBack(nil)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return res
}

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	arr := [][]int{}
	depth := 0
	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}
	order(root, depth)
	return arr
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		levelSize := queue.Len()
		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if i == levelSize-1 {
				res = append(res, node.Val)
			}
		}
	}
	return res
}
