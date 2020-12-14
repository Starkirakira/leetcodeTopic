package main

import (
	"fmt"
	"leetcode/lc"
)

func main()  {

	///***二叉树的序列化与反序列化***///
	//b := lc.Constructor()
	//res := b.Deserialize("1,2,3,null,null,4,null,null,5,null,null")
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
	//fmt.Println(lc.PondSizes(a))

	//***375.猜数字大小***//
	//***对于一个数组[1,n],每次猜错数字需要支付x元。直到猜对为止，求可以赢得游戏至少应该准备多少钱***//
	//n := 10
	//fmt.Println(lc.GetMoneyAmount(n))

	///***368.最大整除子集***///
	///***给出一个无重复的正整数集合，找出其中最大的整除子集，子集中任何一对(Si,Sj)都满足Si%Sj=0或者Sj%Si=0,多子集返回任意一个即可***///
	//nums := []int{1,2,3}
	//fmt.Println(lc.LargestDivisibleSubset(nums))

	///***414.第三大的数***///
	///***给定非空数组，返回此数组里第三大的数，若不存在，则返回最大的数***///
	//nums := []int{1,1,2}
	//fmt.Println(lc.ThirdMax(nums))

	///***1310.有一个正整数数组arr，现给你一个对应的查询数组queries，其中queries[i] = [Li,Ri]。
	//对于每个查询i，请你计算从Li到Ri的XOR值（即arr[Li] xor arr[Li+1] xor ... xor arr[Ri]）作为本次查询的结果。
	//并返回一个包含给定查询queries所有结果的数组。***///
	//arr := []int{1,3,4,8}
	//queries := [][]int{{0,1},{1,2},{0,3},{3,3}}
	//fmt.Println(lc.XorQueries(arr, queries))


	//***
	//在 x 轴上有一个一维的花园。花园长度为n，从点0开始，到点n结束。
	//花园里总共有n + 1 个水龙头，分别位于[0, 1, ..., n] 。
	//给你一个整数n和一个长度为n + 1 的整数数组ranges，其中ranges[i] （下标从 0 开始）表示：如果打开点i处的水龙头，可以灌溉的区域为[i - ranges[i], i + ranges[i]]。
	//请你返回可以灌溉整个花园的最少水龙头数目。如果花园始终存在无法灌溉到的地方，请你返回-1。
	//***//
	//n := 7
	//ranges := []int{1,2,1,0,2,1,0,1}
	//fmt.Println(lc.MinTaps(n , ranges))


	//***
	//一个 Nx N的 board仅由0和1组成。每次移动，你能任意交换两列或是两行的位置。
	//输出将这个矩阵变为 “棋盘” 所需的最小移动次数。“棋盘” 是指任意一格的上下左右四个方向的值均与本身不同的矩阵。如果不存在可行的变换，输出 -1。
	//***//
	//board := [][]int{{0,1,1,0},{0,1,1,0},{1,0,0,1},{1,0,0,1}}
	//fmt.Println(lc.MovesToChessboard(board))

	//***
	//1743.如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。
	//例如，
	//[2,3,4]的中位数是 3
	//[2,3] 的中位数是 (2 + 3) / 2 = 2.5
	//设计一个支持以下两种操作的数据结构：
	//void addNum(int num) - 从数据流中添加一个整数到数据结构中。
	//double findMedian() - 返回目前所有元素的中位数。
	//***//
	//obj := lc.ConstructorMedian()
	//obj.AddNum(1)
	//obj.FindMedian()
	//obj.AddNum(2)
	//fmt.Println(obj.FindMedian())
	//obj.AddNum(3)
	//fmt.Println(obj.FindMedian())
	//obj.AddNum(4)
	//fmt.Println(obj.FindMedian())

	//***
	//最初在一个记事本上只有一个字符 'A'。你每次可以对这个记事本进行两种操作：
	//Copy All (复制全部) : 你可以复制这个记事本中的所有字符(部分的复制是不允许的)。
	//Paste (粘贴) : 你可以粘贴你上一次复制的字符。
	//给定一个数字n。你需要使用最少的操作次数，在记事本中打印出恰好n个 'A'。输出能够打印出n个 'A' 的最少操作次数。
	//***//
	//n := 3
	//fmt.Println(lc.MinSteps(n))

	//***
	//1482.给你一个整数数组 bloomDay，以及两个整数 m 和 k 。
	//现需要制作 m 束花。制作花束时，需要使用花园中 相邻的 k 朵花 。
	//花园中有 n 朵花，第 i 朵花会在 bloomDay[i] 时盛开，恰好 可以用于 一束 花中。
	//请你返回从花园中摘 m 束花需要等待的最少的天数。如果不能摘到 m 束花则返回 -1 。
	//***//
	//bloomDay := []int{1,10,3,10,2}
	//m,k := 3,1
	//fmt.Println(lc.MinDays(bloomDay,m,k))

	//***
	//你有一个日志数组 logs。每条日志都是以空格分隔的字串。
	//对于每条日志，其第一个字为字母与数字混合的 标识符 ，除标识符之外的所有字为这一条日志的 内容 。
	//除标识符之外，所有字均由小写字母组成的，称为 字母日志
	//除标识符之外，所有字均由数字组成的，称为 数字日志
	//题目所用数据保证每个日志在其标识符后面至少有一个字。
	//请按下述规则将日志重新排序：
	//所有 字母日志 都排在 数字日志 之前。
	//字母日志 在内容不同时，忽略标识符后，按内容字母顺序排序；在内容相同时，按标识符排序；
	//数字日志 应该按原来的顺序排列。
	//返回日志的最终顺序。
	//***//
	//logs := []string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}
	//fmt.Println(lc.ReorderLogFiles(logs))

	///***
	//377.给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。
	//***
	//nums := []int{1,2,3}
	//target := 4
	//fmt.Println(lc.CombinationSum4(nums, target))

	//***
	//序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
	//设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
	//编码的字符串应尽可能紧凑。
	//***//
	//root := lc.TreeNodeii{1,&lc.TreeNodeii{2,nil,nil},&lc.TreeNodeii{3,nil,nil}}
	//ser := lc.Constructorii()
	//deser := lc.Constructorii()
	//tree := ser.Serializeii(&root)
	//ans := deser.Deserializeii(tree)
	//fmt.Println(ans)

	//***
	//「HTML实体解析器」 是一种特殊的解析器，它将 HTML 代码作为输入，并用字符本身替换掉所有这些特殊的字符实体。
	//
	//HTML 里这些特殊字符和它们对应的字符实体包括：
	//
	//双引号：字符实体为&quot;，对应的字符是"。
	//单引号：字符实体为&apos;，对应的字符是'。
	//与符号：字符实体为&amp;，对应对的字符是&。
	//大于号：字符实体为&gt;，对应的字符是>。
	//小于号：字符实体为&lt;，对应的字符是<。
	//斜线号：字符实体为&frasl;，对应的字符是/。
	//给你输入字符串text，请你实现一个 HTML实体解析器，返回解析器解析后的结果。
	//***//
	//text := "and I quote: &quot;...&quot;"
	//text := "x &gt; y &amp;&amp; x &lt; y is always fals"
	//fmt.Println(lc.EntityParser(text))


	//***
	//在给定的网格中，每个单元格可以有以下三个值之一：
	//
	//值0代表空单元格；
	//值1代表新鲜橘子；
	//值2代表腐烂的橘子。
	//每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
	//返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回-1。
	//***//
	//grid := [][]int{{2,1,1},{1,1,0},{0,1,1}}
	//fmt.Println(lc.OrangesRotting(grid))

	///***给你一个整数num，请你找出同时满足下面全部要求的两个整数：
	//两数乘积等于 num + 1或num + 2
	//以绝对差进行度量，两数大小最接近
	//你可以按任意顺序返回这两个整数。***///
	//num := 8
	//fmt.Println(lc.ClosestDivisors(num))


	//***
	//给定一个二维网格grid。"."代表一个空房间，"#"代表一堵墙，"@"是起点，（"a","b", ...）代表钥匙，（"A","B", ...）代表锁。
	//我们从起点开始出发，一次移动是指向四个基本方向之一行走一个单位空间。我们不能在网格外面行走，也无法穿过一堵墙。如果途经一个钥匙，我们就把它捡起来。除非我们手里有对应的钥匙，否则无法通过锁。
	//假设 K 为钥匙/锁的个数，且满足1 <= K <= 6，字母表中的前 K 个字母在网格中都有自己对应的一个小写和一个大写字母。换言之，每个锁有唯一对应的钥匙，每个钥匙也有唯一对应的锁。另外，代表钥匙和锁的字母互为大小写并按字母顺序排列。
	//返回获取所有钥匙所需要的移动的最少次数。如果无法获取所有钥匙，返回-1。
	//链接：https://leetcode-cn.com/problems/shortest-path-to-get-all-keys
	//***//
	//grid := []string{"@.a.#","###.#","b.A.B"}
	//fmt.Println(lc.ShortestPathAllKeys(grid))

	//***
	//给你一个整数数组 arr。你可以从中选出一个整数集合，并删除这些整数在数组中的每次出现。
	//返回 至少 能删除数组中的一半整数的整数集合的最小大小。
	//***//
	arr := []int{3,3,3,3,5,5,5,2,2,7}
	fmt.Println(lc.MinSetSize(arr))



}


