package model

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type SendCoinRequest struct {
	ToUser string `json:"toUser" validate:"required"`
	Amount int64  `json:"amount" validate:"required"`
}

type Error struct {
	Errors string `json:"errors"`
}

type InfoResponse struct {
	Coins       int64      `json:"coins"`
	Inventory   []ItemInfo `json:"inventory"`
	CoinHistory History    `json:"coinHistory"`
}

type ItemInfo struct {
	Type     string `json:"type"`
	Quantity int64  `json:"quantity"`
}

type History struct {
	Received []Received `json:"received,omitempty"`
	Sent     []Sent     `json:"sent,omitempty"`
}

type Received struct {
	FromUser string `json:"fromUser"`
	Amount   int64  `json:"amount"`
}

type Sent struct {
	ToUser string `json:"toUser"`
	Amount int64  `json:"amount"`
}
