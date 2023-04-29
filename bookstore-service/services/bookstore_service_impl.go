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

func (h BookStoreServiceImpl) CreateBook(ctx context.Context, request *bookstore.CreateBookRequest) (*bookstore.Book, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Printf("%+v", md)
	}
	book := &bookstore.Book{
		Id:      uuid.New().String(),
		Title:   request.Title,
		Content: request.Content,
	}

	db = append(db, book)

	return book, nil
}

func (h BookStoreServiceImpl) GetBook(_ context.Context, _ *bookstore.GetBookListRequest) (*bookstore.BookList, error) {
	return &bookstore.BookList{
		List: db,
	}, nil
}

func (h BookStoreServiceImpl) mustEmbedUnimplementedBookStoreServiceServer() {
	panic("implement me")
}
