package video

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type videoInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// VideoResponse object for destructuring the required data.
type VideoResponse struct {
	PlayabilityStatus bool      `json:"playabilityStatus"`
	VideoDetails      videoInfo `json:"videoDetails"`
}

// GetVideoInfo function for fetching the video details.
func GetVideoInfo(id string) (*VideoResponse, error) {
	url := fmt.Sprintf("https://www.youtube.com/get_video_info?video_id=%s&eurl=https://youtube.googleapis.com/v/%s", id, id)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	videoData := VideoResponse{}
	err = json.Unmarshal(body, &videoData)

	return &videoData, nil
}
