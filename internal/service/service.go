package service

import (
	"github.com/njslxve/avito-shop/internal/service/auth"
	"github.com/njslxve/avito-shop/internal/service/shop"
)

type Service struct {
	Auth *auth.Auth
	Shop *shop.ShopService
}

func New(auth *auth.Auth, shop *shop.ShopService) *Service {
	return &Service{
		Auth: auth,
		Shop: shop,
	}
}
