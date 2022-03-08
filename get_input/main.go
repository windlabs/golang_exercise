package get_input

import (
	"bufio"
	"fmt"
	"os"
)

func useScan() {
	//var s string
	fmt.Print("请输入内容：")
	//fmt.Scanln(&s)
	//fmt.Printf("你输入的内容是：%s\n", s)
}

func useBufio() {
	var s string
	fmt.Print("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("你的输入是：%s", s)
}
func main() {
	useBufio()
	// useScan()
}
