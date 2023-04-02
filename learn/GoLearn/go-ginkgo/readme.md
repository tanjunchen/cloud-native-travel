# 镜像拉取工具

https://github.com/silenceshell/docker_wrapper

例如
docker_wrapper pull k8s.gcr.io/kube-apiserver:v1.14.1

# e2e 测试

例如使用工具拉取镜像
docker_wrapper.py  pull k8s.gcr.io/kube-cross

## 下载 kubetest 

git clone https://github.com/kubernetes/test-infra.git

kubetest --build
如果缺少相应的镜像包，使用工具下载

## 运行 e2e 测试

参考 https://blog.gmem.cc/kubernetes-e2e-test

例如
kubetest --test --test_args="--ginkgo.focus=\[sig-api-machinery\]  --host=https://192.168.0.32:6443" --provider=local

kubetest --test --test_args="--ginkgo.focus=SSH.to.all     --host=https://192.168.0.32:6443" --provider=local

构建 Kubernetes、启动一个集群、运行测试、清理，这一系列阶段可以通过下面的命令完成

kubetest --build --up --test --down

具体命令
```
# 构建
kubetest --build
 
# 启动空白集群，如果存在先删除之
kubetest --up
 
# 运行所有测试
kubetest --test
 
# 运行匹配的测试
kubetest --test --test_args="--ginkgo.focus=\[Feature:Performance\]" --provider=local
 
# 跳过指定的测试
kubetest --test --test_args="--ginkgo.skip=Pods.*env"
 
# 并行测试，跳过不支持并行的那些用例
GINKGO_PARALLEL=y kubetest --test --test_args="--ginkgo.skip=\[Serial\]"
 
# 指定云提供商
kubetest --provider=aws --build --up --test --down

# 针对临时集群调用 kubectl
kubetest -ctl='get events'
kubetest -ctl='delete pod foobar'

# 清理
kubetest --down
```

测试特定版本
利用kubetest你可以下载任意版本的K8S，包括服务器组件、客户端、测试二进制文件。 

kubetest --extract=版本  --up

使用本地集群
export KUBECONFIG=/path/to/kubeconfig
kubetest --provider=local --test --host="" --kubeconfig=""  

如果使用本地集群进行反复测试，你可能需要周期性的进行某些手工清理：

执行 rm -rf /var/run/kubernetes删除K8S生成的凭证文件，某些情况下上次测试遗留的凭证文件会导致问题
执行 sudo iptables -F清空kube-proxy生成的Iptables规则

指定K8S代码库
使用选项 --repo-root="../../"，可以指定K8S代码库的根目录，E2E测试文件将从中寻找。  


# Ginkgo

顶级的 Describe 容器
Describe块用于组织Specs，其中可以包含任意数量的：
   BeforeEach：在Spec（It块）运行之前执行，嵌套Describe时最外层BeforeEach先执行
   AfterEach：在Spec运行之后执行，嵌套Describe时最内层AfterEach先执行
   JustBeforeEach：在It块，所有BeforeEach之后执行
   Measurement

可以在 Describe 块内嵌套 Describe、Context、When 块
Ginkgo广泛使用闭包（ 闭包不是私有，闭的意思不是“封闭内部状态”，而是“封闭外部状态”！），允许您构建描述性测试套件。
您应该使用 Describe 和 Context 容器来表达性地组织代码的行为
您可以使用 BeforeEach 为您的 Specs 设置状态。您可以使用 It 来指定单个 Spec
为了在 BeforeEach 和 It 之间共享状态，您使用闭包变量，通常在最相关的 Describe 或 Context 容器的顶部定义
我们使用 Gomega 的 Expect 语法来对 CategoryByLength() 方法产生期望值。

logging 日志的使用
Ginkgo提供了一个全局可用的 io.Writer，名为 GinkgoWriter，供您写入
GinkgoWriter 在测试运行时聚合输入，并且只有在测试失败时才将其转储到stdout。
当以详细模式运行时（ginkgo -v或go test -ginkgo.v），GinkgoWriter会立即将其输入重定向到stdout。
当 Ginkgo测试套件中断（通过^C）时，Ginkgo将发出写入GinkgoWriter的任何内容。这样可以更轻松地调试卡住的测试。
当与--progress配对使用时将会特别有用，它指示Ginkgo在运行您的BeforeEaches，Its，AfterEaches等时向GinkgoWriter发出通知

单个规格：It

提取通用步骤 BeforeEach

使用容器组织规格 Describe 和 Context

分离创建和配置 JustBeforeEach JustBeforeEach 是一个很容易被滥用的强大工具。好好利用它。

分离诊断收集和销毁 JustAfterEach JustAfterEach 是一个很容易被滥用的强大工具。好好利用它。

全局设置和销毁 BeforeSuite 和 AfterSuite
有时您希望在整个测试之前运行一些设置代码和在整个测试之后运行一些清理代码。例如，您可能需要启动并销毁外部数据库。
BeforeSuite 函数在任何规格运行之前运行  AfterSuite 函数在所有的规格运行之后运行

文档化复杂的It：By
传递给By的字符串是通过GinkgoWriter发出的。如果测试成功，您将看不到Ginkgo绿点之外的任何输出。
但是，如果测试失败，您将看到失败之前的每个步骤的打印输出。使用ginkgo -v总是输出所有步骤打印。

Pending 态规格
您可以将单个Spec或容器标记为待定。这将阻止Spec（或者容器中的Specs）运行。
您可以在您的Describe, Context, It 和 Measure前面添加一个P或者一个X来实现这一点：PDescribe("some behavior", func() { ... })

重点规格
您可以在Descirbe, Context 和 It前面添加F以编程方式专注于单个规格或者整个容器的规格
您可以使用 --focus = REGEXP 和/或 --skip = REGEXP 标签来传递正则表达式。

规格排列

Ginkgo 命令行
安装 go install github.com/onsi/ginkgo/ginkgo
运行测试  指定运行哪些测试套件  ginkgo  -r  -skipPackage=PACKAGES,TO,SKIP

基准测试
Ginkgo 允许你使用Measure块来测量你的代码的性能。
Measure块可以运行在任何It块可以运行的地方--每一个Meature生成一个规格。
传递给Measure的闭包函数必须使用Benchmarker参数。Benchmarker用于测量运行时间并记录任意数值。
你也必须在该闭包函数之后传递一个整型参数给Measure，它表示Measure将执行的你的代码的样本数。

测量时间
Time(name string, body func (), info ...Interface {}) time.Duration

记录任意值
RecordValue(name string, value float64, info ...Interface{})

共享示例模式
Ginkgo对共享示例（也称为共享行为）没有任何明确的支持，但是您可以使用一些模式来复用套件中的测试。

Ginkgo 与 CI
Ginkgo附带了许多标签，您可能希望在持续集成环境运行时打开这些标签。
ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --cover --trace --race --progress
ginkgo -r
1.递归执行文件夹内的所有测试用例，用于有多层文件夹的情形
2.使用–skipPackages=PACKAGES,TO,SKIP 跳过不需要执行的文件夹，文件夹之间使用逗号间隔

ginkgo -randomizeAllSpecs
随机顺序执行测试集中的所有测试例
Tips：ginkgo执行测试文件时默认会将最顶层的测试集即var _ = Describe()进行随机顺序排列，而其内部的测试例则按顺序一个个执行

ginkgo –randomizeSuites
随机顺序执行测试集
//Tips：ginkgo 执行测试集时，是按照文件夹在文件系统中的存储顺序一次执行

默认值是true，设置是否打印Pending的测试例的信息
//ginkgo –failOnPending

ginkgo –trace
当测试例失败时，会打印详细的错误追踪信息，便于定位

ginkgo -race
显示执行速度

ginkgo -cover
执行测试集后生成覆盖率文件

扩展
Ginkgo附带了核心DSL的扩展。这些可以（可选）点导入以增强Ginkgo的默认DSL。 目前只有一个扩展名：table扩展名。

具体信息可以参考  https: //www.ginkgo.wiki/
参见官网 https://godoc.org/github.com/onsi/ginkgo


Ginkgo 是 Go 语言的一个行为驱动开发（BDD， Behavior-Driven Development）风格的测试框架，
通常和库 Gomega 一起使用。Ginkgo 在一系列的 Specs 中描述期望的程序行为。

Ginkgo 集成了 Go 语言的测试机制，你可以通过 go test 来运行 Ginkgo 测试套件。

Ginkgo 测试结构

It

BeforeEach

AfterEach

Describe/Context

JustBeforeEach

JustAfterEach

BeforeSuite/AfterSuite
支持异步执行，只需要给函数传递一个 Done 参数即可

By
此块用于给逻辑复杂的块添加文档  传递给By的字符串会发送给GinkgoWriter

性能测试
使用Measure块可以进行性能测试，所有It能够出现的地方，都可以使用Measure。
传递给Measure的闭包函数必须具有Benchmarker入参

CLI
运行测试
```
# 运行当前目录中的测试
ginkgo
# 运行其它目录中的测试
ginkgo /path/to/package /path/to/other/package ...
# 递归运行所有子目录中的测试
ginkgo -r ...
```

传递参数
传递参数给测试套件
ginkgo -- PASS-THROUGHS-ARGS
跳过某些包
ginkgo -skipPackage=PACKAGES,TO,SKIP

超时控制
选项 -timeout 用于控制套件的最大运行时间，如果超过此时间仍然没有完成，认为测试失败。默认24小时。

调试信息
选项	说明
--reportPassed	打印通过的测试的详细信息
--v	冗长模式
--trace	打印所有错误的调用栈
--progress	打印进度信息

Spec Runner
Pending Spec
Skiping Spec
Focused Specs
Parallel Specs
如果所有Spec需要共享一个外部进程，则可以利用SynchronizedBeforeSuite、SynchronizedAfterSuite：

# Gomega

这时Ginkgo推荐使用的断言 Matcher 库。

联用 Ginkgo
注册 Fail 处理器即可 
gomega.RegisterFailHandler(ginkgo.Fail)

断言
Ω/Expect

错误处理
对于返回多个值的函数

```
func DoSomethingHard() (string, error) {}
result, err := DoSomethingHard()
断言没有发生错误
Ω(err).ShouldNot(HaveOccurred())
Ω(result).Should(Equal("foo"))

对于仅仅返回一个error的函数
func DoSomethingHard() (string, error) {}
Ω(DoSomethingSimple()).Should(Succeed())
```

简化输出
断言失败时，Gomega 打印牵涉到断言的对象的递归信息，输出可能很冗长。
format 包提供了一些全局变量，调整这些变量可以简化输出。

异步断言
Gomega 提供了两个函数，用于异步断言。
传递给 Eventually、Consistently 的函数，如果返回多个值，
则第一个返回值用于匹配，其它值断言为 nil 或零值。

```
参数是闭包，调用函数
Eventually(func() []int {
    return thing.SliceImMonitoring
}).Should(HaveLen(2))
 
参数是通道，读取通道
Eventually(channel).Should(BeClosed())
Eventually(channel).Should(Receive())
 
参数也可以是普通变量，读取变量
Eventually(myInstance.FetchNameFromNetwork).Should(Equal("archibald"))
 
可以和gexec包的Session配合
Eventually(session).Should(gexec.Exit(0)) 命令最终应当以0退出
Eventually(session.Out).Should(Say("Splines reticulated")) 检查标准输出
```

可以指定超时、轮询间隔
Eventually(func() []int {
    return thing.SliceImMonitoring
}, TIMEOUT, POLLING_INTERVAL).Should(HaveLen(2))

Consistently
检查断言是否在一定时间段内总是通过：

Consistently(func() []int {
    return thing.MemoryUsage()
}, DURATION, POLLING_INTERVAL).Should(BeNumerically("<", 10))
Consistently也可以用来断言最终不会发生的事件，例如下面的例子：

Consistently(channel).ShouldNot(Receive())

内置Matcher
相等性

接口相容

空值/零值

布尔值

错误

通道

文件

字符串

JSON/XML/YML

集合

string、array、map、chan、slice 都属于集合

```
断言为空
Ω(ACTUAL).Should(BeEmpty())
 
断言长度
Ω(ACTUAL).Should(HaveLen(INT))
 
断言容量
Ω(ACTUAL).Should(HaveCap(INT))
 
断言包含元素
Ω(ACTUAL).Should(ContainElement(ELEMENT))
 
断言等于                   其中之一
Ω(ACTUAL).Should(BeElementOf(ELEMENT1, ELEMENT2, ELEMENT3, ...))
 
 
断言元素相同，不考虑顺序
Ω(ACTUAL).Should(ConsistOf(ELEMENT1, ELEMENT2, ELEMENT3, ...))
Ω(ACTUAL).Should(ConsistOf([]SOME_TYPE{ELEMENT1, ELEMENT2, ELEMENT3, ...}))
 
断言存在指定的键，仅用于map
Ω(ACTUAL).Should(HaveKey(KEY))
断言存在指定的键值对，仅用于map
Ω(ACTUAL).Should(HaveKeyWithValue(KEY, VALUE))
```

数字/时间
BeNumerically BeTemporally

Panic
断言

And/Or
自定义Matcher

辅助工具

ghttp
gbytes
gexec
gstruct






