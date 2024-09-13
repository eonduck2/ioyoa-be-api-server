package url

import (
	"path"
)

func UrlJoiner(paths ...string) string {
	return path.Join(paths...)
}