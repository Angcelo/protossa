package main

import (
	"context"
	"fmt"
	pb "github.com/Angcelo/ESB/protoCliente"
)

var ctx = context.TODO()

type Server struct {
	pb.UnimplementedClienteServer
}

func (s *Server) GetPedidos(ctx context.Context, in *pb.vacio) (out *pb.ClienteReponse, error)

func main() {
	fmt.Println("Hello, world.")
}