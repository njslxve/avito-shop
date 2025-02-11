package warehouse

type Warehouse struct {
	wh map[string]int64
}

func New() *Warehouse {
	wh := make(map[string]int64)

	wh["t-shirt"] = 80
	wh["cup"] = 20
	wh["book"] = 50
	wh["pen"] = 10
	wh["powerbank"] = 200
	wh["hoody"] = 300
	wh["umbrella"] = 200
	wh["socks"] = 10
	wh["wallet"] = 50
	wh["pink-hoody"] = 500

	return &Warehouse{wh: wh}
}
