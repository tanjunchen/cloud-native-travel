package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	h bool

	v, V bool
	t, T bool
	q    *bool

	s string
	p string
	c string
	g string

    languages []string
)

func init() {

	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")


	flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}


}

type sliceValue []string

func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

func (s *sliceValue) Set(val string) error {
	*s = strings.Split(val, ",")
	return nil
}

func (s *sliceValue) Get() interface{} { return []string(*s) }

func (s *sliceValue) String() string { return strings.Join([]string(*s), ",") }