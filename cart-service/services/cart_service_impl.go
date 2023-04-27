package services

import (
	"cart-service/services/cart"
	"context"
	"github.com/google/uuid"
)

var db []*cart.Item

type CartServiceImpl struct {
	cart.UnimplementedCartServiceServer
}

func (c CartServiceImpl) AddToCart(_ context.Context, request *cart.AddToCartRequest) (*cart.Item, error) {
	item := &cart.Item{
		Id:      uuid.New().String(),
		Title:   request.Title,
		Content: request.Content,
	}
	db = append(db, item)
	return item, nil
}

func (c CartServiceImpl) GetCart(_ context.Context, _ *cart.GetCartRequest) (*cart.CartList, error) {
	return &cart.CartList{
		List: db,
	}, nil
}

func (c CartServiceImpl) mustEmbedUnimplementedCartServiceServer() {
	panic("implement me")
}

func NewCartService() cart.CartServiceServer {
	return CartServiceImpl{}
}
