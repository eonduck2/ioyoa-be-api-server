package serverList

type Server struct {
	Name string
	Path string
}

func GetServerList() []Server {
	return []Server{
		{"Activity", "../servers/activity-api-server"},
		{"Channel", "../servers/channel-api-server"},
		{"Comment", "../servers/comment-api-server"},
		{"GA", "../servers/ga-api-server"},
		{"Playlist", "../servers/playlist-api-server"},
		{"S3", "../servers/s3-api-server"},
		{"Thumbnail", "../servers/thumbnail-api-server"},
		{"User", "../servers/user-api-server"},
		{"Video", "../servers/video-api-server"},
	}
}
