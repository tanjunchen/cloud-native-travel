package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func test1() {
	r := strings.NewReader("some io.Reader stream to be read\n====\n")

	//if _, err := io.Copy(os.Stdout, r); err != nil {
	//	log.Fatal(err)
	//}
	//
	//r1 := strings.NewReader("first reader\n")
	//r2 := strings.NewReader("second reader\n")
	//buf := make([]byte, 8)
	//
	//// buf is used here...
	//if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
	//	log.Fatal(err)
	//}
	//
	//// ... reused here also. No need to allocate an extra buffer.
	//if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
	//	log.Fatal(err)
	//}

	if _, err := io.CopyN(os.Stdout, r, 6); err != nil {
		log.Fatal(err)
	}
}

func Copy(in io.ReadSeeker, out io.Writer) error {
	w := io.MultiWriter(out, os.Stdout)
	if _, err := io.Copy(w, in); err != nil {
		return err
	}
	in.Seek(0, 0)

	buf := make([]byte, 64)
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	// 打印换行
	fmt.Println()

	return nil
}

func test2() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("stdout on Copy = ")
	if err := Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer =", out.String())
}

func buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)

	var b = new(bytes.Buffer)

	b.Write(rawBytes)

	// 或者 b = bytes.NewBuffer(rawBytes)

	// b = bytes.NewBufferString(rawString)

	return b
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func workWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array ❤️"

	b := buffer(rawString)

	// 使用b.Bytes()可以快速从字节缓冲区获取字节切片
	// 使用b.String()可以快速从字节缓冲区获取字符串
	fmt.Println(b.String())

	// 由于*bytes.Buffer类型的b实现了io Reader 我们可以使用常见的reader函数
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)

	// 可以创建一个 bytes reader 它实现了
	// io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner, and io.RuneScanner
	// 接口
	reader := bytes.NewReader([]byte(rawString))

	// 我们可以使用其创建 scanner 以允许使用缓存读取和建立 token
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanWords)

	// 遍历所有的扫描token
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return nil
}

// searchString 展示了一系列在字符串中进行查询的方法
func searchString() {
	s := "this is a test"

	// 返回 true 表明包含子串
	fmt.Println(strings.Contains(s, "this"))

	// 返回 true 表明包含子串中的任何一字符a或b或c
	fmt.Println(strings.ContainsAny(s, "abc"))

	// 返回 true 表明以该子串开头
	fmt.Println(strings.HasPrefix(s, "this"))

	// 返回 true 表明以该子串结尾
	fmt.Println(strings.HasSuffix(s, "test"))
}

// modifyString 展示了一系列修改字符串的方法
func modifyString() {
	s := "simple string"

	// 输出 [simple string]
	fmt.Println(strings.Split(s, " "))

	// 输出 "Simple String"
	fmt.Println(strings.Title(s))

	// 输出 "simple string" 会移除头部和尾部的空白
	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// stringReader 演示了如何快速创建一个字符串的io.Reader接口
func stringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)

	// 在标准输出上打印 s
	io.Copy(os.Stdout, r)
}

func test3() {
	err := workWithBuffer()
	if err != nil {
		panic(err)
	}

	// each of these print to stdout
	searchString()
	modifyString()
	stringReader()
}

func operate() error {
	// 文件权限 0755 类似于你在命令行中使用的 chown，
	// 这将创建一个目录 /tmp/example，
	// 你也可以使用绝对路径而不是相对路径
	if err := os.Mkdir("example_dir", os.FileMode(0755)); err != nil {
		return err
	}

	// 跳转到 /tmp 目录
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	// f是一个通用的文件对象 它还实现了多个接口，并且如果在打开时设置了正确的方式，则可以用作读取器或写入器
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	// 向文件写入长度已知的数据 并确认写入成功
	value := []byte("hello\n")
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}

	if err := f.Close(); err != nil {
		return err
	}

	// 读取文件
	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		return err
	}

	// 跳转到 /tmp 文件夹
	if err := os.Chdir(".."); err != nil {
		return err
	}

	// 删除建立的文件夹
	// os.RemoveAll如果传递了错误的文件夹路径会返回错误
	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}

	return nil
}

// 读取并转换为大写后复制内容到目标文件
func capitalizer(f1 *os.File, f2 *os.File) error {
	if _, err := f1.Seek(0, 0); err != nil {
		return err
	}

	var tmp = new(bytes.Buffer)

	if _, err := io.Copy(tmp, f1); err != nil {
		return err
	}

	s := strings.ToUpper(tmp.String())

	if _, err := io.Copy(f2, strings.NewReader(s)); err != nil {
		return err
	}
	defer f2.Close()
	defer f1.Close()
	return nil
}

// 建立两个文件 将其中一个的内容转换为大写复制给另一个
func capitalizerExample() error {
	f1, err := os.Create("file1.txt")
	if err != nil {
		return err
	}

	if _, err := f1.Write([]byte(` this file contains a number of words and new lines`)); err != nil {
		return err
	}

	f2, err := os.Create("file2.txt")
	if err != nil {
		return err
	}

	if err := capitalizer(f1, f2); err != nil {
		return err
	}
	if err := os.Remove("file1.txt"); err != nil {
		return err
	}

	if err := os.Remove("file2.txt"); err != nil {
		return err
	}

	return nil
}

func test4() {
	if err := operate(); err != nil {
		panic(err)
	}

	if err := capitalizerExample(); err != nil {
		panic(err)
	}
}

// Movie用来存储CSV解析后的内容
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV 展示了如何处理CSV
// 接收的参数通过io.Reader传入
func readCSV(b io.Reader) ([]Movie, error) {

	//返回的是csv.Reader
	r := csv.NewReader(b)

	// 分隔符和注释是csv.Reader结构体中的字段
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	// 读取并返回一个字符串切片和错误信息
	// 我们也可以将其用于字典键或其他形式的查找
	// 此处忽略了返回的切片 目的是跳过csv首行标题
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// 循环直到全部处理完毕
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}
	return movies, nil
}

// addMoviesFromText 将字符串按 CSV 格式解析
func addMoviesFromText() error {

	in := `
- first our headers
movie title;director;year released

- then some data
Guardians of the Galaxy Vol. 2;James Gunn;2017
Star Wars: Episode VIII;Rian Johnson;2017
`

	b := bytes.NewBufferString(in)
	m, err := readCSV(b)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", m)
	return nil
}

// 结构体Book有Author和Title两个字段
type Book struct {
	Author string
	Title  string
}

// Books是Book的切片类型
type Books []Book

// toCSV将Books写入传进来的 io.Writer
// 返回任何可能发生的错误
func (books *Books) toCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}

	n.Flush()
	return n.Error()
}

// writeCSVOutput 初始化Books并调用ToCSV
// 并写入到标准输出
func writeCSVOutput() error {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}

	return b.toCSV(os.Stdout)
}

// writeCSVBuffer 初始化Books并调用ToCSV
// 并写入到bytes.Buffers
func writeCSVBuffer() (*bytes.Buffer, error) {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}

	w := &bytes.Buffer{}
	err := b.toCSV(w)
	return w, err
}

func test5() {
	if err := addMoviesFromText(); err != nil {
		panic(err)
	}

	if err := writeCSVOutput(); err != nil {
		panic(err)
	}

	buffer, err := writeCSVBuffer()
	if err != nil {
		panic(err)
	}

	fmt.Println("Buffer = ", buffer.String())
}

// 这里展示了临时文件操作
func WorkWithTemp() error {
	// 如果你需要一个临时文件夹，存贮类似与template1-10.html这样的文件
	// 首个参数使用空字符串，意味着会在默认的临时目录中创建以后一个参数为开头名称的文件夹
	// 该函数实际调用了os.TempDir()
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}

	// 这会在整个操作完成后移除该临时文件夹及其中的所有文件
	defer os.RemoveAll(t)

	// 文件夹t必须存在否则将返回错误
	// tf是*os.File类型
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())

	// 通常情况下我们在函数的最后部分删除临时文件
	// 不过通过前面的defer已经完成了这个任务

	return nil
}

func test6() {
	if err := WorkWithTemp(); err != nil {
		panic(err)
	}
}

type YAMLDecoder struct {
	r         io.ReadCloser
	scanner   *bufio.Scanner
	remaining []byte
}

const yamlSeparator = "\n---"

// splitYAMLDocument is a bufio.SplitFunc for splitting YAML streams into individual documents.
func splitYAMLDocument(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	sep := len([]byte(yamlSeparator))
	if i := bytes.Index(data, []byte(yamlSeparator)); i >= 0 {
		// We have a potential document terminator
		i += sep
		after := data[i:]
		if len(after) == 0 {
			// we can't read any more characters
			if atEOF {
				return len(data), data[:len(data)-sep], nil
			}
			return 0, nil, nil
		}
		if j := bytes.IndexByte(after, '\n'); j >= 0 {
			return i + j + 1, data[0 : i-sep], nil
		}
		return 0, nil, nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func test7() {
	path := "go-package/iotest/spark-operator-crds.yaml"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(b) > 64*1024)
	r := ioutil.NopCloser(bytes.NewReader(b))
	scanner := bufio.NewScanner(r)
	scanner.Buffer([]byte{}, 5 * 1024 * 1024)
	scanner.Split(splitYAMLDocument)
	fmt.Println(scanner.Text())
}

func main() {
	test7()
}
