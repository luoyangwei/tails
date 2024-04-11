package main

import (
	"bufio"
	"os"
	"tails/cmd/test/format"
	"time"
)

func main() {
	// 写一个文件
	f, _ := os.Create("./project.log")
	for {
		write := bufio.NewWriter(f)
		_, err := write.WriteString(format.NewJsonString())
		if err != nil {
			panic(err)
		}
		if err = write.Flush(); err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
