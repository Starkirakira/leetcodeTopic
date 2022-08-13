function MaxChunksToSorted(arr: number[]): number {
    let sortedArr = arr.concat()
    let mapCount = new Map()
    let res = 0
    sortedArr.sort((a, b) => a - b)
    for(let i = 0; i < arr.length; i++) {
        //mapCount.set(arr[i],(mapCount.get(arr[i]) || 0) + 1)
        if(mapCount.has(arr[i])) mapCount.set(arr[i],mapCount.get(arr[i]) + 1)
        else mapCount.set(arr[i],1)
        if(mapCount.get(arr[i]) == 0) mapCount.delete(arr[i])
        //mapCount.set(sortedArr[i],(mapCount.get(sortedArr[i]) || 0) - 1)
        if(mapCount.has(sortedArr[i])) mapCount.set(sortedArr[i],mapCount.get(sortedArr[i]) - 1)
        //If not have the key, put the key => value sortedArr[i] => -1, means that the element in sortedArr now dosn't appear in arr, so it is -1
        else mapCount.set(sortedArr[i],-1)
        if(mapCount.get(sortedArr[i]) == 0) mapCount.delete(sortedArr[i])
        if(mapCount.size == 0) res++
        console.log(mapCount)
    }
    return res

    //Better way(maybe), but more difficult to understand
    // const cnt = new Map();
    // let res = 0;
    // const sortedArr = new Array(arr.length).fill(0);
    // sortedArr.splice(0, arr.length, ...arr);
    // sortedArr.sort((a, b) => a - b);
    // for (let i = 0; i < sortedArr.length; i++) {
    //     const x = arr[i], y = sortedArr[i];
    //     cnt.set(x, (cnt.get(x) || 0) + 1);
    //     if (cnt.get(x) === 0) {
    //         cnt.delete(x);
    //     }
    //     cnt.set(y, (cnt.get(y) || 0) - 1);
    //     if (cnt.get(y) === 0) {
    //         cnt.delete(y);
    //     }
    //     if (cnt.size === 0) {
    //         res++;
    //     }
    //     console.log(cnt)
    // }
    // return res;
    
}

MaxChunksToSorted([1,2,1])