package main

import (
	"log"
	"bufio"
	"os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "myapp/show/show"
)

const (
	address1     = "192.168.34.97:8081"
	address2     = "192.168.34.97:8082"
)

func bee_echo(){
	// Set up a connection to the server.
	conn, err := grpc.Dial(address1, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("无法连接: %v", err)
	}
	defer conn.Close()
	
	c := pb.NewHelloClient(conn)
	empty :=""

	for{
		inputReader := bufio.NewReader(os.Stdin)
		empty,err = inputReader.ReadString('\n')
		
		r, err := c.Echo(context.Background(), &pb.RequestStr{Message:empty})
				
		if err != nil {
			log.Fatalf("无法握手: %v", err)
		}

		log.Printf("%s", r.Message)
	}
}

func bee_time(){
		// Set up a connection to the server.
		conn, err := grpc.Dial(address2, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("无法连接: %v", err)
		}
		defer conn.Close()
		
		c := pb.NewHelloClient(conn)
		
		empty :=""
		
		r, err := c.Lock(context.Background(), &pb.RequestTime{Message:empty})
				
		if err != nil {
			log.Fatalf("无法握手: %v", err)
		}
	
		time_cur,err2:=r.Recv()
		time_clone:=time_cur
	
		if err2==nil{
			log.Printf("%s",time_cur.Message)
		}
	
	
		for {
			time_cur,err2=r.Recv()
	
			if err2 == nil && time_cur!=time_clone{
				log.Printf("%s",time_cur.Message)
				time_clone=time_cur
			}
		}
}

func main() {
	go bee_echo()
	bee_time()	
}
