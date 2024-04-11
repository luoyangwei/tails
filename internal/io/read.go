package io

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type TextReader struct {
	fs []*bufio.Reader
}

func NewTextReader(paths []string) *TextReader {

	// open files
	fs := make([]*bufio.Reader, 0)
	for _, path := range paths {
		f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}
		bufread := bufio.NewReader(f)
		fs = append(fs, bufread)
	}

	reader := &TextReader{fs: fs}
	reader.readFiles()
	return reader
}

func (r *TextReader) readFiles() {
	if len(r.fs) == 0 {
		return
	}

	// 首次加载文件打印所有的内容
	for {
		for _, f := range r.fs {

			// ReadString reads until the first occurrence of delim in the input,
			// returning a string containing the data up to and including the delimiter.
			line, err := f.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			fmt.Print(string(line))
			//if err != nil {
			//	return lines, err
			//}
		}
	}
}
