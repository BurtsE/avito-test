package models

var (
	OnModerate EnumRole = &onModerate{}
	Created    EnumRole = &created{}
	Approved   EnumRole = &approved{}
	Declined   EnumRole = &declined{}
)

type Flat struct {
	Id         uint32
	Cost       uint64
	FlatNumber uint32
	Status     FlatStatus
}

type FlatStatus interface{ isEnumStatus() }

type onModerate struct{ EnumRole }
type created struct{ EnumRole }
type approved struct{ EnumRole }
type declined struct{ EnumRole }
