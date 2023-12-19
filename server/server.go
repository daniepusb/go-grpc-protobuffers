package server

import (
	"context"

	"pdaniel.com/go/grpc/models"
	"pdaniel.com/go/grpc/repository"
	"pdaniel.com/go/grpc/studentpb"
)


type server struct {
	repo repository.Repository
	studentpb.UnimplementedStudentServiceServer
}

func NewStudentServer(repo repository.Repository) *server {
	return &server{
		repo: repo,
	}
}