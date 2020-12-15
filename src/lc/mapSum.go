package lc

type MapSum struct {
	val int
	next [26]*MapSum
}

func ConstructorMapSum() MapSum {
	return MapSum{next: [26]*MapSum{}}
}

func (this *MapSum) Insert(key string, val int)  {
	for _,j := range key {
		j := int(j - 'a')
		if this.next[j] == nil {
			this.next[j] = &MapSum{next: [26]*MapSum{}}
		}
		this = this.next[j]
	}
	this.val = val
}

func (this *MapSum) Sum(prefix string) int {
	for _,j := range prefix {
		j := int(j - 'a')
		if this.next[j] == nil {
			return 0
		}
		this = this.next[j]
	}
	ans := 0
	var dfs func(*MapSum); dfs = func(t *MapSum) {
		ans += t.val
		for _,v := range t.next {
			if v !=nil {
				dfs(v)
			}
		}
	}
	dfs(this)
	return ans
}
