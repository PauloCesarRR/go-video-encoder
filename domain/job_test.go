package domain_test

import (
	"encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "Converted", video)

	require.Nil(t, err)
	require.NotNil(t, job)

}
