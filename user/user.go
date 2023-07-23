package user

import "croud-translate-core/constants"

type RatingStats struct {
	RatingList    []int
	RatingOverall int
}

type User struct {
	userId              int
	name                string
	accountStatus       constants.AccountStatus
	role                constants.RoleType
	ratingStats         RatingStats
	translationTextList []int
	commentList         []int
}
