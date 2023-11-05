package main

import "testing"

func TestCompareSlices(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
		name: "EqualSlices",
		args: args{
			slice1: []int{1,2,3},
			slice2: []int{1,2,3},
		},
		want: true,
	},
	{
		name: "DifferentSlices",
		args: args{
			slice1: []int{1,2,3},
			slice2: []int{3,2,1},
		},
		want: true,
	},
	{
		name: "DifferentLength",
		args: args{
			slice1: []int{1,2,3},
			slice2: []int{1,2},
		},
		want: false,
	},
	{
		name: "EmptySlices",
		args: args{
			slice1: []int{},
			slice2: []int{},
		},
		want: true,
	},
}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareSlices(tt.args.slice1, tt.args.slice2); got != tt.want {
				t.Errorf("compareSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}