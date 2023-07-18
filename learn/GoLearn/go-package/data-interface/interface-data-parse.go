package main

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

// ShowConv 演示了一些类型转换
func ShowConv() {
	// int
	var a = 24

	// float 64
	var b = 2.1

	// 将int转换为float64以进行计算
	c := float64(a) * b
	fmt.Println(c)

	// fmt.Sprintf是生成字符串的好方式
	precision := fmt.Sprintf("%.2f", b)

	// 输出值和对应类型
	fmt.Printf("%s - %T\n", precision, precision)
}

// Strconv 展示了字符串转换为基本类型
func Strconv() error {
	s := "1234"
	//指定 10 进制 精度 64 位
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	// 让我们试一下二进制
	res, err = strconv.ParseInt("FF", 16, 64)
	if err != nil {
		return err
	}

	fmt.Println(res)

	// 转换字符串为布尔
	val, err := strconv.ParseBool("true")
	if err != nil {
		return err
	}

	fmt.Println(val)

	return nil
}

// CheckType 演示了类型断言
func CheckType(s interface{}) {
	switch s.(type) {
	case string:
		fmt.Println("It's a string!")
	case int:
		fmt.Println("It's an int!")
	default:
		fmt.Println("not sure what it is...")
	}
}

// Interfaces 演示了如何获得断言操作结果
func Interfaces() {
	CheckType("test")
	CheckType(1)
	CheckType(false)

	var i interface{}
	i = "test"

	// manually check an interface
	if val, ok := i.(string); ok {
		fmt.Println("val is", val)
	}

	// this one should fail
	if _, ok := i.(int); !ok {
		fmt.Println("uh oh! glad we handled this")
	}
}

func test1() {
	ShowConv()
	if err := Strconv(); err != nil {
		panic(err)
	}
	Interfaces()
}

// Examples 演示了math包的基本应用
func Examples() {
	//开平方示例
	i := 25

	// i 是整型，所以需要转型
	result := math.Sqrt(float64(i))

	// 25开方结果是 5
	fmt.Println(result)

	// ceil能够获取大于或等于输入值的最小整数值
	result = math.Ceil(9.5)
	fmt.Println(result)

	// floor能够获取大于或等于输入值的最大整数值
	result = math.Floor(9.5)
	fmt.Println(result)

	// math包同样提供了常用的常数
	fmt.Println("Pi:", math.Pi, "E:", math.E)
}

// 全局变量
var memoize map[int]*big.Int

func init() {
	// 初始化map
	memoize = make(map[int]*big.Int)
}

// Fib打印斐波纳契序列的第n个数字，它将返回1以表示任何<0 ...它是递归计算并使用big.Int因为int64会快速溢出
func Fib(n int) *big.Int {
	if n < 0 {
		return nil
	}

	// 基础条件
	if n < 2 {
		memoize[n] = big.NewInt(1)
	}

	// 检查我们是否存储它之前进行了计算
	if val, ok := memoize[n]; ok {
		return val
	}

	// 使用map存储然后添加前2个fib值
	memoize[n] = big.NewInt(0)
	memoize[n].Add(memoize[n], Fib(n-1))
	memoize[n].Add(memoize[n], Fib(n-2))

	return memoize[n]
}

func test2() {
	Examples()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", Fib(i))
	}
	fmt.Println()
}

// ConvertStringDollarsToPennies 接收美元字符串并转换为int64
func ConvertStringDollarsToPennies(amount string) (int64, error) {
	// 检查传入参数是否合法
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	// 以"."进行分割
	groups := strings.Split(amount, ".")

	// 如果没有"."则取切片中的第一个元素
	result := groups[0]

	r := ""

	// 处理"."后的数据
	if len(groups) == 2 {
		if len(groups[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = groups[1]
		if len(r) > 2 {
			r = r[:2]
		}
	}

	// 填充0
	for len(r) < 2 {
		r += "0"
	}

	result += r

	// 转换为int
	return strconv.ParseInt(result, 10, 64)
}

// ConvertPenniesToDollarString 与上面的例子类似 这是将操作方式逆转
func ConvertPenniesToDollarString(amount int64) string {

	result := strconv.FormatInt(amount, 10)

	negative := false
	if result[0] == '-' {
		result = result[1:]
		negative = true
	}

	for len(result) < 3 {
		result = "0" + result
	}
	length := len(result)

	result = result[0:length-2] + "." + result[length-2:]

	if negative {
		result = "-" + result
	}

	return result
}

func test3() {
	userInput := "15.93"

	pennies, err := ConvertStringDollarsToPennies(userInput)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User input converted to %d pennies\n", pennies)

	pennies += 15

	dollars := ConvertPenniesToDollarString(pennies)

	fmt.Printf("Added 15 cents, new values is %s dollars\n", dollars)
}

const (
	jsonBlob     = `{"name": "Aaron"}`
	fulljsonBlob = `{"name":"Aaron", "age":0}`
)

// Example结构体包含age和name字段
type Example struct {
	Age  int    `json:"age,omitempty"`
	Name string `json:"name"`
}

// BaseEncoding 演示了基本的编码和解码操作
func BaseEncoding() error {
	e := Example{}

	// 注意jsonBlob没有age字段
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	//由于age被设置为omitempty(为空则不输出) 所以显示 Regular Marshal, with no age: {"name":"Aaron"}
	fmt.Println("Regular Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	fmt.Printf("Regular Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	//Regular Marshal, with age = 0: {"name":"Aaron"}
	fmt.Println("Regular Marshal, with age = 0:", string(value))

	return nil
}

// 和上一个例子类似 但是*int类型会出现奇妙的nil
// uses a *Int
type ExamplePointer struct {
	Age  *int   `json:"age,omitempty"`
	Name string `json:"name"`
}

func PointerEncoding() error {

	e := ExamplePointer{}
	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	//Pointer Unmarshal, no age: {Age:<nil> Name:Aaron}
	fmt.Printf("Pointer Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	//Pointer Marshal, with no age: {"name":"Aaron"}
	fmt.Println("Pointer Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	//Pointer Unmarshal, with age = 0: {Age:0xc04200e4b8 Name:Aaron}
	fmt.Printf("Pointer Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	//Pointer Marshal, with age = 0: {"age":0,"name":"Aaron"}
	fmt.Println("Pointer Marshal, with age = 0:", string(value))

	return nil
}

type nullInt64 sql.NullInt64

// 和前面的例子类似 又改变了Age的类型sql.NullInt64
type ExampleNullInt struct {
	Age  *nullInt64 `json:"age,omitempty"`
	Name string     `json:"name"`
}

func (v *nullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

func (v *nullInt64) UnmarshalJSON(b []byte) error {
	v.Valid = false
	if b != nil {
		v.Valid = true
		return json.Unmarshal(b, &v.Int64)
	}
	return nil
}

func NullEncoding() error {
	e := ExampleNullInt{}

	if err := json.Unmarshal([]byte(jsonBlob), &e); err != nil {
		return err
	}
	//nullInt64 Unmarshal, no age: {Age:<nil> Name:Aaron}
	fmt.Printf("nullInt64 Unmarshal, no age: %+v\n", e)

	value, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	//nullInt64 Marshal, with no age: {"name":"Aaron"}
	fmt.Println("nullInt64 Marshal, with no age:", string(value))

	if err := json.Unmarshal([]byte(fulljsonBlob), &e); err != nil {
		return err
	}
	//nullInt64 Unmarshal, with age = 0: {Age:0xc0420623f0 Name:Aaron}
	fmt.Printf("nullInt64 Unmarshal, with age = 0: %+v\n", e)

	value, err = json.Marshal(&e)
	if err != nil {
		return err
	}
	//nullInt64 Marshal, with age = 0: {"age":0,"name":"Aaron"}
	fmt.Println("nullInt64 Marshal, with age = 0:", string(value))

	return nil
}

func test4() {
	if err := BaseEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := PointerEncoding(); err != nil {
		panic(err)
	}
	fmt.Println()

	if err := NullEncoding(); err != nil {
		panic(err)
	}
}

type pos struct {
	X      int
	Y      int
	Object string
}

// GobExample展示了如何使用gob包
func GobExample() error {
	buffer := bytes.Buffer{}

	p := pos{
		X:      10,
		Y:      15,
		Object: "wrench",
	}

	// 注意如果p是个接口我们需要先调用gob.Register
	e := gob.NewEncoder(&buffer)
	if err := e.Encode(&p); err != nil {
		return err
	}

	// 这里是二进制的 所以打印出来的长度可能并不准确
	fmt.Println("Gob Encoded valued length: ", len(buffer.Bytes()))

	p2 := pos{}
	d := gob.NewDecoder(&buffer)
	if err := d.Decode(&p2); err != nil {
		return err
	}

	fmt.Println("Gob Decode value: ", p2)

	return nil
}

// Base64Example 演示了如何使用 base64 包
func Base64Example() error {
	// base64对于不支持以字节/字符串操作的二进制格式的情况很有用

	// 使用辅助函数和URL编码
	value := base64.URLEncoding.EncodeToString([]byte("encoding some data!"))
	fmt.Println("With EncodeToString and URLEncoding: ", value)

	// 解码
	decoded, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return err
	}
	fmt.Println("With DecodeToString and URLEncoding: ", string(decoded))

	return nil
}

// Base64ExampleEncoder 与上面的函数类似，但并没有使用URL编码
func Base64ExampleEncoder() error {

	buffer := bytes.Buffer{}

	// 建立编码器
	encoder := base64.NewEncoder(base64.StdEncoding, &buffer)

	// 确认关闭
	if err := encoder.Close(); err != nil {
		return err
	}
	if _, err := encoder.Write([]byte("encoding some other data")); err != nil {
		return err
	}

	fmt.Println("Using encoder and StdEncoding: ", buffer.String())

	decoder := base64.NewDecoder(base64.StdEncoding, &buffer)
	results, err := ioutil.ReadAll(decoder)
	if err != nil {
		return err
	}

	fmt.Println("Using decoder and StdEncoding: ", string(results))

	return nil
}

func test5() {
	if err := Base64Example(); err != nil {
		panic(err)
	}

	if err := Base64ExampleEncoder(); err != nil {
		panic(err)
	}

	if err := GobExample(); err != nil {
		panic(err)
	}
}

func SerializeStructStrings(s interface{}) (string, error) {
	result := ""
	// reflect.TypeOf 使用传入的接口生成 type 类型
	r := reflect.TypeOf(s)
	// reflect.ValueOf 返回结构体每个字段对应的值
	value := reflect.ValueOf(s)
	// 如果传入的是个结构体的指针 那么可以针对性的对其进行单独处理
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}
	// 循环所有的内部字段

	for i:=0;i<r.NumField();i++{
		field := r.Field(i)
		// 字段的名称
		key := field.Name

		// Lookup 返回与标记字符串中的 key 关联的值。如果密钥存在于标记中，则返回值（可以为空）。
		// 否则返回的值将是空字符串。ok 返回值报告是否在标记字符串中显式设置了值。
		// 如果标记没有传统格式，则 Lookup 返回的值不做指定。

		if serialize,ok := field.Tag.Lookup("serialize");ok{
			// 忽略“ - ”否则整个值成为序列化'键'
			if serialize == "-" {
				continue
			}
			key = serialize
		}
		// 判断每个字段的类型
		switch value.Field(i).Kind() {
		// 当前示例我们仅简单判断字符串
		case reflect.String:
			result += key + ":" + value.Field(i).String() + ";"
		default:
			continue
		}
	}
	return result, nil
}

// DeSerializeStructStrings 反序列化字符串为对应的结构体
func DeSerializeStructStrings(s string, res interface{}) error {
	r := reflect.TypeOf(res)

	// 我们要求传入的必须是指针
	if r.Kind() != reflect.Ptr {
		return errors.New("res must be a pointer")
	}

	// 解指针
	// Elem返回r(Type类型)元素的type
	// 如果该type.Kind不是Array, Chan, Map, Ptr, 或 Slice会产生panic
	r = r.Elem()
	value := reflect.ValueOf(res).Elem()

	// 将传入的序列化字符串分割为map
	vals := strings.Split(s, ";")
	valMap := make(map[string]string)
	for _, v := range vals {
		keyval := strings.Split(v, ":")
		if len(keyval) != 2 {
			continue
		}
		valMap[keyval[0]] = keyval[1]
	}

	// 循环所有的内部字段
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)

		// 检查是否符合预置的 tag
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// 忽略'-'否则整个值成为序列化'键'
			if serialize == "-" {
				continue
			}
			// 判断是否处于map内
			if val, ok := valMap[serialize]; ok {
				value.Field(i).SetString(val)
			}
		} else if val, ok := valMap[field.Name]; ok {
			// 是否是在map中的字段名称
			value.Field(i).SetString(val)
		}
	}
	return nil
}

// 注意Person内个字段的tag标签
type Person struct {
	Name  string `serialize:"name"`
	City  string `serialize:"city"`
	State string
	Misc  string `serialize:"-"`
	Year  int    `serialize:"year"`
}

// EmptyStruct 演示了根据 tag 序列化和反序列化一个空结构体
func EmptyStruct() error {
	p := Person{}

	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}
	fmt.Printf("Empty struct: %#v\n", p)
	fmt.Println("Serialize Results:", res)

	newP := Person{}
	if err := DeSerializeStructStrings(res, &newP); err != nil {
		return err
	}
	fmt.Printf("Deserialize results: %#v\n", newP)
	return nil
}

// FullStruct 演示了根据 tag 序列化和反序列化一个非空结构体
func FullStruct() error {
	p := Person{
		Name:  "Aaron",
		City:  "Seattle",
		State: "WA",
		Misc:  "some fact",
		Year:  2017,
	}
	res, err := SerializeStructStrings(&p)
	if err != nil {
		return err
	}
	fmt.Printf("Full struct: %#v\n", p)
	fmt.Println("Serialize Results:", res)

	newP := Person{}
	if err := DeSerializeStructStrings(res, &newP); err != nil {
		return err
	}
	fmt.Printf("Deserialize results: %#v\n", newP)
	return nil
}

func test6() {
	if err := EmptyStruct(); err != nil {
		panic(err)
	}

	fmt.Println()

	if err := FullStruct(); err != nil {
		panic(err)
	}
}

// WorkWith 会实现集合接口
type WorkWith struct {
	Data    string
	Version int
}

// Filter 是一个过滤函数。
func Filter(ws []WorkWith, f func(w WorkWith) bool) []WorkWith {
	// 初始化返回值
	result := make([]WorkWith, 0)
	for _, w := range ws {
		if f(w) {
			result = append(result, w)
		}
	}
	return result
}

// Map是一个映射函数。
func Map(ws []WorkWith, f func(w WorkWith) WorkWith) []WorkWith {
	// 返回值的长度应该与传入切片长度一致
	result := make([]WorkWith, len(ws))

	for pos, w := range ws {
		newW := f(w)
		result[pos] = newW
	}
	return result
}

// LowerCaseData 将传入WorkWith的data字段变为小写
func LowerCaseData(w WorkWith) WorkWith {
	w.Data = strings.ToLower(w.Data)
	return w
}

// IncrementVersion 将传入WorkWith的Version加1
func IncrementVersion(w WorkWith) WorkWith {
	w.Version++
	return w
}

// OldVersion 返回一个闭包，用于验证版本是否大于指定的值
func OldVersion(v int) func(w WorkWith) bool {
	return func(w WorkWith) bool {
		return w.Version >= v
	}
}

func test7()  {
	ws := []WorkWith{
		{"Example", 1},
		{"Example 2", 2},
	}

	fmt.Printf("Initial list: %#v\n", ws)

	ws = Map(ws, LowerCaseData)
	fmt.Printf("After LowerCaseData Map: %#v\n", ws)

	ws = Map(ws, IncrementVersion)
	fmt.Printf("After IncrementVersion Map: %#v\n", ws)

	ws = Filter(ws, OldVersion(3))
	fmt.Printf("After OldVersion Filter: %#v\n", ws)
}

func main() {
	test7()
}
