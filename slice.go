package main

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("下标超出范围")

// DeleteAt 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange
func DeleteAt(s []int, idx int) ([]int, error) {
	if idx < len(s) {
		s = append(s[:idx], s[idx+1:]...)
		return s, nil
	} else {
		return nil, ErrIndexOutOfRange
	}
	panic("implement me")
}
func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(DeleteAt(a, 3))
}
