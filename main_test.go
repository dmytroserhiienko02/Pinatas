package main

import (
	"reflect"
	"testing"
)

func TestGetMaxCandies(t *testing.T) {
	pinatas1 := []int{1, 1, 4, 5, 1}
	pinatas2 := []int{1, 1, 1, 1, 1}
	pinatas3 := []int{5, 4, 3, 2, 1}
	pinatas4 := []int{1, 2, 3, 4, 5}
	pinatas5 := []int{2, 2, 5, 3, 3}
	pinatas6 := []int{5}
	pinatas7 := []int{}
	type args struct {
		pinatas []int
	}
	tests := []struct {
		name   string
		args   args
		expect int
	}{
		{
			name:   "Random array",
			args:   args{pinatas: pinatas1},
			expect: 40,
		},
		{
			name:   "The same elements",
			args:   args{pinatas: pinatas2},
			expect: 5,
		},
		{
			name:   "By decline (max is first)",
			args:   args{pinatas: pinatas3},
			expect: 110,
		},
		{
			name:   "By growth (max is last)",
			args:   args{pinatas: pinatas4},
			expect: 110,
		},
		{
			name:   "Several identical elements33",
			args:   args{pinatas: pinatas5},
			expect: 95,
		},
		{
			name:   "One element",
			args:   args{pinatas: pinatas6},
			expect: 5,
		},
		{
			name:   "Empty array",
			args:   args{pinatas: pinatas7},
			expect: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxCandies(tt.args.pinatas); !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("getMaxCandiesl() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
