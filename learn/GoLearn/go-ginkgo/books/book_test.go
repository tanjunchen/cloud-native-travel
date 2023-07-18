package books_test

import (
	"flag"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	var (
		// 通过闭包在BeforeEach和It之间共享数据
		longBook  Book
		shortBook Book
	)
	// 此函数用于初始化Spec的状态，在It块之前运行。如果存在嵌套Describe，则最
	// 外面的BeforeEach最先运行
	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("With more than 300 pages", func() {
			// 通过It来创建一个Spec
			It("should be a novel", func() {
				// Gomega的Expect用于断言
				Expect(longBook.CategoryByLength()).To(Equal("NOVEL"))
			})
		})

		Context("With fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(shortBook.CategoryByLength()).To(Equal("SHORT STORY"))
			})
		})

		/**
		Fail(因此Gomega，因为它使用Fail)将为当前的space和panic记录失败。这允许Ginkgo停止其轨道中的当前Spec
		 没有后续的断言（或者任何代码）将被调用。通常情况下，Ginkgo将会补救这个Panic本身然后进行下一步测试。
		然而，如果你的测试启用了goroutine调用Fail（或者，等效地，调用失败的Gomega断言）
		Ginkgo将没有办法补救由Fail引发的Panic.这将导致测试套件出现Panic，并且不会运行后续测试。
		要解决这个问题，你必须使用GinkgoRecover拯救Panic。
		*/
		Context("test panics in a goroutine", func() {
			It("panics in a goroutine", func(done Done) {
				go func() {
					// 如果 doSomething 返回 false 则下面的 defer 会确保从 panic 中恢复
					defer GinkgoRecover()
					// Ω 和 Expect 功能相同
					Ω(doSomething()).Should(BeTrue())

					// 在 Goroutine 中需要关闭 done 通道
					close(done)
				}()
			})
		})

		Context("test flag", func() {
			It("print init flag", func() {
				fmt.Println(myFlag)
			})
		})

	})
})

func doSomething() bool {
	// test return false
	// return false
	return true
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b Book) CategoryByLength() string {
	if b.Pages < 300 {
		return "SHORT STORY"
	}
	return "NOVEL"
}

// 传递参数
var myFlag string

func init() {
	flag.StringVar(&myFlag, "myFlag", "defaultvalue", "myFlag is used to control my behavior")
}
