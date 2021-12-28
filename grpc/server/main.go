package main

import (
	"context"
	"fmt"
	hello "golang_exercise/grpc/pb"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	 hello.UnimplementedHelloGrpcServer
}

func (s *server) SayHi(ctx context.Context, req *hello.Req)(res *hello.Res,err error){
	fmt.Println(req.GetMessage())

	return &hello.Res{Message: "服务端返回信息内容。"},nil
}

func main(){
	l,err :=net.Listen("tcp", ":8888")
	if err!=nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloGrpcServer(s, &server{})

	s.Serve(l)
}
