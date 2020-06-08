/*
面试题29. 顺时针打印矩阵

输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例 1：

输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：

输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

解题思路：
[[1, 1, 1, 1, 1, 1, 1],
 [1, 2, 2, 2, 2, 2, 1],
 [1, 2, 3, 3, 3, 2, 1],
 [1, 2, 2, 2, 2, 2, 1],
 [1, 1, 1, 1, 1, 1, 1]]
*/

package main

import "fmt"

type Matrix [][]int

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	// 切片容量
	size := len(matrix) * len(matrix[0])
	step := 0

	array := make([]int, size)

	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for step < size {
		// 从左到右
		for i := left; i <= right && step < size; i++ {
			fmt.Print(matrix[top][i], " ")
			array[step] = matrix[top][i]
			step++
		}
		top++
		// 从上到下
		for i := top; i <= bottom && step < size; i++ {
			fmt.Print(matrix[i][right], " ")
			array[step] = matrix[i][right]
			step++
		}
		right--
		// 从右到左
		for i := right; i >= left && step < size; i-- {
			fmt.Print(matrix[bottom][i], " ")
			array[step] = matrix[bottom][i]
			step++
		}
		bottom--
		// 从下到上
		for i := bottom; i >= top && step < size; i-- {
			fmt.Print(matrix[i][left], " ")
			array[step] = matrix[i][left]
			step++
		}
		left++
	}
	return array
}

func main() {

	m1 := Matrix{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	m2 := Matrix{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}
	m3 := Matrix{
		[]int{1, 2, 3, 4, 5},
		[]int{5, 6, 7, 8, 9},
		[]int{10, 11, 12, 13, 14},
		[]int{15, 16, 17, 18, 19},
	}
	spiralOrder(m1)
	fmt.Println()
	spiralOrder(m2)
	fmt.Println()
	spiralOrder(m3)
	fmt.Println()
}
