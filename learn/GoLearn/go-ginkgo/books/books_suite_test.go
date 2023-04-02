package books_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestBooks 是一个 testing 测试
// 当您运行 go test 或者 ginkgo 命令时，Go 测试运行器将运行此功能。
func TestBooks(t *testing.T) {
	// 将 Ginkgo 的 Fail 函数传递给 Gomega，Fail 函数用于标记测试失败
	// 这是 Ginkgo 和 Gomega 唯一的交互点
	// 如果 Gomega 断言失败，就会调用 Fail 进行处理
	RegisterFailHandler(Fail)
	// RunSpecs(t *testing.T, suiteDescription string)通知 Ginkgo 启动测试套件。
	// 如果您的任何 specs 失败，Ginkgo 将自动使 testing.T 失败。
	RunSpecs(t, "Books Suite")
}
