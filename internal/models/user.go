package models

var (
	UserRole      EnumRole = &user{}
	ModeratorRole EnumRole = &moderator{}
)

type EnumRole interface{ isEnumRole() }

type user struct{ EnumRole }
type moderator struct{ EnumRole }
