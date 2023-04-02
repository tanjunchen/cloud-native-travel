package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ghodss/yaml"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

// Config存储接收到的标识
type Config struct {
	subject      string
	isAwesome    bool
	howAwesome   int
	countTheWays CountTheWays
}

// Setup 根据传入的标识初始化配置
func (c *Config) Setup() {
	// 你可以使用这样的方式直接初始化标识:
	// var someVar = flag.String("flag_name", "default_val", "description")
	// 但在实际操作中使用结构体来承载会更好一些

	// 完整版
	flag.StringVar(&c.subject, "subject", "", "subject is a string, it defaults to empty")
	// 简写版
	flag.StringVar(&c.subject, "s", "", "subject is a string, it defaults to empty (shorthand)")

	flag.BoolVar(&c.isAwesome, "isawesome", false, "is it awesome or what?")
	flag.IntVar(&c.howAwesome, "howawesome", 10, "how awesome out of 10?")

	// 自定义变量类型
	flag.Var(&c.countTheWays, "c", "comma separated list of integers")
}

// GetMessage 将所有的内部字段拼接成完整的句子
func (c *Config) GetMessage() string {
	msg := c.subject
	if c.isAwesome {
		msg += " is awesome"
	} else {
		msg += " is NOT awesome"
	}

	msg = fmt.Sprintf("%s with a certainty of %d out of 10. Let me count the ways %s", msg, c.howAwesome, c.countTheWays.String())
	return msg
}

// CountTheWays 是一个自定义变量类型
// 我们会从标识中读取到它
type CountTheWays []int

// 想要实现自定义类型的 flag 必须实现 flag.Value 接口
// 该接口包含了
// String() string
// Set(string) error
func (c *CountTheWays) String() string {
	result := ""
	for _, v := range *c {
		if len(result) > 0 {
			result += " ... "
		}
		result += fmt.Sprint(v)
	}
	return result
}

func (c *CountTheWays) Set(value string) error {
	values := strings.Split(value, ",")
	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*c = append(*c, i)
	}
	return nil
}

func test1() {
	// 初始化
	c := Config{}
	c.Setup()

	// 常见的调用方式
	flag.Parse()

	// 将通过命令行输入的flag标识拼接打印
	fmt.Println(c.GetMessage())
}

const version = "hello.0.0"
const usage = `Usage:

%s [command]

Commands:
    Greet
    Version
`

const greetUsage = `Usage:

%s greet name [flag]

Positional Arguments:
    name
        the name to greet

Flags:
`

// MenuConf 保存嵌套命令行的级别参数
type MenuConf struct {
	Goodbye bool
}

// SetupMenu 初始化flag标识
func (m *MenuConf) SetupMenu() *flag.FlagSet {
	menu := flag.NewFlagSet("menu", flag.ExitOnError)
	menu.Usage = func() {
		fmt.Printf(usage, os.Args[0])
		menu.PrintDefaults()
	}
	return menu
}

// GetSubMenu 返回子菜单的 flag 集
func (m *MenuConf) GetSubMenu() *flag.FlagSet {
	submenu := flag.NewFlagSet("submenu", flag.ExitOnError)
	submenu.BoolVar(&m.Goodbye, "goodbye", false, "Say goodbye instead of hello")

	submenu.Usage = func() {
		fmt.Printf(greetUsage, os.Args[0])
		submenu.PrintDefaults()
	}
	return submenu
}

// Greet 由 greet 命令调用
func (m *MenuConf) Greet(name string) {
	if m.Goodbye {
		fmt.Println("Goodbye " + name + "!")
	} else {
		fmt.Println("Hello " + name + "!")
	}
}

// Version 打印存储为const的当前版本值
func (m *MenuConf) Version() {
	fmt.Println("Version: " + version)
}

func test2() {
	c := MenuConf{}
	menu := c.SetupMenu()

	if err := menu.Parse(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing params %s, error: %v", os.Args[1:], err)
		return
	}

	// 在未输出参数的情况下
	// os.Args[0]是执行文件所在的路径
	// len(os.Args) > 1说明输入了命令行参数
	if len(os.Args) > 1 {

		// 根据分支条件打印输出
		switch strings.ToLower(os.Args[1]) {
		case "version":
			c.Version()
		case "greet":
			f := c.GetSubMenu()
			if len(os.Args) < 3 {
				f.Usage()
				return
			}
			if len(os.Args) > 3 {
				if err := f.Parse(os.Args[3:]); err != nil {
					fmt.Fprintf(os.Stderr, "Error parsing params %s, error: %v", os.Args[3:], err)
					return
				}

			}
			c.Greet(os.Args[2])

		default:
			fmt.Println("Invalid command")
			menu.Usage()
			return
		}
	} else {
		menu.Usage()
		return
	}

}

// https://github.com/kelseyhightower/envconfig  环境变量

// LoadConfig将从存储在路径中的json文件中选择加载文件，
// 然后根据envconfig struct标记覆盖这些值。
// envPrefix是我们为环境变量添加的前缀。
func LoadConfig(path, envPrefix string, configFill interface{}) error {
	if path != "" {
		err := LoadFile(path, configFill)
		if err != nil {
			return errors.Wrap(err, "error loading config from file")
		}
	}

	// envconfig.Process 根据环境变量填充指定的结构
	err := envconfig.Process(envPrefix, configFill)
	return errors.Wrap(err, "error loading config from env")
}

// LoadFile 解析一个 json 文件并填充到 config 中
func LoadFile(path string, configFill interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(configFill); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}
	return nil
}

// Config 将保存我们从json文件和环境变量中捕获的配置
type ConfigFill struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func test3() {
	var err error
	// 建立一个临时 json 配置文件
	tf, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	// json 配置文件的内容
	secrets := `{
        "secret": "so so secret"
    }`

	if _, err = tf.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	// 向环境变量中添加变量及对应值
	if err = os.Setenv("EXAMPLE_VERSION", "hello.0.0"); err != nil {
		panic(err)
	}
	if err = os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := ConfigFill{}
	// 从文件中读取配置参数
	if err = LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secrets file contains =", secrets)

	// 获取环境变量参数及对应值
	fmt.Println("EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE =", os.Getenv("EXAMPLE_ISSAFE"))

	// c既保存了json配置文件的参数也保存了环境变量的参数 我们将其打印
	fmt.Printf("Final Config: %#v\n", c)
}

//  go get github.com/BurntSushi/toml
//  go get github.com/go-yaml/yaml
//  其他配置文件的处理方式  ini、cfg、properties、plist、config
func test4() {
	if err := MarshalAll(); err != nil {
		panic(err)
	}

	if err := UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := OtherJSONExamples(); err != nil {
		panic(err)
	}
}

// TOMLData是我们使用TOML结构标记的通用数据结构
type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// ToTOML 将 TOMLData 结构转储为 TOML 格式 bytes.Buffer
func (t *TOMLData) ToTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)

	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return b, nil
}

// Decode 将数据解码为TOMLData
func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}

// YAMLData 是我们使用 YAML 结构标记的通用数据结构
type YAMLData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

// ToYAML 将YAMLData结构转储为YAML格式bytes.Buffer
func (t *YAMLData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(d)
	return b, nil
}

// Decode 将数据解码为 YAML Data
func (t *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}

// JSONData 是我们使用JSON结构标记的通用数据结构
type JSONData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ToJSON 将JSONData结构转储为JSON格式bytes.Buffer
func (t *JSONData) ToJSON() (*bytes.Buffer, error) {
	d, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)

	return b, nil
}

// Decode 将数据解码为JSONData
func (t *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

// OtherJSONExamples 显示对json解析至其他类型的操作
func OtherJSONExamples() error {
	res := make(map[string]string)
	err := json.Unmarshal([]byte(`{"key": "value"}`), &res)
	if err != nil {
		return err
	}
	fmt.Println("We can unmarshal into a map instead of a struct:", res)
	b := bytes.NewReader([]byte(`{"key2": "value2"}`))
	decoder := json.NewDecoder(b)
	if err := decoder.Decode(&res); err != nil {
		return err
	}
	fmt.Println("we can also use decoders/encoders to work with streams:", res)
	return nil
}

// MarshalAll 建立了不同结构类型的数据并将它们转换至对应的格式
func MarshalAll() error {
	t := TOMLData{
		Name: "Name1",
		Age:  20,
	}

	j := JSONData{
		Name: "Name2",
		Age:  30,
	}

	y := YAMLData{
		Name: "Name3",
		Age:  40,
	}

	tomlRes, err := t.ToTOML()
	if err != nil {
		return err
	}

	fmt.Println("TOML Marshal =", tomlRes.String())

	jsonRes, err := j.ToJSON()
	if err != nil {
		return err
	}

	fmt.Println("JSON Marshal=", jsonRes.String())

	yamlRes, err := y.ToYAML()
	if err != nil {
		return err
	}

	fmt.Println("YAML Marshal =", yamlRes.String())
	return nil
}

const (
	exampleTOML = `name="Example1"
age=99
    `

	exampleJSON = `{"name":"Example2","age":98}`

	exampleYAML = `name: Example3
age: 97    
    `
)

// UnmarshalAll 将不同格式的数据转换至对应结构
func UnmarshalAll() error {
	t := TOMLData{}
	j := JSONData{}
	y := YAMLData{}

	if _, err := t.Decode([]byte(exampleTOML)); err != nil {
		return err
	}
	fmt.Println("TOML Unmarshal =", t)

	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal =", j)

	if err := y.Decode([]byte(exampleYAML)); err != nil {
		return err
	}
	fmt.Println("Yaml Unmarshal =", y)
	return nil
}

// 管道
func wordCount(f io.Reader) map[string]int {
	result := make(map[string]int)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		result[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return result
}

func test5() {
	fmt.Printf("string: number_of_occurrences\n\n")
	for key, value := range wordCount(os.Stdin) {
		fmt.Printf("%s: %d\n", key, value)
	}
}

// 信号量

// CatchSig 为SIGINT中断设置一个监听器
func CatchSig(ch chan os.Signal, done chan bool) {
	// 在等待信号时阻塞
	sig := <-ch
	// 当接收到信号时打印
	fmt.Println("\nsig received:", sig)

	// 对信号类型进行处理
	switch sig {
	case syscall.SIGINT:
		fmt.Println("handling a SIGINT now!")
	case syscall.SIGTERM:
		fmt.Println("handling a SIGTERM in an entirely different way!")
	default:
		fmt.Println("unexpected signal received")
	}

	// 终止
	done <- true
}

func test6() {
	// 初始化通道
	signals := make(chan os.Signal)
	done := make(chan bool)

	// 将它们连接到信号lib
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// 如果一个信号被这个go例程捕获，它将写入 done
	go CatchSig(signals, done)

	fmt.Println("Press ctrl-c to terminate...")
	// 程序会持续打印日志直到done通道被写入
	<-done
	fmt.Println("Done!")
}

// ANSI 文本着色

//文本的颜色
type Color int

const (
	// 默认颜色
	ColorNone = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Black Color = -1
)

// ColorText 存储了文本及所属的颜色
type ColorText struct {
	TextColor Color
	Text      string
}

func (r *ColorText) String() string {
	if r.TextColor == ColorNone {
		return r.Text
	}

	value := 30
	if r.TextColor != Black {
		value += int(r.TextColor)
	}
	return fmt.Sprintf("\033[0;%dm%s\033[0m", value, r.Text)
}

func test7() {
	r := ColorText{
		TextColor: Red,
		Text:      "I'm red!",
	}

	fmt.Println(r.String())

	r.TextColor = Green
	r.Text = "Now I'm green!"

	fmt.Println(r.String())

	r.TextColor = ColorNone
	r.Text = "Back to normal..."

	fmt.Println(r.String())
}

func main() {
	test7()
}
