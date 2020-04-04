package model

type Item struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

var Items = map[uint32]Item{
	1: {
		ID:   1,
		Name: "Pele's shoes",
	},
	2: {
		ID:   2,
		Name: "Maradona's socks",
	},
}
