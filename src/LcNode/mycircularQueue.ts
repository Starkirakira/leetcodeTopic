class MyCircularQueue {
    private list: Array<number>
    private f: number
    private n: number
    private c: number
    constructor(k: number) {
        this.c = k + 1
        this.list = new Array<number>(this.c).fill(0)
        this.f = 0
        this.n = 0
        
    }

    enQueue(value: number): boolean {
        if(this.isFull()) return false
        this.list[this.n] = value
        this.n = (this.n + 1) % this.c
        return true
    }

    deQueue(): boolean {
        if (this.isEmpty()) return false
        this.f = (this.f + 1) % this.c
        return true
    }

    Front(): number {
        if (this.isEmpty()) return -1
        return this.list[this.f]
    }

    Rear(): number {
        if (this.isEmpty()) return -1
        return this.list[(this.n - 1 + this.c) % this.c]
    }

    isEmpty(): boolean {
        return this.f == this.n
    }

    isFull(): boolean {
        return this.f == (this.n + 1) % this.c
    }
}

  var obj = new MyCircularQueue(8)
 var param_1 = obj.enQueue(1)

//  var param_2 = obj.deQueue()
//   var param_3 = obj.Front()
//  var param_4 = obj.Rear()
//   var param_5 = obj.isEmpty()
//   var param_6 = obj.isFull()
 console.log(obj)