package services_test

import (
	"log"
	"testing"
	"time"

	"github.com/Andreis3/encoder-video-golang/application/repositories"
	"github.com/Andreis3/encoder-video-golang/application/services"
	"github.com/Andreis3/encoder-video-golang/domain"
	"github.com/Andreis3/encoder-video-golang/framework/database"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "estrelas-neutrons.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {

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

	err = videoService.Finish()
	require.Nil(t, err)
}
