package main

import (
	"context"
	"fmt"
	hello "golang_exercise/grpc/pb"
	"google.golang.org/grpc"
)

func main(){
   conn,e:= grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
   if e!=nil {
	   fmt.Println(e)
   }
   defer conn.Close()

   client := hello.NewHelloGrpcClient(conn)
   res,e:= client.SayHi(context.Background(), &hello.Req{Message: "message from client"})

   if e != nil{
	   fmt.Println(e)
   }

	fmt.Println(res.GetMessage())
}
