package optional

import (
	ytid "ioyoa/types/shared/id/yt"
	ytfields "ioyoa/types/shared/ytFields"
	"ioyoa/types/shared/ytFields/maximum"
	"ioyoa/types/shared/ytFields/page"
	"ioyoa/types/shared/ytFields/region"
)

type TOptional struct {
	hl *region.THl
	maxHeight *maximum.TMaxHeight
	maxResults *maximum.TMaxResults
	maxWidth *maximum.TMaxWidth
	onBehalfOfContentOwner *ytfields.TOnBehalfOfContentOwner
	pageToken *page.TPageToken
	regionCode *region.TRegionCode
	videoCategoryId *ytid.TVideoCategoryId
}