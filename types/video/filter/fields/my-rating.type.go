package fields

import "ioyoa/types/shared/ytFields/rating"

type TMyRating struct {
	like *rating.TLike
	disLike *rating.TDislike
}