// function NumDupDigitsAtMostN(n: number): number {
//     let ans = 0
//     for (let i = 1; i < n+1; i++) {
//         if(isRepeat(i) == true) ans++
//     }
//     return ans
// }

// function isRepeat(n: number):boolean{
//     let digitSet: Set<number> = new Set()
//     let sNumber = n.toString()
//     for(let i = 0; i < sNumber.length;i++) {
//         digitSet.add(+sNumber.charAt(i))
//     }
//     return digitSet.size < sNumber.length? true: false
// }

// 首先我们分析处理的数据量 N \in [1, 10^9]N∈[1,10 
//     9
//      ] ，很明显我们无法去直接一个个数字处理（超时原因）。
    
//     本题要求求有重复数字的正整数个数，其实就是对该可能的数字的排列组合。
    
//     如果我们直接排列重复数字的话，我们需要考虑的情况非常复杂：
    
//     包含 2 个重复数字
//     包含 3 个重复数字
//     ... 包含 k 个重复数字
//     另外我们还需要考虑对于 NN 各个位数上面的限制。
    
//     其实，我们可以反过来思考，假设具有至少 1 位重复数字的正整数个数为 total，则没有重复数字的正整数的个数为 N - totalN−total。
//     而在一个可选的集合（集合内数字不重复）给出所有的排列方式更容易得到，使用组合数学公式即可。
    
//     这里总体思路为，设剩下的位置数为 i, 剩下的数字总数为 j,则不重复数字就是 A_j^iA 
//     j
//     i
//     ​
//       种可能
    
//     这里假设 N= x_1 x_2x_3...x_kN=x 
//     1
//     ​
//      x 
//     2
//     ​
//      x 
//     3
//     ​
//      ...x 
//     k
//     ​
//      。
    
//     我们先考虑位数少于 kk 的情况（也就是高位为 0 的情况）
//     假设我们目前第 ii 位开始非 0（表示 x_1...x_{i-1} 都是 0x 
//     1
//     ​
//      ...x 
//     i−1
//     ​
//      都是0），则该种情况包含的所有不重复位数数字的个数可以通过组合数学求出：
    
//     第 ii 位可以选取 [1, 9][1,9]，x_{i+1}...x_kx 
//     i+1
//     ​
//      ...x 
//     k
//     ​
//      可以选取 [0, 9][0,9] 剩下的，也就是 9 \times A_9^{k-i}9×A 
//     9
//     k−i
//     ​
     
//     因此，我们遍历计算所有的第 ii 非 0 情况即可。
    
//     没有高位 0 的情况，也就是总位数为 kk，这种时候剩下的数字范围 jj 会受到 NN 的位数上数字的影响，比如当前的位数数字为 nn，则j \in [0, n]j∈[0,n]。
//     同样假设我们当前处理到第 ii 位，则剩下的低位可选数字总数为 A_{10-(k-i)}^iA 
//     10−(k−i)
//     i
//     ​
     
    
//     具体例子如下，求 4567 的不重复个数：
    
//     对于情况 1 ：
    
//     4th 3th 2th 1th total
//      0   0   0  1-9 9*A(9,0)
//      0   0  1-9 0-9 9*A(9,1)
//      0  1-9 0-9 0-9 9*A(9,2)
//     对于情况 2：
    
//     4th 3th 2th 1th total
//     1-3 0-9 0-9 0-9 4*A(10-1,3)
//      4  0-4 0-9 0-9 5*A(10-2,2)
//      4   5  0-5 0-9 6*A(10-3,1)
//      4   5   6  0-6 7*A(10-4,0)
//      4   5   6   7  1
//     对于第二种情况，total的是包含部分重复的情况，也就是第 ii 位的选择不能和前面已定的位数数字重复，因此我们需要剔除掉。
    
//     另外需要注意的是，最高位是从 1 开始的，而其他位是 从 0 开始的。
    
    function numDupDigitsAtMostN(n: number): number {
        let digits = []
        let N1 = n
        let total = 0
        let used = new Array(10).fill(0)
        while(N1 > 0) {
            digits.push(N1 % 10)
            N1 = Math.floor(N1 / 10)
        }
        let k = digits.length

        for(let i = k - 1; i > 0; i--) {
            total += 9 * A(9, i - 1)
        }
        
        for(let i = k - 1; i >= 0; i--) {
            let num = digits[i]
            for(let j = (i == k - 1? 1 : 0); j < num; j++) {
                if(used[j] > 0) continue
                total += A(10 - (k - i) , i)
            }
            used[num]++
            if(used[num] > 1) break
            if(i == 0) total++
        }
        return n - total
    }

    function A (i:number,j:number):number {
        return f(i) / f(i - j)
    }

    function f(i:number):number {
        if(i == 1 || i == 0) return 1
        return i * f(i-1)
    }


  