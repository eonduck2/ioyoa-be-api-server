package types

import (
	"ioyoa/types/env/gin"
	"ioyoa/types/env/proxy"
	"ioyoa/types/env/url"
	"ioyoa/types/env/url/aws"
	"ioyoa/types/env/url/google"
	"ioyoa/types/env/url/youtube"
)

type TEnvList struct {
    GIN_MODE   gin.TGIN_MODE
    WL_PROXIES proxy.TWL_PROXIES
	EP_MAIN url.MAIN_URL
	EP_ACTIVITY youtube.ACTIVITY_URL
	EP_CHANNEL youtube.CHANNEL_URL
	EP_COMMENT youtube.COMMENT_URL
	EP_GA google.GA_URL
	EP_PLAYLIST youtube.PLAYLIST_URL
	EP_S3 aws.S3_URL
	EP_THUMBNAIL youtube.THUMBNAIL_URL
	EP_USER youtube.USER_URL
	EP_VIDEO youtube.VIDEO_URL
	EP_SEARCH youtube.SEARCH_URL
}
