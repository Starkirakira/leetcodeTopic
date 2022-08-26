const N = 26
const p: number[] = new Array<number>(N), sz = new Array<number>(N)

function find(x: number): number {
    if (p[x] != x) p[x] = find(p[x])
    return p[x]
}

function union(a: number, b: number): void {
    if (find(a) == find(b)) return
    sz[find(a)] += sz[find(b)]
    p[find(b)] = p[find(a)]
}


function equationsPossible(equations: string[]): boolean {
    for(let i = 0; i < 26; i++) {
        p[i] = i,sz[i] = 1
    }
    for (let v of equations) {
        if(v[1] == '=') {
            union(v.charCodeAt(0) - 'a'.charCodeAt(0), v.charCodeAt(3) - 'a'.charCodeAt(0))
        }
    }

    for(let v of equations) {
        if(v[1] == '!') {
            if(find(v.charCodeAt(0) - 'a'.charCodeAt(0)) == find(v.charCodeAt(3) - 'a'.charCodeAt(0))) return false
        }
    }
    return true
}

console.log(equationsPossible(['a==b','a==c','b!=c']))
