package strev

import (
	"testing"
)

type ReverseTest struct {
	in, out string
}

var ReverseTests = []ReverseTest{
	{"ABCD", "DCBA"},
	{"CVO-AZ", "ZA-OVC"},
	{"Hello 世界", "界世 olleH"},
}

//func verify(t *testing.T, testnum int, testcase, input, output, expected string) {
//	if expected != output {
//		t.Errorf("%d. %s with input = %s: output %s != %s", testnum, testcase, input, output, expected)
//	}
//}
//
//func TestFunction(t *testing.T) {
//	for i, tt := range tests {
//		s := FuncToBeTested(tt.in)
//		verify(t, i, "FuncToBeTested: ", tt.in, s, tt.out)
//	}
//}

// 单元测试 go test
func TestReverse(t *testing.T) {
	/*
		in := "CVO-AZ"
		out := Reverse(in)
		exp := "ZA-OVC"
		if out != exp {
			t.Errorf("Reverse of %s expects %s, but got %s", in, exp, out)
		}
	*/

	// testing with a battery of testdata:
	for _, r := range ReverseTests {
		exp := Reverse(r.in)
		if r.out != exp {
			t.Errorf("Reverse of %s expects %s, but got %s", r.in, exp, r.out)
		}
	}
}

// 基准测试 go test -bench=.
func BenchmarkReverse(b *testing.B) {
	s := "ABCD"
	for i := 0; i < b.N; i++ { // 循环次数
		Reverse(s)
	}
}
