package models

//Items details
type Items struct {
	ID       string  `json:"id, omitempty"`
	Title    string  `json:"title, omitempty"`
	Price    string  `json:"price, omitempty"`
	Quantity *string `json:"qt, omitempty"`
}

type Cart struct {
	Keys string
	Prod Items
}

type Error struct {
	Code int
	Msg  string
}
