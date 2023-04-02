package main

import (
	"bytes"
	"context"
	"fmt"
	apexlog "github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strconv"
)

// ErrorValue创建了包级错误
// 可以采用这样的方式判断： if err == ErrorValue
var ErrorValue = errors.New("this is a typed error")

// TypedError创建了包含错误类型的匿名字段
// 可以采用断言的方式判断：err.(type) == ErrorValue
type TypedError struct {
	error
}

//BasicErrors 演示了错误的创建
func BasicErrors() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorValue
	fmt.Println("value error: ", err)

	err = TypedError{errors.New("typed error")}
	fmt.Println("typed error: ", err)
}

// CustomError 实现了Error接口
type CustomError struct {
	Result string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("there was an error; %s was the result", c.Result)
}

// errors.New fmt.Errorf 还是自定义错误，最重要的是你应该永远不要在代码中忽略错误
// SomeFunc 返回一个 error
func SomeFunc() error {
	c := CustomError{Result: "this"}
	return c
}

func test1() {
	BasicErrors()

	err := SomeFunc()
	fmt.Println("custom error: ", err)
}

// WrappedError 演示了如何对错误进行封装
func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

type ErrorTyped struct {
	error
}

func Wrap() {
	e := errors.New("standard error")

	fmt.Println("Regular Error - ", WrappedError(e))

	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))

	fmt.Println("Nil -", WrappedError(nil))
}

// Unwrap 解除封装并进行断言处理
func Unwrap() {

	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Println("wrapped error: ", err)

	// 处理错误类型
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occurred")
	}
}

// StackTrace 打印错误栈
func StackTrace() {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")

	fmt.Printf("%+v\n", err)
}

func test2()  {
	Wrap()
	fmt.Println()
	Unwrap()
	fmt.Println()
	StackTrace()
}

// 对日志进行设置
func Log() {
	// logger会写入bytes.Buffer类型的数据
	buf := bytes.Buffer{}

	// 第二个参数是前缀最后一个参数是关于选项
	// 配置选项可以用逻辑或符号组合起来
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ldate)

	logger.Println("test")

	logger.SetPrefix("new logger: ")

	logger.Printf("you can also add args(%v) and use Fataln to log and crash", true)

	fmt.Println(buf.String())
}

// OriginalError返回错误的原始信息
func OriginalError() error {
	return errors.New("error occurred")
}

// PassThroughError 调用OriginalError并将其封装
func PassThroughError() error {
	err := OriginalError()
	// 无需检查错误，因为使用该库时此操作适用于nil
	return errors.Wrap(err, "in passthrougherror")
}

// FinalDestination处理错误而不传递它
func FinalDestination() {
	err := PassThroughError()
	if err != nil {
		// 将任何产生的意外记录到日志中
		log.Printf("an error occurred: %s\n", err.Error())
		return
	}
}

func test3()  {
	fmt.Println("basic logging and modification of logger:")
	Log()
	fmt.Println("logging 'handled' errors:")
	FinalDestination()
}

// go get github.com/sirupsen/logrus
// go get github.com/apex/log

// Hook 实现了 logrus 中的 hook 接口
type Hook struct {
	id string
}

// Fire 在每次记录日志时都会触发
func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	return nil
}

// Levels 日志等级
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Logrus 演示了一些基本的 logrus 库操作
func Logrus() {

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{"123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{"Something happened", "Just now"}

	x := logrus.WithFields(fields)
	x.Warn("warning!")
	x.Error("error!")
}

// ThrowError抛出我们将追踪的错误
func ThrowError() error {
	err := errors.New("a crazy failure")
	apexlog.WithField("id", "123").Trace("ThrowError").Stop(&err)
	return err
}

type CustomHandler struct {
	id      string
	handler apexlog.Handler
}

// HandleLog 会对日志进行处理
func (h *CustomHandler) HandleLog(e *apexlog.Entry) error {
	e.WithField("id", h.id)
	return h.handler.HandleLog(e)
}

func Apex() {
	apexlog.SetHandler(&CustomHandler{"123", text.New(os.Stdout)})
	err := ThrowError()
	// WithError可以便利的记录错误
	apexlog.WithError(err).Error("an error occurred")
}

func test4()  {
	fmt.Println("Logrus:")
	Logrus()

	fmt.Println("=========================")

	fmt.Println("Apex:")
	Apex()
}

type key int

const logFields key = 0

func getFields(ctx context.Context) *apexlog.Fields {
	fields, ok := ctx.Value(logFields).(*apexlog.Fields)
	if !ok {
		f := make(apexlog.Fields)
		fields = &f
	}
	return fields
}

// FromContext 接收一个 log.Interface 和 context
// 然后返回由 context 对象填充的 log.Entry 指针
func FromContext(ctx context.Context, l apexlog.Interface) (context.Context, *apexlog.Entry) {
	fields := getFields(ctx)
	e := l.WithFields(fields)
	ctx = context.WithValue(ctx, logFields, fields)
	return ctx, e
}

// WithField 将log.Fielder添加到context
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return WithFields(ctx, apexlog.Fields{key: value})
}

// WithFields 将log.Fielder添加到context
func WithFields(ctx context.Context, fields apexlog.Fielder) context.Context {
	f := getFields(ctx)
	for key, val := range fields.Fields() {
		(*f)[key] = val
	}
	ctx = context.WithValue(ctx, logFields, f)
	return ctx
}

// Initialize 调用 3 个函数来设置，然后在操作完成之前记录日志
func Initialize() {
	apexlog.SetHandler(text.New(os.Stdout))
	//初始化 context
	ctx := context.Background()
	// 将 context 与 log.Log 建立联系
	ctx, e := FromContext(ctx, apexlog.Log)

	ctx = WithField(ctx, "id", "123")
	e.Info("starting")
	gatherName(ctx)
	e.Info("after gatherName")
	gatherLocation(ctx)
	e.Info("after gatherLocation")
}

func gatherName(ctx context.Context) {
	ctx = WithField(ctx, "name", "Go Cookbook")
}

func gatherLocation(ctx context.Context) {
	ctx = WithFields(ctx, apexlog.Fields{"city": "Seattle", "state": "WA"})
}

func test5()  {
	Initialize()
}


// Panic 演示除零恐慌
func Panic() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}

	a := 1 / zero
	fmt.Println("we'll never get here", a)
}

// Catcher 处理恐慌
func Catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred:", r)
		}
	}()
	Panic()
}

func test6()  {
	fmt.Println("before panic")
	Catcher()
	fmt.Println("after panic")
}

func main() {
	test6()
}
