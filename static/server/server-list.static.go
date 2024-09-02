package server

import (
	types "ioyoa/types/static/server"
)

var ServerList = []types.TServer{
	{Name:"Main", Path: "./servers/main.go"},
	{Name:"Activity", Path: "./servers/activity-api-server/activity-main.go"},
	{Name:"Channel", Path:"./servers/channel-api-server/channel-main.go"},
	{Name:"Comment", Path:"./servers/comment-api-server/comment-main.go"},
	{Name:"GA", Path:"./servers/ga-api-server/ga-main.go"},
	{Name:"Playlist", Path:"./servers/playlist-api-server/playlist-main.go"},
	{Name:"S3", Path:"./servers/s3-api-server/s3-main.go"},
	{Name:"Thumbnail", Path:"./servers/thumbnail-api-server/thumbnail-main.go"},
	{Name:"User", Path:"./servers/user-api-server/user-main.go"},
	{Name:"Video", Path:"./servers/video-api-server/video-main.go"},
}
