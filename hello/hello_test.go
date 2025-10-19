package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Gopher!"
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestHello_quote(t *testing.T) {
	want := "Hello, world."
	if got := Hello_quote(); got != want {
		t.Errorf("Hello_quote() = %q, want %q", got, want)
	}
}

// 这是 Go 语言的单元测试代码，用来测试一个叫 Hello() 的函数是否正确工作。
// 测试代码通常和被测试的代码放在同一个包下，比如：
// hello/
// ├── hello.go        // 实现函数 Hello()
// └── hello_test.go   // 测试文件
// testing 是 Go 标准库自带的测试框架包。
// 只要你写的函数名以 Test 开头，并且参数是 *testing.T，就会被 Go 的测试工具自动识别并运行。
// 这就是一个测试函数。
// Go 规定：
// 测试函数必须以 Test 开头；
// 参数是 *testing.T 类型，用来报告测试的失败或日志。
// 当你在终端运行：
// go test时，Go 会自动运行所有这样的函数。
