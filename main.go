package main

import "fmt"

func main()  {
	//a := []string{"1","2","3","null","null","4","5"}
	//b := Constructor()
	//res := b.deserialize("1,2,3,null,null,4,null,null,5,null,null")
	//fmt.Printf("%+v",res)

	a := [][]int{{1,2},{2,1},{1,0},{0,1}}

	fmt.Println(minAreaFreeRect(a))

}



