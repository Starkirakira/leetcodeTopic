function findClosestElements(arr: number[], k: number, x: number): number[] {
    //Sort the array, if one element - x < another element - x, sort it to the front of the array
    // if(x < arr[0]) return k <= arr.length ? arr.slice(0,k+1) : arr
    // if(x > arr[arr.length-1]) return k <= arr.length ? arr.slice(-k) : arr
    // let list = arr.concat()
    // list.sort((a,b) => {
    //     if(Math.abs(a-x) !== Math.abs(b-x)) {
    //         return Math.abs(a-x) - Math.abs(b-x)
    //     } else {
    //         return a - b 
    //     }
    // })
    // let res = list.slice(0,k)
    // res.sort((a,b) => a - b)
    // return res

    //Binary search  + double pointer
    let binartSearch = (arr: number[],x: number) => {
        let low = 0
        let high = arr.length - 1
        if(low < high) {
            let mid = low + Math.floor((high - low) / 2)
            if(arr[mid] >= x) high = mid
            else low = mid + 1
        }
        return low
    }

    let right = binartSearch(arr,x)
    let left = right - 1
    while(k-- > 0) {
        if(left < 0) right++
        else if (right >= arr.length) left--
        else if(x - arr[left] <= arr[right] - x) left--
        else right++
    }
    let res = []
    for(let i = left + 1; i < right; i++) {
        res.push(arr[i])
    }
    return res
}