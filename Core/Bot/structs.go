package Bot

type DataItem struct {
	ID       int `json:"id"`
	ShopData struct {
		ID int `json:"id"`
	} `json:"shop"`
}
