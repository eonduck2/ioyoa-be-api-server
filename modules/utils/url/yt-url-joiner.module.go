package url

import (
	"ioyoa/modules/utils/env"
	staticYtEnv "ioyoa/static/env"
	"path"
)

func YtUrlJoiner(paths ...string) string {
	env.EnvLoader(staticYtEnv.YT_API_KEY)
	return path.Join(paths...)
}