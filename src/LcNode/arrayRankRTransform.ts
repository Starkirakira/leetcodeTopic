function ArrayRankTransform(arr: number[]): number[] {
    //Deep clone to make sure the arr Array will not be changed
    let sortedArr: number[] = arr.concat()
    //First sort
    sortedArr.sort((a, b) =>a - b )
    let sortMap = new Map()
    let count: number = 1
    //Use map restore the key's position, if a key already exist, then the count skip plus 1
    for (let i = 0; i < sortedArr.length; i++) {
        if (sortMap.get(sortedArr[i]) == undefined) {
            sortMap.set(sortedArr[i], count)
            count++
        } 
        
    }
    let result: number[] = new Array(arr.length)
    for (let i = 0; i < arr.length; i++) {
        result[i] = sortMap.get(arr[i])
    }

    return result
}

console.log(ArrayRankTransform([10,1,3,2,5,10,12,12,14,13]))