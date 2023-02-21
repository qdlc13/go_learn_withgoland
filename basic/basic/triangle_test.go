package main

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	//测试数据 表格驱动测试
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{3, 4, 8},
		{30000, 1, 50000},
	}
	//_是索引0-3 tt是 {3 4 5} {5, 12, 13}等
	for _, tt := range tests {
		//fmt.Println(a, tt)
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d);"+
				"got %d; expected %d", tt.a, tt.b, actual, tt.c)

		}
	}

}
