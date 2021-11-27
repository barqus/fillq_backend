package youtube

import (
	"encoding/json"
	"github.com/barqus/fillq_backend/config"
	"github.com/barqus/fillq_backend/internal/common_http"
	"golang.org/x/net/context"
	"os"
)

type Service interface {
	GetChannelsVideos() error
}

type service struct {
	clt     common_http.Client
}

func MustNewService(client common_http.Client) *service {
	return &service{
		client,
	}
}

func (s *service) GetChannelsVideos() error {
	ctx := context.Background()
	apiKey:=os.Getenv("YOUTUBE_API_KEY")
	userTwitchURI := "https://www.googleapis.com/youtube/v3/search?key="+apiKey+"&channelId=UCIWa-X3jQrJqkTBEWnkg6OA&part=snippet,id&order=date&maxResults=20"
	rawResponse, statusCode, err := s.clt.Get(ctx, userTwitchURI, "", "", nil)
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return config.NO_BODY_FOUND
	}

	var responseYoutube YoutubeVideoFromAPI
	err = json.Unmarshal(rawResponse, &responseYoutube)
	if err != nil {
		return err
	}

	allYoutubeVideos := make([]*YoutubeVideo, 0)
	for _, item := range responseYoutube.Items {
		youtubeVideoItem := &YoutubeVideo{
			VideoID: item.ID.VideoID,
			PublishedAt: item.Snippet.PublishTime,
			Thumbnails: item.Snippet.Thumbnails,
		}
		allYoutubeVideos = append(allYoutubeVideos, youtubeVideoItem)
	}

	return nil
}
