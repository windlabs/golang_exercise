package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"

	"socket/common"
)

func main() {
	tcpConn()
	// udpConn()
}
func udpConn() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345})

	if err != nil {
		fmt.Printf("Connect socket failed,err:%v.", err)
		return
	}
	defer conn.Close()

	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("请输入：")
		msg, _ := reader.ReadString('\n')
		conn.Write([]byte(msg))
		n, addr, err := conn.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Printf("Read from udp failed, err:%v", err)
			continue
		}
		fmt.Printf("Recv:%v addr:%v, count:%v\n", string(reply[:]), addr, n)
	}

}

func tcpConn() {
	//1 与Server端建立连接
	conn, err := net.Dial("tcp", net.JoinHostPort("127.0.0.1", "12345"))

	if err != nil {
		fmt.Printf("Connect failed, err:%v", err)
		return
	}
	//2、发送数据
	var msg string
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Println("请输入：")
	// 	msg, _ = reader.ReadString('\n')
	// 	msg = strings.TrimSpace(msg)
	// 	if msg == "exit" {
	// 		break
	// 	}
	// 	conn.Write([]byte(msg))
	// }

	msg = "Hello hi"
	for i := 0; i < 10; i++ {
		msgByte, _ := common.Encode(msg + strconv.Itoa(i))
		//msgByte := []byte(msg + strconv.Itoa(i))
		conn.Write(msgByte)
	}
	conn.Close()
}
