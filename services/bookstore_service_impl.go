package services

import (
	"context"
	"github.com/google/uuid"
	"try-gcp-grpc/services/bookstore"
)

type BookStoreServiceImpl struct {
	bookstore.UnimplementedBookStoreServiceServer
}

var db []*bookstore.Book

func NewHelloService() bookstore.BookStoreServiceServer {
	return BookStoreServiceImpl{}
}

func (h BookStoreServiceImpl) CreateBook(_ context.Context, request *bookstore.CreateBookRequest) (*bookstore.Book, error) {
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
