package data

import context "golang.org/x/net/context"
import "google.golang.org/grpc"

func Register(s *grpc.Server) {
	RegisterProductActorServer(s, &server{})
}

// server is used to implement product.Server
type server struct{}

// Create implements product.Create
func (s *server) Create(ctx context.Context, in *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{Code: "Hello " + in.Code}, nil
}
