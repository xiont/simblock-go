package printfile

import (
	"bufio"
	"os"
)

type FileWriter struct {
	filePath string
	fileObj  *os.File
	writeObj *bufio.Writer
}

func NewFilePrinter(filepath string) *FileWriter {
	fw := FileWriter{
		filePath: filepath,
	}
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	_ = f.Close()

	if fileObj, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err == nil {
		writeObj := bufio.NewWriterSize(fileObj, 4096)

		fw.fileObj = fileObj
		fw.writeObj = writeObj
	} else {
		return nil
	}

	return &fw
}

func (fw *FileWriter) Print(content string) {
	_, _ = fw.writeObj.Write([]byte(content))
}

func (fw *FileWriter) Flush() {
	_ = fw.writeObj.Flush()
}

func (fw *FileWriter) Close() {
	_ = fw.fileObj.Close()
}

func (fw *FileWriter) Println(content string) {
	_, _ = fw.writeObj.Write([]byte(content + "\n"))
}

//func (fw *FileWriter) Print(content string) {
//	WriteWithBufio(fw.filePath, content)
//}
//
//func (fw *FileWriter) Println(content string) {
//	WriteWithBufio(fw.filePath, content+"\n")
//}

//使用bufio包中Writer对象的相关方法进行数据的写入
//func WriteWithBufio(name, content string) {
//	if fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644); err == nil {
//		defer fileObj.Close()
//		writeObj := bufio.NewWriterSize(fileObj, 4096)
//		//
//		//if _,err := writeObj.WriteString(content);err == nil {
//		//	fmt.Println("Successful appending buffer and flush to file with bufio's Writer obj WriteString method",content)
//		//}
//
//		//使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
//		buf := []byte(content)
//		if _, err := writeObj.Print(buf); err == nil {
//			//fmt.Println("Successful appending to the buffer with os.OpenFile and bufio's Writer obj Print method.",content)
//			if err := writeObj.Flush(); err != nil {
//				panic(err)
//			}
//			//fmt.Println("Successful flush the buffer data to file ",content)
//		}
//	}
//}

type NilFileWriter struct {
	filePath string
}

func NewNilFilePrinter(filepath string) *NilFileWriter {
	fw := NilFileWriter{
		filePath: filepath,
	}

	return &fw
}

func (fw *NilFileWriter) Print(content string) {

}

func (fw *NilFileWriter) Flush() {

}

func (fw *NilFileWriter) Close() {

}

func (fw *NilFileWriter) Println(content string) {

}
