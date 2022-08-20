function constructMaximumBinaryTree(nums: number[]): TreeNode | null {
    if(nums.length == 0) return null
    
    let construct = (nums: number[]): TreeNode => {
        let index = findBigestIndex(nums)
        let left = nums.slice(0,index), right = nums.slice(index+1)
        let root = new TreeNode(nums[index])
        if(left.length != 0) root.left =  construct(left)
        if(right.length != 0) root.right = construct(right)
        return root
    }
    return construct(nums)
}

function findBigestIndex(arr: number[]): number{
    let max = Math.max(...arr)
    return arr.indexOf(max)
}

//Monotonic stack
