package models

var (
	UserRole      EnumRole = &user{}
	ModeratorRole EnumRole = &moderator{}
)

type EnumRole interface{ isEnumRole() }

type Role struct{}

type user struct{ EnumRole }
type moderator struct{ EnumRole }

type DummyAuth struct {
	Role *string `json:"role"`
}
type DummyUser struct {
	Role EnumRole `json:"role"`
}

type User struct {
	Id       *string `json:"id"`
	Role     EnumRole
	Password *string `json:"password"`
}
