function DeepestLeavesSum(root: TreeNode | null): number{
    let sum = 0,maxLevel = -1
    let leavesSum = (root: TreeNode | null,level:number) => {
        if(!root) return
        if(level > maxLevel) {
            maxLevel = level
            sum = root.val
        } else if(level == maxLevel) sum += root.val
        leavesSum(root.left,level + 1)
        leavesSum(root.right,level + 1)
    }
    leavesSum(root,0)
    return sum

}