package Chapter03

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func GetSHA256HashCode(message []byte) string {
	hash := sha256.New()

	hash.Write(message)

	bytes := hash.Sum(nil)

	hashCode := hex.EncodeToString(bytes)
	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
	//return hashcode2
	return hashCode
}

func diffrent() {
	var r1 bytes.Buffer
	writer := bufio.NewWriter(&r1)
	fmt.Fprintf(writer, "%08b", sha256.Sum256([]byte("hello")))
	writer.Flush()
	r1Str := r1.String()

	var r2 bytes.Buffer
	writer2 := bufio.NewWriter(&r2)
	fmt.Fprintf(writer2, "%08b", sha256.Sum256([]byte("2")))
	writer2.Flush()
	r2Str := r2.String()

	count := 0
	for i, _ := range r1Str {
		if r1Str[i] != r2Str[i] {
			count += 1
		}
	}
	fmt.Println(count)
}

func Test041(t *testing.T) {
	//message := []byte("hello world")
	//hashCode := GetSHA256HashCode(message)
	//fmt.Println(hashCode)
	diffrent()
}
