package model

type User struct {
	ID       string
	Username string
	Password string
	Coins    int64
}

type Item struct {
	ID    string
	Type  string
	Price int64
}

type Transaction struct {
	Username string
	Amount   int64
}
