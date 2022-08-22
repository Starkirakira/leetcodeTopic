function isNum (a: string): boolean {
    return parseFloat(a).toString() == 'NaN' ? false : true
}

console.log(isNum('1'))