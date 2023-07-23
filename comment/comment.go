package comment

import "croud-translate-core/constants"

type Comment struct {
	commentId int
	authorId  int
	constants.ViewStatus
	translateTextId int
	content         string
}
