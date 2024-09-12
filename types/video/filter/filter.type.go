package filter

import (
	ytid "ioyoa/types/shared/ytId"
	"ioyoa/types/video/filter/fields"
)

type TFilter struct {
	chart *fields.TChart
	id *ytid.TId
	myRating *fields.TMyRating
}