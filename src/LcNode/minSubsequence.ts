function MinSubsequence(nums: number[]): number[] {
     //[1,4,5]
    // let ar = new Array<number>(5).fill(0)
    // ar = nums.slice(1,2)
    // return ar
    nums.sort((a,b) => b - a)
    let ans = new Array<number>()
    let n = nums.length
    let totalSum = nums.reduce((a,b)=>a+b)
    let cur = 0
    for (let i = 0; i < n; i++) {
        cur += nums[i]
        ans.push(nums[i])
        if(cur > totalSum - cur) {
            break
        }
    }
    return ans

}

console.log(MinSubsequence([1,2,3,4,5]))