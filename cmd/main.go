package main

import (
	"flag"
	"strings"
	"tails/internal/io"
)

type pathsFlag []string

func (i *pathsFlag) String() string {
	return strings.Join(*i, ",")
}

func (i *pathsFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var paths pathsFlag

func init() {
	flag.Var(&paths, "f", "文件路径, for example: -f=../file1.log -f=../file2.log")
}

func main() {
	flag.Parse()

	// 接收参数
	io.NewTextReader(paths)
}
