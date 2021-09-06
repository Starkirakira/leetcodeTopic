package main

import (
	"fmt"
	"leetcode/lc"
)

func main() {

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
	//nums := []int{1,3}
	//n := 6
	//fmt.Println(lc.MinPatches(nums, n))

	//***
	//给定一个包括n 个整数的数组nums和 一个目标值target。找出nums中的三个整数，使得它们的和与target最接近。返回这三个数的和。假定每组输入只存在唯一答案。
	//链接：https://leetcode-cn.com/problems/3sum-closest
	//***//
	//nums,target := []int{-1,2,1,-4},1
	//fmt.Println(lc.ThreeSumClosest(nums, target))

	//***
	//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
	//***//
	//root := &lc.TreeNode{
	//	1,
	//	nil,
	//	&lc.TreeNode{
	//		2,
	//		&lc.TreeNode{
	//			3,
	//			nil,
	//			nil,
	//		},
	//		nil,
	//	},
	//}
	//fmt.Println(lc.InorderTraversal(root))

	//***
	//给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。
	//***//
	//fmt.Println(lc.GenerateTrees(3))

	//***
	//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
	//***//
	//fmt.Println(lc.NumTrees(3))

	//***
	//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
	//假设一个二叉搜索树具有如下特征：
	//节点的左子树只包含小于当前节点的数。
	//节点的右子树只包含大于当前节点的数。
	//所有左子树和右子树自身必须也是二叉搜索树。
	//链接：https://leetcode-cn.com/problems/validate-binary-search-tree
	//***//
	//root := &lc.TreeNode{
	//	5,
	//	&lc.TreeNode{
	//		1,
	//		nil,
	//		nil,
	//	},
	//	&lc.TreeNode{
	//		4,
	//		&lc.TreeNode{
	//			3,
	//			nil,
	//			nil,
	//		},
	//		&lc.TreeNode{
	//			6,
	//			nil,nil,
	//		},
	//	},
	//}
	//fmt.Println(lc.IsValidBST(root))

	//***
	//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
	//***//
	//root := &lc.TreeNode{
	//	3,
	//	&lc.TreeNode{
	//		9,nil,nil,
	//	},
	//	&lc.TreeNode{
	//		20,
	//		&lc.TreeNode{
	//			15,nil,nil,
	//		},
	//		&lc.TreeNode{
	//			7,nil,nil,
	//		},
	//	},
	//}
	//fmt.Println(lc.LevelOrder(root))

	//***
	//根据一棵树的前序遍历与中序遍历构造二叉树。
	//注意: 你可以假设树中没有重复的元素。
	//
	//二叉树前序遍历的顺序为： 先遍历根节点； 随后递归地遍历左子树； 最后递归地遍历右子树。
	//二叉树中序遍历的顺序为： 先递归地遍历左子树； 随后遍历根节点； 最后递归地遍历右子树。
	//***//
	//preorder, inorder := []int{3,9,20,15,7}, []int{9,3,15,20,7}
	//fmt.Println(lc.BuildTree(preorder, inorder))

	//***
	//根据一棵树的中序遍历与后序遍历构造二叉树。
	//后序遍历的顺序是每次遍历左孩子，再遍历右孩子，最后遍历根节点
	//***//
	//inorder, postorder := []int{9,3,15,20,7},[]int{9,15,7,20,3}
	//fmt.Println(lc.BuildTreePostorder(inorder, postorder))

	//***
	//给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
	//说明: 叶子节点是指没有子节点的节点。
	//***//
	//root := &lc.TreeNode{
	//	5,
	//	&lc.TreeNode{
	//		4,
	//		&lc.TreeNode{
	//			11,
	//			&lc.TreeNode{7,nil,nil,},
	//			&lc.TreeNode{2,nil,nil,},
	//		},
	//		nil,
	//	},
	//	&lc.TreeNode{
	//		8,
	//		&lc.TreeNode{
	//			13,nil,nil,
	//		},
	//		&lc.TreeNode{
	//			4,
	//			&lc.TreeNode{5,nil,nil,},
	//			&lc.TreeNode{1,nil,nil,},
	//		},
	//	},
	//}
	//sum := 22
	//fmt.Println(lc.PathSum(root,sum))

	//***
	//给定一个二叉树，原地将它展开为一个单链表。
	//***//
	//root := &lc.TreeNode{
	//	1,
	//	&lc.TreeNode{
	//		2,
	//		&lc.TreeNode{3,nil,nil,},
	//		&lc.TreeNode{4,nil,nil,},
	//	},
	//	&lc.TreeNode{
	//		5,
	//		nil,
	//		&lc.TreeNode{6,nil,nil,},
	//	},
	//}
	//lc.Flatten(root)
	//fmt.Println(root)

	//***
	//给定一个二叉树，检查它是否是镜像对称的。
	//***//
	//root := &lc.TreeNode{
	//	1,
	//	&lc.TreeNode{
	//		2,
	//		&lc.TreeNode{3,nil,nil,},
	//		&lc.TreeNode{4,nil,nil,},
	//	},
	//	&lc.TreeNode{
	//		2,
	//		&lc.TreeNode{4,nil,nil,},
	//		&lc.TreeNode{3,nil,nil,},
	//	},
	//}
	//fmt.Println(lc.IsSymmetric(root))

	//***
	//给定一个二叉树，找出其最大深度。
	//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
	//说明: 叶子节点是指没有子节点的节点。
	//***//
	//root := &lc.TreeNode{
	//	3,
	//	&lc.TreeNode{9,nil,nil},
	//	&lc.TreeNode{
	//		20,
	//		&lc.TreeNode{15,nil,nil,},
	//		&lc.TreeNode{7,nil,nil,},
	//	},
	//}
	//fmt.Println(lc.MaxDepth(root))

	//***
	//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
	//***//
	//root := &lc.TreeNode{
	//	3,
	//	&lc.TreeNode{9,nil,nil,},
	//	&lc.TreeNode{
	//		20,
	//		&lc.TreeNode{15,nil,nil,},
	//		&lc.TreeNode{7,nil,nil,},
	//	},
	//}
	//fmt.Println(lc.LevelOrderBottom(root))

	//***
	//有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。
	//省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。
	//给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。
	//返回矩阵中 省份 的数量。
	//链接：https://leetcode-cn.com/problems/number-of-provinces
	//***//
	//isConnected := [][]int{{1,1,0},{1,1,0},{0,0,1}}
	//fmt.Println(lc.FindCircleNum(isConnected))

	//***
	//将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
	//本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
	//***//
	//nums := []int{-10, -3, 0, 5, 9}
	//fmt.Println(lc.SortedArrayToBST(nums))

	//***
	//给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
	//本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1。
	//链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
	//***//
	//list := &lc.ListNode{-10,&lc.ListNode{-3,&lc.ListNode{0, &lc.ListNode{5, &lc.ListNode{9,nil}}}}}
	//fmt.Println(lc.SortedListToBST(list))

	//***
	//寻找最长的回文字串
	//***//
	//fmt.Println(lc.LongestPalindrome("aabbbbbaaa"))

	//***
	//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
	//子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
	//链接：https://leetcode-cn.com/problems/longest-increasing-subsequence
	//***//
	//nums := []int{10,9,2,5,3,7,101,18}
	//fmt.Println(lc.LengthOfLIS(nums))

	//***
	//当 A的子数组A[i], A[i+1], ..., A[j]满足下列条件时，我们称其为湍流子数组：
	//若i <= k < j，当 k为奇数时，A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
	//或 若i <= k < j，当 k 为偶数时，A[k] > A[k+1]，且当 k为奇数时，A[k] < A[k+1]。
	//也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。
	//返回 A 的最大湍流子数组的长度。
	//链接：https://leetcode-cn.com/problems/longest-turbulent-subarray
	//***//
	//arr := []int{9,4,2,10,7,8,8,1,9}
	//fmt.Println(lc.MaxTurbulenceSize(arr))

	//***
	//给定一个有 N 个结点的二叉树的根结点 root，树中的每个结点上都对应有 node.val 枚硬币，并且总共有 N 枚硬币。
	//在一次移动中，我们可以选择两个相邻的结点，然后将一枚硬币从其中一个结点移动到另一个结点。(移动可以是从父结点到子结点，或者从子结点移动到父结点。)。
	//返回使每个结点上只有一枚硬币所需的移动次数。
	//链接：https://leetcode-cn.com/problems/distribute-coins-in-binary-tree
	//***//
	//root := &lc.TreeNode{
	//	3,
	//	&lc.TreeNode{0,nil,nil,},
	//	&lc.TreeNode{0,nil,nil,},
	//}
	//fmt.Println(lc.DistributeCoins(root))

	//***
	//有一个 m x n 的二元网格，其中 1 表示砖块，0 表示空白。砖块 稳定（不会掉落）的前提是：
	//一块砖直接连接到网格的顶部，或者
	//至少有一块相邻（4个方向之一）砖块 稳定 不会掉落时
	//给你一个数组 hits ，这是需要依次消除砖块的位置。每当消除hits[i] = (rowi, coli) 位置上的砖块时，对应位置的砖块（若存在）会消失，然后其他的砖块可能因为这一消除操作而掉落。一旦砖块掉落，它会立即从网格中消失（即，它不会落在其他稳定的砖块上）。
	//返回一个数组 result ，其中 result[i] 表示第 i 次消除操作对应掉落的砖块数目。
	//注意，消除可能指向是没有砖块的空白位置，如果发生这种情况，则没有砖块掉落。
	//链接：https://leetcode-cn.com/problems/bricks-falling-when-hit
	//***//
	//grid := [][]int{{1,0,0,0},{1,1,1,0}}
	//hits := [][]int{{1,0}}
	//fmt.Println(lc.HitBricks(grid, hits))

	//***
	//给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
	//链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
	//***//
	//num := 12258
	//fmt.Println(lc.TranslateNum(num))

	//***
	//在一个XY 坐标系中有一些点，我们用数组coordinates来分别记录它们的坐标，其中coordinates[i] = [x, y]表示横坐标为 x、纵坐标为 y的点。
	//请你来判断，这些点是否在该坐标系中属于同一条直线上，是则返回 true，否则请返回 false。
	//链接：https://leetcode-cn.com/problems/check-if-it-is-a-straight-line
	//***//
	//coordinates := [][]int{{1,2},{2,3},{3,4},{4,5},{5,6}}
	//fmt.Println(lc.CheckStraightLine(coordinates))

	//***
	//给定一个列表 accounts，每个元素 accounts[i]是一个字符串列表，其中第一个元素 accounts[i][0]是名称 (name)，其余元素是 emails 表示该账户的邮箱地址。
	//现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。
	//合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是按字符 ASCII 顺序排列的邮箱地址。账户本身可以以任意顺序返回。
	//链接：https://leetcode-cn.com/problems/accounts-merge
	//***//
	//accounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}
	//fmt.Println(lc.AccountsMerge(accounts))

	//***
	//给你一个points数组，表示 2D 平面上的一些点，其中points[i] = [xi, yi]。
	//连接点[xi, yi] 和点[xj, yj]的费用为它们之间的 曼哈顿距离：|xi - xj| + |yi - yj|，其中|val|表示val的绝对值。
	//请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有一条简单路径时，才认为所有点都已连接。
	//链接：https://leetcode-cn.com/problems/min-cost-to-connect-all-points
	//***//
	//points := [][]int{{0,0},{2,2},{3,10},{5,2},{7,0}}
	//fmt.Println(lc.MinCostConnectPoints(points))

	//***
	//给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
	//***//
	//fmt.Println(lc.MaximumProduct([]int{1,2,3}))

	//***
	//给定一个整数数组A，返回 A中最长等差子序列的长度。
	//回想一下，A的子序列是列表A[i_1], A[i_2], ..., A[i_k] 其中0 <= i_1 < i_2 < ... < i_k <= A.length - 1。并且如果B[i+1] - B[i](0 <= i < B.length - 1) 的值都相同，那么序列B是等差的。
	//链接：https://leetcode-cn.com/problems/longest-arithmetic-subsequence
	//***//
	//fmt.Println(lc.LongestArithSeqLength([]int{3,6,9,12}))

	//***
	//给你一个 n个点的带权无向连通图，节点编号为 0到 n-1，同时还有一个数组 edges，其中 edges[i] = [fromi, toi, weighti]表示在fromi和toi节点之间有一条带权无向边。最小生成树(MST) 是给定图中边的一个子集，它连接了所有节点且没有环，而且这些边的权值和最小。
	//请你找到给定图中最小生成树的所有关键边和伪关键边。如果从图中删去某条边，会导致最小生成树的权值和增加，那么我们就说它是一条关键边。伪关键边则是可能会出现在某些最小生成树中但不会出现在所有最小生成树中的边。
	//请注意，你可以分别以任意顺序返回关键边的下标和伪关键边的下标。
	//链接：https://leetcode-cn.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree
	//***//
	//n := 5
	//edges := [][]int{{0,1,1},{1,2,1},{2,3,2},{0,3,2},{0,4,3},{3,4,3},{1,4,6}}
	//fmt.Println(lc.FindCriticalAndPseudoCriticalEdges(n, edges))

	//***
	//对于非负整数X而言，X的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果X = 1231，那么其数组形式为[1,2,3,1]。
	//给定非负整数 X 的数组形式A，返回整数X+K的数组形式。
	//链接：https://leetcode-cn.com/problems/add-to-array-form-of-integer
	//***//
	//A,K := []int{1,2,0,0}, 34
	//fmt.Println(lc.AddToArrayForm(A,K))

	//***
	//用以太网线缆将n台计算机连接成一个网络，计算机的编号从0到n-1。线缆用connections表示，其中connections[i] = [a, b]连接了计算机a和b。
	//网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。
	//给你这个计算机网络的初始布线connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连通所需的最少操作次数。如果不可能，则返回-1 。
	//链接：https://leetcode-cn.com/problems/number-of-operations-to-make-network-connected
	//***//
	//n := 4
	//connections := [][]int{{0,1},{0,2},{1,2}}
	//fmt.Println(lc.MakeConnected(n ,connections))

	//***
	//给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
	//连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] 就是连续递增子序列。
	//链接：https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence
	//***//
	//nums := []int{1,3,5,4,7}
	//fmt.Println(lc.FindLengthOfLCIS(nums))

	//***
	//在由 1 x 1 方格组成的 N x N 网格grid 中，每个 1 x 1方块由 /、\ 或空格构成。这些字符会将方块划分为一些共边的区域。
	//（请注意，反斜杠字符是转义的，因此 \ 用 "\\"表示。）。
	//返回区域的数目。
	//链接：https://leetcode-cn.com/problems/regions-cut-by-slashes
	//***//
	//fmt.Println(lc.RegionsBySlashes([]string{"/\\","\\/"}))

	//***
	//给你一个由一些多米诺骨牌组成的列表dominoes。
	//如果其中某一张多米诺骨牌可以通过旋转 0度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。
	//形式上，dominoes[i] = [a, b]和dominoes[j] = [c, d]等价的前提是a==c且b==d，或是a==d 且b==c。
	//在0 <= i < j < dominoes.length的前提下，找出满足dominoes[i] 和dominoes[j]等价的骨牌对 (i, j) 的数量。
	//链接：https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs
	//***//
	//dominoes := [][]int{{1,2},{2,1},{3,4},{5,6}}
	//fmt.Println(lc.NumEquivDominoPairs(dominoes))

	//***
	//Alice 和 Bob 共有一个无向图，其中包含 n 个节点和 3 种类型的边：
	//
	//类型 1：只能由 Alice 遍历。
	//类型 2：只能由 Bob 遍历。
	//类型 3：Alice 和 Bob 都可以遍历。
	//给你一个数组 edges ，其中 edges[i] = [typei, ui, vi]表示节点 ui 和 vi 之间存在类型为 typei 的双向边。请你在保证图仍能够被 Alice和 Bob 完全遍历的前提下，找出可
	//以删除的最大边数。如果从任何节点开始，Alice 和 Bob 都可以到达所有其他节点，则认为图是可以完全遍历的。
	//返回可以删除的最大边数，如果 Alice 和 Bob 无法完全遍历图，则返回 -1 。
	//链接：https://leetcode-cn.com/problems/remove-max-number-of-edges-to-keep-graph-fully-traversable
	//***//
	//n := 4
	//edges := [][]int{{3,1,2},{3,2,3},{1,1,3},{1,2,4},{1,1,2},{2,3,4}}
	//fmt.Println(lc.MaxNumEdgesToRemove(n, edges))

	//---
	//channel接收阻塞测试
	//---
	//done := make(chan bool)
	//go myLab.AGoroutine()
	//myLab.Done <- true
	//fmt.Println(myLab.Msg)

	//***
	//给你一个整数数组nums，请编写一个能够返回数组 “中心索引” 的方法。 ···同剑指offer12 2109···
	//数组 中心索引 是数组的一个索引，其左侧所有元素相加的和等于右侧所有元素相加的和。
	//如果数组不存在中心索引，返回 -1 。如果数组有多个中心索引，应该返回最靠近左边的那一个。
	//注意：中心索引可能出现在数组的两端。
	//链接：https://leetcode-cn.com/problems/find-pivot-index
	//***//
	//nums := []int{1,7,3,6,5,6}
	//fmt.Println(lc.PivotIndex(nums))
	//***

	//你准备参加一场远足活动。给你一个二维rows x columns的地图heights，其中heights[row][col]表示格子(row, col)的高度。一开始你在最左上角的格子(0, 0)，且你希望去最右下角的格子(rows-1, columns-1)
	//（注意下标从 0 开始编号）。你每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
	//一条路径耗费的 体力值是路径上相邻格子之间 高度差绝对值的 最大值决定的。
	//请你返回从左上角走到右下角的最小体力消耗值。
	//链接：https://leetcode-cn.com/problems/path-with-minimum-effort
	//***//
	//heights := [][]int{{1,2,2},{3,8,2},{5,3,5}}
	//fmt.Println(lc.MinimumEffortPath(heights))

	//给定两个（单向）链表，判定它们是否相交并返回交点。请注意相交的定义基于节点的引用，而不是基于节点的值。换句话说，如果一个链表的第k个节点与另一个链表的第j个节点是同一节点（引用完全相同），则这两个链表相交。
	//链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists-lcci
	//***//
	//listA := &lc.ListNode{
	//	4,
	//	&lc.ListNode{
	//		1,
	//		&lc.ListNode{
	//			8,
	//			&lc.ListNode{
	//				4,
	//				&lc.ListNode{5,nil},
	//			},
	//		},
	//	},
	//}
	//listB := &lc.ListNode{
	//	5,
	//	&lc.ListNode{
	//		0,
	//		&lc.ListNode{
	//			1,
	//			&lc.ListNode{
	//				8,
	//				&lc.ListNode{4,&lc.ListNode{5,nil}},
	//			},
	//		},
	//	},
	//}
	//fmt.Println(lc.GetIntersectionNode(listA, listB))

	//***
	//如果交换字符串X 中的两个不同位置的字母，使得它和字符串Y 相等，那么称 X 和 Y 两个字符串相似。如果这两个字符串本身是相等的，那它们也是相似的。
	//例如，"tars" 和 "rats" 是相似的 (交换 0 与 2 的位置)；"rats" 和 "arts" 也是相似的，但是 "star" 不与 "tars"，"rats"，或 "arts" 相似。
	//总之，它们通过相似性形成了两个关联组：{"tars", "rats", "arts"} 和 {"star"}。注意，"tars" 和 "arts" 是在同一组中，即使它们并不相似。形式上，对每个组
	//而言，要确定一个单词在组中，只需要这个词和该组中至少一个单词相似。
	//给你一个字符串列表 strs。列表中的每个字符串都是 strs 中其它所有字符串的一个字母异位词。请问 strs 中有多少个相似字符串组？
	//链接：https://leetcode-cn.com/problems/similar-string-groups
	//***//
	//strs := []string{"tars","rats","arts","stars"}
	//fmt.Println(lc.NumSimilarGroups(strs))

	//***
	//爱丽丝和鲍勃有不同大小的糖果棒：A[i] 是爱丽丝拥有的第 i 根糖果棒的大小，B[j] 是鲍勃拥有的第 j 根糖果棒的大小。
	//因为他们是朋友，所以他们想交换一根糖果棒，这样交换后，他们都有相同的糖果总量。（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）
	//返回一个整数数组 ans，其中 ans[0] 是爱丽丝必须交换的糖果棒的大小，ans[1]是 Bob 必须交换的糖果棒的大小。
	//如果有多个答案，你可以返回其中任何一个。保证答案存在。
	//链接：https://leetcode-cn.com/problems/fair-candy-swap
	//***//
	//A, B := []int{1,1}, []int{2,2}
	//fmt.Println(lc.FairCandySwap(A, B))

	//***
	//给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，总共可最多替换k次。在执行上述操作后，找到包含重复字母的最长子串的长度。
	//注意：字符串长度 和 k 不会超过104。
	//链接：https://leetcode-cn.com/problems/longest-repeating-character-replacement
	//***//
	//s := "ABAB"
	//k := 2
	//fmt.Println(lc.CharacterReplacement(s, k))

	//***
	//中位数是有序序列最中间的那个数。如果序列的长度是偶数，则没有最中间的数；此时中位数是最中间的两个数的平均数。
	//
	//例如：
	//
	//[2,3,4]，中位数是3
	//[2,3]，中位数是 (2 + 3) / 2 = 2.5
	//给你一个数组 nums，有一个长度为 k 的窗口从最左端滑动到最右端。窗口中有 k 个数，每次窗口向右移动 1 位。你的任务是找出每次窗口移动后得到的新窗口中元素的中位数，并输出由它们组成的数组。
	//链接：https://leetcode-cn.com/problems/sliding-window-median
	//***//
	//nums := []int{1,3,-1,-3,5,3,6,7}
	//k := 3
	//fmt.Println(lc.MedianSlidingWindow(nums, k))

	//***

	//几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。
	//
	//每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。
	//
	//你的点数就是你拿到手中的所有卡牌的点数之和。
	//
	//给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。
	//
	//来源：力扣（LeetCode）
	//链接：https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards
	//***//
	//cardPoints := []int{1,2,3,4,5,6,1}
	//k := 3
	//fmt.Println(lc.MaxScore(cardPoints, k))

	//给你两个长度相同的字符串，s 和 t。
	//将 s中的第i个字符变到t中的第 i 个字符需要|s[i] - t[i]|的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。
	//用于变更字符串的最大预算是maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。
	//如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。
	//如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。
	//链接：https://leetcode-cn.com/problems/get-equal-substrings-within-budget
	//***//
	//s:="abcd"
	//t:="bcdf"
	//cost:=3
	//fmt.Println(lc.EqualSubstring(s,t,cost))

	//***
	//给你一个长度为n的整数数组，请你判断在 最多 改变1 个元素的情况下，该数组能否变成一个非递减数列。
	//我们是这样定义一个非递减数列的：对于数组中所有的i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。
	//链接：https://leetcode-cn.com/problems/non-decreasing-array
	//***//
	//nums := []int{4,2,3}
	//fmt.Println(lc.CheckPossibility(nums))

	//***
	//给你一个整数数组 nums 和一个整数 target 。
	//请你统计并返回 nums 中能满足其最小元素与最大元素的 和 小于或等于 target 的 非空 子序列的数目。
	//由于答案可能很大，请将结果对 10^9 + 7 取余后返回。
	//链接：https://leetcode-cn.com/problems/number-of-subsequences-that-satisfy-the-given-sum-condition
	//***//
	//nums := []int{3,5,6,7}
	//target := 9
	//fmt.Println(lc.NumSubseq(nums , target))

	//***
	//设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
	//请实现 KthLargest类：
	//KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
	//int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。
	//链接：https://leetcode-cn.com/problems/kth-largest-element-in-a-stream
	//***//
	//pre := lc.ConstructorKth(3,[]int{4,5,8,2})
	//pre.Add(3)
	//pre.Add(5)
	//pre.Add(10)
	//pre.Add(9)
	//pre.Add(4)

	//***
	//给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
	//***//
	//k := 3
	//fmt.Println(lc.GetRow(k))

	//***
	//N 对情侣坐在连续排列的 2N 个座位上，想要牵到对方的手。 计算最少交换座位的次数，以便每对情侣可以并肩坐在一起。 一次交换可选择任意两人，让他们站起来交换座位。
	//人和座位用0到2N-1的整数表示，情侣们按顺序编号，第一对是(0, 1)，第二对是(2, 3)，以此类推，最后一对是(2N-2, 2N-1)。
	//这些情侣的初始座位row[i]是由最初始坐在第 i 个座位上的人决定的。
	//链接：https://leetcode-cn.com/problems/couples-holding-hands
	//***//
	//row := []int{0,2,1,3}
	//fmt.Println(lc.MinSwapsCouples(row))

	//***
	//给定一个二进制数组， 计算其中最大连续1的个数。
	//***//
	//fmt.Println(lc.FindMaxConsecutiveOnes([]int{1,1,0,1,1,1}))

	//***
	//给定长度为2n的整数数组 nums ，你的任务是将这些数分成n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从 1 到n 的 min(ai, bi) 总和最大。
	//返回该 最大总和 。
	//链接：https://leetcode-cn.com/problems/array-partition-i
	//***//
	//fmt.Println(lc.ArrayPairSum([]int{1,4,3,2}))

	//***
	//在MATLAB中，有一个非常有用的函数 reshape，它可以将一个矩阵重塑为另一个大小不同的新矩阵，但保留其原始数据。
	//
	//给出一个由二维数组表示的矩阵，以及两个正整数r和c，分别表示想要的重构的矩阵的行数和列数。
	//
	//重构后的矩阵需要将原始矩阵的所有元素以相同的行遍历顺序填充。
	//
	//如果具有给定参数的reshape操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
	//链接：https://leetcode-cn.com/problems/reshape-the-matrix
	//***//
	//nums := [][]int{{1,2},{3,4}}
	//fmt.Println(lc.MatrixReshape(nums,1,4))

	//***
	//在仅包含 0 和 1 的数组 A 中，一次 K 位翻转包括选择一个长度为 K 的（连续）子数组，同时将子数组中的每个 0 更改为 1，而每个 1 更改为 0。
	//返回所需的 K 位翻转的最小次数，以便数组没有值为 0 的元素。如果不可能，返回 -1。
	//链接：https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips
	//***//
	//A,k := []int{0,1,0}, 1
	//fmt.Println(lc.MinKBitFlips(A,k))

	//***
	//给定一个由若干 0 和 1 组成的数组 A，我们最多可以将 K 个值从 0 变成 1 。
	//返回仅包含 1 的最长（连续）子数组的长度。
	//***//
	//A := []int{1,1,1,0,0,0,1,1,1,1,0};k := 2
	//fmt.Println(lc.LongestOnes(A,k))

	//***
	//找出数组中重复的数字
	//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。
	//链接：https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof
	//***//
	//fmt.Println(lc.FindRepeatNumbe([]int{2,3,1,0,2,5,3}))

	//***
	//给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。
	//如果不存在满足条件的子数组，则返回 0 。
	//链接：https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit
	//***//
	//nums := []int{8,2,4,7}
	//limit := 4
	//fmt.Println(lc.LongestSubarray(nums,limit))

	//***
	//今天，书店老板有一家店打算试营业customers.length分钟。每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。
	//在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。
	//书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续X 分钟不生气，但却只能使用一次。
	//请你返回这一天营业下来，最多有多少客户能够感到满意的数量。
	//链接：https://leetcode-cn.com/problems/grumpy-bookstore-owner
	//***//
	//customers := []int{1,0,1,2,1,1,7,5}
	//grumpy := []int{0,1,0,1,0,1,0,1}
	//x := 3
	//fmt.Println(lc.MaxSatisfied(customers, grumpy, x))

	//***
	//给定一个二进制矩阵A，我们想先水平翻转图像，然后反转图像并返回结果。
	//水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转[1, 1, 0]的结果是[0, 1, 1]。
	//反转图片的意思是图片中的0全部被1替换，1全部被0替换。例如，反转[0, 1, 1]的结果是[1, 0, 0]。
	//链接：https://leetcode-cn.com/problems/flipping-an-image
	//***//
	//A := [][]int{{1,1,0},{1,0,1},{0,0,0}}
	//fmt.Println(lc.FlipAndInvertImage(A))

	//***
	//给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。
	//矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
	//***//
	//matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	//fmt.Println(lc.Transpose(matrix))

	//***
	//外国友人仿照中国字谜设计了一个英文版猜字谜小游戏，请你来猜猜看吧。
	//字谜的迷面puzzle 按字符串形式给出，如果一个单词word符合下面两个条件，那么它就可以算作谜底：
	//单词word中包含谜面puzzle的第一个字母。
	//单词word中的每一个字母都可以在谜面puzzle中找到。
	//例如，如果字谜的谜面是 "abcdefg"，那么可以作为谜底的单词有 "faced", "cabbage", 和 "baggage"；而 "beefed"（不含字母 "a"）以及"based"（其中的 "s" 没有出现在谜面中）都不能作为谜底。
	//返回一个答案数组answer，数组中的每个元素answer[i]是在给出的单词列表 words 中可以作为字谜迷面puzzles[i]所对应的谜底的单词数目。
	//链接：https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle
	//***//

	//***
	//给你一个字符串 s 和一个整数 k ，请你找出 s 中的最长子串， 要求该子串中的每一字符出现次数都不少于 k 。返回这一子串的长度。
	//***//
	//s := "aaabb"
	//k := 3
	//fmt.Println(lc.LongestSubstring(s, k))

	//***
	//给你两个长度相同的整数数组target和arr。
	//
	//每一步中，你可以选择arr的任意 非空子数组并将它翻转。你可以执行此过程任意次。
	//
	//如果你能让 arr变得与 target相同，返回 True；否则，返回 False 。
	//链接：https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays
	//***//
	//target := []int{1,2,3,4}
	//arr := []int{2,4,1,3}
	//fmt.Println(lc.CanBeEqual(target, arr))

	//***
	//给定一个整数数组 nums，求出数组从索引i到j（i≤j）范围内元素的总和，包含i、j两点。
	//实现 NumArray 类：
	//NumArray(int[] nums) 使用数组 nums 初始化对象
	//int sumRange(int i, int j) 返回数组 nums 从索引i到j（i≤j）范围内元素的总和，包含i、j两点（也就是 sum(nums[i], nums[i + 1], ... , nums[j])）
	//链接：https://leetcode-cn.com/problems/range-sum-query-immutable
	//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
	//***//
	//obj := lc.ConstructorRange([]int{1,2,3,4})
	//param := obj.SumRange(1,2)
	//fmt.Println(param)

	//***
	//给定一个二维矩阵，计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为 (row2, col2) 。
	//***//
	//matrix := [][]int{
	//{3, 0, 1, 4, 2},
	//{5, 6, 3, 2, 1},
	//{1, 2, 0, 1, 5},
	//{4, 1, 0, 1, 7},
	//{1, 0, 3, 0, 5}}
	//obj := lc.ConstructorRegion(matrix)
	//param_1 := obj.SumRegion(2,1,4,3)
	//fmt.Println(param_1)

	//***
	//给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
	//***//
	//fmt.Println(lc.CountBits(2))

	//***
	//给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
	//当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
	//请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
	//注意：不允许旋转信封。
	//链接：https://leetcode-cn.com/problems/russian-doll-envelopes
	//***//
	//envelopes := [][]int{{5,4},{6,4},{6,7},{2,3}}
	//fmt.Println(lc.MaxEnvelopes(envelopes))

	//***
	//实现一个基本的计算器来计算一个简单的字符串表达式 s 的值。
	//***//
	//fmt.Println(lc.Calculate("2-(1+1)"))

	//***
	//给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
	//***//
	//fmt.Println(lc.Calculateii("1+3 / 2"))

	//***
	//不使用任何内建的哈希表库设计一个哈希集合（HashSet）。
	//实现 MyHashSet 类：
	//void add(key) 向哈希集合中插入值 key 。
	//bool contains(key) 返回哈希集合中是否存在这个值 key 。
	//void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
	//链接：https://leetcode-cn.com/problems/design-hashset
	//****//
	//set := lc.Constructoriii()
	//set.Add(222)
	//set.Add(333)
	//set.Remove(222)
	//fmt.Println(set.Contains(222))

	//不使用任何内建的哈希表库设计一个哈希映射（HashMap）。
	//实现 MyHashMap 类：
	//void add(key) 向哈希集合中插入值 key 。
	//bool contains(key) 返回哈希集合中是否存在这个值 key 。
	//void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
	//set := lc.Constructoriiii()
	//set.Put(222,222)
	//set.Get(222)
	//set.Remove(222)
	//fmt.Println(set)

	//***
	//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
	//***//
	//matrix := [][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}
	//fmt.Println(lc.SpiralOrder(matrix))

	//***
	//给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
	//***//
	//fmt.Println(lc.GenerateMatrix(3))

	//***
	//给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
	//字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE"是"ABCDE"的一个子序列，而"AEC"不是）
	//题目数据保证答案符合 32 位带符号整数范围。
	//链接：https://leetcode-cn.com/problems/distinct-subsequences
	//***//
	//s := "rabbbit"
	//t := "rabbit"
	//fmt.Println(lc.NumDistinct(s, t))

	//***
	//给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。
	//如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。
	//进阶：很容易想到时间复杂度为 O(n^2) 的解决方案，你可以设计一个时间复杂度为 O(n logn) 或 O(n) 的解决方案吗？
	//链接：https://leetcode-cn.com/problems/132-pattern
	//***//
	//nums := []int{1,2,3,4}
	//fmt.Println(lc.Find132pattern(nums))

	//***
	//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
	//返回同样按升序排列的结果链表。
	//***//
	//list := &lc.ListNode{1,&lc.ListNode{1,&lc.ListNode{2,&lc.ListNode{2,nil}}}}
	//fmt.Println(lc.DeleteDuplicates(list))

	//***
	//已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
	//若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
	//若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
	//注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
	//给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
	//链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii
	//***//
	//nums := []int{1,3,5}
	//fmt.Println(lc.FindMin(nums))

	//***
	//给你两个整数，n 和 start 。
	//数组 nums 定义为：nums[i] = start + 2*i（下标从 0 开始）且 n == nums.length 。
	//请返回 nums 中所有元素按位异或（XOR）后得到的结果。
	//链接：https://leetcode-cn.com/problems/xor-operation-in-an-array
	//***//
	//n := 5
	//start := 0
	//fmt.Println(lc.XorOperation(n, start))

	//*** easy
	//给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
	//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
	//你可以按任意顺序返回答案。
	//链接：https://leetcode-cn.com/problems/two-sum
	//***//
	//nums := []int{2, 7, 11, 15}
	//target := 9
	//fmt.Println(lc.TwoSum(nums, target))

	//***
	//给定一个三角形 triangle ，找出自顶向下的最小路径和。
	//每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
	//链接：https://leetcode-cn.com/problems/triangle
	//***//
	//triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	//fmt.Println(lc.MinimumTotal(triangle))

	//***
	//剑指offer 01
	//给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%'。
	//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8以及truncate(-2.7335) = -2
	//假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,231−1]。本题中，如果除法结果溢出，则返回 231− 1
	//链接：https://leetcode-cn.com/problems/xoh6Oh
	//***//
	//fmt.Println(lc.Divide(12, 3))

	//***
	//剑指offer 02
	//给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
	//输入为 非空 字符串且只包含数字 1 和 0。
	//***//
	//fmt.Println(lc.AddBinary("11", "10"))

	//***
	//剑指offer 03 计算前N位数二进制的1数量
	//给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
	//***//
	//lc.CountBits(2)

	//***
	//剑指offer 04 只出现一次的数字
	//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
	//***//
	//nums := []int{2,2,3,2}
	//fmt.Println(lc.SingleNumber(nums))

	//***
	//剑指offer 05 单词长度最大乘积
	//给定一个字符串数组words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
	//链接：https://leetcode-cn.com/problems/aseY1I
	//***//
	//words := []string{"abcw", "baz", "foo", "bar", "fxyz", "abcdef"}
	//fmt.Println(lc.MaxProduct(words))

	//***
	//剑指offer 06 排序数组中两个数字之和
	//给定一个已按照 升序排列 的整数数组numbers ，请你从数组中找出两个数满足相加之和等于目标数target 。
	//函数应该以长度为 2 的整数数组的形式返回这两个数的下标值。numbers 的下标 从 0开始计数 ，所以答案数组应当满足 0<= answer[0] < answer[1] <numbers.length。
	//假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。
	//链接：https://leetcode-cn.com/problems/kLl5u1
	//***//
	//numbers, target := []int{1, 2, 4, 6, 10}, 8
	//fmt.Println(lc.TwoNumberSum(numbers, target))

	//***
	//剑指offer 07 排序数组中两个数字之和
	//给定一个包含 n 个整数的数组nums，判断nums中是否存在三个元素a ，b ，c ，使得a + b + c = 0 ？请找出所有和为 0 且不重复的三元组。
	//链接：https://leetcode-cn.com/problems/1fGaJU
	//***//
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//fmt.Println(lc.ThreeSum(nums))

	//***
	//剑指offer 08 和大于等于target的最短子数组
	//给定一个含有n个正整数的数组和一个正整数 target 。
	//找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
	//链接：https://leetcode-cn.com/problems/2VG8Kg
	//target, nums := 7, []int{1,1,1,1,1}
	//fmt.Println(lc.MinSubArrayLen(target, nums))

	//***
	//剑指offer 09 乘积小于K的连续子数组数量
	//给定一个正整数数组 nums和整数 k ，请找出该数组内乘积小于 k 的连续的子数组的个数。
	//nums := []int{10,5,2,6}
	//k := 100
	//fmt.Println(lc.NumsSubarrayProductLessThank(nums,k))

	//***
	//剑指offer 10 和为k的子数组
	//给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数。
	//nums := []int{1, 2, 3}
	//k := 3
	//fmt.Println(lc.SubarraySum(nums, k))

	//***
	//剑指offer 11 0和1个数相同的子数组
	//给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
	//***//
	//nums := []int{1, 0, 1}
	//fmt.Println(lc.FindMaxLength(nums))

	//***
	//剑指offer 12 左右两边子数组和相等 ···2109···
	//给你一个整数数组nums ，请计算数组的 中心下标 。
	//数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
	//如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
	//如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。
	//链接：https://leetcode-cn.com/problems/tvdfij
	//***//
	//nums := []int{1,7,3,6,5,6}
	//fmt.Println(lc.PivotIndex(nums))

	//***
	//剑指offer 13 二维子矩阵的和
	//给定一个二维矩阵 matrix，以下类型的多个请求：
	//
	//计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1,col1) ，右下角为 (row2,col2) 。
	//实现 NumMatrix 类：
	//
	//NumMatrix(int[][] matrix)给定整数矩阵 matrix 进行初始化
	//int sumRegion(int row1, int col1, int row2, int col2)返回左上角 (row1,col1)、右下角(row2,col2)的子矩阵的元素总和。
	//链接：https://leetcode-cn.com/problems/O4NDxx
	//***//
	//obj := lc.Constructor2110([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})
	//fmt.Println(obj.SumRegion(2, 1, 4, 3))
	//fmt.Println(obj.SumRegion(1, 1, 2, 2))
	//fmt.Println(obj.SumRegion(1, 2, 2, 4))

	//***
	//剑指 Offer II 014. 字符串中的变位词
	//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
	//换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
	//***//
	//s1, s2 := "ab", "eidbaooo"
	//fmt.Println(lc.CheckInclusion(s1, s2))

	//***
	//剑指 Offer II 015 字符串中所有的变位词
	//给定两个字符串s和p，找到s中所有 p 的变位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
	//变位词 指字母相同，但排列不同的字符串。
	//链接：https://leetcode-cn.com/problems/VabMRr
	//***//
	//s, p := "cbaebabacd", "abc"
	//fmt.Println(lc.FindAnagrams(s, p))

	//***
	//剑指offer 16 不含重复字符的最长子字符串
	//给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
	//***//
	//s := "abcabcbb"
	//fmt.Println(lc.LengthOfLongestSubstring(s))

	//***
	//剑指 Offer II 017 含有所有字符串的最短字符串
	//给定两个字符串 s 和t 。返回 s 中包含t的所有字符的最短子字符串。如果 s 中不存在符合条件的子字符串，则返回空字符串 "" 。
	//如果 s 中存在多个符合条件的子字符串，返回任意一个。
	//注意： 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
	//链接：https://leetcode-cn.com/problems/M1oyTv
	//***//
	//s, t := "ADOBECODEBANC", "ABC"
	//fmt.Println(lc.MinWindow(s, t))

	//***
	////剑指 Offer II 018 有效回文
	//给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
	//本题中，将空字符串定义为有效的 回文串 。
	//***//
	//s := "0P2112P0"
	//fmt.Println(lc.IsPalindrome(s))

	//***
	//剑指 Offer II 019 最多删除一个字符得到回文串
	//给定一个非空字符串 s，请判断如果 最多 从字符串中删除一个字符能否得到一个回文字符串。
	//***//
	//s := "abca"
	//fmt.Println(lc.ValidPalindrome(s))

	//***
	//剑指 Offer II 020 计算回文子字符串的个数 2118
	//给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
	//
	//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
	//***
	//s := "abc"
	//fmt.Println(lc.CountSubstrings(s))

	//***
	//剑指 Offer II 021. 删除链表的倒数第 n 个结点 2125
	//给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
	//***//
	//head := &lc.ListNode{1, &lc.ListNode{2, &lc.ListNode{3, &lc.ListNode{4, &lc.ListNode{5, nil}}}}}
	//n := 2
	//fmt.Println(lc.RemoveNthFromEnd(head, n))

	//***
	//剑指 Offer II 022 链表中环的入口
	//给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。如果链表无环，则返回 null。
	//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
	//说明：不允许修改给定的链表。
	//链接：https://leetcode-cn.com/problems/c32eOV
	//***//
	//head := &lc.ListNode{3, &lc.ListNode{3, &lc.ListNode{3, &lc.ListNode{3, nil}}}}
	//fmt.Println(lc.DetectCycle(head))

	//***
	//剑指 Offer II 023. 两个链表的第一个重合节点
	//给定两个单链表的头节点 headA 和 headB ，请找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
	//***//
	//A := &lc.ListNode{4, &lc.ListNode{1, &lc.ListNode{8, &lc.ListNode{4, &lc.ListNode{5, nil}}}}}
	//B := &lc.ListNode{5, &lc.ListNode{0, &lc.ListNode{1, &lc.ListNode{8, &lc.ListNode{4, &lc.ListNode{5, nil}}}}}}
	//fmt.Println(lc.GetIntersectionNode(A, B))

	//***
	//剑指 Offer II 024. 反转链表
	//给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。
	//***//
	//head := &lc.ListNode{1, &lc.ListNode{2, &lc.ListNode{3, &lc.ListNode{4, &lc.ListNode{5, nil}}}}}
	//fmt.Println(lc.ReverseList(head))

	//***
	//剑指 Offer II 025. 链表中的两数相加
	//给定两个 非空链表 l1和 l2来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
	//可以假设除了数字 0 之外，这两个数字都不会以零开头。
	//链接：https://leetcode-cn.com/problems/lMSNwu
	//***//
	//l1 := &lc.ListNode{7, &lc.ListNode{2, &lc.ListNode{4, &lc.ListNode{3, nil}}}}
	//l2 := &lc.ListNode{5, &lc.ListNode{6, &lc.ListNode{4, nil}}}
	//fmt.Println(lc.AddTwoNumbers(l1, l2))

	//***
	//剑指 Offer II 026. 重排链表
	//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
	//L0→ L1→ … → Ln-1→ Ln
	//请将其重新排列后变为：
	//L0→Ln→L1→Ln-1→L2→Ln-2→ …
	//不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
	//链接：https://leetcode-cn.com/problems/LGjMqU
	//***//
	//l1 := &lc.ListNode{7, &lc.ListNode{2, &lc.ListNode{4, &lc.ListNode{3, nil}}}}
	//lc.ReorderList(l1)
	//fmt.Println(l1)

	//***
	//剑指 Offer II 027. 回文链表
	//给定一个链表的 头节点 head ，请判断其是否为回文链表。
	//如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
	//***//
	head := &lc.ListNode{1, &lc.ListNode{2, &lc.ListNode{2, &lc.ListNode{1, nil}}}}
	fmt.Println(lc.IsListPalindrome(head))

	//***
	//剑指offer ii 28 展开多级双向链表
	//多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。
	//给定位于列表第一级的头节点，请扁平化列表，即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。
	//链接：https://leetcode-cn.com/problems/Qv1Da2
	//***//
	//懒得构筑双链表

	//***
	//剑指 Offer II 029. 排序的循环链表 2133
	//给定循环升序列表中的一个点，写一个函数向这个列表中插入一个新元素insertVal ，使这个列表仍然是循环升序的。
	//给定的可以是这个列表中任意一个顶点的指针，并不一定是这个列表中最小元素的指针。
	//如果有多个满足条件的插入位置，可以选择任意一个位置插入新的值，插入后整个列表仍然保持有序。
	//如果列表为空（给定的节点是 null），需要创建一个循环有序列表并返回这个节点。否则。请返回原先给定的节点。
	//链接：https://leetcode-cn.com/problems/4ueAj6
	//***//
	//懒得构建循环链表
	//lc.InsertListNode()

	//***
	//剑指 Offer II 030. 插入、删除和随机访问都是 O(1) 的容器
	//设计一个支持在平均时间复杂度 O(1)下，执行以下操作的数据结构：
	//insert(val)：当元素 val 不存在时返回 true，并向集合中插入该项，否则返回 false 。
	//remove(val)：当元素 val 存在时返回 true，并从集合中移除该项，否则返回 false。
	//getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。
	//链接：https://leetcode-cn.com/problems/FortPu
	//***//
	obj := lc.ConstructorOFFER30()

	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Remove(2))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.Insert(2))
	obj.GetRandom()
	fmt.Println(obj)
}
