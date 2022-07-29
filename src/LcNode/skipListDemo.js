"use strict";
const level = 10;
class TNode {
    constructor(_val) {
        this.ne = new Array(level);
        this.val = _val;
    }
}
class SkiplistNode {
    constructor() {
        this.he = new TNode(-1);
    }
    find(t, ns) {
        let cur = this.he;
        for (let i = level - 1; i >= 0; i--) {
            while (cur.ne[i] != null && cur.ne[i].val < t)
                cur = cur.ne[i];
            ns[i] = cur;
        }
    }
    search(t) {
        let ns = new Array(level);
        this.find(t, ns);
        return ns[0].ne[0] != null && ns[0].ne[0].val == t;
    }
    add(t) {
        let ns = new Array(level);
        this.find(t, ns);
        const node = new TNode(t);
        for (let i = 0; i < level; i++) {
            node.ne[i] = ns[i].ne[i];
            ns[i].ne[i] = node;
            if (Math.round(Math.random()) == 0)
                break;
        }
    }
    erase(t) {
        let ns = new Array(level);
        this.find(t, ns);
        const node = ns[0].ne[0];
        if (node == null || node.val != t)
            return false;
        for (let i = 0; i < level && ns[i].ne[i] == node; i++)
            ns[i].ne[i] = ns[i].ne[i].ne[i];
        return true;
    }
}
