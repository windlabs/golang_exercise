package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"socket/common"
)

//Tcp server

func connProcess(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := common.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("message decode failed, err:%v", err)
		}
		fmt.Println(msg)
	}

}

func udpConn() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345})

	if err != nil {
		fmt.Printf("listen failed,err %v.", err)
		return
	}

	defer listen.Close()

	var data [1024]byte
	for {
		n, addr, err := listen.ReadFromUDP(data[:])
		fmt.Println(n, err)
		if err == io.EOF {
			continue
		}
		if err != nil {
			fmt.Printf("read udp failed.err:%v.", err)
			continue
		}

		fmt.Printf("data:%#v, addr:%v, count:%v \n", strings.TrimSpace(string(data[:n])), addr, n)
		//发送数据
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Printf("Write to udp failed, err:%v.", err)
			continue
		}
	}
}

func tcpConn() {
	// 1 、 Server端启动
	listener, err := net.Listen("tcp", net.JoinHostPort("127.0.0.1", "12345"))
	if err != nil {
		fmt.Printf("Start server failed, err:%v.", err)
		return
	}

	for {
		//2、等待连接
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("Accept failed, err:%v\n", err)
			return
		}
		//3、通信
		go connProcess(conn)
	}
}

func main() {
	tcpConn()

	// udpConn()
}
