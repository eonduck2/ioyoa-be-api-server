package fields

import (
	id "ioyoa/types/shared/id"
	ytfields "ioyoa/types/shared/ytFields"
	"ioyoa/types/shared/ytFields/core"
	"ioyoa/types/shared/ytFields/details"
	"ioyoa/types/shared/ytFields/location"
	"ioyoa/types/shared/ytFields/status"
)

type TPart struct {
	ContentDetails       *details.TContentDetails
	FileDetails          *details.TFileDetails
	Id                   *id.TId
	LiveStreamingDetails *details.TLiveStreamingDetails
	Localizations        *location.TLocalizations
	Player               *ytfields.TPlayer
	ProcessingDetails    *details.TProcessingDetails
	RecordingDetails     *details.TRecordingDetails
	Snippet              *core.TSnippet
	Statistics           *ytfields.TStatistics
	Status               *status.TStatus
	Suggestions          *ytfields.TSuggestions
	TopicDetails         *details.TTopicDetails
}