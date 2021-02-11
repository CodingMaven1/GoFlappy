package main

import (
	"fmt"
	"os"

	"github.com/CodingMaven1/GoFlappy/video"
)

func main() {
	response := &video.VideoResponse{}
	response, err := video.GetVideoInfo("XXYlFuWEuKI")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(*response)
}
