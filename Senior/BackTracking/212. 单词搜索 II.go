package BackTracking

/*
给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words， 返回所有二维网格上的单词 。
单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单
元格。同一个单元格内的字母在一个单词中不允许被重复使用。

示例 1：
输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]],
words = ["oath","pea","eat","rain"]
输出：["eat","oath"]

示例 2：
输入：board = [["a","b"],["c","d"]], words = ["abcb"]
输出：[]
*/

// Trie 前缀树结构体 在单词的最后一个节点记录当前word是什么
type Trie struct {
	children [26]*Trie
	word     string
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.word = word
}

var directions = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) []string {
	// 声明一个前缀树实例
	t := &Trie{}
	// 初始化前缀树 将word添加到前缀树中
	for _, word := range words {
		t.Insert(word)
	}
	m, n := len(board), len(board[0])
	// 用于记录当前这个格子有没有被访问过
	seen := map[string]bool{}

	var dfs func(node *Trie, x, y int)
	// x,y 分别代表横纵坐标
	dfs = func(node *Trie, x, y int) {
		// ch为当前遍历到的字符
		ch := board[x][y]
		// node更新为当前字符所在的节点
		node = node.children[ch-'a']
		// 节点为空证明不是前缀
		if node == nil {
			return
		}
		// 如果node.word不是空白字符 证明找到了目标字符
		if node.word != "" {
			seen[node.word] = true
		}
		// board[x][y]已经被遍历过了
		board[x][y] = '#'
		for _, d := range directions {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
				dfs(node, nx, ny)
			}
		}
		// 恢复现场
		board[x][y] = ch
	}
	// 遍历board 从每个格子出发 进行深度优先遍历
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}
	ans := make([]string, 0, len(seen))
	for s := range seen {
		ans = append(ans, s)
	}
	return ans
}

// 根据自己理解写的
func findWords1(board [][]byte, words []string) (ans []string) {
	// 实例化一个Trie
	t := &Trie{}
	m, n := len(board), len(board[0])
	// 标记被访问的格子
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}
	// 将word添加进入前缀树
	for _, word := range words {
		t.Insert(word)
	}
	seen := map[string]bool{}
	var dfs func(node *Trie, row, column int)
	dfs = func(node *Trie, row, column int) {
		if row < 0 || row >= m || column < 0 || column >= n || used[row][column] {
			return
		}
		ch := board[row][column]
		node = node.children[ch-'a']
		// 空节点 证明不是word
		if node == nil {
			return
		}
		// word字段不为空 证明找到了word
		if node.word != "" {
			seen[node.word] = true
		}
		used[row][column] = true
		dfs(node, row-1, column)
		dfs(node, row+1, column)
		dfs(node, row, column-1)
		dfs(node, row, column+1)
		used[row][column] = false
	}
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}
	for s := range seen {
		if seen[s] {
			ans = append(ans, s)
		}
	}
	return
}
