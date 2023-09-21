package main

import "fmt"

type sliceStruct struct {
	slice []int
}

func NewIntSlice(sl []int) *sliceStruct {
	return &sliceStruct{slice: sl}
}

func (is *sliceStruct) Sort() {
	for i := 0; i < len(is.slice)-1; i++ {
		for j := 0; j < len(is.slice)-i-1; j++ {
			if is.slice[j] > is.slice[j+1] {
				is.slice[j], is.slice[j+1] = is.slice[j+1], is.slice[j]
			}
		}
	}
}

func compareSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	// Сортируем оба слайса
	sliceCopy1 := NewIntSlice(slice1)
	sliceCopy1.Sort()
	sliceCopy2 := NewIntSlice(slice2)
	sliceCopy2.Sort()

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(compareSlices([]int{1,2,3},[]int{2,3,1}))
}