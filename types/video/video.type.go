package video

import (
	"ioyoa/types/video/essential"
	"ioyoa/types/video/filter"
	"ioyoa/types/video/optional"
)

type TVideo struct {
	essential essential.TEssential
	filter filter.TFilter
	optional optional.TOptional
}