package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFileByFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v.", err)
		return
	}
	defer fileObj.Close()
	//读文件
	var fileTmp [128]byte

	for {
		n, err := fileObj.Read(fileTmp[:])
		if err == io.EOF {
			fmt.Printf("read complete.")
			return
		}
		if err != nil {
			fmt.Printf("read from file failed, err:%v.", err)
			return
		}
		if n == 0 {
			return
		}
		fmt.Println(string(fileTmp[:n]))
	}
}

func readFileByBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v.", err)
		return
	}

	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v.", err)
			return
		}
		if line == "\n" {
			continue
		}
		fmt.Printf("%v", line)
		// fmt.Printf("%#v, %T", line, line)
	}

}

func readFileByIoutil() {

	fileObj, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v.", err)
		return
	}
	fmt.Print(string(fileObj))
}

func writeFile() {
	fileObj, err := os.OpenFile("./readme.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Open file failed, err:%v.", err)
		return
	}
	defer fileObj.Close()

	fileObj.Write([]byte("#readme\n"))
	fileObj.WriteString("golang 写文件操作")
	fmt.Printf("%b", 1&4)
}

func writeFileByBufio() {
	fileObj, err := os.OpenFile("./readme.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Open file failed, err:%v.", err)
		return
	}
	defer fileObj.Close()

	writer := bufio.NewWriter(fileObj)
	size, err := writer.Write([]byte("#go file write readme\n"))
	writer.Flush()
	if err != nil {
		fmt.Printf("write file failed. err:%v.", err)
	} else {
		fmt.Printf("Write file successed. 字节数：%v.", size)
	}

}

func writeFileByUtil() {
	content := "#Golang 文件操作示例\n"
	err := ioutil.WriteFile("./readme.md", []byte(content), 0664)
	if err != nil {
		fmt.Printf("Write file failed, err:%v.", err)
	} else {
		fmt.Printf("Write file successed.")
	}

}

func CopyFile(srcFile, dstFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("open %s failed, err: %v.\n", srcFile, err)
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open %s failed,  err:%v.", dstFile, err)
		return
	}

	defer dst.Close()

	return io.Copy(dst, src)

}

func Cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}

		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func fileSeek() {
	fileObj, err := os.OpenFile("./readme.md", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Open file failed, err:%v.", err)
		return
	}
	defer fileObj.Close()

	fileObj.Seek(-20, 2)

	var ret [1]byte
	_, err = fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("Read file failed, err:%v", err)
		return
	}

	fmt.Println(string(ret[:]))

}
func main() {
	fileSeek()
	// readFileByFile()
	// readFileByBufio()
	//readFileByIoutil()

	// writeFileByBufio()
	// writeFileByUtil()
	// CopyFile("./readme.md", "./copy.md")

	// flag.Parse()
	// if flag.NArg() == 0 {
	// 	Cat(bufio.NewReader(os.Stdin))
	// }
	// for i := 0; i < flag.NArg(); i++ {
	// 	f, err := os.Open(flag.Arg(i))
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v.", flag.Arg(i), err)
	// 		continue
	// 	}
	// 	Cat(bufio.NewReader(f))
	// }
}
