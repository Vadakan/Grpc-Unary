package main

import (
	"context"
	"fmt"
	"github.com/hello_service/stub/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

const(
 address = ":50051"
)
type Server struct {
  pb.UnimplementedGreeterServer
}

func (Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse,error) {
	fmt.Println("Input recieved",in.GetName())

	return &pb.HelloResponse{Message : "hello"+in.GetName()},nil
}


func main(){

  lis,_ := net.Listen("tcp",address)

   s := grpc.NewServer()


   pb.RegisterGreeterServer(s,&Server{})


   err := s.Serve(lis)
   if err != nil{
      log.Fatalln("server error!!!!")
   }
}
