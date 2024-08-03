package models

type HouseBuilder struct {
	Address          *string `json:"address"`
	ConstructionDate *int    `json:"year"`
	Developer        *string `json:"developer,omitempty"`
}

type FlatBuilder struct {
	HouseId *uint64 `json:"house_id"`
	Price   *uint64 `json:"price"`
	Rooms   *byte   `json:"rooms"`
}

type FlatStatus struct {
	Id    *uint64 `json:"id"`
	Value *string `json:"status"`
}
