package DoublePointer

/*
给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
*/
func maxArea(height []int) (ans int) {
	left, right := 0, len(height)-1
	for left < right {
		ans = max(min(height[left], height[right])*(right-left), ans)
		if left < right {
			left++
		} else {
			right--
		}
	}
	return
}
