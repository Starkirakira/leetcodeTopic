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


	///***
	//给定一个整数数组和一个整数k，你需要在数组里找到不同的k-diff 数对，并返回不同的 k-diff 数对 的数目。
	//
	//这里将k-diff数对定义为一个整数对 (nums[i], nums[j])，并满足下述全部条件：
	//
	//0 <= i, j < nums.length
	//i != j
	//|nums[i] - nums[j]| == k
	//注意，|val| 表示 val 的绝对值。
	//链接：https://leetcode-cn.com/problems/k-diff-pairs-in-an-array
	//***///
	//num := []int{1,2,4,4,3,3,0,9,2,3}
	//num := []int{1,3,1,5,40}
	//k := 0
	//fmt.Println(lc.FindPairs(num , k))

	//***
	//电子游戏“辐射4”中，任务“通向自由”要求玩家到达名为“Freedom Trail Ring”的金属表盘，并使用表盘拼写特定关键词才能开门。
	//给定一个字符串ring，表示刻在外环上的编码；给定另一个字符串key，表示需要拼写的关键词。您需要算出能够拼写关键词中所有字符的最少步数。
	//最初，ring的第一个字符与12:00方向对齐。您需要顺时针或逆时针旋转 ring 以使key的一个字符在 12:00 方向对齐，然后按下中心按钮，以此逐个拼写完key中的所有字符。
	//旋转ring拼出 key 字符key[i]的阶段中：
	//您可以将ring顺时针或逆时针旋转一个位置，计为1步。旋转的最终目的是将字符串ring的一个字符与 12:00 方向对齐，并且这个字符必须等于字符key[i] 。
	//如果字符key[i]已经对齐到12:00方向，您需要按下中心按钮进行拼写，这也将算作1 步。按完之后，您可以开始拼写key的下一个字符（下一阶段）, 直至完成所有拼写。
	//链接：https://leetcode-cn.com/problems/freedom-trail
	//***//
	//ring := "godding"
	//key := "gd"
	//fmt.Println(lc.FindRotateSteps(ring , key))

	//***给你一个长度固定的整数数组arr，请你将该数组中出现的每个零都复写一遍，并将其余的元素向右平移。
	//
	//注意：请不要在超过该数组长度的位置写入元素。
	//
	//要求：请对输入的数组就地进行上述修改，不要从函数返回任何东西。
	//链接：https://leetcode-cn.com/problems/duplicate-zeros ***//
	//arr := []int{1,0,2,3,0,4,5,0}
	//lc.DuplicateZeros(arr)

	//***
	//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
	//***//
	//n := 3
	//fmt.Println(lc.GenerateParenthesis(n))

	//***
	//给你一个整数数组 arr。你可以从中选出一个整数集合，并删除这些整数在数组中的每次出现。
	//返回 至少 能删除数组中的一半整数的整数集合的最小大小。
	//***//
	//arr := []int{3,3,3,3,5,5,5,2,2,7}
	//fmt.Println(lc.MinSetSize(arr))

	//***
	//实现一个 MapSum 类，支持两个方法，insert和sum：
	//MapSum() 初始化 MapSum 对象
	//void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
	//int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。
	//链接：https://leetcode-cn.com/problems/map-sum-pairs
	//***//
	//obj := lc.ConstructorMapSum()
	//obj.Insert("apple", 3)
	//param_2 := obj.Sum("ap")
	//fmt.Println(param_2)
	//obj.Insert("app", 2)
	//param_3 := obj.Sum("ap")
	//fmt.Println(param_3)

	//***
	//给你一个整数数组nums 和一个整数 k。
	//如果某个 连续 子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
	//请返回这个数组中「优美子数组」的数目。
	//链接：https://leetcode-cn.com/problems/count-number-of-nice-subarrays
	//***//
	//nums := []int{2,2,2,1,2,2,1,2,2,2}
	//k := 2
	//fmt.Println(lc.NumberOfSubarrays(nums, k))

	//***
	//给定一种规律 pattern和一个字符串str，判断 str 是否遵循相同的规律。
	//这里的遵循指完全匹配，例如，pattern里的每个字母和字符串str中的每个非空单词之间存在着双向连接的对应规律。
	//链接：https://leetcode-cn.com/problems/word-pattern
	//***//
	//pattern := "abba"
	//str := "dog cat cat dog"
	//fmt.Println(lc.WordPattern(pattern,str))

	//***
	//给定一个整数数组prices，其中第i个元素代表了第i天的股票价格 ；非负整数fee 代表了交易股票的手续费用。
	//你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
	//返回获得利润的最大值。
	//注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
	//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee
	//***//
	//prices := []int{1,3,2,8,4,9}
	//fee := 2
	//fmt.Println(lc.MaxProfit(prices, fee))

	//***
	//小A 和 小B 在玩猜数字。小B 每次从 1, 2, 3 中随机选择一个，小A 每次也从 1, 2, 3 中选择一个猜。他们一共进行三次这个游戏，请返回 小A 猜对了几次？
	//输入的guess数组为 小A 每次的猜测，answer数组为 小B 每次的选择。guess和answer的长度都等于3。
	//链接：https://leetcode-cn.com/problems/guess-numbers
	//***//
	//guess := []int{1,2,3}
	//answer := []int{1,2,3}
	//fmt.Println(lc.Game(guess, answer))

	//***
	//给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用补码运算方法。
	//注意:
	//十六进制中所有字母(a-f)都必须是小写。
	//十六进制字符串中不能包含多余的前导零。如果要转化的数为0，那么以单个字符'0'来表示；对于其他情况，十六进制字符串中的第一个字符将不会是0字符。
	//给定的数确保在32位有符号整数范围内。
	//不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。
	//链接：https://leetcode-cn.com/problems/convert-a-number-to-hexadecimal
	//***//
	//num := 26
	//fmt.Println(lc.ToHex(num))

	//***
	//给定两个字符串 s 和 t，它们只包含小写字母。
	//字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
	//请找出在 t 中被添加的字母。
	//***//
	//s := "abcd"
	//t := "abcde"
	//fmt.Println(lc.FindTheDifference(s, t))

	//***
	//给定一个 n×n 的二维矩阵表示一个图像。
	//将图像顺时针旋转 90 度。
	//说明：
	//你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
	//链接：https://leetcode-cn.com/problems/rotate-image
	//***//
	//matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//lc.Rotate(matrix)
	//fmt.Println(matrix)

	///***
	//请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。
	//链接：https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof
	//***///
	//head := &lc.NodeRandomList{
	//	1,
	//	&lc.NodeRandomList{
	//		2,
	//		&lc.NodeRandomList{},
	//		&lc.NodeRandomList{
	//			2,
	//			&lc.NodeRandomList{},
	//			&lc.NodeRandomList{},
	//		},
	//	},
	//	&lc.NodeRandomList{
	//		2,
	//		&lc.NodeRandomList{},
	//		&lc.NodeRandomList{},
	//	}}
	//fmt.Println(lc.CopyRandomList(head))

	//***
	//给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
	//***//
	//s := "bcabc"
	//fmt.Println(lc.RemoveDuplicateLetters(s))

	//***
	//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
	//***//
	//nums := []int{1,3,5,6}
	//target := 5
	//fmt.Println(lc.SearchInsert(nums, target))

	//***
	//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
	//链接：https://leetcode-cn.com/problems/container-with-most-water
	//***//
	//height := []int{1,8,6,2,5,4,8,3,7}
	//fmt.Println(lc.MaxArea(height))

	//***
	//数组的每个索引作为一个阶梯，第i个阶梯对应着一个非负数的体力花费值cost[i](索引从0开始)。
	//每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。
	//您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。
	//链接：https://leetcode-cn.com/problems/min-cost-climbing-stairs
	//***//
	//cots := []int{10,15,20}
	//fmt.Println(lc.MinCostClimbingStairs(cots))

	//***
	//给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
	//	  3
	//   / \
	//  9  20
	//    /  \
	//   15   7
	//[[3],[20,9],[15,7]]
	//***//
	//root := &lc.TreeNode{
	//	3,
	//	&lc.TreeNode{
	//		9,
	//		nil,
	//		nil,
	//	},
	//	&lc.TreeNode{
	//		20,
	//		&lc.TreeNode{
	//			15,
	//			nil,
	//			nil,
	//		},
	//		&lc.TreeNode{
	//			7,
	//			nil,
	//			nil,
	//		},
	//	},
	//}
	//fmt.Println(lc.ZigzagLevelOrder(root))

	//***
	//给定字符串J代表石头中宝石的类型，和字符串S代表你拥有的石头。S中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。
	//
	//J中的字母不重复，J和S中的所有字符都是字母。字母区分大小写，因此"a"和"A"是不同类型的石头。
	//链接：https://leetcode-cn.com/problems/jewels-and-stones
	//***//
	//J := "aA"
	//S := "aAAbbbb"
	//fmt.Println(lc.NumJewelsInStones(J, S))

	//***
	//给定两个字符串 S 和 T，每个字符串代表一个非负有理数，只有当它们表示相同的数字时才返回 true；否则，返回 false。字符串中可以使用括号来表示有理数的重复部分。
	//通常，有理数最多可以用三个部分来表示：整数部分<IntegerPart>、小数非重复部分<NonRepeatingPart>和小数重复部分<(><RepeatingPart><)>。数字可以用以下三种方法之一来表示：
	//<IntegerPart>（例：0，12，123）
	//<IntegerPart><.><NonRepeatingPart> （例：0.5，2.12，2.0001）
	//<IntegerPart><.><NonRepeatingPart><(><RepeatingPart><)>（例：0.1(6)，0.9(9)，0.00(1212)）
	//十进制展开的重复部分通常在一对圆括号内表示。例如：
	//1 / 6 = 0.16666666... = 0.1(6) = 0.1666(6) = 0.166(66)
	//0.1(6) 或0.1666(6) 或0.166(66) 都是1 / 6 的正确表示形式。
	//链接：https://leetcode-cn.com/problems/equal-rational-numbers
	//***//
	//S := "0.(52)"
	//T := "0.5(25)"
	//fmt.Println(lc.IsRationalEqual(S, T))

	//***
	//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
	//***//
	//s := "leetcode"
	//fmt.Println(lc.FirstUniqChar(s))

	//***
	//给你一个长度为 偶数 n 的整数数组 nums 和一个整数 limit 。每一次操作，你可以将 nums 中的任何整数替换为1到limit 之间的另一个整数。
	//如果对于所有下标 i（下标从 0 开始），nums[i] + nums[n - 1 - i]都等于同一个数，则数组 nums 是 互补的 。例如，数组 [1,2,3,4] 是互补的，因为对于所有下标i ，nums[i] + nums[n - 1 - i] = 5 。
	//返回使数组 互补 的 最少操作次数。
	//链接：https://leetcode-cn.com/problems/minimum-moves-to-make-array-complementary
	//***//
	//nums,limit := []int{1,2,2,1},2
	//fmt.Println(lc.MinMoves(nums, limit))

	//***
	//老师想给孩子们分发糖果，有 N个孩子站成了一条直线，老师会根据每个孩子的表现，预先给他们评分。
	//你需要按照以下要求，帮助老师给这些孩子分发糖果：
	//每个孩子至少分配到 1 个糖果。
	//相邻的孩子中，评分高的孩子必须获得更多的糖果。
	//那么这样下来，老师至少需要准备多少颗糖果呢？
	//链接：https://leetcode-cn.com/problems/candy
	//***//
	//ratings := []int{1,0,2}
	//fmt.Println(lc.Candy(ratings))

	//***
	//假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
	//对每个孩子 i，都有一个胃口值g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j]。如果 s[j]>= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。
	//你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。
	//链接：https://leetcode-cn.com/problems/assign-cookies
	//***//
	//g,s := []int{1,2,3},[]int{1,1}
	//fmt.Println(lc.FindContentChildren(g,s))

	//***
	//给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
	//两棵树重复是指它们具有相同的结构以及相同的结点值。
	//***//
	//root := &lc.TreeNode{
	//	1,
	//	&lc.TreeNode{
	//		2,
	//		&lc.TreeNode{
	//			4,
	//			nil,
	//			nil,
	//		},
	//		nil,
	//	},
	//	&lc.TreeNode{
	//		3,
	//		&lc.TreeNode{
	//			2,
	//			&lc.TreeNode{
	//				4,
	//				nil,
	//				nil,
	//			},
	//			nil,
	//
	//		},
	//		&lc.TreeNode{
	//			4,
	//			nil,
	//			nil,
	//		},
	//	},
	//}
	//fmt.Println(lc.FindDuplicateSubtrees(root))

	//***给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。***//
	//matrix := [][]byte{{1,0,1,0,0},{1,0,1,1,1},{1,1,1,1,1},{1,0,0,1,0}}
	//fmt.Println(lc.MaximalRectangle(matrix))

	//***
	//给定一个方形整数数组A，我们想要得到通过 A 的下降路径的最小和。
	//下降路径可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列。
	//链接：https://leetcode-cn.com/problems/minimum-falling-path-sum
	//***//
	//A := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//fmt.Println(lc.MinFallingPathSum(A))

	//***
	//给定两个字符串s和t，判断它们是否是同构的。
	//如果s中的字符可以被替换得到t，那么这两个字符串是同构的。
	//所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
	//链接：https://leetcode-cn.com/problems/isomorphic-strings
	//***//
	//s := "egg"
	//t := "add"
	//fmt.Println(lc.IsIsomorphic(s,t))

	//***
	//给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。
	//设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
	//注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
	//链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv
	//***//
	//k,prices := 2,[]int{2,4,1}
	//fmt.Println(k,prices)

	//***
	//给定一个已排序的正整数数组 nums，和一个正整数n 。从[1, n]区间内选取任意个数字补充到nums中，使得[1, n]区间内的任何
	//数字都可以用nums中某几个数字的和来表示。请输出满足上述要求的最少需要补充的数字个数。
	//链接：https://leetcode-cn.com/problems/patching-array
	//***//
	nums := []int{1,3}
	n := 6
	fmt.Println(lc.MinPatches(nums, n))



































}


