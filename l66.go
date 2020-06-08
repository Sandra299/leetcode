/*
面试题66. 构建乘积数组
给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

示例:

输入: [1,2,3,4,5]
输出: [120,60,40,30,24]

提示：

所有元素乘积之和不会溢出 32 位整数
a.length <= 100000

https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
*/
package main

import "fmt"

func constructArr(a []int) []int {
	if a == nil || len(a) == 0 {
		return nil
	}
	size := len(a)

	upTriangle, downTriangle := make([]int, size), make([]int, size)
	// 从上往下
	upTriangle[0] = 1
	for i := 1; i < size; i++ {
		upTriangle[i] = upTriangle[i-1] * a[i-1]
	}
	// 从下往上
	downTriangle[size-1] = 1
	for i := size - 2; i >= 0; i-- {
		downTriangle[i] = downTriangle[i+1] * a[i+1]
	}
	b := make([]int, len(a))
	for i := 0; i < size; i++ {
		b[i] = upTriangle[i] * downTriangle[i]
	}
	return b
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	array := make([]int, len(a))
	array[0] = a[0]
	for i := 1; i < len(a); i++ {
		array[i] = array[i-1] * a[i]
	}
	fmt.Println(array)
	fmt.Println(constructArr(a))
}
