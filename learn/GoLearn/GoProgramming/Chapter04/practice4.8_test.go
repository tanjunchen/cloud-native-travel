package Chapter03

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
	"unicode"
	"unicode/utf8"
)

type CharCount struct {
	zhCount    int
	numCount   int
	spaceCount int
	otherCount int
}

func MainStart() {
	file, err := os.Open("D:/GoLearn/GoProgramming/Chapter04/file.txt")
	if err != nil {
		fmt.Printf("open file err=%v \n", err)
		return
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range []rune(str) {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.zhCount++
			case v == ' ' || v == '\t':
				count.spaceCount++
			case v >= '0' && v <= '9':
				count.numCount++
			default:
				count.otherCount++
			}
		}
	}
	fmt.Printf("字符的个数为：%v，数字的个数为：%v，空格的个数为：%v，其它字符的个数为：%v",
		count.zhCount, count.numCount, count.spaceCount, count.otherCount)
}

func charCount() {
	//统计输入中每个Unicode码点出现的次数
	counts := make(map[rune]int)
	typeCounts := make(map[string]int)
	var uftlen [utf8.UTFMax + 1]int
	invalid := 0
	input := bufio.NewReader(os.Stdin)
	for {
		r, n, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		if n == 1 && r == unicode.ReplacementChar {
			invalid++
			continue
		}
		switch {
		case unicode.IsControl(r):
			typeCounts["Control"]++
		case unicode.IsLetter(r):
			typeCounts["Letter"]++
		case unicode.IsNumber(r):
			typeCounts["Number"]++
		case unicode.IsSpace(r):
			typeCounts["Space"]++
		default:
			typeCounts["Other"]++
		}
		counts[r]++
		uftlen[n]++
	}
	fmt.Printf("Rune\tCount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("\nLen\tCount\n")
	for i, n := range uftlen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("\nType\tCount\n")
	for key, value := range typeCounts {
		fmt.Printf("%s\t%d\n", key, value)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UFT-8 characters\n", invalid)
	}

}
func Test048(t *testing.T) {
	MainStart()
}

func countCharacter() {
	in := bufio.NewReader(os.Stdin)
	var control, digit, graphic, letter, lower, mark, number, print, punct, space, symbol, title, upper int
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if unicode.IsControl(r) {
			control++
		}
		if unicode.IsDigit(r) {
			digit++
		}
		if unicode.IsGraphic(r) {
			graphic++
		}
		if unicode.IsLetter(r) {
			letter++
		}
		if unicode.IsLower(r) {
			lower++
		}
		if unicode.IsMark(r) {
			mark++
		}
		if unicode.IsNumber(r) {
			number++
		}
		if unicode.IsPrint(r) {
			print++
		}
		if unicode.IsPunct(r) {
			punct++
		}
		if unicode.IsSpace(r) {
			space++
		}
		if unicode.IsSpace(r) {
			symbol++
		}
		if unicode.IsTitle(r) {
			title++
		}
		if unicode.IsUpper(r) {
			upper++
		}
	}
	fmt.Printf("control: %d\n", control)
	fmt.Printf("digit: %d\n", digit)
	fmt.Printf("graphic: %d\n", graphic)
	fmt.Printf("letter: %d\n", letter)
	fmt.Printf("lower: %d\n", lower)
	fmt.Printf("mark: %d\n", mark)
	fmt.Printf("number: %d\n", number)
	fmt.Printf("print: %d\n", print)
	fmt.Printf("punct: %d\n", punct)
	fmt.Printf("space: %d\n", space)
	fmt.Printf("symbol: %d\n", symbol)
	fmt.Printf("title: %d\n", title)
	fmt.Printf("upper: %d\n", upper)
}
