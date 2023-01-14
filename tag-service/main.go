package main

import (
	"flag"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "gorm_pro/tag-service/proto"
	"gorm_pro/tag-service/server"
	"log"
	"net"
	"net/http"
)

var grpcPort string
var httpPort string
var port string

func init() {
	flag.StringVar(&grpcPort, "grpc_port", "8001", "gRPC 启动端口号")
	flag.StringVar(&httpPort, "http_port", "9001", "HTTP 启动端口号")
	flag.StringVar(&port, "port", "8003", "启动端口号")
	flag.Parse()
}

func main() {
	l, err := RunTCPServer(port)
	if err != nil {
		log.Fatalf("Run TCP Server err: %v", err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := RunGrpcServer()
	httpS := RunHttpServer(port)
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	err = m.Serve()
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}
func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp-server", ":"+port)
}

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}
