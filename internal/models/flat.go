package models

import (
	"encoding/json"
	"errors"
)

var (
	OnModerate ModerationStatus = &onModerate{}
	Created    ModerationStatus = &created{}
	Approved   ModerationStatus = &approved{}
	Declined   ModerationStatus = &declined{}
)

type Flat struct {
	Id          uint64           `json:"id"`
	UnitNumber  uint64           `json:"unit_number"`
	HouseId     uint64           `json:"house_id"`
	Price       uint64           `json:"price"`
	RoomNumber  byte             `json:"rooms"`
	Status      ModerationStatus `json:"status"`
	ModeratorId string           `json:"moderator_id"`
}

func (f *Flat) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id         uint64 `json:"id"`
		UnitNumber uint64 `json:"unit_number"`
		HouseId    uint64 `json:"house_id"`
		Price      uint64 `json:"price"`
		RoomNumber byte   `json:"rooms"`
		Status     string `json:"status"`
	}{
		Id:         f.Id,
		UnitNumber: f.UnitNumber,
		HouseId:    f.HouseId,
		Price:      f.Price,
		RoomNumber: f.RoomNumber,
		Status:     f.Status.String(),
	})
}
func (f *Flat) UnmarshalJSON(data []byte) error {
	flat := struct {
		Id          uint64 `json:"id"`
		UnitNumber  uint64 `json:"unit_number"`
		HouseId     uint64 `json:"house_id"`
		Price       uint64 `json:"price"`
		RoomNumber  byte   `json:"rooms"`
		Status      string `json:"status"`
		ModeratorId string `json:"moderator_id"`
	}{}
	err := json.Unmarshal(data, &flat)
	if err != nil {
		return err
	}
	switch flat.Status {
	case "on moderate":
		f.Status = OnModerate
	case "created":
		f.Status = Created
	case "approved":
		f.Status = Approved
	case "declined":
		f.Status = Declined
	default:
		return errors.New("unknown status")
	}
	f.Id = flat.Id
	f.HouseId = flat.HouseId
	f.Price = flat.Price
	f.RoomNumber = flat.RoomNumber
	f.ModeratorId = flat.ModeratorId
	f.UnitNumber = flat.UnitNumber
	return nil
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
