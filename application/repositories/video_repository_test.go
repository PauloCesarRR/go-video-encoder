package repositories_test

import (
	"github.com/PauloCesarRR/go-video-encoder/application/repositories"
	"github.com/PauloCesarRR/go-video-encoder/domain"
	"github.com/PauloCesarRR/go-video-encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)
	v, err := repo.Find(video.ID)

	if err != nil {
		t.Fatal(err)
	}

	require.NotEmpty(t, video.ID)
	require.Equal(t, video.ID, v.ID)
	require.Equal(t, video.FilePath, v.FilePath)

}
