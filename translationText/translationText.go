package translationText

import "croud-translate-core/constants"

type TranslationText struct {
	authorId          int
	ratingId          int
	translationTextId int
	formattedContent  string
	sourceContent     string
	viewStatus        constants.ViewStatus
}
