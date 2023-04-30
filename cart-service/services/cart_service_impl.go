package services

import (
	"cart-service/services/cart"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

var db []*cart.Item

type CartServiceImpl struct {
	cart.UnimplementedCartServiceServer
}

func (c CartServiceImpl) AddToCart(ctx context.Context, request *cart.AddToCartRequest) (*cart.AddToCartResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	uid := md.Get("x-user-id")[0]

	item := &cart.Item{
		Id:      uuid.New().String(),
		Title:   request.Title,
		Content: request.Content,
		OwnerId: uid,
	}
	db = append(db, item)
	return &cart.AddToCartResponse{
		Message: fmt.Sprintf("success, login as %s", uid),
		Data:    item,
	}, nil
}

func (c CartServiceImpl) GetCart(ctx context.Context, _ *cart.GetCartRequest) (*cart.GetCartResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	uid := md.Get("x-user-id")[0]

	filtered := make([]*cart.Item, 0)

	for _, d := range db {
		if d.OwnerId == uid {
			filtered = append(filtered, d)
		}
	}

	return &cart.GetCartResponse{
		Message: fmt.Sprintf("success, login as %s", uid),
		Data:    filtered,
	}, nil
}

func (c CartServiceImpl) mustEmbedUnimplementedCartServiceServer() {
	panic("implement me")
}

func NewCartService() cart.CartServiceServer {
	return CartServiceImpl{}
}
