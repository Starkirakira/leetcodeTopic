function printTree(root: TreeNode): string[][] {
    //Rules: rows = deepth + 1
    //cols = (2^deepth+1) -1
    //root.val in res[0][(n-1)/2]
    //res[r][c],root.left = res[r+1][c-(2^deepth-r-1)],root.right = res[r+1][c+2^(deepth-r-1)]
    //nul => ''
    let deepth = 0
    let deepthFunc = (root: TreeNode | null): number => {
        let h = 0
        if(root?.left) h = Math.max(h, deepthFunc(root.left) + 1)
        if(root?.right) h = Math.max(h, deepthFunc(root.right) + 1)
        return h
    }
    let setFunc = (res: string[][],root: TreeNode,r: number,c: number,height: number) => {
        res[r][c] = root.val.toString()
        if(root.left) setFunc(res, root.left,r + 1, c - (1 << (height - r - 1)), height)
        if(root.right) setFunc(res, root.right,r + 1, c + (1 << (height - r - 1)), height)
    }
    deepth = deepthFunc(root)
    let m = deepth + 1
    let n = (1 << (deepth + 1)) - 1
    let res: Array<Array<string>> = new Array(m).fill(0).map(()=> new Array(n).fill(''))
    setFunc(res, root, 0, Math.floor((n-1) / 2), deepth)
    return res
    

    
}