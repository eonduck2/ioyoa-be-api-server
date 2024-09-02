package types

import (
	"ioyoa/types/env/gin"
	"ioyoa/types/env/proxy"
	"ioyoa/types/env/url"
)

// TEnvList 구조체를 정의합니다.
type TEnvList struct {
    GIN_MODE   gin.TGIN_MODE
    WL_PROXIES proxy.TWL_PROXIES
	MAIN_URL url.MAIN_URL
}
