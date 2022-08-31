function validateStackSequences(pushed: number[], poped: number[]): boolean {
    let n = pushed.length
    let validateList: Array<number> = new Array<number>(n)
    for(let i = 0,j = 0; i < n; i++) {
        validateList.push(pushed[i])
        while(validateList.length && validateList[validateList.length - 1] == poped[j]) {
            validateList.pop()
            j++
        }
    }

    return validateList.length == 0

}