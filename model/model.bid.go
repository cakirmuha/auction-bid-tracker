package model

type Bid struct {
	UserID   uint32 `json:"user_id"`
	UserName string `json:"user_name,omitempty"`
	ItemID   uint32 `json:"item_id"`
	ItemName string `json:"item_name,omitempty"`
	Amount   uint64 `json:"amount"`
	BidTime  uint64 `json:"bid_time,omitempty"`
}
