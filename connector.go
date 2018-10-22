package main

import (
	// "context"
	"log"
	"net"
	"time"

	"./data"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9091"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// server is used to implement helloworld.GreeterServer.
type product_server struct{}

// SayHello implements helloworld.GreeterServer
func (s *product_server) Create(ctx context.Context, in *data.ProductRequest) (*data.ProductResponse, error) {
	// To-Do: code to create the data.
	return &data.ProductResponse{Code: "Hello " + in.Code}, nil
}

func main() {
	// Docker need time to startup the database for the app to connect
	// this needs to be hanlded in a better way.
	// https://docs.docker.com/compose/startup-order/
	time.Sleep(2 * time.Second)

	db, err := gorm.Open("postgres", "host=db port=5432 user=in_user dbname=in_db password=in_password sslmode=disable")
	if err != nil {
		print(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	data.RegisterGreeterServer(s, &product_server)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
