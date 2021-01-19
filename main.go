package main

import (
	"context"
	pb "github.com/trifoliumproj/nh/api/generated/github.com/trifoliumproj/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedAPIServer
}

func (s *Server) Gallery(ctx context.Context, in *pb.GalleryRequest) (*pb.GalleryResponse, error) {
	return &pb.GalleryResponse{
		Id:           0,
		MediaId:      0,
		Title:        nil,
		Images:       nil,
		Scanlator:    "",
		UploadDate:   "",
		Tags:         nil,
		NumPages:     0,
		NumFavorites: 0,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAPIServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
