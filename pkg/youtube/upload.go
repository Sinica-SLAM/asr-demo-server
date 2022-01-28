package youtube

import (
	"io"

	"google.golang.org/api/googleapi"
	youtube "google.golang.org/api/youtube/v3"
)

func (s Service) UploadVideo(title string, file io.Reader) (*youtube.Video, error) {
	video := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title: title,
			Tags:  []string{"test"},
		},
		Status: &youtube.VideoStatus{PrivacyStatus: "private", MadeForKids: false},
	}

	response, err := s.youtubeService.Videos.Insert([]string{"snippet", "status"}, video).Media(file, googleapi.ContentType("application/octet-stream")).Do()
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s Service) UploadCaptions(videoId string, file io.Reader) (*youtube.Caption, error) {
	caption := &youtube.Caption{
		Snippet: &youtube.CaptionSnippet{
			VideoId:  videoId,
			Language: "zh-TW",
			Name:     "zh-TW",
		},
	}
	response, err := s.youtubeService.Captions.Insert([]string{"snippet"}, caption).Media(file).Do()
	if err != nil {
		return nil, err
	}

	return response, nil
}
