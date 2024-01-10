package main

import (
	"log"
	"net"

	"github.com/pinpointdev90/go-gRPC/article/pb"
	"github.com/pinpointdev90/go-gRPC/article/repository"
	"github.com/pinpointdev90/go-gRPC/article/service"
	"google.golang.org/grpc"
)

func main() {

	// articleサーバーに接続
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	defer lis.Close()

	// RepositoryとServiceを作成
	repository, err := repository.NewsqliteRepo()
	if err != nil {
		log.Fatalf("Failed to create sqlite repository: %v\n", err)
	}
	service := service.NewService(repository)

	//サーバーにarticleサービスを登録
	server := grpc.NewServer()
	pb.RegisterArticleServiceServer(server, service)

	//articleサーバーを起動
	log.Println("Listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
