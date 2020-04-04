package model

type User struct {
	ID   uint32 `json:"id"`
	Name string `json:"name"`
}

var Users = map[uint32]User{
	1: {
		ID:   1,
		Name: "fabregas",
	},
	2: {
		ID:   2,
		Name: "walcott",
	},
	3: {
		ID:   3,
		Name: "ronaldo",
	},
	4: {
		ID:   4,
		Name: "messi",
	},
}
