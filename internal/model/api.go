package model

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type SendCoinRequest struct {
	ToUser string `json:"to_user" validate:"required"`
	Amount int64  `json:"amount" validate:"required"`
}

type Error struct {
	Errors string `json:"errors"`
}
