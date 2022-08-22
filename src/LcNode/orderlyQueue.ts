function OrderlyQueue (s: string, k: number): string {
    //If k equal 1, then it will return the smallest lexicographical sort
    if (k == 1) {
        let ans = s
        for(let i = 0; i < s.length - 1; i++){
            let n = s.length
            s = s.substring(1,n) + s[0]
            ans = ans < s ? ans : s
        }
        return ans
    }
    //If k > 1, whatever the string is, the string can switch itself to the sorted string, and this is the answer.
    //Three ways to make a string to a sorted string
    //return [...s].sort().join('')
    //return Array.from(s).sort().join('')
    return s.split('').sort().join('')
}

console.log(OrderlyQueue('abcdefg',2) > 'abcdefghi')