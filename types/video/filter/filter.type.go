package filter

import (
	id "ioyoa/types/shared/id"
	"ioyoa/types/video/filter/fields"
)

type TFilter struct {
	chart *fields.TChart
	id *id.TId
	myRating *fields.TMyRating
}