package cmd

import (
	"auth-service/db"
	pb_auth "auth-service/pb/auth-service"
	"auth-service/repositories"
	"auth-service/services"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()
	repository := repositories.NewRepository(db.InitDatabase())

	service := services.NewService(repository)

	grpcServer := grpc.NewServer()

	pb_auth.RegisterAuthServiceServer(grpcServer, service)

	listener, err := net.Listen("tcp", os.Getenv(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))))
	if err != nil{
		log.Fatal(err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
