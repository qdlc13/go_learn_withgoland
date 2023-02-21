package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//norlmal cases
		{"abcabcbb", 3},
		{"pppppp", 1},
		{"pwwkew", 3},
		//edge cases
		{"", 0},
		{"abcabcabcd", 4},
		//chinese support
		{"我是初学者", 5},
	}
	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}

	}
}

func BenchmarkSubstr(b *testing.B) {
	s, ans := "我是初学者", 5
	for i := 0; i <= 6; i++ {
		s = s + s
	}
	b.Logf("len(s)=%d", len(s))
	//重新设置测试时间 以上时间不算
	b.ResetTimer()

	//b.N系统内部有算法确定数量
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
