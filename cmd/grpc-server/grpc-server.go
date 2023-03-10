package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/aria3ppp/url-shortener-grpc/internal/core/usecase"
	"github.com/aria3ppp/url-shortener-grpc/internal/gapi"
	"github.com/aria3ppp/url-shortener-grpc/internal/generator"
	"github.com/aria3ppp/url-shortener-grpc/internal/pb"
	"github.com/aria3ppp/url-shortener-grpc/internal/repository"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	repo := repository.NewRepository(db)
	generator := generator.NewRandomStringGenerator(6)

	serviceUseCases := usecase.NewService(repo, generator)

	server := grpc.NewServer()
	reflection.Register(server) // enable server reflection
	pb.RegisterUrlShortenerServer(server, gapi.NewServer(serviceUseCases))

	l, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVER_PORT"))
	if err != nil {
		panic(err)
	}

	if err := server.Serve(l); err != nil {
		panic(err)
	}
}
