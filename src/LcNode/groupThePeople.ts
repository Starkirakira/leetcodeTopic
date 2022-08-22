function GroupThePeople(groupSizes: number[]): number[][] {
    let n = groupSizes.length
    let groupMap = new Map<number,Array<number>>()
    let res: Array<Array<number>> = new Array<Array<number>>()
    
    
    for(let i = 0; i < n; i++) {
        if(groupMap.has(groupSizes[i])) groupMap.get(groupSizes[i])?.push(i)
        else groupMap.set(groupSizes[i],[i])
    }
    console.log(groupMap)

    groupMap.forEach((v,k) => {
        if(v.length > 1) {
            let n = v.length / k //need split to number n
            //Waht i need know is that how much nums should be in one group
            //So it need k numbers in a group, just push these group to the return array
            for(let i = 0, len = v.length; i < len; i += k) {
                res.push(v.slice(i,i+k))
            }
            
        }else res.push([v[0]])
    })
    return res
}
//Stupid solve way

console.log(GroupThePeople([2,2,1,1,1,1,1,1]))

//Better way

function GroupThePeopleBetter(groupSizes: number[]): number[][]{
    const groups = new Map();
    const n = groupSizes.length;
    for (let i = 0; i < n; i++) {
        const size = groupSizes[i];
        if (!groups.has(size)) {
            groups.set(size, []);
        }
        groups.get(size).push(i);
    }
    const groupList = [];
    for (const [size, people] of groups.entries()) {
        const groupCount = Math.floor(people.length / size);
        for (let i = 0; i < groupCount; i++) {
            const group = [];
            const start = i * size;
            for (let j = 0; j < size; j++) {
                group.push(people[start + j]);
            }
            groupList.push(group);
        }
    }
    return groupList;

}