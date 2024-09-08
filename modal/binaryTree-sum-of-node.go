package modal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type WalkThru struct {
	Curr int
	Next string
	// ChildOf int
}
