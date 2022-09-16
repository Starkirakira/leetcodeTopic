function missingElement(nums: number[], k: number): number {
    let n = nums.length
    let sumDiff: Array<number> = new Array<number>(n).fill(0)
    for(let i = 1; i < n; i++) {
        sumDiff[i] = nums[i] - nums[i-1] +sumDiff[i - 1] - 1
    }
    if(k > sumDiff[n-1]) return nums[n-1] + k - sumDiff[n-1]

    let idx = 1
    while(sumDiff[idx] < k) idx++
    return nums[idx-1] + k - sumDiff[idx-1]
}

console.log(missingElement([1,2,4,6,7,9,10,12,14,18,22],7))