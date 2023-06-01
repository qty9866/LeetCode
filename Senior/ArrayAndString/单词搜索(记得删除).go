package ArrayAndString

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
	// 在结束的这个节点记录一下这个单词
	node.word = word
}

func findWords(board [][]byte, words []string) (ans []string) {
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
