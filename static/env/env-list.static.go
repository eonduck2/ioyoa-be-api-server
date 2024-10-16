package env

import types "ioyoa/types/env"

var EnvListUsedByServer types.TEnvList = types.TEnvList{
	GIN_MODE: "GIN_MODE",
	WL_PROXIES: "WL_PROXIES",
	EP_MAIN: "EP_MAIN",
	EP_ACTIVITY: "EP_ACTIVITY",
	EP_CHANNEL: "EP_CHANNEL",
	EP_COMMENT: "EP_COMMENT",
	EP_GA: "EP_GA",
	EP_PLAYLIST: "EP_PLAYLIST",
	EP_S3: "EP_S3",
	EP_THUMBNAIL: "EP_THUMBNAIL",
	EP_USER: "EP_USER",
	EP_VIDEO: "EP_VIDEO",
	EP_SEARCH: "EP_SEARCH",
	EP_REDIS: "EP_REDIS",
}

const (
	YT_API_URL = "YT_API_URL"

	YT_API_KEY = "YT_API_KEY"
)

const (
	REDIS_HOST = "REDIS_HOST"

	REDIS_PORT = "REDIS_PORT"

	REDIS_PW = "REDIS_PW"
)