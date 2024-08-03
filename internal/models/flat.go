package models

import (
	"encoding/json"
)

var (
	OnModerate ModerationStatus = &onModerate{}
	Created    ModerationStatus = &created{}
	Approved   ModerationStatus = &approved{}
	Declined   ModerationStatus = &declined{}
)

type Flat struct {
	Id         uint64           `json:"id"`
	HouseId    uint64           `json:"house_id"`
	Price      uint64           `json:"price"`
	RoomNumber byte             `json:"rooms"`
	Status     ModerationStatus `json:"status"`
}

func (m *Flat) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         uint64 `json:"id"`
		HouseId    uint64 `json:"house_id"`
		Price      uint64 `json:"price"`
		RoomNumber byte   `json:"rooms"`
		Status     string `json:"status"`
	}{
		Id:         m.Id,
		HouseId:    m.HouseId,
		Price:      m.Price,
		RoomNumber: m.RoomNumber,
		Status:     m.Status.String(),
	})
}

type ModerationStatus interface {
	String() string
	isEnumStatus()
}

type onModerate struct{ ModerationStatus }

func (s *onModerate) String() string {
	return "on moderate"
}

type created struct{ ModerationStatus }

func (s *created) String() string {
	return "created"
}

type approved struct{ ModerationStatus }

func (s *approved) String() string {
	return "approved"
}

type declined struct{ ModerationStatus }

func (s *declined) String() string {
	return "declined"
}
