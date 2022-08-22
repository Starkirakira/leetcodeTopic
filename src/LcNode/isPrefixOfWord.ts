function isPrefixOfWord(sentence: string, searchWord: string): number {
    let words = sentence.split(' ')
    for(let i = 0; i < words.length; i++) {
        if(words[i].length < searchWord.length) continue
        let ok = true
        for(let j = 0; j < searchWord.length && ok; j++) {
            if(searchWord[j] != words[i][j]) ok = false
        }
        if(ok) return i + 1
    }
    return -1
}