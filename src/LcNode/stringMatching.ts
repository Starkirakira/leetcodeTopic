function StringMatching(words: string[]): string[] {
    let res = new Array<string>()
    for(let i = 0; i < words.length;i++) {
        for(let j = 0; j < words.length; j++) {
            if (i != j && words[j].search(words[i]) != -1) {
                res.push(words[i])
                break
            }
        }
    }
    return res

}

console.log(StringMatching(["aaa","bb","cc","dd"]))