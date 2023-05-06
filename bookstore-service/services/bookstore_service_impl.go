package services

import (
	"bookstore-service/services/bookstore"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

type BookStoreServiceImpl struct {
	bookstore.UnimplementedBookStoreServiceServer
}

var db []*bookstore.Book

func NewHelloService() bookstore.BookStoreServiceServer {
	return BookStoreServiceImpl{}
}

func (h BookStoreServiceImpl) CreateBook(ctx context.Context, request *bookstore.CreateBookRequest) (*bookstore.CreateBookResponse, error) {

	md, _ := metadata.FromIncomingContext(ctx)
	uid := md.Get("x-user-id")[0]
	book := &bookstore.Book{
		Id:      uuid.New().String(),
		Title:   request.Title,
		Content: request.Content,
		OwnerId: uid,
	}
	db = append(db, book)

	return &bookstore.CreateBookResponse{
		Message: fmt.Sprintf("success, login as %s", uid),
		Data:    book,
	}, nil
}

func (h BookStoreServiceImpl) GetBook(ctx context.Context, _ *bookstore.GetBookListRequest) (*bookstore.GetBookListResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	uid := md.Get("x-user-id")[0]

	return &bookstore.GetBookListResponse{
		Message: fmt.Sprintf("success, login as %s", uid),
		Data:    db,
	}, nil
}

func (h BookStoreServiceImpl) mustEmbedUnimplementedBookStoreServiceServer() {
	panic("implement me")
}
