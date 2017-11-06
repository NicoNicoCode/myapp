package main

import (
	"log"
	"net"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "myapp/show/show"
	"google.golang.org/grpc/reflection"

	_ "myapp/routers"
	"github.com/astaxie/beego"

)

const(
	port1 = "localhost:8081"
	port2 = "localhost:8082"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}


func (s *server) Echo(ctx context.Context, in *pb.RequestStr) (*pb.ReplyStr,error){
	return &pb.ReplyStr{Message:in.Message}, nil
}

func (s *server) Lock(in *pb.RequestTime,stream pb.Hello_LockServer) error {
	
	x := time.Now().Format("2006-01-02 15:04:05")
	
	for{
		timer := time.NewTimer(time.Minute*10)
			x = time.Now().Format("2006-01-02 15:04:05")
			stream.Send(&pb.ReplyTime{Message:x})
		<-timer.C
	}

	return nil

}

func start_echo(){
	lis, err := net.Listen("tcp", port1)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("开启服务失败: %v", err)
	}
}

func start_time(){
	lis, err := net.Listen("tcp", port2)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("开启服务失败: %v", err)
	}
}



func main() {

	go start_echo()
	go start_time()

	beego.Run()
}
