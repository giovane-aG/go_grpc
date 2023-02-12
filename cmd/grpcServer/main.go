package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/giovane-aG/go_grpc/internal/database"
	"github.com/giovane-aG/go_grpc/internal/pb"
	"github.com/giovane-aG/go_grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {
	const PostgresDriver = "postgres"
	const User = "postgres"
	const Host = "localhost"
	const Port = "5433"
	const Password = "postgres"
	const DbName = "postgres"
	var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

	db, err := sql.Open(PostgresDriver, DataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
