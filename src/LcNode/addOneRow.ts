function AddOneRow(root: TreeNode | null, val: number, depth: number): TreeNode | null {
    if (root == null) return null
    if (depth == 1) return new TreeNode(val,root,null)
    if (depth == 2) {
        root.left = new TreeNode(val,root.left,null)
        root.right = new TreeNode(val,null,root.right)
    } else {
        root.left = AddOneRow(root.left,val,depth-1)
        root.right = AddOneRow(root.right,val,depth-1)
    }
    
    return root
}