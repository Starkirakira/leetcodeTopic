function groupAnagrams(strs: string[]): string[][] {
    const map = new Map<string,Array<string>>()
    for(let str of strs) {
        let s = [...str].sort().toString()
        let list = map.get(s) ?? new Array<string>()
        list.push(str)
        map.set(s,list)
    }
    return Array.from(map.values())
}
console.log(groupAnagrams(['aaa','bbb','aca','aac']))

