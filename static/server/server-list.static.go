package server

import (
	types "ioyoa/types/static/server"
)

var ServerList = []types.TServer{
	{Name:"Main", Path: "./servers/main.go"},
	{Name:"Activity", Path: "./servers/youtube/activity/activity-main.go"},
	{Name:"Channel", Path:"./servers/youtube/channel/channel-main.go"},
	{Name:"Comment", Path:"./servers/youtube/comment/comment-main.go"},
	{Name:"GA", Path:"./servers/google/ga/ga-main.go"},
	{Name:"Playlist", Path:"./servers/youtube/playlist/playlist-main.go"},
	{Name:"S3", Path:"./servers/aws/s3/s3-main.go"},
	{Name:"Thumbnail", Path:"./servers/youtube/thumbnail/thumbnail-main.go"},
	{Name:"User", Path:"./servers/youtube/user/user-main.go"},
	{Name:"Video", Path:"./servers/youtube/video/video-main.go"},
	{Name:"Search", Path:"./servers/youtube/search/search-main.go"},
	{Name:"Redis", Path:"./servers/redis/redis-main.go"},

}
