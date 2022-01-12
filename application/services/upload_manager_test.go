package services_test

import (
	"log"
	"os"
	"testing"

	"github.com/Andreis3/encoder-video-golang/application/services"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestVideoServiceUploadManager(t *testing.T) {

	video, repo := prepare()
	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("encoderas3test")

	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	videoUpload := services.NewVideoUpload()
	videoUpload.OutputBucket = "encoderas3test"
	videoUpload.VideoPath = os.Getenv("localStoragePath") + "/" + video.ID

	doneUpload := make(chan string)
	concurrency := 50
	go videoUpload.ProcessUpload(concurrency, doneUpload)

	result := <-doneUpload

	require.Equal(t, result, "upload completed")
}
