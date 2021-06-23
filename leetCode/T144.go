package leetCode

/*
*  @author liqiqiorz
*  @data 2021/6/23 17:06
 */
func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
