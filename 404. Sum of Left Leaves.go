func isLeaf(root * TreeNode) bool{
	if root == nil || root.Left!=nil || root.Right!=nil{
		return false
	}
	return true;
}
func sumOfLeftLeaves(root *TreeNode) int {
    res := 0
	if root != nil{
		if isLeaf(root.Left){
			res += root.Left.Val
		}else{
			res += sumOfLeftLeaves(root.Left)
		}
		res += sumOfLeftLeaves(root.Right)
	}
	return res
}
