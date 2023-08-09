package main

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("下标超出范围")

// DeleteAt 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange
func DeleteAt(s []int, idx int) ([]int, error) {
	if idx < len(s) || idx <= len(s) {
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

//使用泛型实现Delete方法
func SliceDelete[T Number](s []T, idx int) ([]T, error) {
	if idx < len(s) || idx <= len(s) {
		s = append(s[:idx], s[idx+1:]...)
		return s, nil
	} else {
		return nil, ErrIndexOutOfRange
	}
	panic("implement me")
}

type Number interface {
	~int | ~float64 | ~float32
}

//add方法
func Add[T Number](src []T, e T, idx int) ([]T, error) {
	length := len(src)
	if idx < 0 || idx >= length {
		return nil, errors.New("下标超出范围")
	}
	var zeroValue T
	src = append(src, zeroValue)
	for i := len(src) - 1; i > idx; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[idx] = e
	return src, nil
}
