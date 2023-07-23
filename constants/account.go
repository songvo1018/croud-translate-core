package constants

type AccountStatus string

type RoleType string

var (
	ACTIVE    = AccountStatus("ACTIVE")
	BLOCKED   = AccountStatus("BLOCKED")
	MODERATOR = RoleType("MODERATOR")
	AUTHOR    = RoleType("AUTHOR")
)
