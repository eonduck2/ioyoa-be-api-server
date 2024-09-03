package server

import (
	types "ioyoa/types/static/server"
)

var ServerList = []types.TServer{
	{Name:"Main", Path: "./servers/main.go"},
	{Name:"Activity", Path: "./servers/activityApiServer/activity-main.go"},
	{Name:"Channel", Path:"./servers/channelApiServer/channel-main.go"},
	{Name:"Comment", Path:"./servers/commentApiServer/comment-main.go"},
	{Name:"GA", Path:"./servers/gaApiServer/ga-main.go"},
	{Name:"Playlist", Path:"./servers/playlistApiServer/playlist-main.go"},
	{Name:"S3", Path:"./servers/s3ApiServer/s3-main.go"},
	{Name:"Thumbnail", Path:"./servers/thumbnailApiServer/thumbnail-main.go"},
	{Name:"User", Path:"./servers/userApiServer/user-main.go"},
	{Name:"Video", Path:"./servers/videoApiServer/video-main.go"},
}
