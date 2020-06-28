package printfile

import (
	"bufio"
	"os"
)

type FileWriter struct {
	filePath string
}

func NewFileWriter(filepath string) *FileWriter {
	fw := FileWriter{
		filepath,
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	return &fw
}

func (fw *FileWriter) Write(content string) {
	WriteWithBufio(fw.filePath, content)
}

//使用bufio包中Writer对象的相关方法进行数据的写入
func WriteWithBufio(name, content string) {
	if fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err == nil {
		defer fileObj.Close()
		writeObj := bufio.NewWriterSize(fileObj, 4096)
		//
		//if _,err := writeObj.WriteString(content);err == nil {
		//	fmt.Println("Successful appending buffer and flush to file with bufio's Writer obj WriteString method",content)
		//}

		//使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
		buf := []byte(content)
		if _, err := writeObj.Write(buf); err == nil {
			//fmt.Println("Successful appending to the buffer with os.OpenFile and bufio's Writer obj Write method.",content)
			if err := writeObj.Flush(); err != nil {
				panic(err)
			}
			//fmt.Println("Successful flush the buffer data to file ",content)
		}
	}
}
