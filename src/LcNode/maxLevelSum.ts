function MaxLevelSum(root: TreeNode): number {
    const sum: number[] = []
    const dfs = (node: TreeNode, level: number) => {
        if (level == sum.length) {
            sum.push(node.val)
        } else {
            sum.splice(level,1, sum[level] + node.val)
        }
        if (node.left) {
            dfs(node.left, level + 1)
        }
        if (node.right) {
            dfs(node.right, level + 1)
        }
    }
    dfs(root,0)
    let ans = 0
    for (let i = 0; i < sum.length; i++ ) {
        if (sum[i] > sum[ans]) {
            ans = i
        }
    }
    return ans + 1
}