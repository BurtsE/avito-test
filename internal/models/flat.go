package models

var (
	OnModerate ModerationStatus = &onModerate{}
	Created    ModerationStatus = &created{}
	Approved   ModerationStatus = &approved{}
	Declined   ModerationStatus = &declined{}
)

type Flat struct {
	Id         uint64
	HouseId    uint64
	Price      uint64
	RoomNumber byte
	Status     ModerationStatus
}

type ModerationStatus interface{ isEnumStatus() }

type onModerate struct{ ModerationStatus }
type created struct{ ModerationStatus }
type approved struct{ ModerationStatus }
type declined struct{ ModerationStatus }
