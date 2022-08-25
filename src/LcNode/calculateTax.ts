function calculateTax(brackets: number[][], income: number): number{
    let evt = 0, res = 0
    for(let i = 0; i < brackets.length; i++) {
        i == 0 ? evt = brackets[0][0] : evt = brackets[i][0]-brackets[i-1][0]
        if(income < evt) {
            res += brackets[i][1] * income * 0.01
            break
        }else {
            res += brackets[i][1] * evt * 0.01
            income -= evt
        }
    }
    return Number(res.toFixed(4)) 
}