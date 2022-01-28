package youtube

import (
	"context"

	"go.uber.org/fx"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

type Service struct {
	youtubeService *youtube.Service
}

func NewService(lifecycle fx.Lifecycle) (*Service, error) {
	service, err := youtube.NewService(context.Background(), option.WithHTTPClient(getClient([]string{youtube.YoutubeUploadScope, youtube.YoutubeForceSslScope, youtube.YoutubepartnerScope})))
	if err != nil {
		return nil, err
	}

	return &Service{youtubeService: service}, nil
}
