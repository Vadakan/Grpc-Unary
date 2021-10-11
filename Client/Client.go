package main

import (
	"context"
	"fmt"
	"github.com/hello_service/stub/pb"
	"google.golang.org/grpc"
	"time"
)

const(
   target_Server_address = "localhost:50051"
)

func main(){

      conn, _ := grpc.Dial(target_Server_address,grpc.WithInsecure(),grpc.WithBlock())
	  defer conn.Close()

	  c := pb.NewGreeterClient(conn)

	  ctx,cancel := context.WithTimeout(context.Background(),time.Second)

	  defer cancel()

	  h1 := new(pb.HelloRequest)
	  h1.Name="sundar"
	  r,_ := c.SayHello(ctx,h1)
      fmt.Println(r.GetMessage())


}
