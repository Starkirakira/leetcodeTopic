package main

import "fmt"

func main()  {

	///***二叉树的序列化与反序列化***///
	//b := Constructor()
	//res := b.deserialize("1,2,3,null,null,4,null,null,5,null,null")
	//fmt.Printf("%+v",res)

	//***1838水域大小***//
	//***你有一个用于表示一片土地的整数矩阵land，该矩阵中每个点的值代表对应地点的海拔高度。若值为0则表示水域。由垂直、水平或对角
	//连接的水域为池塘。池塘的大小是指相连接的水域的个数。编写一个方法来计算矩阵中所有池塘的大小，返回值需要从小到大排序。 ***//
	//###input:
	//[
	//[0,2,1,0],
	//[0,1,0,1],
	//[1,1,0,1],
	//[0,1,1,1]
	//]
	//target:[1,2,4]###//
	//a := [][]int{{0,2,1,0},{0,1,0,1},{1,1,0,1},{0,1,0,1}}
	//fmt.Println(pondSizes(a))

	//***375.猜数字大小***//
	//***对于一个数组[1,n],每次猜错数字需要支付x元。直到猜对为止，求可以赢得游戏至少应该准备多少钱***//
	//n := 10
	//fmt.Println(getMoneyAmount(n))

	///***368.最大整除子集***///
	///***给出一个无重复的正整数集合，找出其中最大的整除子集，子集中任何一对(Si,Sj)都满足Si%Sj=0或者Sj%Si=0,多子集返回任意一个即可***///
	//nums := []int{1,2,3}
	//fmt.Println(largestDivisibleSubset(nums))

	///***414.第三大的数***///
	///***给定非空数组，返回此数组里第三大的数，若不存在，则返回最大的数***///
	//nums := []int{1,1,2}
	//fmt.Println(thirdMax(nums))

	///***1310.有一个正整数数组arr，现给你一个对应的查询数组queries，其中queries[i] = [Li,Ri]。
	//对于每个查询i，请你计算从Li到Ri的XOR值（即arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。
	//并返回一个包含给定查询queries所有结果的数组。***///
	//arr := []int{1,3,4,8}
	//queries := [][]int{{0,1},{1,2},{0,3},{3,3}}
	//fmt.Println(xorQueries(arr, queries))

	//***
	//在 x 轴上有一个一维的花园。花园长度为n，从点0开始，到点n结束。
	//花园里总共有n + 1 个水龙头，分别位于[0, 1, ..., n] 。
	//给你一个整数n和一个长度为n + 1 的整数数组ranges，其中ranges[i] （下标从 0 开始）表示：如果打开点i处的水龙头，可以灌溉的区域为[i - ranges[i], i + ranges[i]]。
	//请你返回可以灌溉整个花园的最少水龙头数目。如果花园始终存在无法灌溉到的地方，请你返回-1。
	//***//
	//n := 7
	//ranges := []int{1,2,1,0,2,1,0,1}
	//fmt.Println(minTaps(n , ranges))

	//***
	//一个 Nx N的 board仅由0和1组成。每次移动，你能任意交换两列或是两行的位置。
	//输出将这个矩阵变为 “棋盘” 所需的最小移动次数。“棋盘” 是指任意一格的上下左右四个方向的值均与本身不同的矩阵。如果不存在可行的变换，输出 -1。
	//***//
	//board := [][]int{{0,1,1,0},{0,1,1,0},{1,0,0,1},{1,0,0,1}}
	//fmt.Println(movesToChessboard(board))

	//***
	//1743.如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
	//例如，
	//[2,3,4]的中位数是 3
	//[2,3] 的中位数是 (2 + 3) / 2 = 2.5
	//设计一个支持以下两种操作的数据结构：
	//void addNum(int num) - 从数据流中添加一个整数到数据结构中。
	//double findMedian() - 返回目前所有元素的中位数。
	//***//
	obj := ConstructorMedian();
	obj.AddNum(1)
	obj.FindMedian()
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
	obj.AddNum(4)
	fmt.Println(obj.FindMedian())


}


